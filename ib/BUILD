load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "go_default_library",
    srcs = [
        "contract_utils.go",
        "error_utils.go",
        "ib.go",
        "ib_env.go",
        "ib_wrapper.go",
        "structs.go",
    ],
    importpath = "ibGoClient/ib",
    visibility = ["//visibility:public"],
    deps = [
        "@com_github_hadrianl_ibapi//:go_default_library",
        "@org_uber_go_zap//:go_default_library",
    ],
)
