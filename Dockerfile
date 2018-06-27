FROM golang:1.8
RUN mkdir /code
COPY . /code
RUN cd /code && go get -u github.com/astaxie/beego
RUN cd /code/src/myapp && go run main.go

