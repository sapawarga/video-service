package usecase

import (
	"context"
	"errors"
	"math"

	"github.com/sapawarga/video-service/lib/converter"
	"github.com/sapawarga/video-service/model"
	"github.com/sapawarga/video-service/repository"

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
		limit = converter.GetInt64FromPointer(req.Limit)
	}
	if req.Page != nil && *req.Page > 1 {
		offset = (converter.GetInt64FromPointer(req.Page) - 1) * limit
	}

	request := &model.GetListVideoRepoRequest{
		RegencyID:  req.RegencyID,
		Offset:     &offset,
		Limit:      &limit,
		CategoryID: req.CategoryID,
		Title:      req.Title,
		Search:     req.Search,
		SortBy:     req.SortBy,
		OrderBy:    req.SortOrder,
	}

	resp, err := v.repo.GetListVideo(ctx, request)
	if err != nil {
		level.Error(logger).Log("error_get_list", err)
		return nil, err
	}
	videos, err := v.appendVideoData(ctx, resp)
	if err != nil {
		level.Error(logger).Log("error_append_list", err)
		return nil, err
	}
	meta := &model.Metadata{}

	if req.Page != nil {
		total, err := v.repo.GetMetadataVideo(ctx, request)
		if err != nil {
			level.Error(logger).Log("error_get_metadata", err)
			return nil, err
		}

		meta.Page = limit
		meta.TotalPage = math.Ceil(float64(converter.GetInt64FromPointer(total)) / float64(limit))
		meta.Total = converter.GetInt64FromPointer(total)
		meta.CurrentPage = converter.GetInt64FromPointer(req.Page)
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

	if resp == nil {
		return nil, nil
	}

	result := &model.VideoDetail{
		ID:                 resp.ID,
		Title:              resp.Title,
		Source:             resp.Source,
		VideoURL:           resp.VideoURL,
		Status:             resp.Status,
		Sequence:           resp.Sequence.Int64,
		TotalLikes:         resp.TotalLikes.Int64,
		IsPushNotification: model.BoolFromInt[resp.IsPushNotification.Int64],
		CreatedAt:          converter.SetPointerInt64(resp.CreatedAt),
		UpdatedAt:          converter.SetPointerInt64(resp.UpdatedAt),
		CreatedBy:          converter.SetPointerInt64(resp.CreatedBy),
		UpdatedBy:          converter.SetPointerInt64(resp.UpdatedBy),
	}

	videoData, err := v.appendVideoData(ctx, []*model.VideoResponse{resp})
	if err != nil {
		level.Error(logger).Log("error_get_detail", err)
		return nil, err
	}

	result.Category = videoData[0].Category
	result.Regency = videoData[0].Regency

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
		if _, err = v.repo.GetLocationByID(ctx, converter.GetInt64FromPointer(req.RegencyID)); err != nil {
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

	data, err := v.repo.GetDetailVideo(ctx, converter.GetInt64FromPointer(req.ID))
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
		if _, err := v.repo.GetCategoryNameByID(ctx, converter.GetInt64FromPointer(req.CategoryID)); err != nil {
			return err
		}
	}

	if req.RegencyID != nil {
		if _, err := v.repo.GetLocationByID(ctx, converter.GetInt64FromPointer(req.RegencyID)); err != nil {
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

func (v *Video) appendVideoData(ctx context.Context, data []*model.VideoResponse) ([]*model.Video, error) {
	result := make([]*model.Video, 0)
	for _, video := range data {
		categoryName, err := v.repo.GetCategoryNameByID(ctx, video.CategoryID)
		if err != nil {
			return nil, err
		}

		var location *model.Location
		if video.RegencyID.Valid {
			location, err = v.repo.GetLocationByID(ctx, video.RegencyID.Int64)
			if err != nil {
				return nil, err
			}
		}

		video := &model.Video{
			ID:    video.ID,
			Title: video.Title,
			Category: &model.Category{
				ID:   video.CategoryID,
				Name: converter.GetStringFromPointer(categoryName),
			},
			Source:             video.Source,
			VideoURL:           video.VideoURL,
			Regency:            location,
			IsPushNotification: converter.ConvertBoolFromInteger(video.IsPushNotification.Int64),
			TotalLikes:         video.TotalLikes.Int64,
			Status:             video.Status,
			CreatedAt:          video.CreatedAt,
			UpdatedAt:          video.UpdatedAt,
			CreatedBy:          video.CreatedBy,
			UpdatedBy:          video.UpdatedBy,
		}
		result = append(result, video)
	}
	return result, nil
}
