### grpc环境安装
1.确保go module是开启的状态

linux: export GO111MODULE=on

windows: set GO111MODULE=on

2.创建一个项目

go mod init <module-name>

3.安装相关包

go get google.golang.org/grpc

go get github.com/golang/protobuf/protoc-gen-go

4、安装protoc工具

https://github.com/protocolbuffers/protobuf/releases
进入官网下载对应的版本，把执行路径加到系统环境变量

#### rpc代码编写

1、写proto文件

注意带上option go_package

2、根据proto文件生成go代码

protoc --go_out=. --go-grpc_out=. proto\hello.proto

3、写对应的服务端代码，以及客户端的调用代码即可

4、运行代码

go run server.go

go run main.go