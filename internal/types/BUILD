load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "types",
    srcs = ["types.go"],
    importpath = "github.com/levintp/observer/internal/types",
    visibility = ["//observer:__subpackages__"],
    deps = ["@in_gopkg_yaml_v3//:yaml_v3"],
)
