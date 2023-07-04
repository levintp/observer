load("@io_bazel_rules_go//go:def.bzl", "go_binary", "go_library")

go_library(
    name = "observerdbd_lib",
    srcs = ["observerdbd.go"],
    importpath = "github.com/levintp/observer/cmd/observerdbd",
    visibility = ["//visibility:private"],
    deps = [
        "//internal/config",
        "//internal/log",
    ],
)

go_binary(
    name = "observerdbd",
    embed = [":observerdbd_lib"],
    visibility = ["//visibility:public"],
)
