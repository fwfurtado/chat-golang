package grpc

import (
	"chat/packages/chat/core/conversation/ports"
	pb "chat/packages/protos"
	"context"
	"fmt"
)

func (s server) PullUnreadMessages(ctx context.Context, filter *pb.FilterConversationByOwner) (*pb.ListMessagesResponse, error) {

	params := &ports.ConversationFilter{
		ConversationID: filter.ConversationId,
	}

	unread, err := s.controller.PullUnread(params)

	if err != nil {
		return nil, fmt.Errorf("error pulling unread messages: %w", err)
	}

	grpcMessages := make([]*pb.Message, len(unread))

	for i, message := range unread {
		grpcMessages[i] = messageViewToGrpc(message)
	}

	return &pb.ListMessagesResponse{
		Messages: grpcMessages,
	}, nil

}
