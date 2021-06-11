package main

import (
	"context"
	"fmt"
	"grpc-app/proto"
	"io"
	"log"
	"net"
	"time"

	"google.golang.org/grpc"
)

type server struct{}

func (s *server) Add(ctx context.Context, req *proto.CalculatorRequest) (*proto.CalculatorResponse, error) {
	x := req.GetX()
	y := req.GetY()
	fmt.Printf("Request received for adding %d and %d\n", x, y)
	result := x + y
	response := &proto.CalculatorResponse{Result: result}
	return response, nil
}

func (s *server) Average(reqStream proto.AppServices_AverageServer) error {
	sum := int64(0)
	count := int64(0)
	for {
		req, err := reqStream.Recv()
		if err == io.EOF {
			//send the response
			fmt.Println("[Average] finished receiving all the data from the client")
			avg := sum / count
			res := &proto.AverageResponse{Result: avg}
			reqStream.SendAndClose(res)
			break
		}
		if err != nil {
			log.Fatalln(err)
		}
		no := req.GetNo()
		fmt.Printf("[Average] Average req with No : %d\n", no)
		count++
		sum += no
	}
	return nil
}

//server streaming
func (s *server) GeneratePrime(req *proto.PrimeNumberRequest, stream proto.AppServices_GeneratePrimeServer) error {
	rangeStart := req.GetRangeStart()
	rangeEnd := req.GetRangeEnd()
	for no := rangeStart; no <= rangeEnd; no++ {
		if isPrime(no) {
			fmt.Println("Sending prime no : ", no)
			res := &proto.PrimeNumberResponse{No: no}
			stream.Send(res)
			time.Sleep(500 * time.Millisecond)
		}
	}
	return nil
}

func isPrime(no int32) bool {
	for i := int32(2); i < (no / 2); i++ {
		if no%i == 0 {
			return false
		}
	}
	return true
}

//bidirectional streaming
func (s *server) GreetEveryone(stream proto.AppServices_GreetEveryoneServer) error {
	fmt.Printf("GreetEveryone function was invoked with a streaming request\n")

	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			log.Fatalf("Error while reading client stream: %v", err)
			return err
		}
		firstName := req.GetGreeting().GetFirstName()
		result := "Hello " + firstName + "! "

		sendErr := stream.Send(&proto.GreetEveryoneResponse{
			Result: result,
		})
		if sendErr != nil {
			log.Fatalf("Error while sending data to client: %v", sendErr)
			return sendErr
		}
	}

}

func main() {
	listener, err := net.Listen("tcp", "localhost:50000")
	if err != nil {
		log.Fatalln(err)
	}
	grpcServer := grpc.NewServer()
	proto.RegisterAppServicesServer(grpcServer, &server{})
	e := grpcServer.Serve(listener)
	if e != nil {
		log.Fatalln(e)
	}
}
