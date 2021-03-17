package endpoint

import (
	"context"

	"github.com/go-kit/kit/endpoint"
	"github.com/sapawarga/video-service/model"
	"github.com/sapawarga/video-service/usecase"
)

func MakeGetListVideo(ctx context.Context, fs usecase.UsecaseI) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(*GetVideoRequest)
		resp, err := fs.GetListVideo(ctx, &model.GetListVideoRequest{
			RegencyID: req.RegencyID,
			Page:      req.Page,
		})
		if err != nil {
			return nil, err
		}

		return &VideoResponse{
			Data:     resp.Data,
			Metadata: resp.Metadata,
		}, nil
	}
}

func MakeGetDetailVideo(ctx context.Context, fs usecase.UsecaseI) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(*RequestID)
		resp, err := fs.GetDetailVideo(ctx, req.ID)
		if err != nil {
			return nil, err
		}

		return &VideoDetail{
			ID:           resp.ID,
			Title:        resp.Title,
			CategoryID:   resp.CategoryID,
			CategoryName: resp.CategoryName,
			Source:       resp.Source,
			VideoURL:     resp.VideoURL,
			RegencyID:    resp.RegencyID,
			RegencyName:  resp.RegencyName,
			Status:       resp.Status,
			CreatedAt:    resp.CreatedAt,
			UpdatedAt:    resp.UpdatedAt,
			CreatedBy:    resp.CreatedBy,
			UpdatedBy:    resp.UpdatedBy,
		}, nil
	}
}
