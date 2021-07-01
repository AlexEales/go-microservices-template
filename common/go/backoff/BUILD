go_library(
    name = "backoff",
    srcs = ["backoff.go"],
    visibility = ["PUBLIC"],
)

go_test(
    name = "backoff_test",
    srcs = ["backoff_test.go"],
    deps = [
        ":backoff",
        "//third_party/go:testify",
    ],
)
