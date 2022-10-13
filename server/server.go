package server

import (
	"context"

	"github.com/hulb/grpc-test/pb"
	"go.uber.org/zap"
	"google.golang.org/protobuf/types/known/structpb"
)

type server struct {
	pb.UnimplementedGreeterServer
	lg *zap.Logger
}

func (s *server) SayHi(ctx context.Context, req *pb.Request) (*pb.Response, error) {
	s.lg.Info("SayHi called", zap.String("msg", req.Msg), zap.Time("time", req.Time.AsTime()))
	return &pb.Response{
		Back: "back,hi",
		Value: &structpb.Struct{
			Fields: map[string]*structpb.Value{
				"a": structpb.NewStringValue("value"),
				"b": structpb.NewNumberValue(100),
			},
		},
	}, nil
}

func NewServer(lg *zap.Logger) pb.GreeterServer {
	return &server{lg: lg}
}
