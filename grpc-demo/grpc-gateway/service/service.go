package service

import (
	"context"
	"fmt"
	v1 "grpcgateway/helloworld/v1"
	"log"

	"google.golang.org/genproto/googleapis/api/httpbody"
	"google.golang.org/protobuf/types/known/emptypb"
)

type service struct {
	v1.UnimplementedGreeterServer
}

func NewService() *service {
	return &service{}
}

func (s *service) SayHello(ctx context.Context, in *v1.HelloRequest) (*v1.HelloReply, error) {
	return &v1.HelloReply{Message: in.Name + " world"}, nil
}

func (s *service) Hello1(ctx context.Context, in *emptypb.Empty) (*httpbody.HttpBody, error) {
	return &httpbody.HttpBody{
		ContentType: "text/html",
		Data:        []byte("Hello1"),
	}, nil
}

func (s *service) Download(_ *emptypb.Empty, stream v1.Greeter_DownloadServer) error {
	msgs := []*httpbody.HttpBody{
		{
			ContentType: "text/html",
			Data:        []byte("Hello 1"),
		},
		{
			ContentType: "text/html",
			Data:        []byte("Hello 2"),
		},
	}

	for _, msg := range msgs {
		if err := stream.Send(msg); err != nil {
			return err
		}
	}

	return nil
}

func (s *service) UpdateV2(ctx context.Context, in *v1.UpdateV2Request) (*emptypb.Empty, error) {
	log.Println("uuid", in.Abe.Uuid)
	log.Println("string value", in.Abe.StringValue)
	log.Println(*in.UpdateMask)
	return &emptypb.Empty{}, nil
}

func (s *service) UpdateV2A(ctx context.Context, in *v1.UpdateV2Request) (*emptypb.Empty, error) {
	log.Println("uuid", in.Abe.Uuid)
	log.Println("string value", in.Abe.StringValue)
	log.Println("single", *in.Abe.SingleNested)
	log.Println(*in.UpdateMask)
	return &emptypb.Empty{}, nil
}

func (s *service) GetMessage(ctx context.Context, in *v1.GetMessageReq) (*v1.GetMessageResp, error) {
	if in.MessageId == "" {
		return nil, fmt.Errorf("name is empty")
	}
	log.Println("MessageId", in.MessageId)
	log.Println("UserId", in.UserId)

	log.Println("Revision", in.Revision)
	if in.Sub != nil {
		log.Println("Subfield", in.Sub.Subfield)
	}

	return &v1.GetMessageResp{
		Text: "reply " + in.MessageId,
	}, nil
}

func (s *service) UpdateMessage(ctx context.Context, in *v1.UpdateMessageRequest) (*v1.UpdateMessageResp, error) {
	log.Println("MessageId", in.MessageId)
	if in.Message != nil {
		log.Println("text", in.Message.Text)
	}

	return &v1.UpdateMessageResp{
		Text: in.MessageId,
	}, nil
}

func (s *service) UpdateMessage1(ctx context.Context, in *v1.UpdateMessage1Request) (*v1.UpdateMessage1Resp, error) {
	log.Println("MessageId", in.MessageId)
	log.Println("Text", in.Text)

	return &v1.UpdateMessage1Resp{
		Text: in.MessageId,
	}, nil
}
