FROM golang:latest

WORKDIR /EliFuchsmanBE

COPY . .

CMD ["./bazel-bin/EliFuchsmanBE_/EliFuchsmanBE"]
