FROM golang:latest

WORKDIR /EliFuchsmanBE

COPY . .

COPY bazel-bin-docker /EliFuchsmanBE

EXPOSE 80

CMD ["./bazel-bin-docker/EliFuchsmanBE"]

