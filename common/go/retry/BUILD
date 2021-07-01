go_library(
    name = "retry",
    srcs = ["retry.go"],
    visibility = ["PUBLIC"],
    deps = [
        "//common/go/backoff",
    ],
)

go_test(
    name = "retry_test",
    srcs = ["retry_test.go"],
    deps = [
        ":retry",
        "//third_party/go:testify",
    ],
)
