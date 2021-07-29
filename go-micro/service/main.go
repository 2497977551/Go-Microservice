package main

import (
	"context"
	"fmt"
	proto "github.com/MicroService-Su/go-micro/proto"
	"github.com/micro/go-micro"
	"log"
)

type Greeter struct {

}
func (g *Greeter)Hello(ctx context.Context, req *proto.HelloRequest, rsp *proto.HelloResponse) error{
	rsp.Greeting = "Hello" + req.Name
	return nil
}
func main() {
	// 创建服务，定义服务名以及版本，除了服务名其它选项可加可不加
	service := micro.NewService(
		micro.Name("joshua"),
		micro.Version("latest"),
		)

	service.Init()
	fmt.Println(proto.RegisterGreeterHandler(service.Server(),&Greeter{}))
	log.Fatal(service.Run())
}
