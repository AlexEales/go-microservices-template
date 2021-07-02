go_binary(
    name = "deploy-vault",
    srcs = ["deploy_vault.go"],
    deps = [
        "//common/go/backoff",
        "//common/go/retry",
        "//services/helm",
        "//services/k8s",
        "//services/minikube",
        "//third_party/go:flags",
        "//third_party/go:logrus",
        "//tools/helm",
    ],
)
