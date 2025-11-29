Tool to programmatically generate git repositories from structured assets.

See [gl-exA-src](https://github.com/mj41/gl-exA-src) for an example of assets used to generate the `gl-exA` repository.

## How to generate gl-exA repository

Expected directory structure:

```
gl-exA/ - empty directory
gl-exA-scr/ - this repository
    assets/ - assets used to generate gl-exA repository
    rgen-conf.json - git-rgen configuration file
git-rgen - tool to generate git repositories programatically
```

Command to run:
```sh
go-run ./git-rgen/cmd/git-rgen --conf ./gl-exA-src/rgen-conf.json
cd ../gl-exA
git log --oneline --decorate --graph
```
