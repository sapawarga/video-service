package usecase

import (
	"context"
	"errors"
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

	var limit, offset int64 = 10, 0
	if req.Limit != nil {
		limit = helper.GetInt64FromPointer(req.Limit)
	}
	if req.Page != nil && *req.Page > 1 {
		offset = (helper.GetInt64FromPointer(req.Page) - 1) * limit
	}

	request := &model.GetListVideoRepoRequest{
		RegencyID: helper.SetPointerInt64(*req.RegencyID),
		Offset:    &offset,
		Limit:     &limit,
	}

	resp, err := v.repo.GetListVideo(ctx, request)
	if err != nil {
		level.Error(logger).Log("error_get_list", err)
		return nil, err
	}
	videos := v.appendVideoData(ctx, resp)

	meta := &model.Metadata{}

	if req.Page != nil {
		total, err := v.repo.GetMetadataVideo(ctx, request)
		if err != nil {
			level.Error(logger).Log("error_get_metadata", err)
			return nil, err
		}

		totalPage := int64(math.Ceil(float64(helper.GetInt64FromPointer(total) / limit)))

		meta.Page = helper.GetInt64FromPointer(req.Page)
		meta.TotalPage = totalPage
		meta.Total = helper.GetInt64FromPointer(total)
	}

	return &model.VideoWithMetadata{
		Data:     videos,
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

	if req.RegencyID != nil {
		if _, err = v.repo.GetLocationNameByID(ctx, helper.GetInt64FromPointer(req.RegencyID)); err != nil {
			level.Error(logger).Log("error_get_regency", err)
			return err
		}
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

	data, err := v.repo.GetDetailVideo(ctx, helper.GetInt64FromPointer(req.ID))
	if err != nil {
		level.Error(logger).Log("error_get_detail", err)
		return err
	}
	if data != nil {
		if err = v.updateRegencyOrCategory(ctx, req, data); err != nil {
			level.Error(logger).Log("error_update_category_or_regency", err)
			return err
		}
	}

	return nil
}

func (v *Video) updateRegencyOrCategory(ctx context.Context, req *model.UpdateVideoRequest, data *model.VideoResponse) error {
	if req.CategoryID != nil {
		if _, err := v.repo.GetCategoryNameByID(ctx, helper.GetInt64FromPointer(req.CategoryID)); err != nil {
			return err
		}
	}

	if req.RegencyID != nil {
		if _, err := v.repo.GetLocationNameByID(ctx, helper.GetInt64FromPointer(req.RegencyID)); err != nil {
			return err
		}
	}

	if err := v.repo.Update(ctx, req); err != nil {
		return err
	}
	return nil
}

func (v *Video) DeleteVideo(ctx context.Context, id int64) error {
	logger := kitlog.With(v.logger, "method", "DeleteVideo")
	data, err := v.repo.GetDetailVideo(ctx, id)
	if err != nil {
		level.Error(logger).Log("error_get_detail", err)
		return err
	}

	if data != nil {
		if err = v.repo.Delete(ctx, id); err != nil {
			level.Error(logger).Log("error_delete", err)
			return err
		}
	}

	return nil
}

func (v *Video) CheckHealthReadiness(ctx context.Context) error {
	logger := kitlog.With(v.logger, "method", "CheckHealthReadiness")
	if err := v.repo.HealthCheckReadiness(ctx); err != nil {
		level.Error(logger).Log("error", errors.New("service_not_ready"))
		return errors.New("service_not_ready")
	}
	return nil
}

func (v *Video) appendVideoData(ctx context.Context, data []*model.VideoResponse) []*model.Video {
	result := make([]*model.Video, 0)
	for _, v := range data {
		video := &model.Video{
			ID:         v.ID,
			Title:      v.Title.String,
			CategoryID: v.CategoryID.Int64,
			Source:     v.Source.String,
			VideoURL:   v.VideoURL.String,
			RegencyID:  v.RegencyID.Int64,
			Status:     v.Status.Int64,
			CreatedAt:  v.CreatedAt.Time,
			UpdatedAt:  v.UpdatedAt.Time,
			CreatedBy:  v.CreatedBy.Int64,
			UpdatedBy:  v.UpdatedBy.Int64,
		}
		result = append(result, video)
	}
	return result
}
