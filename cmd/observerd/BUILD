load("@io_bazel_rules_go//go:def.bzl", "go_binary", "go_library")

go_library(
    name = "observerd_lib",
    srcs = ["observerd.go"],
    importpath = "github.com/levintp/observer/cmd/observerd",
    visibility = ["//visibility:private"],
    deps = [
        "//observer/internal/config",
        "//observer/internal/data_stream",
        "//observer/internal/log",
    ],
)

go_binary(
    name = "observerd",
    embed = [":observerd_lib"],
    visibility = ["//visibility:public"],
)
