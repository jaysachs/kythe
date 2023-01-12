package(default_visibility = ["//visibility:public"])

package_group(
    name = "default_visibility",
    packages = ["//kythe/..."],
)

package_group(
    name = "proto_visibility",
    includes = [":default_visibility"],
)

filegroup(
    name = "srcs",
    srcs = glob(
        ["**"],
        exclude = [
            "web/site/.bundle/**",
            "web/site/_vendor/**",
        ],
    ),
    visibility = ["//visibility:private"],
)

load("@npm//:defs.bzl", "npm_link_all_packages")

npm_link_all_packages(name = "node_modules")
