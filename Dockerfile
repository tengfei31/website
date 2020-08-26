FROM scratch

ENV GOPROXY https://goproxy.cn,direct
WORKDIR $GOPATH/src/github.com/tengfei31/website
COPY . $GOPATH/src/github.com/tengfei31/website

EXPOSE 8000
CMD ["./website"]



