package main

import (
	"github.com/golang/protobuf/proto"
	"log"
	"net"

	"google.golang.org/grpc"
	pb "github.com/HanFa/learn-go/grpc-example/route"
	"context"
)


func main() {
	lis,err := net.Listen("tcp","localhost:8878")

	if err != nil {
		log.Fatalln(err.Error())
	}
	gRpcServer := grpc.NewServer()
	pb.RegisterRouteGuideServer(gRpcServer,newServer())

	err = gRpcServer.Serve(lis)
	if err != nil {
		log.Fatalln(err.Error())
	}
}


type routeGuideServer struct{
	db []*pb.Feature
	pb.UnimplementedRouteGuideServer
}

func (r *routeGuideServer) GetFeature(ctx context.Context, point *pb.Point) (*pb.Feature, error){
	for _, v := range r.db {
		if proto.Equal(v.Location,point){
			return v,nil
		}
	}
	return nil,nil
}

func (r *routeGuideServer) ListFeatures(*pb.Rectangle, pb.RouteGuide_ListFeaturesServer) error{
	return nil
}

func (r *routeGuideServer) RecordRoute(pb.RouteGuide_RecordRouteServer) error{
	return nil
}

func (r *routeGuideServer) Recommend(pb.RouteGuide_RecommendServer) error{
	return nil
}

func newServer() *routeGuideServer{
	return &routeGuideServer{
		db:[]*pb.Feature{
			{
				Name: "xxx",
				Location: &pb.Point{
					Latitude: 123123,
					Longitude: 232323,
				},
			},
		},
	}
}