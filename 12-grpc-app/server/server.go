package main

import (
	"context"
	"fmt"
	"grpc-app/proto"
	"io"
	"log"
	"net"

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
