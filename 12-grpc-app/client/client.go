package main

import (
	"context"
	"fmt"
	"grpc-app/proto"
	"log"
	"time"

	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:50000", grpc.WithInsecure())
	if err != nil {
		log.Fatalln(err)
	}
	client := proto.NewAppServicesClient(conn)

	//add operation (request & response)
	/*
		addRequest := &proto.CalculatorRequest{X: 100, Y: 200}
		addResult, reqErr := client.Add(context.Background(), addRequest)
		if reqErr != nil {
			log.Fatalln(reqErr)
		}
		fmt.Println("Add Result = ", addResult.GetResult())
	*/

	//average operation (client streaming)
	fmt.Println("[Average Operation starts")
	data := []int64{10, 20, 30, 40}
	avgClientStream, err := client.Average(context.Background())
	if err != nil {
		log.Fatalln(err)
	}
	for _, no := range data {
		fmt.Println("Sending ", no)
		avgReq := &proto.AverageRequest{No: no}
		avgClientStream.Send(avgReq)
		time.Sleep(500 * time.Millisecond)
	}
	avgResponse, responseErr := avgClientStream.CloseAndRecv()
	if responseErr != nil {
		log.Fatalln(responseErr)
	}
	fmt.Println("Average : ", avgResponse.GetResult())
}
