package model

import (
	"database/sql"
	"time"
)

type VideoResponse struct {
	ID         int64         `db:"id"`
	Title      string        `db:"title"`
	CategoryID int64         `db:"category_id"`
	Source     string        `db:"source"`
	VideoURL   string        `db:"video_url"`
	RegencyID  sql.NullInt64 `db:"kabkota_id"`
	Status     int64         `db:"status"`
	CreatedAt  time.Time     `db:"created_at"`
	UpdatedAt  time.Time     `db:"updated_at"`
	CreatedBy  int64         `db:"created_by"`
}
