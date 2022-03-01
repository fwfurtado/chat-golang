package grpc

import (
	"chat/packages/chat/core/conversation/ports"
	pb "chat/packages/protos"
	"context"
	"fmt"
)

func (s server) ListConversations(ctx context.Context, filter *pb.FilterByParticipant) (*pb.ListConversationsResponse, error) {
	participantID := filter.ParticipantId
	params := &ports.CurrentUserParam{
		UserID: participantID,
	}

	conversations, err := s.controller.ListConversations(params)

	if err != nil {
		return nil, fmt.Errorf("failed to list conversations: %w", err)
	}

	grpcConversations := make([]*pb.ConversationSummary, len(conversations))

	for i, conversation := range conversations {
		grpcConversations[i] = conversationSummaryToGrpc(conversation)
	}

	response := &pb.ListConversationsResponse{
		Conversations: grpcConversations,
	}

	return response, nil
}
