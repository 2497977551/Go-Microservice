package main

import (
	"context"
	"fmt"
	pb "github.com/MicroService-Su/demo/test"
	"google.golang.org/grpc"
	"log"
	"net"
)

type Product struct {
	pb.UnimplementedShopServerServer
}
func (p *Product)Shop(ctx context.Context,request *pb.ShopRequest)(*pb.ShopResponse,error){
	fmt.Println("服务端收到请求！")
	if request != nil {
		if request.IsVip && request.GetProductNumber()!= 0{
			return &pb.ShopResponse{
					ProductId:     request.GetProductId(),
					ProductName:   request.GetProductName(),
					ProductPrice:  request.GetProductPrice() * 0.9 * float32(request.GetProductNumber()),
					ProductNumber: request.GetProductNumber(),
					IsVip:         true,
					Remark:        "会员打九折，欢迎下次光临！",
				},
				nil
		}else {
			return &pb.ShopResponse{
					ProductId:     request.GetProductId(),
					ProductName:   request.GetProductName(),
					ProductPrice:  request.GetProductPrice(),
					ProductNumber: request.GetProductNumber(),
					IsVip:         true,
					Remark:        "非会员不打折，欢迎下次光临！",
				},
				nil
		}
	}
	return nil,nil
}
func main() {
	listen,err := net.Listen("tcp",":8000")
	if err != nil {
		log.Fatalln("Listen Error:",err.Error())
	}
	gRpcServer := grpc.NewServer()
	pb.RegisterShopServerServer(gRpcServer, &Product{})
	err = gRpcServer.Serve(listen)
	if err != nil {
		log.Fatalln(err.Error())
	}
}
