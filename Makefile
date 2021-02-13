# Make will use bash instead of sh
SHELL := /usr/bin/env bash
CC=clang # required by bazel

# GNU make man page
# http://www.gnu.org/software/make/manual/make.html


.PHONY: help
help:
	@echo ''
	@echo '    make all   			Installs, configures, builds, and runs the project in one command.'
	@echo '    make setup   		Installs all requirements and tools.'
	@echo '    make configure		Configures bazel and gazelle to manage dependencies semi-automatically.'
	@echo '    make build		    	Build the project with bazel.'
	@echo '    make run			Runs the binary with bazel.'

.PHONY: all
all: setup  configure build run

.PHONY: setup
setup:
	@source scripts/setup.sh

.PHONY: configure
configure:
	@source scripts/configure.sh

.PHONY: build
build:
	@source scripts/build.sh

.PHONY: run
run:
	@source scripts/run.sh