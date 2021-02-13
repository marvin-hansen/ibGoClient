load("@io_bazel_rules_go//go:def.bzl", "go_binary", "go_library")

package(default_visibility = ["//visibility:public"])

load("@bazel_gazelle//:def.bzl", "DEFAULT_LANGUAGES", "gazelle", "gazelle_binary")

gazelle_binary(
    name = "gazelle_binary",
    languages = DEFAULT_LANGUAGES + ["@golink//gazelle/go_link:go_default_library"],
    visibility = ["//visibility:public"],
)

# gazelle:prefix ibGoClient/
gazelle(
    name = "gazelle",
    gazelle = "//:gazelle_binary",
)

go_library(
    name = "go_default_library",
    srcs = ["main.go"],
    importpath = "ibGoClient/",
    deps = [
        "//ib:go_default_library",
        "@com_github_hadrianl_ibapi//:go_default_library",
    ],
)

go_binary(
    name = "ibGoClient",
    embed = [":go_default_library"],
)
