package client

import (
	"context"
	"google.golang.org/grpc"
	"log"

	"github.com/raphaelluethy/go-metrics/backend/proto"
)

var opts []grpc.DialOption
var serverAddr = "localhost:8089"

func RunClient() {
	conn, err := grpc.Dial(serverAddr, opts...)
	if err != nil {
		log.Fatalf("cannot dial server: %s", err)
	}

	client := proto.NewInvoicerClient(conn)
	res, err := client.CreateStonkz(context.Background(), &proto.CreateRequest{
		Name: "Hello World",
	})
	if err != nil {
		log.Fatalf("cannot create stonkz: %s", err)
	}

	log.Printf("stonkz created: %s", res.Id)

	defer func(conn *grpc.ClientConn) {
		err := conn.Close()
		if err != nil {

		}
	}(conn)
}
