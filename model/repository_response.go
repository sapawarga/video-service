package model

import (
	"database/sql"
)

// VideoResponse ...
type VideoResponse struct {
	ID                 int64          `db:"id"`
	Title              sql.NullString `db:"title"`
	CategoryID         sql.NullInt64  `db:"category_id"`
	Source             sql.NullString `db:"source"`
	VideoURL           sql.NullString `db:"video_url"`
	RegencyID          sql.NullInt64  `db:"kabkota_id"`
	IsPushNotification sql.NullInt64  `db:"is_push_notification"`
	TotalLikes         sql.NullInt64  `db:"total_likes"`
	Status             sql.NullInt64  `db:"status"`
	Sequence           sql.NullInt64  `db:"seq"`
	CreatedAt          sql.NullInt64  `db:"created_at"`
	UpdatedAt          sql.NullInt64  `db:"updated_at"`
	CreatedBy          sql.NullInt64  `db:"created_by"`
	UpdatedBy          sql.NullInt64  `db:"updated_by"`
}

// VideoStatistic ...
type VideoStatistic struct {
	ID    int64          `db:"id"`
	Name  sql.NullString `db:"name"`
	Count int64          `db:"count"`
}

// Location ...
type Location struct {
	ID      int64  `db:"id" json:"id" `
	BPSCode string `db:"code_bps" json:"code_bps" `
	Name    string `db:"name" json:"name" `
}
