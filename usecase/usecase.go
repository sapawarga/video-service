package usecase

import (
	"context"
	"math"

	"github.com/sapawarga/video-service/model"
	"github.com/sapawarga/video-service/repository"

	"github.com/sapawarga/video-service/helper"

	kitlog "github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
)

type Video struct {
	repo   repository.DatabaseI
	logger kitlog.Logger
}

func NewVideo(repo repository.DatabaseI, logger kitlog.Logger) *Video {
	return &Video{
		repo:   repo,
		logger: logger,
	}
}

func (v *Video) GetListVideo(ctx context.Context, req *model.GetListVideoRequest) (*model.VideoWithMetadata, error) {
	logger := kitlog.With(v.logger, "method", "GetListVideo")

	var limit, offset int64 = 0, 0
	if req.Page != nil {
		limit = 10
		offset = (helper.GetInt64FromPointer(req.Page) - 1) * limit
	}
	request := &model.GetListVideoRepoRequest{
		RegencyID: req.RegencyID,
		Offset:    helper.SetPointerInt64(offset),
	}
	if limit > 0 {
		request.Limit = helper.SetPointerInt64(limit)
	}

	resp, err := v.repo.GetListVideo(ctx, request)
	if err != nil {
		level.Error(logger).Log("error_get_list", err)
		return nil, err
	}

	meta := &model.Metadata{}

	if req.Page != nil {
		total, err := v.repo.GetMetadataVideo(ctx, request)
		if err != nil {
			level.Error(logger).Log()
			return nil, err
		}

		totalPage := int64(math.Floor(float64(helper.GetInt64FromPointer(total) / limit)))

		meta.Page = helper.GetInt64FromPointer(req.Page)
		meta.TotalPage = totalPage
		meta.Total = helper.GetInt64FromPointer(total)
	}

	return &model.VideoWithMetadata{
		Data:     resp,
		Metadata: meta,
	}, nil
}
