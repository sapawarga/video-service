package grpc

import (
	"context"

	"github.com/sapawarga/video-service/endpoint"
	"github.com/sapawarga/video-service/helper"
	"github.com/sapawarga/video-service/usecase"

	kitgrpc "github.com/go-kit/kit/transport/grpc"
	transportVideo "github.com/sapawarga/proto-file/video"
)

func MakeHandler(ctx context.Context, fs usecase.UsecaseI) transportVideo.VideoHandlerServer {
	videoGetListHandler := kitgrpc.NewServer(
		endpoint.MakeGetListVideo(ctx, fs),
		decodingGetListVideoRequest,
		encodingGetListVideoResponse,
	)

	videoGetDetailHandler := kitgrpc.NewServer(
		endpoint.MakeGetDetailVideo(ctx, fs),
		decodingRequestID,
		encodingGetDetailResponse,
	)

	videoGetStatisticHandler := kitgrpc.NewServer(
		endpoint.MakeGetVideoStatistic(ctx, fs),
		decodingNoRequest,
		encodingGetVideoStatisticResponse,
	)

	videoCreateHandler := kitgrpc.NewServer(
		endpoint.MakeCreateNewVideo(ctx, fs),
		decodingCreateNewVideoRequest,
		encodingStatusResponse,
	)

	videoUpdateHandler := kitgrpc.NewServer(
		endpoint.MakeUpdateVideo(ctx, fs),
		decodingUpdateVideo,
		encodingStatusResponse,
	)

	videoDeleteHandler := kitgrpc.NewServer(
		endpoint.MakeDeleteVideo(ctx, fs),
		decodingRequestID,
		encodingStatusResponse,
	)

	return &grpcServer{
		videoGetListHandler,
		videoGetDetailHandler,
		videoGetStatisticHandler,
		videoCreateHandler,
		videoUpdateHandler,
		videoDeleteHandler,
	}
}

func decodingGetListVideoRequest(ctx context.Context, r interface{}) (interface{}, error) {
	req := r.(*transportVideo.GetListVideoRequest)
	return &endpoint.GetVideoRequest{
		RegencyID: helper.SetPointerInt64(req.GetRegencyId()),
		Page:      helper.SetPointerInt64(req.GetPage()),
	}, nil
}

func encodingGetListVideoResponse(ctx context.Context, r interface{}) (interface{}, error) {
	resp := r.(*endpoint.VideoResponse)
	videoResp := make([]*transportVideo.VideoList, 0)

	for _, v := range resp.Data {
		video := &transportVideo.VideoList{
			Id:         v.ID,
			Title:      v.Title,
			CategoryId: v.Category.ID,
			Source:     v.Source,
			VideoUrl:   v.VideoURL,
			Status:     v.Status,
			CreatedBy:  v.CreatedBy,
		}
		videoResp = append(videoResp, video)
	}

	metadata := &transportVideo.Metadata{
		Page:  resp.Metadata.Page,
		Total: resp.Metadata.Total,
	}

	return &transportVideo.GetListVideoResponse{
		Data:     videoResp,
		Metadata: metadata,
	}, nil
}

func decodingRequestID(ctx context.Context, r interface{}) (interface{}, error) {
	req := r.(*transportVideo.RequestID)
	return &endpoint.RequestID{
		ID: req.GetId(),
	}, nil
}

func encodingGetDetailResponse(ctx context.Context, r interface{}) (interface{}, error) {
	resp := r.(*endpoint.VideoDetail)
	return &transportVideo.GetDetailVideoResponse{
		Id:           resp.ID,
		Title:        resp.Title,
		CategoryId:   resp.Cateogry.ID,
		CategoryName: resp.Cateogry.Name,
		Source:       resp.Source,
		VideoUrl:     resp.VideoURL,
		Status:       resp.Status,
		CreatedBy:    helper.GetInt64FromPointer(resp.CreatedBy),
		UpdatedBy:    helper.GetInt64FromPointer(resp.UpdatedBy),
	}, nil

}

func decodingNoRequest(ctx context.Context, r interface{}) (interface{}, error) {
	return r, nil
}

func encodingGetVideoStatisticResponse(ctx context.Context, r interface{}) (interface{}, error) {
	resp := r.(*endpoint.VideoStatisticResponse)
	result := make([]*transportVideo.VideoStatistic, 0)
	for _, v := range resp.Data {
		result = append(result, &transportVideo.VideoStatistic{
			Id:    v.ID,
			Name:  v.Name,
			Count: v.Count,
		})
	}

	return &transportVideo.GetStatisticVideoResponse{
		Data: result,
	}, nil
}

func decodingCreateNewVideoRequest(ctx context.Context, r interface{}) (interface{}, error) {
	req := r.(*transportVideo.CreateVideoRequest)

	return &endpoint.CreateVideoRequest{
		Title:      helper.SetPointerString(req.GetSource()),
		Source:     helper.SetPointerString(req.GetSource()),
		CategoryID: helper.SetPointerInt64(req.GetCategoryId()),
		RegencyID:  helper.SetPointerInt64(req.GetRegencyId()),
		VideoURL:   helper.SetPointerString(req.GetVideoUrl()),
		Status:     &req.Status,
	}, nil
}

func encodingStatusResponse(ctx context.Context, r interface{}) (interface{}, error) {
	resp := r.(*endpoint.StatusResponse)

	return &transportVideo.StatusResponse{
		Code:    resp.Code,
		Message: resp.Message,
	}, nil
}

func decodingUpdateVideo(ctx context.Context, r interface{}) (interface{}, error) {
	req := r.(*transportVideo.UpdateVideoRequest)

	return &endpoint.UpdateVideoRequest{
		ID:         helper.SetPointerInt64(req.GetId()),
		Title:      helper.SetPointerString(req.GetTitle()),
		Source:     helper.SetPointerString(req.GetSource()),
		CategoryID: helper.SetPointerInt64(req.GetCategoryId()),
		RegencyID:  helper.SetPointerInt64(req.GetRegencyId()),
		VideoURL:   helper.SetPointerString(req.GetVideoUrl()),
		Status:     helper.SetPointerInt64(req.GetStatus()),
	}, nil
}
