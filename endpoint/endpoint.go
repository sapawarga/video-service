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
