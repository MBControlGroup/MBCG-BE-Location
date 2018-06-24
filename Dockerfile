FROM golang:1.8

COPY . "$GOPATH/src/myapp/"

WORKDIR $GOPATH/src/myapp

CMD ["go", "run", "main.go"]
