package service

import (
	"context"
	"grpc-server/proto"
)

type Demo struct {
}

func (c Demo) GetData(ctx context.Context, msg *proto.DemoReq) (*proto.DemoRsp, error) {
	params := msg.A
	return &proto.DemoRsp{
		Rel: params,
	}, nil
}