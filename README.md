# ibGoClient

Go client that connects to the IB Gateway and logs about everything to stdout. 
Sample usage in example folder.  

## Conventions
* File encoding: UTF-8
* File separator: \n LF (Unix / Mac style) 


## Build requirements

* Ubuntu 20.04 LTS (Focal Fossa) environment i.e. Docker / WSL etc.  
* Bazel v4 (LTS) 
* Go 1.15
* Make 
* clang 

All build requirements will be checked and installed with "make setup"

## Bazel setup

Run only once: 
```
./scripts/bash-setup.sh
```

This appends "export CC=clang" to .bachrc as bazel depends on the clang toolchain. 
Bazel itself will be installed during  "make setup" if its not already present. 

## Getting started:

1. git clone repo 
2. cd ibGoClient
3. make all 

Ensure IB Gateway already runs on localhost:4003 

### Run with make

```
 make all                    Installs, configures, builds and runs the project in one command.
 make setup                  Installs all requirements and tools.
 make configure              Configures bazel and gazelle to manage dependencies semi-automatically.
 make build                  Build the project with bazel.
 make run                    Runs the binary with bazel.
```

## Run with plain old go 

```
cd example
go build main.go  
go run main.go 
```

## Run with bazel 

```
 bazel run //example:main
```
 
## Known issues

* GoLand may shows some syntax errors but project builds & runs just fine.
* IbWrapper not implement. Just the client & an example how to use it.  
* Missing graceful shutdown in case of failed gateway connect. 