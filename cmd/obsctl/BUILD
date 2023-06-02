load("@io_bazel_rules_go//go:def.bzl", "go_binary", "go_library")

go_library(
    name = "obsctl_lib",
    srcs = ["obsctl.go"],
    importpath = "github.com/levintp/observer/cmd/obsctl",
    visibility = ["//visibility:private"],
)

go_binary(
    name = "obsctl",
    embed = [":obsctl_lib"],
    visibility = ["//visibility:public"],
)
