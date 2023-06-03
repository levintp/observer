load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "config",
    srcs = ["config.go"],
    importpath = "github.com/levintp/observer/internal/config",
    visibility = ["//observer:__subpackages__"],
    deps = [
        "//observer/internal/types",
        "@com_github_ilyakaznacheev_cleanenv//:cleanenv",
        "@in_gopkg_yaml_v3//:yaml_v3",
    ],
)
