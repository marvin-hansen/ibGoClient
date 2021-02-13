# bin/bash
set -o errexit
set -o nounset
set -o pipefail

echo ""
echo "==============================="
echo " Configure build requirements :"
echo "==============================="
echo ""


# Check for system requirements silently and install missing if needed.
command sudo apt-get -qqq update

command make --version >/dev/null 2>&1 || {
  command sudo apt-get -qqq -y install gnupg
}

command curl --version >/dev/null 2>&1 || {
  command sudo apt-get -qqq -y install curl
}

command wget --version >/dev/null 2>&1 || {
  command sudo apt-get -qqq -y install wget
}

command gpg --version >/dev/null 2>&1 || {
  command sudo apt-get -qqq -y install gnupg
}

command git version >/dev/null 2>&1 || {
  echo "Install Git"
  command sudo apt-get -qqq -y install git
}
echo "* Git installed"

command docker --version >/dev/null 2>&1 || {
  echo "Install Docker"
  command sudo apt-get -qqq -y install docker.io
}
echo "* Docker installed"


command go version >/dev/null 2>&1 || {
  echo "Install Go-lang"
  command sudo tar -C /usr/local -xzf go1.15.8.linux-amd64.tar.gz
  command export PATH=$PATH:/usr/local/go/bin
  echo "Done"
  go version
}
echo "* Go-lang installed"

command java -version >/dev/null 2>&1 || {
  # https://www.digitalocean.com/community/tutorials/how-to-install-java-with-apt-on-ubuntu-20-04
  echo "Download open jdk"
  command sudo apt-get -qqq update
  echo "Install open jdk"
  command sudo apt-get -qqq -y install default-jdk
  echo "Done"
  command java -version
}
echo "* Java installed"

command bazel version >/dev/null 2>&1 || {
  # https://docs.bazel.build/versions/master/install-ubuntu.html
  echo "Download bazel"
  command curl -fsSL https://bazel.build/bazel-release.pub.gpg | gpg --dearmor >bazel.gpg
  command sudo mv bazel.gpg /etc/apt/trusted.gpg.d/
  command echo "deb [arch=amd64] https://storage.googleapis.com/bazel-apt stable jdk1.8" | sudo tee /etc/apt/sources.list.d/bazel.list
  command sudo apt-get -qqq -y update
  echo "Install bazel"
  command sudo apt-get -qqq -y install -y bazel
  echo "Done"
  bazel -version
}
echo "* Bazel installed"

echo ""
echo "==============================="
echo "All dependencies installed."
echo "==============================="
echo ""

echo "To configure the project if its not yet, "
echo "please run the following command: "
echo "==============================="
echo "* make configure "
echo "==============================="
echo ""
