package main

import (
	"fmt"
	"log"
	"net"
	"time"

	pb "go_grpc/pb/notification"

	"google.golang.org/grpc"
)

// ServerServerSide is server
type ServerServerSide struct {
	pb.UnimplementedNotificationServer
}

// Notification is
func (s *ServerServerSide) Notification(req *pb.NotificationRequest, stream pb.Notification_NotificationServer) error {
	fmt.Println("リクエスト受け取った")
	for i := int32(0); i < req.GetNum(); i++ {
		message := fmt.Sprintf("%d", i)
		if err := stream.Send(&pb.NotificationReply{
			Message: message,
		}); err != nil {
			return err
		}
		time.Sleep(time.Second * 1)
	}
	return nil
}

func main() {
	fmt.Println("起動")

	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal(err)
	}
	s := grpc.NewServer()
	var server ServerServerSide
	pb.RegisterNotificationServer(s, &server)
	if err := s.Serve(lis); err != nil {
		log.Fatal(err)
	}

}
