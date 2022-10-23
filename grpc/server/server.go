package main

import (
	"context"
	"fmt"
	demo "food_delivery/pb"
	"log"
	"net"

	"google.golang.org/grpc"
)


type server struct {
	demo.UnimplementedRestaurantLikeServiceServer
}

func (s *server) GetRestaurantLikeStat(ctx context.Context, req *demo.RestaurantLikeStatRequest)(
	*demo.RestaurantLikeStatResponse, error) {

	
	return &demo.RestaurantLikeStatResponse{
		Result: map[int32]int32{1: 2, 2 : 4},
	}, nil
}

func main() {
	address := "0.0.0.0:50051"

	lis, err := net.Listen("tcp", address)

	if err != nil {
		log.Fatalf("Err %v", err)
	}

	fmt.Println("Server is serving on port",address)

	s := grpc.NewServer()

	demo.RegisterRestaurantLikeServiceServer(s, &server{})
	s.Serve(lis)
}