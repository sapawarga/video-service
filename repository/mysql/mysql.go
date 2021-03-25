package mysql

import (
	"bytes"
	"context"

	"github.com/sapawarga/video-service/helper"
	"github.com/sapawarga/video-service/model"

	"github.com/jmoiron/sqlx"
)

type VideoRepository struct {
	conn *sqlx.DB
}

func NewVideoRepository(db *sqlx.DB) *VideoRepository {
	return &VideoRepository{
		conn: db,
	}
}

func (r *VideoRepository) GetListVideo(ctx context.Context, req *model.GetListVideoRepoRequest) ([]*model.VideoResponse, error) {
	var query bytes.Buffer
	var queryParams []interface{}
	var result []*model.VideoResponse
	var err error

	query.WriteString(`
		category_id, title, source, video_url, kabkota_id, status, created_at, updated_at, created_by, updated_by
		FROM videos
	`)
	if req.RegencyID != nil {
		query.WriteString(" WHERE kabkota_id = ? ")
		queryParams = append(queryParams, req.RegencyID)
	}
	if req.Limit != nil && req.Offset != nil {
		query.WriteString(" LIMIT ?, ? ")
		queryParams = append(queryParams, req.Offset, req.Limit)
	}

	if ctx != nil {
		err = r.conn.SelectContext(ctx, &result, query.String(), queryParams...)
	} else {
		err = r.conn.Select(&result, query.String(), queryParams...)
	}

	if err != nil {
		return nil, err
	}

	return result, nil
}

func (r *VideoRepository) Insert(ctx context.Context, params *model.CreateVideoRequest) error {
	var query bytes.Buffer
	var err error
	_, current := helper.GetCurrentTimeUTC()

	query.WriteString("INSERT INTO videos")
	query.WriteString(`
		(category_id, title, source, video_url, kabkota_id, seq, status, created_at, updated_at)`)
	query.WriteString(`VALUES(
		:category_id, :title, :source, :video_url, :kabkota_id, 1, :status, :created_at, :updated_at)`)
	queryParams := map[string]interface{}{
		"category_id": params.CategoryID,
		"title":       params.Title,
		"source":      params.Source,
		"video_url":   params.VideoURL,
		"kabkota_id":  params.RegencyID,
		"status":      params.Status,
		"created_at":  current,
		"updated_at":  current,
	}

	if ctx != nil {
		_, err = r.conn.NamedExecContext(ctx, query.String(), queryParams)
	} else {
		_, err = r.conn.NamedExec(query.String(), queryParams)
	}

	if err != nil {
		return err
	}

	return nil
}
