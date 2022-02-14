package main

import (
	"context"
	"io"
	"log"

	"grpc_streaming1/pb"

	"google.golang.org/grpc"
)

func main(){
	clientConn, err := grpc.Dial("127.0.0.1:8060", grpc.WithInsecure())
	if err != nil{
		log.Fatal(err)
	}

	productServiceClient := pb.NewProductServiceClient(clientConn)

	/*
	res, err := productServiceClient.CreateProduct(context.Background(), &pb.CreateProductRequest{ 
		Product: &pb.Product{
			Id: "",
			Title: "Sneakers - Reebok - WSX-1",
			Price: 4990,
			Category: "Shoes",
			Rating: 0.0,
			Description: "Nice choise for every day :D",
		},
	})
	if err != nil{
		log.Println(err)
	}
	log.Println(res.Id)
	*/

	stream, err := productServiceClient.SearchProduct(context.Background(), &pb.SearchProductRequest{
		Filter: &pb.ProductFilter{
			MaxPrice: 5900,
			Category: "Shoes",
			Rating: 0.0,
		},
	})
	if err != nil{
		log.Println(err)
	}
	waitc := make(chan struct{})
	go func(){
		for{
			res, err := stream.Recv()
			if err == io.EOF{
				close(waitc)
				return
			}
			if err != nil{
				log.Println(err)
				return
			}
			product := res.GetProduct()
			log.Println(product)
		}
	}()

	<-waitc
}	
