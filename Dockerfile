FROM golang:latest

ADD sources.list /etc/apt/

ENV GOPROXY https://goproxy.cn

WORKDIR /www

COPY ./ /www

RUN apt-get update && apt-get install net-tools

RUN go mod download

RUN go get github.com/githubnemo/CompileDaemon

EXPOSE 8100

ENTRYPOINT CompileDaemon --build="go build main.go" --command=./main