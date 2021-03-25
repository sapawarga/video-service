package mysql

import (
	"bytes"
	"context"
	"database/sql"

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
		SELECT
			id, category_id, title, source, video_url, kabkota_id, status, FROM_UNIXTIME(created_at) as created_at, 
			FROM_UNIXTIME(updated_at) as updated_at, created_by, updated_by
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

func (r *VideoRepository) GetMetadataVideo(ctx context.Context, req *model.GetListVideoRepoRequest) (*int64, error) {
	var query bytes.Buffer
	var queryParams []interface{}
	var total *int64
	var err error

	query.WriteString(`
		SELECT COUNT(1) FROM videos
	`)
	if req.RegencyID != nil {
		query.WriteString(" WHERE kabkota_id = ? ")
		queryParams = append(queryParams, req.RegencyID)
	}

	if ctx != nil {
		err = r.conn.GetContext(ctx, &total, query.String(), queryParams...)
	} else {
		err = r.conn.Get(&total, query.String(), queryParams...)
	}

	if err != nil {
		return nil, err
	}

	return total, nil
}

func (r *VideoRepository) GetDetailVideo(ctx context.Context, id int64) (*model.VideoResponse, error) {
	var query bytes.Buffer
	var result = &model.VideoResponse{}
	var err error

	query.WriteString(`
	SELECT
		id, category_id, title, source, video_url, kabkota_id, status, FROM_UNIXTIME(created_at) as created_at, 
		FROM_UNIXTIME(updated_at) as updated_at, created_by, updated_by
	FROM videos
	`)
	query.WriteString(" WHERE id = ? ")

	if ctx != nil {
		err = r.conn.GetContext(ctx, result, query.String(), id)
	} else {
		err = r.conn.Get(result, query.String(), id)
	}

	if err == sql.ErrNoRows {
		return nil, sql.ErrNoRows
	}

	if err != nil {
		return nil, err
	}

	return result, nil
}

func (r *VideoRepository) GetCategoryNameByID(ctx context.Context, id int64) (*string, error) {
	var query bytes.Buffer
	var result *string
	var err error

	query.WriteString(` SELECT name from categories WHERE id = ? AND type = 'video' AND status = 10 `)
	if ctx != nil {
		err = r.conn.GetContext(ctx, &result, query.String(), id)
	} else {
		err = r.conn.Get(&result, query.String(), id)
	}

	if err == sql.ErrNoRows {
		return nil, sql.ErrNoRows
	}

	if err != nil {
		return nil, err
	}

	return result, nil
}

func (r *VideoRepository) GetLocationNameByID(ctx context.Context, id int64) (*string, error) {
	var query bytes.Buffer
	var result *string
	var err error

	query.WriteString(` SELECT name from areas WHERE id = ?`)
	if ctx != nil {
		err = r.conn.GetContext(ctx, &result, query.String(), id)
	} else {
		err = r.conn.Get(&result, query.String(), id)
	}

	if err == sql.ErrNoRows {
		return nil, sql.ErrNoRows
	}

	if err != nil {
		return nil, err
	}

	return result, nil
}

func (r *VideoRepository) GetVideoStatistic(ctx context.Context) ([]*model.VideoStatistic, error) {
	var query bytes.Buffer
	var result = make([]*model.VideoStatistic, 0)
	var err error

	query.WriteString(`
		SELECT  c.id, c.name , COUNT(v.category_id) as count 
		FROM sapawarga.videos v 
		JOIN sapawarga.categories c  
		ON c.id = v.category_id 
		GROUP BY 1, 2
	`)

	if ctx != nil {
		err = r.conn.SelectContext(ctx, &result, query.String())
	} else {
		err = r.conn.Select(&result, query.String())
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

func (r *VideoRepository) Update(ctx context.Context, params *model.UpdateVideoRequest) error {
	var query bytes.Buffer
	var queryParams = make(map[string]interface{})
	var first = true
	var err error
	_, unixTime := helper.GetCurrentTimeUTC()

	query.WriteString(` UPDATE videos SET`)
	if params.CategoryID != nil {
		query.WriteString(` category_id = :category_id`)
		queryParams["category_id"] = params.CategoryID
		first = false
	}
	if params.Title != nil {
		if !first {
			query.WriteString(" , ")
		}
		query.WriteString(" title = :title ")
		queryParams["title"] = params.Title
		first = false
	}
	if params.Source != nil {
		if !first {
			query.WriteString(" , ")
		}
		query.WriteString(" source = :source ")
		queryParams["source"] = params.Source
		first = false
	}
	if params.VideoURL != nil {
		if !first {
			query.WriteString(" , ")
		}
		query.WriteString(" video_url = :video_url ")
		queryParams["video_url"] = params.VideoURL
		first = false
	}
	if params.Status != nil {
		if !first {
			query.WriteString(" , ")
		}
		query.WriteString(" status = :status")
		queryParams["status"] = params.Status
		first = false
	}
	if !first {
		query.WriteString(" , ")
	}
	query.WriteString(" kabkota_id = :regency_id ,  created_at = :updated_at, updated_at = :updated_at WHERE id = :id")
	queryParams["regency_id"] = params.RegencyID
	queryParams["updated_at"] = unixTime
	queryParams["id"] = params.ID

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

func (r *VideoRepository) Delete(ctx context.Context, id int64) error {
	var query bytes.Buffer
	var params = make(map[string]interface{})
	var err error

	query.WriteString(" UPDATE videos SET status = :status WHERE id = :id ")
	params["status"] = helper.DELETED
	params["id"] = id
	if ctx != nil {
		_, err = r.conn.NamedExecContext(ctx, query.String(), params)
	} else {
		_, err = r.conn.NamedExec(query.String(), params)
	}

	if err != nil {
		return err
	}

	return nil
}
