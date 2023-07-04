load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "log",
    srcs = ["log.go"],
    importpath = "github.com/levintp/observer/internal/log",
    visibility = ["//:__subpackages__"],
    deps = [
        "@org_uber_go_zap//:zap",
        "@org_uber_go_zap//zapcore",
    ],
)
