package mysql

import (
	"bytes"
	"context"
	"database/sql"
	"fmt"

	"github.com/sapawarga/video-service/lib/constants"
	"github.com/sapawarga/video-service/lib/converter"
	"github.com/sapawarga/video-service/lib/generator"
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
			id, category_id, title, source, video_url, kabkota_id, status, total_likes, is_push_notification, created_at, seq,
			updated_at, created_by, updated_by
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
		id, category_id, title, source, video_url, kabkota_id, status, created_at, 
		updated_at, created_by, updated_by
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
		return nil, nil
	}

	if err != nil {
		return nil, err
	}

	return result, nil
}

func (r *VideoRepository) GetLocationByID(ctx context.Context, id int64) (*model.Location, error) {
	var query bytes.Buffer
	var result *model.Location
	var err error

	query.WriteString(`  SELECT id, name, code_bps FROM areas WHERE id = ? `)
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
	SELECT c.id, c.name, IFNULL(count, 0) AS count FROM categories c
	LEFT JOIN (
		SELECT c.id, c.name, COUNT(v.id) AS count
		FROM videos v
		LEFT JOIN categories c ON v.category_id = c.id
		WHERE v.status <> -1
		GROUP BY c.id
	) AS statistic
	ON c.id = statistic.id
	WHERE c.type = 'video'
	AND c.status <> -1;
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
	_, current := generator.GetCurrentTimeUTC()
	// TODO: actor is from authenticator
	actor := 1
	query.WriteString("INSERT INTO videos")
	query.WriteString(`
		(category_id, title, source, video_url, kabkota_id, seq, status, created_by, created_at, updated_by, updated_at)`)
	query.WriteString(`VALUES(
		:category_id, :title, :source, :video_url, :kabkota_id, :seq, :status, :actor, :created_at, :actor, :updated_at)`)
	queryParams := map[string]interface{}{
		"category_id": params.CategoryID,
		"title":       params.Title,
		"source":      params.Source,
		"video_url":   params.VideoURL,
		"kabkota_id":  params.RegencyID,
		"status":      params.Status,
		"created_at":  current,
		"actor":       actor,
		"updated_at":  current,
		"seq":         params.Sequence,
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
	var err error
	_, unixTime := generator.GetCurrentTimeUTC()

	query.WriteString(` UPDATE videos SET`)
	if params.CategoryID != nil {
		query.WriteString(` category_id = :category_id`)
		queryParams["category_id"] = converter.GetInt64FromPointer(params.CategoryID)
	}
	if params.Title != nil {
		query.WriteString(updateNext(ctx, "title"))
		queryParams["title"] = converter.GetStringFromPointer(params.Title)
	}
	if params.Source != nil {
		query.WriteString(updateNext(ctx, "source"))
		queryParams["source"] = converter.GetStringFromPointer(params.Source)
	}
	if params.VideoURL != nil {
		query.WriteString(updateNext(ctx, "video_url"))
		queryParams["video_url"] = converter.GetStringFromPointer(params.VideoURL)
	}
	if params.Status != nil {
		query.WriteString(updateNext(ctx, "status"))
		queryParams["status"] = converter.GetInt64FromPointer(params.Status)
	}
	if params.Sequence != nil {
		query.WriteString(updateNext(ctx, "seq"))
		queryParams["seq"] = converter.GetInt64FromPointer(params.Sequence)
	}
	query.WriteString(" kabkota_id = :kabkota_id ,  created_at = :updated_at, updated_at = :updated_at WHERE id = :id")
	queryParams["kabkota_id"] = converter.GetInt64FromPointer(params.RegencyID)
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
	params["status"] = constants.DELETED
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

func (r *VideoRepository) HealthCheckReadiness(ctx context.Context) error {
	var err error
	if ctx != nil {
		err = r.conn.PingContext(ctx)
	} else {
		err = r.conn.Ping()
	}

	if err != nil {
		return err
	}

	return nil
}

func updateNext(ctx context.Context, field string) string {
	var query bytes.Buffer
	query.WriteString(fmt.Sprintf(" , %s = :%s ", field, field))
	return query.String()
}
