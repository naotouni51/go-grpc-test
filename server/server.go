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

	num := req.GetNum()
	log.Printf("リクエスト受信 %d秒通信します", num)

	for i := int32(0); i < req.GetNum(); i++ {
		message := fmt.Sprintf("サーバータイム: %s", time.Now())
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
	fmt.Println("サーバーを起動しました")

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
