# bin/bash
set -o errexit
set -o nounset
set -o pipefail

command bazel build //ib:go_default_library



