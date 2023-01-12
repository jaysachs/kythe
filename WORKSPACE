workspace(name = "io_kythe")

load("@bazel_tools//tools/build_defs/repo:http.bzl", "http_archive")
load("@bazel_tools//tools/build_defs/repo:utils.bzl", "maybe")
load("//:version.bzl", "MAX_VERSION", "MIN_VERSION", "check_version")

# Check that the user has a version between our minimum supported version of
# Bazel and our maximum supported version of Bazel.
check_version(MIN_VERSION, MAX_VERSION)


load("@bazel_tools//tools/build_defs/repo:http.bzl", "http_archive")
http_archive(
    name = "aspect_rules_ts",
    sha256 = "6406905c5f7c5ca6dedcca5dacbffbf32bb2a5deb77f50da73e7195b2b7e8cbc",
    strip_prefix = "rules_ts-1.0.5",
    url = "https://github.com/aspect-build/rules_ts/archive/refs/tags/v1.0.5.tar.gz",
)

##################
# rules_ts setup #
##################
# Fetches the rules_ts dependencies.
# If you want to have a different version of some dependency,
# you should fetch it *before* calling this.
# Alternatively, you can skip calling this function, so long as you've
# already fetched all the dependencies.
load("@aspect_rules_ts//ts:repositories.bzl", "rules_ts_dependencies")

rules_ts_dependencies(
    # This keeps the TypeScript version in-sync with the editor, which is typically best.
    ts_version_from = "//:package.json",

    # Alternatively, you could pick a specific version, or use
    # load("@aspect_rules_ts//ts:repositories.bzl", "LATEST_VERSION")
    # ts_version = LATEST_VERSION
)

# Fetch and register node, if you haven't already
load("@rules_nodejs//nodejs:repositories.bzl", "DEFAULT_NODE_VERSION", "nodejs_register_toolchains")

nodejs_register_toolchains(
    name = "node",
    node_version = DEFAULT_NODE_VERSION,
)

load("@aspect_rules_js//npm:npm_import.bzl", "npm_translate_lock")

npm_translate_lock(
    name = "npm",
    pnpm_lock = "//:pnpm-lock.yaml",
)

load("@npm//:repositories.bzl", "npm_repositories")

npm_repositories()



load("//:setup.bzl", "kythe_rule_repositories", "remote_java_repository")

kythe_rule_repositories()

# gazelle:repository_macro external.bzl%_go_dependencies
load("//:external.bzl", "kythe_dependencies")

kythe_dependencies()

load("//tools/build_rules/external_tools:external_tools_configure.bzl", "external_tools_configure")

external_tools_configure()

# load("@npm//@bazel/labs:package.bzl", "npm_bazel_labs_dependencies")

# npm_bazel_labs_dependencies()

load("@maven//:compat.bzl", "compat_repositories")

compat_repositories()

load("@maven//:defs.bzl", "pinned_maven_install")

pinned_maven_install()

load(
    "@bazelruby_rules_ruby//ruby:defs.bzl",
    "ruby_bundle",
)

ruby_bundle(
    name = "website_bundle",
    bundler_version = "2.1.4",
    gemfile = "//kythe/web/site:Gemfile",
    gemfile_lock = "//kythe/web/site:Gemfile.lock",
)

load("@rules_rust//crate_universe:defs.bzl", "crate", "crates_repository", "render_config")

# Run `CARGO_BAZEL_REPIN=1 bazel sync --only=crate_index` after updating
crates_repository(
    name = "crate_index",
    cargo_lockfile = "//:Cargo.Bazel.lock",
    lockfile = "//:cargo-bazel-lock.json",
    packages = {
        "anyhow": crate.spec(
            version = "1.0.58",
        ),
        "base64": crate.spec(
            version = "0.13.0",
        ),
        "clap": crate.spec(
            features = ["derive"],
            version = "3.1.6",
        ),
        "colored": crate.spec(
            version = "2.0.0",
        ),
        "glob": crate.spec(
            version = "0.3.0",
        ),
        "hex": crate.spec(
            version = "0.4.3",
        ),
        "lazy_static": crate.spec(
            version = "1.4.0",
        ),
        "quick-error": crate.spec(
            version = "2.0.1",
        ),
        "path-clean": crate.spec(
            version = "0.1.0",
        ),
        "rayon": crate.spec(
            version = "1.5.3",
        ),
        "regex": crate.spec(
            version = "1.5.6",
        ),
        "rls-analysis": crate.spec(
            version = "0.18.3",
        ),
        "rls-data": crate.spec(
            version = "0.19.1",
        ),
        "serde": crate.spec(
            version = "1.0.137",
        ),
        "serde_json": crate.spec(
            version = "1.0.64",
        ),
        "sha2": crate.spec(
            version = "0.10.2",
        ),
        "tempfile": crate.spec(
            version = "3.3.0",
        ),
        "zip": crate.spec(
            version = "0.5.11",
        ),
        # Dev dependency for fuchsia extractor
        "serial_test": crate.spec(
            version = "0.6.0",
        ),
        # Dependencies for our Rust protobuf toolchain
        "protobuf": crate.spec(
            features = ["with-bytes"],
            version = "=2.27.1",
        ),
        "protobuf-codegen": crate.spec(
            version = "=2.27.1",
        ),
    },
    render_config = render_config(
        default_package_name = "",
    ),
)

load("@crate_index//:defs.bzl", "crate_repositories")

crate_repositories()

# Register our Rust protobuf toolchain from the BUILD file
register_toolchains(
    ":rust_proto_toolchain",
)

# Bazel does not yet ship with JDK-19 compatible JDK, so configure our own.
remote_java_repository(
    name = "remotejdk19_linux",
    prefix = "remotejdk",
    sha256 = "2ac8cd9e7e1e30c8fba107164a2ded9fad698326899564af4b1254815adfaa8a",
    strip_prefix = "zulu19.30.11-ca-jdk19.0.1-linux_x64",
    target_compatible_with = [
        "@platforms//os:linux",
    ],
    urls = [
        "https://mirror.bazel.build/cdn.azul.com/zulu/bin/zulu19.30.11-ca-jdk19.0.1-linux_x64.tar.gz",
        "https://cdn.azul.com/zulu/bin/zulu19.30.11-ca-jdk19.0.1-linux_x64.tar.gz",
    ],
    version = "19",
)

register_toolchains("//buildenv/java:all")
