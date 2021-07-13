package model

type VideoWithMetadata struct {
	Data     []*Video  `json:"items"`
	Metadata *Metadata `json:"_meta"`
}

type Metadata struct {
	Page        int64   `json:"perPage"`
	TotalPage   float64 `json:"pageCount"`
	CurrentPage int64   `json:"currentPage"`
	Total       int64   `json:"totalCount"`
}

type Video struct {
	ID                 int64     `json:"id"`
	Title              string    `json:"title"`
	Category           *Category `json:"category"`
	Source             string    `json:"source"`
	VideoURL           string    `json:"video_url"`
	Regency            *Location `json:"kabkota"`
	IsPushNotification bool      `json:"is_push_notification"`
	TotalLikes         int64     `json:"total_likes"`
	Status             int64     `json:"status"`
	Sequence           int64     `json:"seq"`
	CreatedAt          int64     `json:"created_at"`
	UpdatedAt          int64     `json:"updated_at"`
	CreatedBy          int64     `json:"created_by"`
	UpdatedBy          int64     `json:"updated_by"`
}

type Category struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

// VideoDetail ...
type VideoDetail struct {
	ID                 int64
	Title              string
	Category           *Category
	Source             string
	VideoURL           string
	Regency            *Location
	Status             int64
	Sequence           int64
	CreatedAt          *int64
	UpdatedAt          *int64
	CreatedBy          *int64
	UpdatedBy          *int64
	TotalLikes         int64
	IsPushNotification bool
}

// VideoStatisticWithMetadata ...
type VideoStatisticWithMetadata struct {
	Data     []*VideoStatistic
	Metadata *Metadata
}

// VideoStatisticUC ...
type VideoStatisticUC struct {
	ID    int64  `json:"id"`
	Name  string `json:"name"`
	Count int64  `json:"count"`
}

// BoolFromInt ...
var BoolFromInt = map[int64]bool{
	0: false,
	1: true,
}
