load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "plugins",
    srcs = ["plugins.go"],
    importpath = "github.com/levintp/observer/internal/plugins",
    visibility = ["//observer:__subpackages__"],
    deps = [
        "//observer/internal/config",
        "//observer/internal/types",
    ],
)
