load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "go_default_library",
    srcs = [
        "contract_utils.go",
        "error_utils.go",
    ],
    importpath = "quantum/ib",
    visibility = ["//visibility:public"],
    deps = ["@com_github_hadrianl_ibapi//:go_default_library"],
)
