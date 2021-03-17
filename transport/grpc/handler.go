package grpc

import (
	"context"

	kitgrpc "github.com/go-kit/kit/transport/grpc"
	transportVideo "github.com/sapawarga/proto-file/video"
)

type grpcServer struct {
	getList kitgrpc.Handler
}

func (g *grpcServer) GetListVideo(ctx context.Context, req *transportVideo.GetListVideoRequest) (*transportVideo.GetListVideoResponse, error) {
	_, resp, err := g.getList.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp.(*transportVideo.GetListVideoResponse), nil
}
