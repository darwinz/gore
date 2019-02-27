# gore [![Travis Build Status](https://travis-ci.org/motemen/gore.svg?branch=master)](https://travis-ci.org/motemen/gore)
### Yet another Go REPL that works nicely. Featured with line editing, code completion, and more.

![Screencast](doc/screencast.gif)

(Screencast taken with [cho45/KeyCast](https://github.com/cho45/KeyCast))

## Usage

```sh
gore
```
After a prompt is shown, enter any Go expressions/statements or commands described below.

To quit the session, type `Ctrl-D`.

## Features

- Line editing with history
- Multiline inputs
- Package importing with completion
- Evaluates any expressions or statements
- No "evaluated but not used"
- Code completion (requires [gocode](https://github.com/mdempsky/gocode))
- Pretty printing ([pp](https://github.com/k0kubun/pp) or
  [spew](https://github.com/davecgh/go-spew) recommended)
- Showing documents (requires [godoc](https://golang.org/x/tools/cmd/godoc))
- Auto-importing

## REPL Commands

Some functionalities are provided as colon-commands:

```
:import <package path>  Import package
:type <expr>            Print the type of expression
:print                  Show current source
:write [<filename>]     Write out current source to file
:clear                  Clear the codes
:doc <expr or pkg>      Show document (requires godoc)
:help                   List commands
:quit                   Quit the session
```

## Installation

gore uses Go toolchains, so I don’t provide binaries.

```sh
go get -u github.com/motemen/gore
```

Make sure `$GOPATH/bin` is in your `$PATH`.

Also recommended:

```sh
go get -u github.com/mdempsky/gocode # for code completion
go get -u github.com/k0kubun/pp # or github.com/davecgh/go-spew/spew
go get -u golang.org/x/tools/cmd/godoc # for using with the :doc colon-command
```

## FAQ/Caveats

- If you see `too many arguments in call to mainScope.LookupParent`
  while installing gore, run `go get -u golang.org/x/tools/go/types`.
- gore runs code using `go run` for each input. If you have entered
  time-consuming code, gore will run it for each input and take some
  time. Also errors shows the line where the actual code run is.
- To import a local package, first fetch it with `go get my/package`,
  then `:import` will work properly

## TODO

- Undoing input
- Configuration
- `:write` completion
- Direct editing of code
- Using external sources
- API

## License

[The MIT License](./LICENSE).

## Author

motemen &lt;<motemen@gmail.com>&gt;
