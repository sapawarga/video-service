package repository

import (
	"context"

	"github.com/sapawarga/video-service/model"
)

type DatabaseI interface {
	GetListVideo(ctx context.Context, req *model.GetListVideoRepoRequest) ([]*model.VideoResponse, error)
	GetMetadataVideo(ctx context.Context, req *model.GetListVideoRepoRequest) (*int64, error)
}
