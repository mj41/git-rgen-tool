package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"sort"
	"strings"
	"time"
)

// Config represents the rgen-conf.json structure
type Config struct {
	Defaults DefaultsConfig `json:"defaults"`
	Src      SrcConfig      `json:"src"`
	Dest     DestConfig     `json:"dest"`
}

type DefaultsConfig struct {
	Branch            string `json:"branch"`
	CommitMsg         string `json:"commit_msg"`
	AuthorName        string `json:"author_name"`
	AuthorEmail       string `json:"author_email"`
	AuthorDate        string `json:"author_date"`
	AuthorDateDiff    string `json:"author_date_diff"`
	CommitterName     string `json:"committer_name"`
	CommitterEmail    string `json:"committer_email"`
	CommitterDate     string `json:"committer_date"`
	CommitterDateDiff string `json:"committer_date_diff"`
}

type SrcConfig struct {
	Path string `json:"path"`
}

type DestConfig struct {
	Path string `json:"path"`
	Git  string `json:"git"`
}

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}

func run() error {
	// 1. Parse flags
	var configFile string
	flag.StringVar(&configFile, "conf", "rgen-conf.json", "path to configuration file")
	flag.Parse()

	// 2. Read config
	configData, err := os.ReadFile(configFile)
	if err != nil {
		return fmt.Errorf("failed to read config file %s: %w", configFile, err)
	}

	var config Config
	if err := json.Unmarshal(configData, &config); err != nil {
		return fmt.Errorf("failed to parse config file: %w", err)
	}

	// Validate config
	if config.Dest.Path == "" {
		return fmt.Errorf("dest.path is required")
	}
	if config.Src.Path == "" {
		return fmt.Errorf("src.path is required")
	}

	// Resolve paths
	// Paths are relative to the config file location
	configDir := filepath.Dir(configFile)

	resolvePath := func(p string) (string, error) {
		if filepath.IsAbs(p) {
			return p, nil
		}
		return filepath.Abs(filepath.Join(configDir, p))
	}

	destPath, err := resolvePath(config.Dest.Path)
	if err != nil {
		return fmt.Errorf("failed to resolve dest path: %w", err)
	}
	srcPath, err := resolvePath(config.Src.Path)
	if err != nil {
		return fmt.Errorf("failed to resolve src path: %w", err)
	}

	// 2. Check destination
	if exists, err := dirExists(destPath); err != nil {
		return err
	} else if exists {
		isEmpty, err := isDirEmpty(destPath)
		if err != nil {
			return err
		}
		if !isEmpty {
			return fmt.Errorf("destination path %s exists and is not empty", destPath)
		}
	} else {
		if err := os.MkdirAll(destPath, 0755); err != nil {
			return fmt.Errorf("failed to create destination directory: %w", err)
		}
	}

	// 3. Check source
	if exists, err := dirExists(srcPath); err != nil {
		return err
	} else if !exists {
		return fmt.Errorf("source path %s does not exist", srcPath)
	}

	// 4. Init git repo
	if err := runGit(destPath, "init"); err != nil {
		return fmt.Errorf("failed to git init: %w", err)
	}

	// Set default branch if specified
	if config.Defaults.Branch != "" {
		if err := runGit(destPath, "branch", "-M", config.Defaults.Branch); err != nil {
			return fmt.Errorf("failed to set branch name: %w", err)
		}
	}

	if config.Dest.Git != "" {
		if err := runGit(destPath, "remote", "add", "origin", config.Dest.Git); err != nil {
			return fmt.Errorf("failed to add remote origin: %w", err)
		}
	}

	// 5. Iterate assets
	entries, err := os.ReadDir(srcPath)
	if err != nil {
		return fmt.Errorf("failed to read src dir: %w", err)
	}

	var assetDirs []string
	for _, entry := range entries {
		if entry.IsDir() {
			assetDirs = append(assetDirs, entry.Name())
		}
	}
	sort.Strings(assetDirs)

	// Parse initial dates
	authorTime, err := time.Parse(time.RFC3339, config.Defaults.AuthorDate)
	if err != nil {
		return fmt.Errorf("invalid author_date: %w", err)
	}
	committerTime, err := time.Parse(time.RFC3339, config.Defaults.CommitterDate)
	if err != nil {
		return fmt.Errorf("invalid committer_date: %w", err)
	}

	authorDiff, err := time.ParseDuration(config.Defaults.AuthorDateDiff)
	if err != nil {
		return fmt.Errorf("invalid author_date_diff: %w", err)
	}
	committerDiff, err := time.ParseDuration(config.Defaults.CommitterDateDiff)
	if err != nil {
		return fmt.Errorf("invalid committer_date_diff: %w", err)
	}

	for i, dirName := range assetDirs {
		assetDir := filepath.Join(srcPath, dirName)
		fmt.Printf("Processing %s...\n", dirName)

		// Calculate dates
		if i > 0 {
			authorTime = authorTime.Add(authorDiff)
			committerTime = committerTime.Add(committerDiff)
		}

		// Check for overrides
		currentAuthorName := readFileOr(filepath.Join(assetDir, "author_name"), config.Defaults.AuthorName)
		currentAuthorEmail := readFileOr(filepath.Join(assetDir, "author_email"), config.Defaults.AuthorEmail)
		currentCommitterName := readFileOr(filepath.Join(assetDir, "committer_name"), config.Defaults.CommitterName)
		currentCommitterEmail := readFileOr(filepath.Join(assetDir, "committer_email"), config.Defaults.CommitterEmail)

		if val, ok := readFile(filepath.Join(assetDir, "author_date")); ok {
			if t, err := time.Parse(time.RFC3339, val); err == nil {
				authorTime = t
			}
		}
		if val, ok := readFile(filepath.Join(assetDir, "committer_date")); ok {
			if t, err := time.Parse(time.RFC3339, val); err == nil {
				committerTime = t
			}
		}

		commitMsg := readFileOr(filepath.Join(assetDir, "commit_msg"), config.Defaults.CommitMsg)

		// Copy tree
		treeDir := filepath.Join(assetDir, "tree")
		if exists, _ := dirExists(treeDir); exists {
			if err := copyDir(treeDir, destPath); err != nil {
				return fmt.Errorf("failed to copy tree from %s: %w", dirName, err)
			}
		}

		// Delete files
		deletedFilesPath := filepath.Join(assetDir, "deleted_files")
		if exists, _ := fileExists(deletedFilesPath); exists {
			content, err := os.ReadFile(deletedFilesPath)
			if err != nil {
				return fmt.Errorf("failed to read deleted_files in %s: %w", dirName, err)
			}
			lines := strings.Split(string(content), "\n")
			for _, line := range lines {
				line = strings.TrimSpace(line)
				if line == "" {
					continue
				}
				fileToDelete := filepath.Join(destPath, line)
				if err := os.RemoveAll(fileToDelete); err != nil {
					return fmt.Errorf("failed to delete file %s: %w", line, err)
				}
			}
		}

		// Git add all
		if err := runGit(destPath, "add", "."); err != nil {
			return fmt.Errorf("failed to git add: %w", err)
		}

		// Git add -u to stage deletions
		if err := runGit(destPath, "add", "-u"); err != nil {
			return fmt.Errorf("failed to git add -u: %w", err)
		}

		// Git commit
		env := os.Environ()
		env = append(env, fmt.Sprintf("GIT_AUTHOR_NAME=%s", currentAuthorName))
		env = append(env, fmt.Sprintf("GIT_AUTHOR_EMAIL=%s", currentAuthorEmail))
		env = append(env, fmt.Sprintf("GIT_AUTHOR_DATE=%s", authorTime.Format(time.RFC3339)))
		env = append(env, fmt.Sprintf("GIT_COMMITTER_NAME=%s", currentCommitterName))
		env = append(env, fmt.Sprintf("GIT_COMMITTER_EMAIL=%s", currentCommitterEmail))
		env = append(env, fmt.Sprintf("GIT_COMMITTER_DATE=%s", committerTime.Format(time.RFC3339)))

		if err := runGitWithEnv(destPath, env, "commit", "-m", commitMsg); err != nil {
			return fmt.Errorf("failed to git commit in %s: %w", dirName, err)
		}
	}

	return nil
}

func dirExists(path string) (bool, error) {
	info, err := os.Stat(path)
	if err == nil {
		return info.IsDir(), nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

func fileExists(path string) (bool, error) {
	info, err := os.Stat(path)
	if err == nil {
		return !info.IsDir(), nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

func isDirEmpty(path string) (bool, error) {
	f, err := os.Open(path)
	if err != nil {
		return false, err
	}
	defer f.Close()

	_, err = f.Readdirnames(1)
	if err == io.EOF {
		return true, nil
	}
	return false, err
}

func runGit(dir string, args ...string) error {
	return runGitWithEnv(dir, os.Environ(), args...)
}

func runGitWithEnv(dir string, env []string, args ...string) error {
	cmd := exec.Command("git", args...)
	cmd.Dir = dir
	cmd.Env = env
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

func copyDir(src, dst string) error {
	// Recursive copy
	return filepath.Walk(src, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		relPath, err := filepath.Rel(src, path)
		if err != nil {
			return err
		}
		if relPath == "." {
			return nil
		}
		destPath := filepath.Join(dst, relPath)

		if info.IsDir() {
			return os.MkdirAll(destPath, info.Mode())
		}

		return copyFile(path, destPath, info.Mode())
	})
}

func copyFile(src, dst string, mode os.FileMode) error {
	sourceFile, err := os.Open(src)
	if err != nil {
		return err
	}
	defer sourceFile.Close()

	destFile, err := os.OpenFile(dst, os.O_RDWR|os.O_CREATE|os.O_TRUNC, mode)
	if err != nil {
		return err
	}
	defer destFile.Close()

	_, err = io.Copy(destFile, sourceFile)
	return err
}

func readFileOr(path, def string) string {
	if val, ok := readFile(path); ok {
		return val
	}
	return def
}

func readFile(path string) (string, bool) {
	content, err := os.ReadFile(path)
	if err != nil {
		return "", false
	}
	return strings.TrimSpace(string(content)), true
}
