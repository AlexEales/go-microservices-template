subinclude("//build/defs:docker", "//build/defs:k8s")

k8s_config(
    name = "postgres",
    srcs = [
        "deployment.yaml",
        "secrets.yaml",
        "service.yaml",
        "volumes.yaml",
    ],
)
