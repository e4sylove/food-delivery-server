package main

import (
	"context"
	"fmt"
	demo "food_delivery/pb"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	
	opts := grpc.WithTransportCredentials(insecure.NewCredentials())
	cc, err := grpc.Dial("localhost:50051", opts)

	if err != nil {
		log.Fatal(err)
	}

	defer cc.Close()

	client := demo.NewRestaurantLikeServiceClient(cc)
	request := &demo.RestaurantLikeStatRequest{ResIds: []int32{1,2,3}}

	for i := 1; i <= 5; i++ {
		resp, _ := client.GetRestaurantLikeStat(context.Background(), request)

		fmt.Println("Received response", resp)
	}
}