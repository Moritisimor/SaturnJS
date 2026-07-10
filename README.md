# SaturnJS
A JS Runtime intended for System-Scripting. Written in Go using the Goja engine

## What is this project about?
SaturnJS is a JavaScript runtime based on the [goja](https://github.com/dop251/goja) JS engine.

Unlike most other JS runtimes, SaturnJS is intended to be used for scripts that interact with the OS/FS.

It is still in development, and although the API is not supposed to break, it will be extended upon.

## Cloning & Compilation
### Prerequisites
- Latest [Go](https://go.dev) Compiler
- [Git](https://git-scm.com/install/)

### Script
```bash
git clone https://github.com/Moritisimor/SaturnJS
cd SaturnJS
go build -ldflags="-s -w" -o saturn cmd/saturn/main.go
```

## Usage
### Examples
You can find a collection of example scripts [here](https://github.com/Moritisimor/SaturnJS/tree/main/examples).

### REPL
To get started with the REPL, simply execute `saturn`:
```bash
./saturn
```

### Running Scripts
To run a script, simply supply the script path as a CLI argument:
```bash
./saturn my_script.js
```

You could run the `httpGetDemo.js` like this:
```bash
./saturn examples/httpGetDemo.js
```
