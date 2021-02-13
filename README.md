# ibGoClient

## Conventions
* File encoding: UTF-8
* File separator: \n LF (Unix / Mac style) 


## Build requirements

* Ubuntu 20.04 LTS (Focal Fossa) environment i.e. Docker / WSL etc.  
* Bazel v4 (LTS) 
* Go 1.15
* Make 
* clang 
* export CC=clang in .bachrc as bazel depends on clang toolchain 

Make setup will install all build requirements, build & run the project. 

## Getting started:

1. git clone repo 
2. cd ibGoClient
3. make all 

Ensure IB Gateway runs on localhost:4003 

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
 

