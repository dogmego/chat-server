package main

import (
	"context"
	"fmt"
	"github.com/brianvoe/gofakeit"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"google.golang.org/protobuf/types/known/emptypb"
	"log"
	"net"

	desc "github.com/dogmego/MicroservicesPractice/chat-server/pkg/note_v1"
)

const grpcPort = 50051

type server struct {
	desc.UnimplementedNoteV1Server
}

func (s *server) Create(ctx context.Context, req *desc.CreateChatRequest) (*desc.CreateChatResponse, error) {
	log.Printf("chat created with usernames: %v", req.Usernames)

	return &desc.CreateChatResponse{
		Id: gofakeit.Int64(),
	}, nil
}

func (s *server) Delete(ctx context.Context, req *desc.DeleteChatRequest) (*emptypb.Empty, error) {
	log.Printf("delete user with id: %d",
		req.Id)

	return nil, nil
}

func (s *server) SendMessage(ctx context.Context, req *desc.SendMessageRequest) (*emptypb.Empty, error) {
	log.Printf("messadge was sent from: %s; text: %s; time: %v",
		req.From, req.Text, req.Timestamp)

	return nil, nil
}

func main() {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", grpcPort))
	if err != nil {
		log.Fatalf("failed to listen: %s", err)
	}

	s := grpc.NewServer()
	reflection.Register(s)
	desc.RegisterNoteV1Server(s, &server{})

	log.Printf("server listening at %v", lis.Addr())
	if err = s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}
