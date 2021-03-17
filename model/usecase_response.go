package model

import "time"

type VideoWithMetadata struct {
	Data     []*VideoResponse
	Metadata *Metadata
}

type Metadata struct {
	Page      int64
	TotalPage int64
	Total     int64
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
