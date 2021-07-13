package model

import (
	"database/sql"
)

// VideoResponse ...
type VideoResponse struct {
	ID                 int64         `db:"id"`
	Title              string        `db:"title"`
	CategoryID         int64         `db:"category_id"`
	Source             string        `db:"source"`
	VideoURL           string        `db:"video_url"`
	RegencyID          sql.NullInt64 `db:"kabkota_id"`
	IsPushNotification sql.NullInt64 `db:"is_push_notification"`
	TotalLikes         sql.NullInt64 `db:"total_likes"`
	Status             int64         `db:"status"`
	Sequence           sql.NullInt64 `db:"seq"`
	CreatedAt          int64         `db:"created_at"`
	UpdatedAt          int64         `db:"updated_at"`
	CreatedBy          int64         `db:"created_by"`
	UpdatedBy          int64         `db:"updated_by"`
}

// VideoStatistic ...
type VideoStatistic struct {
	ID    int64          `db:"id"`
	Name  sql.NullString `db:"name"`
	Count int64          `db:"count"`
}

// Location ...
type Location struct {
	ID      int64          `db:"id" json:"id" `
	BPSCode sql.NullString `db:"code_bps" json:"code_bps" `
	Name    sql.NullString `db:"name" json:"name" `
}
