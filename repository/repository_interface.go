package repository

import (
	"context"

	"github.com/sapawarga/video-service/model"
)

type DatabaseI interface {
	// query get
	GetListVideo(ctx context.Context, req *model.GetListVideoRepoRequest) ([]*model.VideoResponse, error)
	GetMetadataVideo(ctx context.Context, req *model.GetListVideoRepoRequest) (*int64, error)
	GetDetailVideo(ctx context.Context, id int64) (*model.VideoResponse, error)
	GetCategoryNameByID(ctx context.Context, id int64) (*string, error)
	GetLocationByID(ctx context.Context, id int64) (*model.Location, error)
	GetVideoStatistic(ctx context.Context) ([]*model.VideoStatistic, error)
	// query insert
	Insert(ctx context.Context, params *model.CreateVideoRequest) error
	// query update
	Update(ctx context.Context, params *model.UpdateVideoRequest) error
	// query delete
	Delete(ctx context.Context, id int64) error
	// health check readiness
	HealthCheckReadiness(ctx context.Context) error
}
