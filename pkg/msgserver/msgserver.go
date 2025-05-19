package msgserver

import (
	"context"
	"github.com/sayhellolexa/grpc-message-server/pkg/api"
	"google.golang.org/protobuf/types/known/emptypb"
)

// GRPCServer - структура сервиса
type GRPCServer struct {
	api.UnimplementedMessageServiceServer
}

// messages - хранилище сообщений
var messages []*api.Message

// Send - метод отправки и сохранения сообщений
func (s *GRPCServer) Send(ctx context.Context, req *api.MessageRequest) (*emptypb.Empty, error) {
	for _, msg := range req.Message {
		messages = append(messages, msg)
	}

	return &emptypb.Empty{}, nil
}

// Messages - метод получения накопленных сообщений
func (s *GRPCServer) Messages(ctx context.Context, req *emptypb.Empty) (*api.MessageResponse, error) {
	return &api.MessageResponse{Message: messages}, nil
}
