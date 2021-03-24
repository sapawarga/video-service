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
			level.Error(logger).Log("error_get_metadata", err)
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

func (v *Video) GetDetailVideo(ctx context.Context, id int64) (*model.VideoDetail, error) {
	logger := kitlog.With(v.logger, "method", "GetDetailVideo")
	resp, err := v.repo.GetDetailVideo(ctx, id)
	if err != nil {
		level.Error(logger).Log("error_get_detail", err)
		return nil, err
	}

	result := &model.VideoDetail{
		ID:        resp.ID,
		Title:     resp.Title.String,
		Source:    resp.Source.String,
		VideoURL:  resp.VideoURL.String,
		Status:    resp.Status.Int64,
		CreatedAt: helper.SetPointerTime(resp.CreatedAt.Time),
		UpdatedAt: helper.SetPointerTime(resp.UpdatedAt.Time),
		CreatedBy: helper.SetPointerInt64(resp.CreatedBy.Int64),
		UpdatedBy: helper.SetPointerInt64(resp.UpdatedBy.Int64),
	}

	if resp.CategoryID.Valid {
		name, err := v.repo.GetCategoryNameByID(ctx, resp.CategoryID.Int64)
		if err != nil {
			level.Error(logger).Log("error_get_category", err)
			return nil, err
		}
		result.CategoryID = helper.SetPointerInt64(resp.CategoryID.Int64)
		result.CategoryName = name
	}
	if resp.RegencyID.Valid {
		name, err := v.repo.GetLocationNameByID(ctx, resp.RegencyID.Int64)
		if err != nil {
			level.Error(logger).Log("error_get_location", err)
			return nil, err
		}
		result.RegencyID = helper.SetPointerInt64(resp.RegencyID.Int64)
		result.RegencyName = name
	}

	return result, nil
}

func (v *Video) GetStatisticVideo(ctx context.Context) ([]*model.VideoStatisticUC, error) {
	logger := kitlog.With(v.logger, "method", "GetStatisticVideo")
	resp, err := v.repo.GetVideoStatistic(ctx)
	if err != nil {
		level.Error(logger).Log("error_get_video_statistic", err)
		return nil, err
	}

	result := make([]*model.VideoStatisticUC, 0)
	for _, v := range resp {
		result = append(result, &model.VideoStatisticUC{
			ID:    v.ID,
			Name:  v.Name.String,
			Count: v.Count,
		})
	}

	return result, nil
}

func (v *Video) CreateNewVideo(ctx context.Context, req *model.CreateVideoRequest) error {
	logger := kitlog.With(v.logger, "method", "CreateNewVideo")
	var err error
	if _, err = v.repo.GetCategoryNameByID(ctx, req.CategoryID); err != nil {
		level.Error(logger).Log("error_get_category", err)
		return err
	}

	if _, err = v.repo.GetLocationNameByID(ctx, req.RegencyID); err != nil {
		level.Error(logger).Log("error_get_regency", err)
		return err
	}

	if err = v.repo.Insert(ctx, req); err != nil {
		level.Error(logger).Log("error_insert_video", err)
		return err
	}
	return nil
}

func (v *Video) UpdateVideo(ctx context.Context, req *model.UpdateVideoRequest) error {
	logger := kitlog.With(v.logger, "method", "UpdateVideo")
	var err error
	if _, err = v.repo.GetCategoryNameByID(ctx, req.CategoryID); err != nil {
		level.Error(logger).Log("error_get_category", err)
		return err
	}

	if _, err = v.repo.GetLocationNameByID(ctx, req.RegencyID); err != nil {
		level.Error(logger).Log("error_get_regency", err)
		return err
	}

	if err = v.repo.Update(ctx, req); err != nil {
		level.Error(logger).Log("error_update_video", err)
		return err
	}
	return nil

}
