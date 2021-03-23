package model

import (
	"database/sql"
)

type VideoResponse struct {
	ID         int64          `db:"id"`
	Title      sql.NullString `db:"title"`
	CategoryID sql.NullInt64  `db:"category_id"`
	Source     sql.NullString `db:"source"`
	VideoURL   sql.NullString `db:"video_url"`
	RegencyID  sql.NullInt64  `db:"kabkota_id"`
	Status     sql.NullInt64  `db:"status"`
	CreatedAt  sql.NullTime   `db:"created_at"`
	UpdatedAt  sql.NullTime   `db:"updated_at"`
	CreatedBy  sql.NullInt64  `db:"created_by"`
	UpdatedBy  sql.NullInt64  `db:"updated_by"`
}

type VideoStatistic struct {
	ID    int64          `db:"id"`
	Name  sql.NullString `db:"name"`
	Count int64          `db:"count"`
}
