bazel run //:gazelle
bazel build :all

docker-compose up --build
