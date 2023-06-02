load("@bazel_gazelle//:deps.bzl", "go_repository")

def observer_go_dependencies():
    go_repository(
        name = "com_github_ohler55_ojg",
        build_file_proto_mode = "disable_global",
        importpath = "github.com/ohler55/ojg",
        sum = "h1:tzn5LJtkSyXowCo8SlGieU0zEc7WF4143Ri9MYlQy7Q=",
        version = "v1.18.5",
    )
    go_repository(
        name = "com_github_pelletier_go_toml",
        build_file_proto_mode = "disable_global",
        importpath = "github.com/pelletier/go-toml",
        sum = "h1:4yBQzkHv+7BHq2PQUZF3Mx0IYxG7LsP222s7Agd3ve8=",
        version = "v1.9.5",
    )
