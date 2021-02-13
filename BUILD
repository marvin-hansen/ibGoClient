workspace(name = "ibGoClient")

# These rules are built-into Bazel but we need to load them first to download more rules
# https://docs.bazel.build/versions/master/be/workspace.html#http_archive
# https://docs.bazel.build/versions/master/repo/git.html
load("@bazel_tools//tools/build_defs/repo:http.bzl", "http_archive")
load("@bazel_tools//tools/build_defs/repo:git.bzl", "git_repository")

################################################################################
# Container dependencies
################################################################################

http_archive(
    name = "io_bazel_rules_docker",
    sha256 = "1698624e878b0607052ae6131aa216d45ebb63871ec497f26c67455b34119c80",
    strip_prefix = "rules_docker-0.15.0",
    urls = ["https://github.com/bazelbuild/rules_docker/releases/download/v0.15.0/rules_docker-v0.15.0.tar.gz"],
)

# Instantiate the "repositories" Bazel rule in rules_docker as "container_repositories"
load(
    "@io_bazel_rules_docker//repositories:repositories.bzl",
    container_repositories = "repositories",
)

container_repositories()

load(
    "@io_bazel_rules_docker//repositories:deps.bzl",
    container_deps = "deps",
)

container_deps()

################################################################################
# Docker authentication to push into any container registry
################################################################################
# Load the macro that allows you to customize the docker toolchain configuration.
load(
    "@io_bazel_rules_docker//toolchains/docker:toolchain.bzl",
    docker_toolchain_configure = "toolchain_configure",
)

# For authentication, run
# gcloud auth login
# gcloud auth configure-docker
# https://cloud.google.com/container-registry/docs/advanced-authentication#gcloud-helper
# docker custom configuration
# https://github.com/bazelbuild/rules_docker#container_push-custom-client-configuration
docker_toolchain_configure(
    name = "docker_config",
    # Replace this with an absolute path to a directory which has a custom docker client config.json.
    # Docker allows you to specify custom authentication credentials in the client configuration JSON file.
    # See https://docs.docker.com/engine/reference/commandline/cli/#configuration-files
    # client_config = "/path/to/.docker",
    # OPTIONAL: Path to the docker binary.
    # Should be set explicitly for remote execution.
    docker_path = "/usr/bin/docker",
    # OPTIONAL: Path to the gzip binary.
    # gzip_path="<enter absolute path to the gzip binary (in the remote exec env) here>",
)

################################################################################
# Containers to pull required to build.
################################################################################
load(
    "@io_bazel_rules_docker//container:container.bzl",
    "container_pull",
)

# docker pull debian:buster-slim
# https://hub.docker.com/_/debian
container_pull(
    name = "debian_slim",
    registry = "index.docker.io",
    repository = "library/debian",
    tag = "buster-slim",
)

# docker pull golang:buster
# https://hub.docker.com/_/golang
container_pull(
    name = "go_buster",
    registry = "index.docker.io",
    repository = "library/golang",
    tag = "buster",
)

# docker pull swift:5.3.2-focal
# https://hub.docker.com/_/swift
container_pull(
    name = "swift",
    registry = "index.docker.io",
    repository = "library/swift",
    tag = "5.3.2-focal",
)


################################################################################
#  gazelle
################################################################################

http_archive(
    name = "bazel_gazelle",
    urls = [
        "https://storage.googleapis.com/bazel-mirror/github.com/bazelbuild/bazel-gazelle/releases/download/0.18.2/bazel-gazelle-0.18.2.tar.gz",
        "https://github.com/bazelbuild/bazel-gazelle/releases/download/0.18.2/bazel-gazelle-0.18.2.tar.gz",
    ],
    sha256 = "7fc87f4170011201b1690326e8c16c5d802836e3a0d617d8f75c3af2b23180c4",
)

load("@bazel_gazelle//:deps.bzl", "gazelle_dependencies", "go_repository")

gazelle_dependencies()

################################################################################
#  go_repositories
################################################################################

