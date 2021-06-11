package main

import (
	"context"
	"fmt"
	"grpc-app/proto"
	"log"

	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:50000", grpc.WithInsecure())
	if err != nil {
		log.Fatalln(err)
	}
	client := proto.NewAppServicesClient(conn)

	//add operation (request & response)
	addRequest := &proto.CalculatorRequest{X: 100, Y: 200}
	addResult, reqErr := client.Add(context.Background(), addRequest)
	if reqErr != nil {
		log.Fatalln(reqErr)
	}
	fmt.Println("Add Result = ", addResult.GetResult())
}
