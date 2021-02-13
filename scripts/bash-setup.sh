# bin/bash
set -o errexit
set -o nounset
set -o pipefail

command echo "export CC=clang" >>  ~/.bashrc