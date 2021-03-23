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

	return &grpcServer{
		videoGetListHandler,
		videoGetDetailHandler,
		videoGetStatisticHandler,
	}
}

func decodingGetListVideoRequest(ctx context.Context, r interface{}) (interface{}, error) {
	req := r.(*transportVideo.GetListVideoRequest)
	var regencyID, page *int64

	if req.Page != 0 {
		page = helper.SetPointerInt64(req.Page)
	}
	if req.RegencyId != 0 {
		regencyID = helper.SetPointerInt64(req.RegencyId)
	}

	return &endpoint.GetVideoRequest{
		RegencyID: regencyID,
		Page:      page,
	}, nil
}

func encodingGetListVideoResponse(ctx context.Context, r interface{}) (interface{}, error) {
	resp := r.(*endpoint.VideoResponse)
	videoResp := make([]*transportVideo.VideoList, 0)

	for _, v := range resp.Data {
		video := &transportVideo.VideoList{
			Id:         v.ID,
			Title:      v.Title.String,
			CategoryId: v.CategoryID.Int64,
			Source:     v.Source.String,
			VideoUrl:   v.VideoURL.String,
			RegencyId:  v.RegencyID.Int64,
			Status:     v.Status.Int64,
			CreatedAt:  v.CreatedAt.Time.String(),
			UpdatedAt:  v.UpdatedAt.Time.String(),
			CreatedBy:  v.CreatedBy.Int64,
			UpdatedBy:  v.UpdatedBy.Int64,
		}
		videoResp = append(videoResp, video)
	}

	metadata := &transportVideo.Metadata{
		Page:      resp.Metadata.Page,
		TotalPage: resp.Metadata.TotalPage,
		Total:     resp.Metadata.Total,
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
		CategoryId:   helper.GetInt64FromPointer(resp.CategoryID),
		CategoryName: helper.GetStringFromPointer(resp.CategoryName),
		Source:       resp.Source,
		VideoUrl:     resp.VideoURL,
		RegencyId:    helper.GetInt64FromPointer(resp.RegencyID),
		RegencyName:  helper.GetStringFromPointer(resp.RegencyName),
		Status:       resp.Status,
		CreatedAt:    resp.CreatedAt.String(),
		UpdatedAt:    resp.UpdatedAt.String(),
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
