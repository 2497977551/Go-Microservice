package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"log"

	pb "github.com/MicroService-Su/demo/test"
)

func main() {
	// 链接服务器
	conn,err := grpc.Dial(":8000",grpc.WithBlock(),grpc.WithInsecure())
	if err != nil {
		log.Fatalln("Client Dial Err:",err.Error())
	}
	defer conn.Close()

	// 建立gRPC链接
	client := pb.NewShopServerClient(conn)
	sr := pb.ShopRequest{
		ProductId:     1,
		ProductName:   "中华牌香烟",
		ProductPrice:  45,
		ProductNumber: 2,
		IsVip:         true,
	}
	res,err := client.Shop(context.Background(),&sr)
	if err != nil {
		log.Fatalln("ShopResponse Err:",err.Error())
	}
	fmt.Println("产品名称：",res.ProductName)
	fmt.Printf("支付费用：%v 人民币\n",res.ProductPrice)
	fmt.Println("购买数量：",res.ProductNumber)
	fmt.Println("备注：",res.Remark)
}
