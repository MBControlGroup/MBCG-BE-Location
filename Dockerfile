FROM golang:1.8
#设置工作目录
WORKDIR $GOPATH/src/myapp
#将服务器的go工程代码加入到docker容器中
COPY . $GOPATH/src/myapp
COPY . $GOPATH/src/bin
COPY . $GOPATH/src/pkg
#go构建可执行文件
RUN cd "$GOPATH/src/myapp" && go get -u github.com/astaxie/beego && go build .
#暴露端口
EXPOSE 8090
#最终运行docker的命令
ENTRYPOINT  ["./myapp"]
