package main

import (
	"context"
	"fmt"
	proto "github.com/MicroService-Su/go-micro/proto"
	"github.com/micro/go-micro"
)

func main() {
	// 定义服务
	service := micro.NewService(micro.Name("client"))
	// 初始化
	service.Init()
	// 创建客户端
	greeter := proto.NewGreeterService("joshua",service.Client())

	// 调用服务
	res,err := greeter.Hello(context.TODO(),&proto.HelloRequest{Name: "Athena"})
	if err!=nil{
		fmt.Println(err)
		return
	}
	fmt.Println(res.Greeting)
}
