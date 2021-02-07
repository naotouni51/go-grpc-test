package chat

import (
	"context"
	"log"
)

type Server struct {
}

func (s *Server) SayHello(ctx context.Context, message *Message) (*Message, error) {
	log.Printf("%s", message.Body)
	return &Message{Body: "こちらはサーバーです"}, nil
}
