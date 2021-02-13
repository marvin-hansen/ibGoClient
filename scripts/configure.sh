# bin/bash
set -o errexit
set -o nounset
set -o pipefail

command bazel run //:gazelle -- update-repos -from_file=go.mod

command  bazel run //:gazelle


echo "To build the project, "
echo "please run the following command: "
echo "==============================="
echo "* make build "
echo "==============================="
echo ""
