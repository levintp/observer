load("@io_bazel_rules_go//go:def.bzl", "go_binary", "go_library")

go_library(
    name = "observerctld_lib",
    srcs = ["observerctld.go"],
    importpath = "github.com/levintp/observer/cmd/observerctld",
    visibility = ["//visibility:private"],
    deps = [
        "//observer/internal/config",
        "//observer/internal/log",
    ],
)

go_binary(
    name = "observerctld",
    embed = [":observerctld_lib"],
    visibility = ["//visibility:public"],
)
