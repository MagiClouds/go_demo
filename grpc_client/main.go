package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"start/grpc_client/services"
)

func main() {
	conn, err := grpc.Dial(":8080", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	ctx := context.Background()

	prodClient := services.NewProdServiceClient(conn)
	//prodRes, err := prodClient.GetProdStock(ctx, &services.ProdRequest{ProdId:10})
	//if err != nil {
	//	log.Fatal(err)
	//}

	res, err := prodClient.GetProdStocks(ctx, &services.QuerySize{Size:3})
	if err != nil {
		log.Fatal(err)
	}

	for _, i := range res.ProdInfo {
		fmt.Println(i)
	}

}