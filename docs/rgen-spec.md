# Specification

This document specifies the structure of assets used by `git-rgen` tool to programmatically generate git repositories.

## Source repository structure

The source repository should have the following structure:

See [gl-exA-src](https://github.com/mj41/gl-exA-src) for an example.

[rgen-conf.json](https://github.com/mj41/gl-exA-src/blob/main/rgen-conf.json) is the main configuration file. All referenced paths are relative to its location.

[assets/](https://github.com/mj41/gl-exA-src/tree/main/assets) directory contains the assets used to generate the git repository commits:
- [assets/001-init-repo/](https://github.com/mj41/gl-exA-src/tree/main/assets/001-init-repo) - assets for the initial commit
- [assets/002-commit-spec/](https://github.com/mj41/gl-exA-src/tree/main/assets/002-commit-spec) - assets for the second commit
- [assets/003-more-files/](https://github.com/mj41/gl-exA-src/tree/main/assets/003-more-files) - assets for the third commit
- ...

Directories are processed in lexicographical order to determine the commit order. Prefix numbers (e.g. `001-`, `002-`, ...) are recommended to ensure correct ordering.

### Commit assets

Each commit assets directory (e.g. `assets/001-init-repo/`) should contain:
- `tree/` directory - the file tree to be committed
- `commit_msg` file - the commit message
- `deleted_files` file - a list of files to be deleted in this commit (relative to the repository root)

Files in `tree/` are copied over the existing repo. Existing files not present in tree/ are preserved. The file in `tree/` fully replaces the existing file. Partial updates (patching) are not supported.

Optionally it can also contain these files to override defaults from `rgen-conf.json`:
- author_name
- author_email
- author_date
- committer_name
- committer_email
- committer_date

Files like `author_name` or `commit_msg` in asset directories are whitespace-trimmed.

All defaults must be specified in the `rgen-conf.json` configuration file.

File permissions are copied from the source files in `tree/` to the generated repository.

Empty directories in `tree/` are not created in the git repo unless they contain a placeholder file (e.g., `.gitkeep`).

### deleted_files file

Example content of `deleted_files` file:
```
file1.txt
dir1/file2.txt
```

Exact relative paths from the repo root, one per line. Wildcards/globs are not supported.

## rgen-conf.json configuration file

The `rgen-conf.json` file specifies the configuration for `git-rgen` tool. Example:

```json
{
    "defaults": {
        "branch": "main",
        "commit_msg": "Auto generate commit",
        "author_name": "GitRgen Bot",
        "author_email": "git-rgen-bot@mj41.cz",
        "author_date": "2025-11-22T08:16:32Z",
        "author_date_diff": "25h1m1s",
        "committer_name": "GitRgen Bot",
        "committer_email": "git-rgen-bot@mj41.cz",
        "committer_date": "2025-11-22T09:18:36Z",
        "committer_date_diff": "25h1m1s"
    },
    "src": {
        "path": "assets/"
    },
    "dest": {
        "path": "../my-generated-repo1",
        "git": "git@github.com:ghuser/my-generated-repo1"
    }
}
```

All fields are required.

## Generation process

`git-rgen` 
- reads the `rgen-conf.json` configuration file.
- checks that destination paths do not already exist or is empty.
- checks that source paths exist.
- creates destination directories if they do not exist.
- initializes git repositories in the destination paths.
- iterates over the commit assets directories in the source path (e.g. `assets/001-init-repo/`, `assets/002-commit-spec/`, ...).
    - For each commit assets directory:
    - It copies files from the `tree/` directory to the repository.
    - It deletes files listed in the `deleted_files` file.
    - It creates a commit with the specified or default author, committer, and commit message.

If configuration option `dest.git` is present, the tool also configures the remote URL for the generated repository.

### Date handling

The tool calculates the date for commit `N` by adding `author_date_diff`/`committer_date_diff` to the date of commit `N−1`. The first commit uses the `author_date`/`committer_date` from the config.

author_date_diff` and `committer_date_diff` format follow the Go time.ParseDuration format (e.g., 25h1m1s, 3s).

## Error handling

The tool should provide clear error messages and exit with a non-zero status code.

## Out of scope

These features are not yet specified nor supported:
- Empty commits.
- File patches.
- Branching.
- Tags.
- Content generators.
- Templating.
- Extension or plugin system.
- Symlinks.
- GPG signing of commits.

Pull requests are welcome to add these features!
