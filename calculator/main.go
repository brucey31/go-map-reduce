package main

import (
	pb "calculator/protos/gen/calculator/v1"
	"calculator/internal"
	"context"
	"flag"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var (
	port = flag.Int("port", 50051, "The server port")
)

type server struct {
	pb.UnimplementedCalculatorServer
}

func (s *server) CalculateDataPoint(ctx context.Context, in *pb.CalculationInput) (*pb.CalculationResult, error) {
	log.Printf("Received contents of basket owner: %v", in.BasketOwner)
	log.Printf("items = %v", in.Item)
	aggregated_items := internal.CountItems(in.Item)
	outputs := internal.CastCounts(aggregated_items)
	return &pb.CalculationResult{BasketOwner: in.BasketOwner, Output: outputs}, nil
}

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()

	pb.RegisterCalculatorServer(s, &server{})
	log.Printf("server listening at %v", lis.Addr())
	// Register reflection service on gRPC server.
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
