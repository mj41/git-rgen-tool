# git-rgen-tool

Tool to programmatically generate git repositories from structured assets.

## Example

The [gl-exA](https://github.com/mj41/gl-exA) repository (including the full git history) was generated from [gl-exA-src](https://github.com/mj41/gl-exA-src) assets.

Generate the repository:

```bash
cd git-rgen-tool
go run ./cmd/rgen --conf ../gl-exA-src/rgen-conf.json
```

See [docs/exA-log.md](docs/exA-log.md) for the full commit log of the generated repository.

## Specification

See [docs/rgen-spec.md](docs/rgen-spec.md) for the specification of the assets structure and how git-rgen works.
