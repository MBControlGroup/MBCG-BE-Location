FROM golang:1.8
RUN mkdir -p /go/src/myapp
ADD . /go/src/myapp
WORKDIR /go/src/myapp
CMD ["go", "run", "main.go"]