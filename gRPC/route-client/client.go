package main

import (
	"context"
	"fmt"
	pb "github.com/HanFa/learn-go/grpc-example/route"

	"google.golang.org/grpc"
	"log"
)

func main() {
	// grpc.Dial 第一个参数是服务器的地址，第二个和第三个参数表示忽略证书验证
	conn,err := grpc.Dial("localhost:8878",grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalln("conn Dial Error",err.Error())
	}
	defer  conn.Close()
	client:=pb.NewRouteGuideClient(conn)
	run(client)
}

func run(client pb.RouteGuideClient){
	feature,err := client.GetFeature(context.Background(),&pb.Point{
		Latitude:  123123,
		Longitude: 232323,
	})
	if err != nil {
		log.Fatalln(err.Error())
	}
	fmt.Println(feature)
}
