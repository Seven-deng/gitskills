grpc(python-go)

step1：工具准备
1、download python
2、download go
3、download pycharm
4、download goland
5、download protoc
6、配置环境变量

python端
1、创建项目
2、配置虚拟环境
3、安装gRPC    pip install grpcio
4、安装gRPC-tools    pip install grpcio-tools
5、配置proto文件
6、###proto文件生成.py文件
    python -m grpc_tools.protoc -I./protos --python_out=./rpc_package --grpc_python_out=./rpc_package ./protos/helloworld.proto
7、编写客户端和服务端



go端
1、创建项目
2、### 修改代理镜像
    https://xueyuanjun.com/post/19887.html
### modlue模式下载依赖包
    go mod tidy
    > 成功后复制依赖
    go mod vendor
    > 查找依赖来构建程序
    go build -mod=vendor
    
### 安装protocol buffer compiler

    > mac安装
    brew install protobuf
    protoc --version  # Ensure compiler version is 3+
    
    > windows安装
    
    > liunx安装
    apt install -y protobuf-compiler
    protoc --version  # Ensure compiler version is 3+
    
### 安装GO plugin

     export GO111MODULE=on  # Enable module mode
     export PATH="$PATH:$(go env GOPATH)/bin"
     
### proto文件生成.go文件
    protoc --go_out=plugins=grpc:./helloworld/ ./helloworld/helloworld.proto

编写客户端和服务端
go run main.go

