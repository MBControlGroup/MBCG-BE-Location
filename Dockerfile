FROM golang:latest   
  
WORKDIR $GOPATH/src/myapp  
ADD ./myapp/ $GOPATH/src/myapp  
#WORKDIR $GOPATH/src/myapp 
  
RUN go get github.com/astaxie/beedb && go get github.com/astaxie/beego && go get github.com/astaxie/beego/orm && go get github.com/astaxie/beego/toolbox   
RUN go get github.com/ziutek/mymysql/godrv  
  
RUN go build ./main.go  
  
EXPOSE 8090   
  
#ENTRYPOINT ["./main"]  
