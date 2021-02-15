# ibGoClient

Go client based on [GOLang IB API](https://github.com/hadrianl/ibapi) that connects to the IB Gateway and logs about everything to stdout. 
Sample usage in example folder.  

## Credit & reference 
* [GOLang IB API](https://github.com/hadrianl/ibapi)
* [IB API] (https://interactivebrokers.github.io/tws-api/)


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
 
## Dependencies management. 

After each go get package, run 

```
make configure
```

Which auto-converts go mod dependencies into bazel dependencies, auto-updates all BUILD files, 
and ensures build remain functional. For details, see docs folder. In general, go mod & bazel dependencies 
should remain in sync to ensure the project can be build  with plain go in case bazel isn't an option for any reason. 


Note, bazel can also build docker containers, publish images to registry, and deploy to kubernetes, but
none of that has been configured. Search for the bazel docker and bazel k8s rules for details when needed. 

 
## Known issues

* GoLand may shows some syntax errors but project builds & runs just fine.
* IbWrapper not implement. Just the client & an example how to use it.  
* Missing graceful shutdown in case of failed gateway connect. 