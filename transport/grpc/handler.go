package grpc

import (
	"context"

	kitgrpc "github.com/go-kit/kit/transport/grpc"
	transportVideo "github.com/sapawarga/proto-file/video"
)

type grpcServer struct {
	getList      kitgrpc.Handler
	getDetail    kitgrpc.Handler
	getStatistic kitgrpc.Handler
}

func (g *grpcServer) GetListVideo(ctx context.Context, req *transportVideo.GetListVideoRequest) (*transportVideo.GetListVideoResponse, error) {
	_, resp, err := g.getList.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp.(*transportVideo.GetListVideoResponse), nil
}

func (g *grpcServer) GetDetailVideo(ctx context.Context, req *transportVideo.RequestID) (*transportVideo.GetDetailVideoResponse, error) {
	_, resp, err := g.getDetail.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp.(*transportVideo.GetDetailVideoResponse), nil
}

func (g *grpcServer) GetStatisticVideo(ctx context.Context, req *transportVideo.NonRequest) (*transportVideo.GetStatisticVideoResponse, error) {
	_, resp, err := g.getStatistic.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp.(*transportVideo.GetStatisticVideoResponse), nil
}
