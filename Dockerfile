FROM golang:1.8

COPY . "$GOPATH/src/myapp/"
COPY . "$GOPATH/src/github.com/beego/"
RUN cd "$GOPATH/src/myapp" && go get -v && go install -v

WORKDIR $GOPATH/src/myapp

CMD ["go", "run", "main.go"]
