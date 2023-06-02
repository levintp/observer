load("@bazel_gazelle//:def.bzl", "gazelle")

# gazelle:prefix github.com/levintp/observer
gazelle(
    name = "gazelle",
)

gazelle(
    name = "gazelle-update-repos",
    args = [
        "-from_file=observer/go.mod",
        "-to_macro=observer/deps.bzl%observer_go_dependencies",
        "-prune",
        "-build_file_proto_mode=disable_global",
    ],
    command = "update-repos",
)
