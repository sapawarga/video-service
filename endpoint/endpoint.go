package endpoint

import (
	"context"

	"github.com/go-kit/kit/endpoint"
	"github.com/sapawarga/video-service/helper"
	"github.com/sapawarga/video-service/model"
	"github.com/sapawarga/video-service/usecase"
)

// MakeGetListVideo ...
func MakeGetListVideo(ctx context.Context, fs usecase.UsecaseI) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(*GetVideoRequest)
		resp, err := fs.GetListVideo(ctx, &model.GetListVideoRequest{
			RegencyID: req.RegencyID,
			Page:      req.Page,
			Limit:     req.Limit,
		})
		if err != nil {
			return nil, err
		}

		return resp, nil
	}
}

// MakeGetDetailVideo ...
func MakeGetDetailVideo(ctx context.Context, fs usecase.UsecaseI) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(*RequestID)
		resp, err := fs.GetDetailVideo(ctx, req.ID)
		if err != nil {
			return nil, err
		}

		data := &VideoDetail{
			ID:                 resp.ID,
			Title:              resp.Title,
			TotalLikes:         resp.TotalLikes,
			IsPushNotification: resp.IsPushNotification,
			Sequence:           resp.Sequence,
			Cateogry:           resp.Category,
			Source:             resp.Source,
			VideoURL:           resp.VideoURL,
			Status:             resp.Status,
			CreatedAt:          resp.CreatedAt,
			UpdatedAt:          resp.UpdatedAt,
			CreatedBy:          resp.CreatedBy,
			UpdatedBy:          resp.UpdatedBy,
			StatusLabel:        GetStatusLabel[resp.Status]["id"],
		}

		if resp.Regency != nil {
			data.RegencyID = helper.SetPointerInt64(resp.Regency.ID)
			data.Regency = resp.Regency
		}

		result := map[string]interface{}{
			"data": data,
		}
		return result, nil
	}
}

// MakeGetVideoStatistic ...
func MakeGetVideoStatistic(ctx context.Context, fs usecase.UsecaseI) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		resp, err := fs.GetStatisticVideo(ctx)
		if err != nil {
			return nil, err
		}

		meta := &model.Metadata{
			Page:        20,
			TotalPage:   1,
			CurrentPage: 1,
			Total:       int64(len(resp)),
		}

		data := &VideoStatisticResponse{
			Data:     resp,
			Metadata: meta,
		}
		return map[string]interface{}{
			"data": data,
		}, nil
	}
}

// MakeCreateNewVideo ...
func MakeCreateNewVideo(ctx context.Context, fs usecase.UsecaseI) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(*CreateVideoRequest)
		if err := ValidateInputs(req); err != nil {
			return nil, err
		}

		if err = fs.CreateNewVideo(ctx, &model.CreateVideoRequest{
			Title:      helper.GetStringFromPointer(req.Source),
			Source:     helper.GetStringFromPointer(req.Source),
			CategoryID: helper.GetInt64FromPointer(req.CategoryID),
			RegencyID:  req.RegencyID,
			Sequence:   helper.GetInt64FromPointer(req.Sequence),
			VideoURL:   helper.GetStringFromPointer(req.VideoURL),
			Status:     helper.GetInt64FromPointer(req.Status),
		}); err != nil {
			return nil, err
		}

		return &StatusResponse{
			Code:    helper.STATUS_CREATED,
			Message: "video_has_created_successfully",
		}, nil
	}
}

func MakeUpdateVideo(ctx context.Context, fs usecase.UsecaseI) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(*UpdateVideoRequest)
		if err := ValidateInputs(req); err != nil {
			return nil, err
		}

		if err = fs.UpdateVideo(ctx, &model.UpdateVideoRequest{
			ID:         req.ID,
			Title:      req.Title,
			Source:     req.Source,
			CategoryID: req.CategoryID,
			RegencyID:  req.RegencyID,
			VideoURL:   req.VideoURL,
			Status:     req.Status,
			Sequence:   req.Sequence,
		}); err != nil {
			return nil, err
		}
		return &StatusResponse{
			Code:    helper.STATUS_UPDATED,
			Message: "video_has_updated_successfully",
		}, nil
	}
}

func MakeDeleteVideo(ctx context.Context, fs usecase.UsecaseI) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(*RequestID)

		if err := fs.DeleteVideo(ctx, req.ID); err != nil {
			return nil, err
		}

		return &StatusResponse{
			Code:    helper.STATUS_DELETED,
			Message: "video_has_deleted_successfully",
		}, nil
	}
}

func MakeCheckHealthy(ctx context.Context) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		return &StatusResponse{
			Code:    helper.STATUS_OK,
			Message: "service_is_ok",
		}, nil
	}
}

func MakeCheckReadiness(ctx context.Context, fs usecase.UsecaseI) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		if err := fs.CheckHealthReadiness(ctx); err != nil {
			return nil, err
		}
		return &StatusResponse{
			Code:    helper.STATUS_OK,
			Message: "service_is_ready",
		}, nil
	}
}
