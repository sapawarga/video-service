package endpoint

import (
	"time"

	"github.com/sapawarga/video-service/model"
)

type VideoResponse struct {
	Data     []*model.VideoResponse `json:"data"`
	Metadata *model.Metadata        `json:"metadata"`
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

type VideoStatisticResponse struct {
	Data []*model.VideoStatisticUC `json:"data"`
}
