package model

import "time"

type VideoWithMetadata struct {
	Data     []*Video
	Metadata *Metadata
}

type Metadata struct {
	Page      int64
	TotalPage int64
	Total     int64
}

type Video struct {
	ID         int64     `json:"id"`
	Title      string    `json:"title,omitempty"`
	CategoryID int64     `json:"category_id,omitempty"`
	Source     string    `json:"source,omitempty"`
	VideoURL   string    `json:"video_url,omitempty"`
	RegencyID  int64     `json:"regency_id,omitempty"`
	Status     int64     `json:"status"`
	CreatedAt  time.Time `json:"created_at,omitempty"`
	UpdatedAt  time.Time `json:"updated_at,omitempty"`
	CreatedBy  int64     `json:"created_by,omitempty"`
	UpdatedBy  int64     `json:"updated_by,omitempty"`
}

type VideoDetail struct {
	ID           int64
	Title        string
	CategoryID   *int64
	CategoryName *string
	Source       string
	VideoURL     string
	RegencyID    *int64
	RegencyName  *string
	Status       int64
	CreatedAt    *time.Time
	UpdatedAt    *time.Time
	CreatedBy    *int64
	UpdatedBy    *int64
}

type VideoStatisticWithMetadata struct {
	Data     []*VideoStatistic
	Metadata *Metadata
}

type VideoStatisticUC struct {
	ID    int64
	Name  string
	Count int64
}
