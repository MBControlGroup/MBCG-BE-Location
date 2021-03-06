FROM golang:1.8
#将服务器的go工程代码加入到docker容器中
COPY . "$GOPATH/"
#go构建可执行文件
RUN cd "$GOPATH/src/myapp" && go get -u github.com/astaxie/beego
#设置工作目录
WORKDIR $GOPATH/src/myapp
#最终运行docker的命令
CMD ["go", "run", "main.go"]

