package endpoint

import (
	"time"

	"github.com/sapawarga/video-service/model"
)

type VideoResponse struct {
	Data     []*model.Video  `json:"items"`
	Metadata *model.Metadata `json:"_meta"`
}

type VideoDetail struct {
	ID                 int64           `json:"id"`
	Title              string          `json:"title"`
	Cateogry           *model.Category `json:"category"`
	Source             string          `json:"source"`
	VideoURL           string          `json:"video_url"`
	TotalLikes         int64           `json:"total_likes,omitempty"`
	IsPushNotification int64           `json:"is_push_notification"`
	RegencyID          *int64          `json:"kabkota_id,omitempty"`
	RegencyName        *string         `json:"kabkota,omitempty"`
	Status             int64           `json:"status"`
	CreatedAt          *time.Time      `json:"created_at"`
	UpdatedAt          *time.Time      `json:"updated_at"`
	CreatedBy          *int64          `json:"created_by"`
	UpdatedBy          *int64          `json:"updated_by"`
}

// VideoWithMeta ...
type VideoWithMeta struct {
	Data *VideoResponse `json:"data"`
}

type VideoStatisticResponse struct {
	Data []*model.VideoStatisticUC `json:"data"`
}

type StatusResponse struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}
