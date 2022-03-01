package grpc

import (
	"chat/packages/chat/core/conversation/ports"
	pb "chat/packages/protos"
	"context"
	"fmt"
)

func (s server) CreateConversation(ctx context.Context, request *pb.CreateConversationRequest) (*pb.ConversationSummary, error) {

	params := &ports.CreateConversationParam{
		Name:    request.Name,
		OwnerID: request.OwnerId,
	}

	conversation, err := s.controller.Create(params)

	if err != nil {
		return nil, fmt.Errorf("error creating conversation: %w", err)
	}

	return conversationSummaryToGrpc(conversation), nil
}
