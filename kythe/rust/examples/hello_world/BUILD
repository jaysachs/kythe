load("@rules_rust//rust:rust.bzl", "rust_binary")

rust_binary(
    name = "hello_world",
    srcs = ["src/main.rs"],
    edition = "2018",
    deps = ["//kythe/rust/cargo:colored"],
)
