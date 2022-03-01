package grpc

import (
	"chat/packages/chat/core/conversation/ports"
	pb "chat/packages/protos"
)

type server struct {
	controller ports.ConversationController
}

func New(controller ports.ConversationController) pb.ConversationServiceServer {
	return &server{
		controller: controller,
	}
}
