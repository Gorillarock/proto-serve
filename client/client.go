package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	pb "github.com/Gorillarock/proto-serv/protos"
	"google.golang.org/grpc"
)

const (
	address = "localhost:50051"
)

func SetupCloseHandler(cancel context.CancelFunc) {
	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		fmt.Println("\r- Ctrl+C pressed in Terminal")
		cancel()
		os.Exit(0)
	}()
}

func main() {
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v\n", err)
	}
	defer conn.Close()

	c := pb.NewCurrencyClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	SetupCloseHandler(cancel)
	defer cancel()

	r, err := c.GetRate(ctx, &pb.RateRequest{Base: "USD", Destination: "EUR"})
	if err != nil {
		log.Fatalf("could not get rate: %v\n", err)
	}
	log.Printf("rate: %v\n", r.GetRate())
}
