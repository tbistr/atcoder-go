# atcoder-go

This is go-library and cli-tool for [atcoder.jp](https://atcoder.jp/).
*Under development.*

## Architecture and develop environment

Directoriy -> description

- `/.devcontainer` -> This prj supports devcontainer. Please use [this](https://github.com/tbistr/golang-vscode-devcontainer).
- `/atcodergo` -> Library for golang. First, you should see [`client.go`](https://github.com/tbistr/atcoder-go/blob/main/atcodergo/client.go).
- `/cmd` -> CLI tool which uses atcodergo. `/cmd/*.go` only parse flags, initialize configs and call handlers.
  - `/handler` -> `/handler/*.go` takes responsiblity for core procedure.
- `/example` -> Example program. Run `go run ./example/main.go`. It has some unused functions, so edit (comments out) for show more complex ones.

## TODOs

- add tests
- set log level
  - and actualy output logs
- support custom template
  - use external program
  - generate boilerplate for each tasks
  - [input signeture] -> [custom program] -> [template]
- show submit results
  - lib-level and cli-level
- switch language (jp or en)
