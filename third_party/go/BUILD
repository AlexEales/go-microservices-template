package(default_visibility = ["PUBLIC"])

go_module(
    name = "aec",
    module = "github.com/morikuni/aec",
    version = "v1.0.0",
)

go_module(
    name = "backoff",
    module = "github.com/cenkalti/backoff/v4",
    version = "v4.1.1",
)

go_module(
    name = "cobra",
    module = "github.com/spf13/cobra",
    version = "v1.1.3",
    deps = [
        ":pflag",
    ],
)

go_module(
    name = "containerd",
    module = "github.com/containerd/containerd",
    install = [
        "errdefs",
        "platforms",
    ],
    version = "v1.5.2",
    deps = [
        ":grpc",
        ":image-spec",
        ":logrus",
        ":pkg-errors",
    ],
)

go_module(
    name = "continuity",
    module = "github.com/containerd/continuity",
    install = ["pathdriver"],
    version = "v0.1.0",
)

go_module(
    name = "difflib",
    module = "github.com/pmezard/go-difflib",
    install = ["difflib"],
    version = "v1.0.0",
)

go_module(
    name = "docker",
    module = "github.com/docker/docker",
    install = [
        "api",
        "api/types/...",
        "client",
        "errdefs",
        "pkg/homedir",
        "pkg/system",
        "registry",
    ],
    version = "v20.10.7",
    deps = [
        ":aec",
        ":containerd",
        ":docker-distribution",
        ":go-connections",
        ":go-units",
        ":gogo-protobuf",
        ":image-spec",
        ":logrus",
        ":moby-term",
        ":moby-mount",
        ":xtime",
    ],
)

go_module(
    name = "docker-cli",
    module = "github.com/docker/cli",
    install = [
        "cli",
        "cli/compose/loader",
    ],
    version = "v20.10.7",
    deps = [
        ":cobra",
        ":docker",
        ":docker-credential-helper",
        ":docker-distribution",
        ":go-digest",
        ":gojsonschema",
        ":logrus",
        ":mapstructure",
        ":mergo",
        ":notary",
        ":pflag",
        ":pkg-errors",
        ":shlex",
        ":sortorder",
        ":xsys",
        ":yamlv2",
    ],
)

go_module(
    name = "docker-credential-helper",
    module = "github.com/docker/docker-credential-helpers",
    install = [
        "credentials",
        "client",
    ],
    version = "v0.6.4",
    deps = [
        ":xsys",
    ],
)

go_module(
    name = "docker-distribution",
    module = "github.com/docker/distribution",
    install = [
        ".",
        "manifest/...",
        "reference",
        "registry/api/errcode",
        "registry/api/v2",
        "registry/client",
        "registry/client/auth",
        "registry/client/auth/challenge",
        "registry/client/transport",
        "uuid",
    ],
    version = "v2.7.1",
    deps = [
        ":docker-metrics",
        ":go-digest",
        ":image-spec",
        ":libtrust",
        ":logrus",
        ":mux",
    ],
)

go_module(
    name = "docker-go",
    module = "github.com/docker/go",
    install = ["canonical/json"],
    version = "v1.5.1-1",
)

go_module(
    name = "docker-metrics",
    module = "github.com/docker/go-metrics",
    version = "v0.0.1",
    deps = [
        ":prometheus",
    ],
)

go_module(
    name = "dockertest",
    module = "github.com/ory/dockertest/v3",
    version = "v3.7.0",
    deps = [
        ":backoff",
        ":continuity",
        ":docker-cli",
        ":go-connections",
        ":go-units",
        ":gotestyourself",
        ":gotty",
        ":image-spec",
        ":libcontainer-user",
        ":pkg-errors",
    ],
)

go_module(
    name = "env",
    module = "github.com/Netflix/go-env",
    version = "v0.0.0-20210215222557-e437a7e7f9fb",
)

go_module(
    name = "flags",
    module = "github.com/jessevdk/go-flags",
    version = "v1.5.0",
    deps = [
        ":xsys",
    ],
)

go_module(
    name = "go-cmp",
    module = "github.com/google/go-cmp",
    install = ["cmp/..."],
    version = "v0.5.6",
)

go_module(
    name = "go-connections",
    module = "github.com/docker/go-connections",
    install = [
        "nat",
        "sockets",
        "tlsconfig",
    ],
    version = "v0.4.0",
    deps = [
        ":pkg-errors",
        ":xnet",
    ],
)

go_module(
    name = "go-digest",
    module = "github.com/opencontainers/go-digest",
    version = "v1.0.0",
)

go_module(
    name = "go-units",
    module = "github.com/docker/go-units",
    version = "v0.4.0",
)

go_module(
    name = "gogo-protobuf",
    module = "github.com/gogo/protobuf",
    install = ["proto"],
    version = "v1.3.2",
)

go_module(
    name = "gojsonpointer",
    module = "github.com/xeipuuv/gojsonpointer",
    version = "v0.0.0-20190905194746-02993c407bfb",
)

go_module(
    name = "gojsonreference",
    module = "github.com/xeipuuv/gojsonreference",
    version = "v0.0.0-20180127040603-bd5ef7bd5415",
    deps = [
        ":gojsonpointer",
    ],
)

go_module(
    name = "gojsonschema",
    module = "github.com/xeipuuv/gojsonschema",
    version = "v1.2.0",
    deps = [
        ":gojsonreference",
    ],
)

go_module(
    name = "google-protobuf",
    module = "google.golang.org/protobuf",
    install = ["..."],
    version = "v1.27.1",
    deps = [
        ":go-cmp",
    ],
)

go_module(
    name = "grpc",
    module = "google.golang.org/grpc",
    install = [
        ".",
        "codes",
        "credentials",
        "status",
    ],
    version = "v1.39.0",
    deps = [
        ":protobuf",
        ":rpcstatus",
        ":xnet",
        ":xsys",
    ],
)

go_module(
    name = "gotestyourself",
    module = "github.com/gotestyourself/gotestyourself",
    version = "v1.3.0",
)

go_module(
    name = "gotty",
    module = "github.com/Nvveen/Gotty",
    version = "v0.0.0-20170406111628-a8b993ba6abd",
)

go_module(
    name = "image-spec",
    module = "github.com/opencontainers/image-spec",
    install = ["specs-go/v1"],
    version = "v1.0.1",
    deps = [
        ":go-digest",
    ],
)

go_module(
    name = "libcontainer-user",
    module = "github.com/opencontainers/runc",
    install = ["libcontainer/user"],
    version = "v1.0.0",
    deps = [
        ":xsys",
    ],
)

go_module(
    name = "libtrust",
    module = "github.com/docker/libtrust",
    version = "v0.0.0-20160708172513-aabc10ec26b7",
)

go_module(
    name = "logrus",
    module = "github.com/sirupsen/logrus",
    version = "v1.8.1",
    deps = [
        ":xsys"
    ],
)

go_module(
    name = "mapstructure",
    module = "github.com/mitchellh/mapstructure",
    version = "v1.4.1",
)

go_module(
    name = "mergo",
    module = "github.com/imdario/mergo",
    version = "v0.3.12",
)

go_module(
    name = "moby-mountinfo",
    module = "github.com/moby/sys/mountinfo",
    version = "v0.4.1",
    deps = [
        ":xsys",
    ],
)

go_module(
    name = "moby-term",
    module = "github.com/moby/term",
    version = "v0.0.0-20201216013528-df9cb8a40635",
    deps = [
        ":xsys",
    ],
)

go_module(
    name = "moby-mount",
    module = "github.com/moby/sys/mount",
    version = "v0.2.0",
    deps = [
        ":moby-mountinfo",
        ":xsys",
    ],
)

go_module(
    name = "mux",
    module = "github.com/gorilla/mux",
    version = "v1.8.0",
)

go_module(
    name = "notary",
    module = "github.com/theupdateframework/notary",
    install = [
        ".",
        "client",
        "passphrase",
        "storage",
        "trustmanager",
        "trustpinning",
        "tuf/data",
        "tuf/signed",
    ],
    version = "v0.7.0",
    deps = [
        ":docker-distribution",
        ":docker-go",
        ":logrus",
        ":xcrypto",
        ":xterm",
    ],
)

go_module(
    name = "objx",
    module = "github.com/stretchr/objx",
    version = "v0.3.0",
)

go_module(
    name = "pbutil",
    module = "github.com/matttproud/golang_protobuf_extensions",
    install = ["pbutil"],
    version = "v1.0.1",
    deps = [
        ":protobuf",
    ],
)

go_module(
    name = "pflag",
    module = "github.com/spf13/pflag",
    version = "v1.0.5",
)

go_module(
    name = "pkg-errors",
    module = "github.com/pkg/errors",
    version = "v0.9.1",
)

go_module(
    name = "pq",
    module = "github.com/lib/pq",
    version = "v1.10.2",
)

go_module(
    name = "prometheus",
    module = "github.com/prometheus/client_golang",
    install = [
        "prometheus",
        "prometheus/promhttp",
    ],
    version = "v1.11.0",
    deps = [
        ":prometheus-common",
        ":procfs",
        ":quantile",
        ":xxhash",
    ],
)

go_module(
    name = "prometheus-client-model",
    module = "github.com/prometheus/client_model",
    install = [
        "go",
    ],
    version = "v0.2.0",
    deps = [
        ":protobuf",
    ],
)

go_module(
    name = "prometheus-common",
    module = "github.com/prometheus/common",
    install = [
        "expfmt",
        "model",
    ],
    version = "v0.29.0",
    deps = [
        ":pbutil",
        ":prometheus-client-model",
    ],
)

go_module(
    name = "procfs",
    module = "github.com/prometheus/procfs",
    version = "v0.6.0",
    deps = [
        ":xsys",
    ],
)

go_module(
    name = "protobuf",
    module = "github.com/golang/protobuf",
    install = [
        "proto",
        "ptypes",
        "ptypes/timestamp",
        "ptypes/empty"
    ],
    version = "v1.5.2",
    deps = [
        ":google-protobuf",
    ],
)

go_module(
    name = "quantile",
    module = "github.com/beorn7/perks",
    install = ["quantile"],
    version = "v1.0.1",
)

go_module(
    name = "rpcstatus",
    module = "google.golang.org/genproto",
    install = [
        "googleapis/rpc/status",
    ],
    version = "v0.0.0-20210629200056-84d6f6074151",
    deps = [
        ":google-protobuf",
    ],
)

go_module(
    name = "shlex",
    module = "github.com/google/shlex",
    version = "v0.0.0-20191202100458-e7afc7fbc510",
)

go_module(
    name = "spew",
    module = "github.com/davecgh/go-spew",
    install = ["spew"],
    version = "v1.1.1",
)

go_module(
    name = "sqlx",
    module = "github.com/jmoiron/sqlx",
    version = "v1.3.4",
)

go_module(
    name = "sortorder",
    module = "github.com/fvbommel/sortorder",
    version = "v1.0.2",
)

go_module(
    name = "testify",
    module = "github.com/stretchr/testify",
    install = ["..."],
    version = "v1.7.0",
    deps = [
        ":difflib",
        ":objx",
        ":spew",
        ":yamlv3",
    ],
)

go_module(
    name = "uuid",
    module = "github.com/gofrs/uuid",
    version = "v4.0.0",
)

go_module(
    name = "xcrypto",
    module = "golang.org/x/crypto",
    install = [
        "ed25519",
        "pbkdf2",
        "ssh/terminal",
    ],
    version = "v0.0.0-20200622213623-75b288015ac9",
    deps = [
        ":xsys",
    ],
)

go_module(
    name = "xerrors",
    module = "golang.org/x/xerrors",
    install = ["..."],
    version = "v0.0.0-20200804184101-5ec99f83aff1",
)

go_module(
    name = "xmod",
    module = "golang.org/x/mod",
    install = ["semver"],
    version = "v0.3.0",
)

go_module(
    name = "xnet",
    module = "golang.org/x/net",
    install = ["..."],
    version = "v0.0.0-20201021035429-f5854403a974",
    deps = [
        ":xcrypto",
        ":xtext",
    ],
)

go_module(
    name = "xsys",
    module = "golang.org/x/sys",
    install = ["..."],
    version = "v0.0.0-20210616094352-59db8d763f22",
)

go_module(
    name = "xterm",
    module = "golang.org/x/term",
    version = "v0.0.0-20210615171337-6886f2dfbf5b",
    deps = [
        ":xsys",
    ],
)

go_module(
    name = "xtext",
    module = "golang.org/x/text",
    install = ["..."],
    version = "v0.3.3",
    deps = [
        ":xtools",
    ],
)

go_module(
    name = "xtime",
    module = "golang.org/x/time",
    install = ["rate"],
    version = "v0.0.0-20210611083556-38a9dc6acbc6"
)

go_module(
    name = "xtools",
    module = "golang.org/x/tools",
    install = [
        "go/callgraph",
        "go/callgraph/cha",
        "go/loader",
        "go/ssa",
        "go/ssa/ssautil",
        "go/buildutil",
    ],
    version = "v0.0.0-20210106214847-113979e3529a",
    deps = [
        ":xerrors",
        ":xmod",
    ],
)

go_module(
    name = "xxhash",
    module = "github.com/cespare/xxhash/v2",
    version = "v2.1.1",
)

go_module(
    name = "yamlv2",
    module = "gopkg.in/yaml.v2",
    version = "v2.4.0",
)

go_module(
    name = "yamlv3",
    module = "gopkg.in/yaml.v3",
    version = "v3.0.0-20200313102051-9f266ea9e77c",
)
