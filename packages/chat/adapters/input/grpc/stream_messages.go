package grpc

import (
	"chat/packages/chat/core/conversation/ports"
	pb "chat/packages/protos"
	"fmt"
	"google.golang.org/grpc/metadata"
	"strconv"
)

var (
	counter = 0
)

func (s server) ListenMessages(filter *pb.FilterMessages, stream pb.ConversationService_ListenMessagesServer) error {

	params := &ports.ConversationFilter{
		ConversationID: filter.GetConversationId(),
		UserID:         filter.GetParticipantId(),
	}

	messageChannel, err := s.controller.StreamMessages(params)

	if err != nil {
		return fmt.Errorf("error to create messages channel: %w", err)
	}

	for message := range messageChannel {
		grpcMessage := messageViewToGrpc(message)

		counter++

		headerErr := stream.SendHeader(metadata.Pairs("my_custom_header", strconv.Itoa(counter)))

		if headerErr != nil {
			return fmt.Errorf("error to send header: %w", headerErr)
		}

		if sendErr := stream.Send(grpcMessage); sendErr != nil {
			return fmt.Errorf("error to send message: %w", sendErr)
		}
	}

	return nil
}
