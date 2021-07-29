package main

import (
	"context"
	"google.golang.org/grpc"
	"log"
	"net"

	pb "github.com/MicroService-Su/gRPC/demo/route"
)

type SimpleService struct {
	
}


const (
	// Address 监听地址
	Address string = ":8000"
	// Network 网络通信协议
	Network string = "tcp"
)
func main() {
	// 监听本地端口
	listener, err := net.Listen(Network, Address)
	if err != nil {
		log.Fatalf("net.Listen err: %v", err)
	}
	log.Println(Address + " net.Listing...")
	// 新建gRPC服务器实例
	gRpcServer := grpc.NewServer()
	// 在gRPC服务器注册我们的服务
	pb.RegisterSimpleServer(gRpcServer, &SimpleService{})

	//用服务器 Serve() 方法以及我们的端口信息区实现阻塞等待，直到进程被杀死或者 Stop() 被调用
	err = gRpcServer.Serve(listener)
	if err != nil {
		log.Fatalf("grpcServer.Serve err: %v", err)
	}
}

func (s SimpleService)Route(ctx context.Context, req *pb.SimpleRequest)(*pb.SimpleResponse,error){

	if req.Data == "admin" {
		res := pb.SimpleResponse{
			Code:  200,
			Value: "OK",
		}
		return &res,nil
	}
	return &pb.SimpleResponse{
		Code:  404,
		Value: "无权限!",
	},nil
}


