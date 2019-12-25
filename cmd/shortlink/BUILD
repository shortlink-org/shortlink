load("@io_bazel_rules_go//go:def.bzl", "go_binary")

go_binary(
    name = "shortlink",
    srcs = ["shortlink.go"],
    embed = [":go_repository"],
    importpath = "github.com/batazor/shortlink/cmd/shortlink",
    visibility = ["//visibility:public"],
)
