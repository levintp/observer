load("@io_bazel_rules_go//go:def.bzl", "go_binary", "go_library")

go_library(
    name = "observerapid_lib",
    srcs = ["observerapid.go"],
    importpath = "github.com/levintp/observer/observer/cmd/observerapid",
    visibility = ["//visibility:private"],
)

go_binary(
    name = "observerapid",
    embed = [":observerapid_lib"],
    visibility = ["//visibility:public"],
)
