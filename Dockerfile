FROM golang:latest

WORKDIR /app

COPY bazel-bin/EliFuchsmanBE_/EliFuchsmanBE.runfiles/ .

CMD ["./EliFuchsmanBE_/EliFuchsmanBE.runfiles/_main/EliFuchsmanBE_/EliFuchsmanBE"]
