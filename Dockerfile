FROM golang:latest

RUN mkdir /app
WORKDIR /app
ADD . /app

#TODO
