package main

import (
	pb "github.com/MicroService-Su/gRPC/StreamDemo/route"
	"google.golang.org/grpc"
	"log"
	"net"
	"strconv"
)
const (
	// Address 监听地址
	Address string = ":8000"
	// Network 网络通信协议
	Network string = "tcp"
)
type StreamService struct {
	pb.UnimplementedStreamServerServer
}

func (s StreamService)ListValue(sr *pb.SimpleRequest, lvs pb.StreamServer_ListValueServer) error{
	for n := 0; n < 5; n++ {
		// 向流中发送消息， 默认每次send送消息最大长度为`math.MaxInt32`bytes
		err := lvs.Send(&pb.StreamResponse{
			StreamValue: sr.Data + strconv.Itoa(n),
		})
		if err != nil {
			return err
		}
	}
	return nil
}

func main() {
	// 监听本地端口
	listener, err := net.Listen(Network, Address)
	if err != nil {
		log.Fatalf("net.Listen err: %v", err)
	}
	log.Println(Address + " net.Listing...")
	// 新建gRPC服务器实例
	// 默认单次接收最大消息长度为`1024*1024*4`bytes(4M)，单次发送消息最大长度为`math.MaxInt32`bytes
	// gRpcServer := grpc.NewServer(grpc.MaxRecvMsgSize(1024*1024*4), grpc.MaxSendMsgSize(math.MaxInt32))
	gRpcServer := grpc.NewServer()
	// 在gRPC服务器注册我们的服务
	pb.RegisterStreamServerServer(gRpcServer, &StreamService{})

	//用服务器 Serve() 方法以及我们的端口信息区实现阻塞等待，直到进程被杀死或者 Stop() 被调用
	err = gRpcServer.Serve(listener)
	if err != nil {
		log.Fatalf("grpcServer.Serve err: %v", err)
	}
}


