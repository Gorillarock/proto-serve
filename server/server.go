package main

import (
	"context"
	"log"
	"net"

	pb "github.com/Gorillarock/proto-serv/protos"
	"google.golang.org/grpc"
)

const (
	port = ":50051"
)

type CurrencyServer struct {
	pb.UnimplementedCurrencyServer
}

func (s *CurrencyServer) GetRate(ctx context.Context, in *pb.RateRequest) (*pb.RateResponse, error) {
	log.Printf("Received: request with Base = %v and Dest. = %v\n", in.GetBase(), in.GetDestination())
	return &pb.RateResponse{
		Rate: 1.25,
	}, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v\n", err)
	}

	s := grpc.NewServer()
	pb.RegisterCurrencyServer(s, &CurrencyServer{})
	log.Printf("server listening at: %v\n", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v\n", err)
	}
}
