load("//rules/typescript:index.bzl", "ts_library")

package(default_visibility = ["//visibility:public"])

ts_library(
    name = "footer",
    srcs = glob(["*.tsx"]),
    deps = [
        "//app/capabilities",
        "@npm//@types/react",
        "@npm//react",
    ],
)
