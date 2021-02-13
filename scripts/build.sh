# bin/bash
set -o errexit
set -o nounset
set -o pipefail

command bazel build //ib:go_default_library

command bazel build //example:go_default_library
