load("@bazel_tools//tools/build_defs/repo:http.bzl", "http_archive")

http_archive(
    name = "io_bazel_rules_go",
    integrity = "sha256-fHbWI2so/2laoozzX5XeMXqUcv0fsUrHl8m/aE8Js3w=",
    urls = [
        "https://mirror.bazel.build/github.com/bazelbuild/rules_go/releases/download/v0.44.2/rules_go-v0.44.2.zip",
        "https://github.com/bazelbuild/rules_go/releases/download/v0.44.2/rules_go-v0.44.2.zip",
    ],
)

http_archive(
  name = "io_bazel_rules_pkg",
  sha256 = "a880a793b916b0312e8d3e6ebd26128a9b64e38aa48dd20ef2bcfe25420a22a7",
  urls = ["https://github.com/bazelbuild/rules_pkg/releases/download/0.1.0/rules_pkg-0.1.0.tar.gz"],
)

http_archive(
    name = "bazel_gazelle",
    integrity = "sha256-MpOL2hbmcABjA1R5Bj2dJMYO2o15/Uc5Vj9Q0zHLMgk=",
    urls = [
        "https://mirror.bazel.build/github.com/bazelbuild/bazel-gazelle/releases/download/v0.35.0/bazel-gazelle-v0.35.0.tar.gz",
        "https://github.com/bazelbuild/bazel-gazelle/releases/download/v0.35.0/bazel-gazelle-v0.35.0.tar.gz",
    ],
)

http_archive(
  name = "logrus",
  sha256 = "a3c60bbde616e7579dbd104e47f5c0e9153f4e672b88e8d3a87f7f0f8852ad4a",
  strip_prefix = "logrus-1.9.3",
  urls = ["https://github.com/sirupsen/logrus/archive/v1.9.3.zip"],
)

load(
  "@io_bazel_rules_go//go:deps.bzl",
  "go_register_toolchains",
  "go_rules_dependencies",
)
load(
  "@bazel_gazelle//:deps.bzl",
  "gazelle_dependencies",
  "go_repository",
)

go_repository(
  name = "com_github_joho_godotenv",
  importpath = "github.com/joho/godotenv",
  sum = "h1:7eLL/+HRGLY0ldzfGMeQkb7vMd0as4CfYvUVzLqw0N0=",
  version = "v1.5.1",
)

go_repository(
  name = "com_github_sirupsen_logrus",
  importpath = "github.com/sirupsen/logrus",
  sum = "h1:dueUQJ1C2q9oE3F7wvmSGAaVtTmUizReu6fjN8uqzbQ=",
  version = "v1.9.3",
)

go_repository(
  name = "com_github_gorilla_mux",
  importpath = "github.com/gorilla/mux",
  sum = "h1:i40aqfkR1h2SlN9hojwV5ZA91wcXFOvkdNIeFDP5koI=",
  version = "v1.8.0",
)

go_repository(
  name = "com_github_aws_aws_sdk_go",
  importpath = "github.com/aws/aws-sdk-go",
  sum = "h1:FaXvNwHG3Ri1paUEW16Ahk9zLVqSAdqa1M3phjZR35Q=",
  version = "v1.50.6",
)

go_repository(
  name = "com_github_jmespath_go_jmespath",
  importpath = "github.com/jmespath/go-jmespath",
  sum = "h1:BEgLn5cpjn8UN1mAw4NjwDrS35OdebyEtFe+9YPoQUg=",
  version = "v0.4.0",
)

go_rules_dependencies()

go_register_toolchains(version = "1.20.7")

gazelle_dependencies()
