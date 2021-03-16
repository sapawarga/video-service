package usecase

import (
	"context"

	"github.com/sapawarga/video-service/model"
)

type UsecaseI interface {
	GetListVideo(ctx context.Context, req *model.GetListVideoRequest) (interface{}, error)
	// GetDetailVideo(ctx context.Context, id int64) (interface{}, error)
	// GetStatisticVideo(ctx context.Context, req interface{}) (interface{}, error)
}
