package endpoint

import (
	"context"

	"github.com/go-kit/kit/endpoint"
	"github.com/sapawarga/video-service/lib/constants"
	"github.com/sapawarga/video-service/lib/converter"
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
			Category:           resp.Category,
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
			data.RegencyID = converter.SetPointerInt64(resp.Regency.ID)
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
			Page:        constants.PER_PAGE,
			TotalPage:   float64(constants.TOTAL_PAGE),
			CurrentPage: constants.TOTAL_PAGE,
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
			Title:      converter.GetStringFromPointer(req.Source),
			Source:     converter.GetStringFromPointer(req.Source),
			CategoryID: converter.GetInt64FromPointer(req.CategoryID),
			RegencyID:  req.RegencyID,
			Sequence:   converter.GetInt64FromPointer(req.Sequence),
			VideoURL:   converter.GetStringFromPointer(req.VideoURL),
			Status:     converter.GetInt64FromPointer(req.Status),
		}); err != nil {
			return nil, err
		}

		return &StatusResponse{
			Code:    constants.STATUS_CREATED,
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
			Code:    constants.STATUS_UPDATED,
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
			Code:    constants.STATUS_DELETED,
			Message: "video_has_deleted_successfully",
		}, nil
	}
}

func MakeCheckHealthy(ctx context.Context) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		return &StatusResponse{
			Code:    constants.STATUS_OK,
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
			Code:    constants.STATUS_OK,
			Message: "service_is_ready",
		}, nil
	}
}
