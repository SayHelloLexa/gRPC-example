package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/types/known/emptypb"

	"github.com/sayhellolexa/grpc-message-server/pkg/api"
)

func main() {
	conn, err := grpc.NewClient(
		"localhost:8080",
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithIdleTimeout(5*time.Second),
	)
	defer conn.Close()
	if err != nil {
		log.Fatal(err)
	}

	ctx := context.Background()

	client := api.NewMessageServiceClient(conn)
	msgs := []*api.Message{
		{Id: "1", Time: "15:00", Text: "Hello, admin"},
		{Id: "2", Time: "16:00", Text: "Hello, admin"},
		{Id: "3", Time: "17:00", Text: "Hello, admin"},
		{Id: "4", Time: "18:00", Text: "Good evening, Admin"},
	}

	_, err = client.Send(ctx, &api.MessageRequest{Message: msgs})
	if err != nil {
		log.Fatal(err)
	}

	resp, err := client.Messages(ctx, &emptypb.Empty{})
	if err != nil {
		log.Fatal(err)
	}

	for _, msg := range resp.Message {
		fmt.Println(msg)
	}
}
