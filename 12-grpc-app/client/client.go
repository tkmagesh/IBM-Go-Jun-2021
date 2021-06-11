package main

import (
	"context"
	"fmt"
	"grpc-app/proto"
	"io"
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
	/*
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
	*/

	//server streaming
	/* fmt.Println("Generating Prime numbers from 25 to 75 (server streaming)")
	primeNoReq := &proto.PrimeNumberRequest{RangeStart: 25, RangeEnd: 75}
	primeNoStream, err := client.GeneratePrime(context.Background(), primeNoReq)
	if err != nil {
		log.Fatalln(err)
	}
	for {
		res, e := primeNoStream.Recv()
		if e == io.EOF {
			fmt.Println("All the prime numbers are received")
			break
		}
		if e != nil {
			log.Fatalln(e)
		}
		fmt.Println("Prime No received : ", res.GetNo())
	} */

	doBiDiStreaming(client)
}

func doBiDiStreaming(c proto.AppServicesClient) {
	fmt.Println("Starting to do a BiDi Streaming RPC...")

	// we create a stream by invoking the client
	stream, err := c.GreetEveryone(context.Background())
	if err != nil {
		log.Fatalf("Error while creating stream: %v", err)
		return
	}

	requests := []*proto.GreetEveryoneRequest{
		&proto.GreetEveryoneRequest{
			Greeting: &proto.Greeting{
				FirstName: "Stephane",
			},
		},
		&proto.GreetEveryoneRequest{
			Greeting: &proto.Greeting{
				FirstName: "John",
			},
		},
		&proto.GreetEveryoneRequest{
			Greeting: &proto.Greeting{
				FirstName: "Lucy",
			},
		},
		&proto.GreetEveryoneRequest{
			Greeting: &proto.Greeting{
				FirstName: "Mark",
			},
		},
		&proto.GreetEveryoneRequest{
			Greeting: &proto.Greeting{
				FirstName: "Piper",
			},
		},
	}

	waitc := make(chan struct{})
	// we send a bunch of messages to the client (go routine)
	go func() {
		// function to send a bunch of messages
		for _, req := range requests {
			fmt.Printf("Sending message: %v\n", req)
			stream.Send(req)
			time.Sleep(1000 * time.Millisecond)
		}
		stream.CloseSend()
	}()
	// we receive a bunch of messages from the client (go routine)
	go func() {
		// function to receive a bunch of messages
		for {
			res, err := stream.Recv()
			if err == io.EOF {
				break
			}
			if err != nil {
				log.Fatalf("Error while receiving: %v", err)
				break
			}
			fmt.Printf("Received: %v\n", res.GetResult())
		}
		close(waitc)
	}()

	// block until everything is done
	<-waitc
}
