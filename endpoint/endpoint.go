package endpoint

import (
	"context"
	"errors"

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
			RegencyID:  req.RegencyID,
			Page:       req.Page,
			Limit:      req.Limit,
			CategoryID: req.CategoryID,
			Title:      req.Title,
		})
		if err != nil {
			return nil, err
		}

		data := encodeResponse(resp.Data)

		videoWithMeta := &VideoResponse{
			Data:     data,
			Metadata: resp.Metadata,
		}

		return map[string]interface{}{
			"data": videoWithMeta,
		}, nil
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

		data := &VideoDetail{}

		if resp != nil {
			data.ID = resp.ID
			data.Title = resp.Title
			data.TotalLikes = resp.TotalLikes
			data.IsPushNotification = resp.IsPushNotification
			data.Sequence = resp.Sequence
			data.Category = resp.Category
			data.Source = resp.Source
			data.VideoURL = resp.VideoURL
			data.Status = resp.Status
			data.CreatedAt = resp.CreatedAt
			data.UpdatedAt = resp.UpdatedAt
			data.CreatedBy = resp.CreatedBy
			data.UpdatedBy = resp.UpdatedBy
			data.StatusLabel = GetStatusLabel[resp.Status]["id"]

			if resp.Regency != nil {
				data.RegencyID = converter.SetPointerInt64(resp.Regency.ID)
				data.Regency = resp.Regency
			}
		} else {
			data = nil
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

		if ok := containsStatus(Status, req.Status); !ok {
			return nil, errors.New("status: must be a valid value")
		}

		if err = fs.CreateNewVideo(ctx, &model.CreateVideoRequest{
			Title:      req.Title,
			Source:     req.Source,
			CategoryID: req.CategoryID,
			RegencyID:  req.RegencyID,
			Sequence:   req.Sequence,
			VideoURL:   req.VideoURL,
			Status:     req.Status,
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
