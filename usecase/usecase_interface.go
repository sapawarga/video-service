package usecase

import (
	"context"

	"github.com/sapawarga/video-service/model"
)

type UsecaseI interface {
	GetListVideo(ctx context.Context, req *model.GetListVideoRequest) (*model.VideoWithMetadata, error)
	GetDetailVideo(ctx context.Context, id int64) (*model.VideoDetail, error)
	GetStatisticVideo(ctx context.Context) ([]*model.VideoStatisticUC, error)
	CreateNewVideo(ctx context.Context, req *model.CreateVideoRequest) error
}
