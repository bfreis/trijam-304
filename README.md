# github.com/bfreis/trijam-304

## Single Button Maze

This is an experiment: how long would it take for me to direct an AI to create a very simple game?

Around 95% of the code was written by Cursor AI, with a few iterations to make sure it got everything correct. The other 5% were manual fixes for things it couldn't get right.

This little game has been submitted to Trijam #304, in Itch.io: https://itch.io/jam/trijam-304/rate/3234536

## Pre-requisites

You need a modern version of Go (at least Go 1.21), ideally with `GOTOOLCHAIN=auto` so your Go toolchain downloads a more modern one if necessary.

If you wish to build a browser distribution (webassembly), the easiest way is to ensure you have [Task](https://taskfile.dev/) installed.

For rapid iteration with a local webserver, you may also want to have Python installed.

## How to run?

On your terminal, run:

```
$ go run .
```

If this is the first time building an Ebitengine app, it may take a while.

Once the process is done, the game should open on a new window.

This process has been tested on macOS Sequoia 15.2 on arm64. It should also run on any other OS and platforms supported by Ebitengine (i.e., Windows, Linux, amd64, arm64).

## How to run on the browser?

If all you want is to run the game from the browser without worrying about build artifacts and bundling and whatnot, run:

```
$ task run:web
```

This command will show the address where the game is running (typically `(http://127.0.0.1:8000/`).


### More detailed build and bundle process

To build the WebAssembly and prepare the necessary artifacts, assuming you have [Task](https://taskfile.dev/) installed, run:

```
$ task build:web
```

The resulting files will be placed in `platforms/web/*`.

If you wish to bundle a zip file with the game, run:

```
$ task bundle:web
```

The bundle will be placed in `platforms/web/bundle.zip` and contain the other artifacts.
