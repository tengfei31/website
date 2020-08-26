FROM golang:latest

ENV GOPROXY https://goproxy.cn,direct
WORKDIR $GOPATH/src/github.com/tengfei31/website
COPY . $GOPATH/src/github.com/tengfei31/website
RUN go build .

EXPOSE 8000
ENTRYPOINT ["./website"]




