load("@io_bazel_rules_go//go:def.bzl", "go_binary", "go_library")

go_library(
    name = "go_default_library",
    srcs = ["main.go"],
    importpath = "ibGoClient/example",
    visibility = ["//visibility:private"],
    deps = [
        "//ib:go_default_library",
        "@com_github_hadrianl_ibapi//:go_default_library",
        "@org_golang_x_net//context:go_default_library",
    ],
)

go_binary(
    name = "main",
    embed = [":go_default_library"],
    gc_linkopts = [
        "-linkmode",
        "external",
        "-extldflags",
        "-static",
        "-w",
        "-s",
    ],
    goarch = "amd64",
    goos = "linux",
    importpath = "services/smx",
    pure = "off",
    static = "on",
    visibility = ["//visibility:public"],
)
