# ibGoClient

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
 

