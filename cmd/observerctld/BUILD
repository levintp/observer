load("@io_bazel_rules_go//go:def.bzl", "go_binary", "go_library")

go_library(
    name = "observerctld_lib",
    srcs = ["observerctld.go"],
    importpath = "github.com/levintp/observer/observer/cmd/observerctld",
    visibility = ["//visibility:private"],
)

go_binary(
    name = "observerctld",
    embed = [":observerctld_lib"],
    visibility = ["//visibility:public"],
)
