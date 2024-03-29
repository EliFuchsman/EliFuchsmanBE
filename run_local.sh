bazel run //:gazelle
bazel build :all
cp -R bazel-bin/EliFuchsmanBE_/EliFuchsmanBE bazel-bin-docker

docker-compose up --build
