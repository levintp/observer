load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "data_stream",
    srcs = ["data_stream.go"],
    importpath = "github.com/levintp/observer/internal/data_stream",
    visibility = ["//:__subpackages__"],
    deps = [
        "//internal/config",
        "//internal/log",
        "//internal/plugins",
        "//internal/types",
    ],
)
