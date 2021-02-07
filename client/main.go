package main

import (
	"context"
	"io"
	"log"

	pb "go_grpc_client/pb/notification"

	"google.golang.org/grpc"
)

func request(client pb.NotificationClient, num int32) error {
	req := &pb.NotificationRequest{
		Num: num,
	}

	stream, err := client.Notification(context.Background(), req)
	if err != nil {
		log.Fatal(err)
	}

	for {
		reply, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}
		log.Println("通信中", reply.GetMessage())
	}
	return nil
}

func main() {
	num := int32(60)

	address := "localhost:8080"
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	client := pb.NewNotificationClient(conn)

	err = request(client, num)
	if err != nil {
		log.Fatal(err)
	}

}
