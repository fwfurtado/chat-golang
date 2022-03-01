package grpc

import (
	"chat/packages/chat/core/conversation/ports"
	pb "chat/packages/protos"
	"fmt"
	"google.golang.org/protobuf/types/known/timestamppb"
	"strings"
	"time"
)

func conversationSummaryToGrpc(summary *ports.ConversationSummary) *pb.ConversationSummary {
	return &pb.ConversationSummary{
		Id:          summary.ID,
		Name:        summary.Name,
		CreatedAt:   timestamppb.New(time.Now()), // TODO: get create at from summary
		IsNew:       false,                       // TODO: get is new from summary
		UnreadCount: summary.UnreadCount,
		LastMessage: nil, // TODO: get last message from summary
	}
}

func messageViewToGrpc(view *ports.MessageView) *pb.Message {
	return &pb.Message{
		Id:        view.ID,
		Sender:    participantFromId(view.SenderID),
		State:     messageStateFromString(view.State),
		Text:      view.Content,
		Timestamp: timestamppb.New(view.CreatedAt),
	}
}

func participantFromId(id int64) *pb.Participant {
	return &pb.Participant{Id: id}
}

func messageStateFromString(state string) pb.MessageState {

	grpcState := fmt.Sprintf("MESSAGE_STATE_%s", strings.ToUpper(state))

	value, ok := pb.MessageState_value[grpcState]

	if !ok {
		return pb.MessageState_MESSAGE_STATE_UNKNOWN
	}

	return pb.MessageState(value)
}
