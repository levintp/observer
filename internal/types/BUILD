load("@io_bazel_rules_go//go:def.bzl", "go_library", "go_test")

go_library(
    name = "types",
    srcs = [
        "agent_spec.go",
        "api_spec.go",
        "category.go",
        "controller_spec.go",
        "database_spec.go",
        "metric_data.go",
        "metric_spec.go",
        "module_spec.go",
        "node_spec.go",
        "queue.go",
        "sample.go",
        "sampling_func.go",
        "stream_spec.go",
        "threshold_spec.go",
    ],
    importpath = "github.com/levintp/observer/internal/types",
    visibility = ["//:__subpackages__"],
)

go_test(
    name = "types_test",
    srcs = [
        "node_spec_test.go",
        "queue_test.go",
    ],
    deps = [":types"],
)
