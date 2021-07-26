package endpoint

import (
	"github.com/sapawarga/video-service/lib/converter"
	"github.com/sapawarga/video-service/model"
)

// VideoResponse ...
type VideoResponse struct {
	Data     []*Video        `json:"items"`
	Metadata *model.Metadata `json:"_meta"`
}

// VideoDetail ...
type VideoDetail struct {
	ID                 int64           `json:"id"`
	Title              string          `json:"title"`
	CategoryID         int64           `json:"category_id"`
	Category           *model.Category `json:"category"`
	Source             string          `json:"source"`
	VideoURL           string          `json:"video_url"`
	TotalLikes         int64           `json:"total_likes"`
	IsPushNotification bool            `json:"is_push_notification"`
	RegencyID          *int64          `json:"kabkota_id"`
	Regency            *model.Location `json:"kabkota"`
	Sequence           int64           `json:"seq"`
	Status             int64           `json:"status"`
	StatusLabel        string          `json:"status_label"`
	CreatedAt          *int64          `json:"created_at"`
	UpdatedAt          *int64          `json:"updated_at"`
	CreatedBy          *int64          `json:"created_by"`
	UpdatedBy          *int64          `json:"updated_by"`
}

// VideoStatisticResponse ...
type VideoStatisticResponse struct {
	Data     []*model.VideoStatisticUC `json:"items"`
	Metadata *model.Metadata           `json:"_meta"`
}

// StatusResponse ...
type StatusResponse struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

// Video ...
type Video struct {
	ID                 int64           `json:"id"`
	Title              string          `json:"title"`
	CategoryID         *int64          `json:"category_id"`
	Category           *model.Category `json:"category"`
	Source             string          `json:"source"`
	VideoURL           string          `json:"video_url"`
	RegencyID          *int64          `json:"kabkota_id"`
	Regency            *Location       `json:"kabkota"`
	TotalLike          int64           `json:"total_likes"`
	IsPushNotification bool            `json:"is_push_notification"`
	Sequence           int64           `json:"seq"`
	Status             int64           `json:"status"`
	StatusLabel        string          `json:"status_label"`
	CreatedAt          int64           `json:"created_at"`
	UpdatedAt          int64           `json:"updated_at"`
	CreatedBy          int64           `json:"created_by"`
}

// Location ...
type Location struct {
	ID      int64  `json:"id"`
	CodeBPS string `json:"code_bps"`
	Name    string `json:"name"`
}

func encodeResponse(data []*model.Video) (result []*Video) {
	if len(data) > 0 {
		for _, v := range data {
			encodeData := &Video{
				ID:                 v.ID,
				Title:              v.Title,
				CategoryID:         converter.SetPointerInt64(v.Category.ID),
				Category:           v.Category,
				Source:             v.Source,
				VideoURL:           v.VideoURL,
				TotalLike:          v.TotalLikes,
				IsPushNotification: v.IsPushNotification,
				Sequence:           v.Sequence,
				Status:             v.Status,
				StatusLabel:        GetStatusLabel[v.Status]["id"],
				CreatedAt:          v.CreatedAt,
				UpdatedAt:          v.UpdatedAt,
				CreatedBy:          v.CreatedBy,
			}
			if v.Regency != nil {
				location := &Location{
					ID:      v.Regency.ID,
					CodeBPS: v.Regency.BPSCode.String,
					Name:    v.Regency.Name.String,
				}
				encodeData.RegencyID = converter.SetPointerInt64(v.Regency.ID)
				encodeData.Regency = location
			}

			result = append(result, encodeData)
		}
	}
	return result
}

// GetStatusLabel ...
var GetStatusLabel = map[int64]map[string]string{
	-1: {"en": "status deleted", "id": "Dihapus"},
	0:  {"en": "Not Active", "id": "Tidak Aktif"},
	10: {"en": "Active", "id": "Aktif"},
}
