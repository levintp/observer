load("@io_bazel_rules_go//go:def.bzl", "go_binary", "go_library")

go_library(
    name = "observerd_lib",
    srcs = ["observerd.go"],
    importpath = "github.com/levintp/observer/observer/cmd/observerd",
    visibility = ["//visibility:private"],
    deps = [
        "@com_github_ohler55_ojg//jp:go_default_library",
        "@com_github_ohler55_ojg//oj:go_default_library",
    ],
)

go_binary(
    name = "observerd",
    embed = [":observerd_lib"],
    visibility = ["//visibility:public"],
)
