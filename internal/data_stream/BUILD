load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "streams",
    srcs = ["streams.go"],
    importpath = "github.com/levintp/observer/internal/streams",
    visibility = ["//observer:__subpackages__"],
    deps = [
        "//observer/internal/config",
        "//observer/internal/log",
        "//observer/internal/plugins",
        "//observer/internal/types",
    ],
)

go_library(
    name = "data_stream",
    srcs = ["data_stream.go"],
    importpath = "github.com/levintp/observer/internal/data_stream",
    visibility = ["//observer:__subpackages__"],
    deps = [
        "//observer/internal/config",
        "//observer/internal/log",
        "//observer/internal/plugins",
        "//observer/internal/types",
    ],
)
