package main

import (
	"context"
	pb "github.com/MicroService-Su/gRPC/demo/route"
	"google.golang.org/grpc"
	"log"
)

const Address string = ":8000"
func main() {
	// 连接服务器
	conn,err := grpc.Dial(Address,grpc.WithInsecure(),grpc.WithBlock())
	if err != nil {
		log.Fatalf("net.Connect err: %v", err)
	}
	defer conn.Close()


	// 建立gRPC连接
	gRpcClient := pb.NewSimpleClient(conn)
	// 创建发送结构体
	req := pb.SimpleRequest{
		Data: "admin1",
	}
	// 调用我们的服务(Route方法)
	// 同时传入了一个 context.Context ，在有需要时可以让我们改变RPC的行为，比如超时/取消一个正在运行的RPC
	res, err := gRpcClient.Route(context.Background(), &req)
	if err != nil {
		log.Fatalf("Call Route err: %s", err.Error())
	}
	// 打印返回值
	log.Println(res)

}
