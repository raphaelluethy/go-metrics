package main

import (
	"context"
	"github.com/raphaelluethy/go-metrics/backend/client"
	"github.com/raphaelluethy/go-metrics/backend/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"net"
)

type myStonkzServer struct {
	proto.UnimplementedInvoicerServer
}

func (s myStonkzServer) CreateStonkz(ctx context.Context, in *proto.CreateRequest) (*proto.CreateResponse, error) {

	return &proto.CreateResponse{
		Id: "1234",
	}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":8089")
	if err != nil {
		log.Fatalf("cannot create listener: %s", err)
	}

	s := grpc.NewServer()
	grpc.WithTransportCredentials(insecure.NewCredentials())
	service := &myStonkzServer{}

	proto.RegisterInvoicerServer(s, service)

	client.RunClient()

	err = s.Serve(lis)
	if err != nil {
		log.Fatalf("cannot serve: %s", err)
	}
}
