package main

import (
	"context"
	pb "github.com/castaneai/grpc-broadcast-example"
	"google.golang.org/grpc"
	"log"
	"os"
	"time"
)

func main() {
	addr := "localhost:50051"
	conn, err := grpc.Dial(addr, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewChatRoomClient(conn)

	clientName := os.Args[1]

	ctx := context.Background()

	stream, err := c.Chat(ctx)
	if err != nil {
		log.Fatal(err)
	}
	go func() {
		for {
			mes := "hello, I'm "+clientName
			if err := stream.SendMsg(&pb.ChatRequest{Message: mes}); err != nil {
				log.Fatal(err)
			}
			log.Printf("sent: %s", mes)
			time.Sleep(1*time.Second)
		}
	}()
	for {
		resp, err := stream.Recv()
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("recv: %s", resp.Message)
	}
}
