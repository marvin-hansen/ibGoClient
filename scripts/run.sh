# bin/bash
set -o errexit
set -o nounset
set -o pipefail


command bazel run //example:main


