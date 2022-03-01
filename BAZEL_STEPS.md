# Steps to use the bazel with go

1. Install the [Bazel](https://bazel.build/) tool. (preferably via [ASDF](https://asdf-vm.com/))
2. Install the buildifier tool: `go install github.com/bazelbuild/buildtools/buildifier@latest`
 
## Setup dependencies in WORKSPACE file
1. Create WORKSPACE file at root of your project.
2. Load `http_archive` rule in WORKSPACE file.
3. Using `http_archive` rule, download the `io_bazel_rule_go` and `bazel_gazelle` rules.
4. Load `go_register_toolchains` and `go_rules_dependencies` functions from `io_bazel_rule_go` rule in WORKSPACE file.
5. Load `gazelle_dependencies` and `go_repository` functions from `bazel_gazelle` rule in WORKSPACE file.

## Setup gazelle in BUILD file
1. Create BUILD file at root of your project.
2. Load the gazelle rule for code generation: `load("@io_bazel_rule_go//go:def.bzl", "gazelle_dependencies")`
3. Set the gazelle prefix (module base name): `# gazelle:prefix github.com/your_company/your_project_base`
4. Call gazelle `gazelle(name = "gazelle")` to generate the BUILD file for each submodule.