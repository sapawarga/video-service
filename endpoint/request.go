package endpoint

import (
	"errors"
	"regexp"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/sapawarga/video-service/lib/constants"
)

type GetVideoRequest struct {
	Search     *string `json:"search"`
	RegencyID  *int64  `json:"regency_id"`
	Page       *int64  `json:"page"`
	Limit      *int64  `json:"limit"`
	Title      *string `json:"title"`
	CategoryID *int64  `json:"category_id"`
	SortBy     string  `json:"sort_by"`
	SortOrder  string  `json:"sort_order"`
}

type RequestID struct {
	ID int64 `httpquery:"id" json:"id"`
}

type CreateVideoRequest struct {
	Title              string `json:"title"`
	Source             string `json:"source"`
	CategoryID         int64  `json:"category_id"`
	RegencyID          *int64 `json:"kabkota_id"`
	VideoURL           string `json:"video_url"`
	Status             int64  `json:"status"`
	Sequence           int64  `json:"seq"`
	IsPushNotification bool   `json:"is_push_notification"`
}

type UpdateVideoRequest struct {
	ID                 *int64  `json:"id"`
	Title              *string `json:"title"`
	Source             *string `json:"source"`
	CategoryID         *int64  `json:"category_id"`
	RegencyID          *int64  `json:"kabkota_id"`
	VideoURL           *string `json:"video_url"`
	Status             *int64  `json:"status"`
	Sequence           *int64  `json:"seq"`
	IsPushNotification *bool   `json:"is_push_notification"`
}

func ValidateInputs(in interface{}) error {
	if obj, ok := in.(*CreateVideoRequest); ok {
		return validation.ValidateStruct(obj,
			validation.Field(&obj.Title, validation.Required, validation.Length(10, 0)),
			validation.Field(&obj.Source, validation.Required, validation.In("youtube")),
			validation.Field(&obj.CategoryID, validation.Required),
			validation.Field(&obj.VideoURL, validation.Required, validation.Match(regexp.MustCompile("^(https://www.youtube.com/).+$"))),
		)
	} else if obj, ok := in.(*UpdateVideoRequest); ok {
		return validation.ValidateStruct(obj,
			validation.Field(&obj.ID, validation.Required),
			validation.Field(&obj.Title, validation.Length(10, 0)),
			validation.Field(&obj.Source, validation.In("youtube")),
			validation.Field(&obj.CategoryID),
			validation.Field(&obj.VideoURL, validation.Match(regexp.MustCompile("^(https://www.youtube.com/).+$"))),
			validation.Field(&obj.Status, validation.In(constants.DELETED, constants.INACTIVED, constants.ACTIVED)),
		)
	}
	return errors.New("format_struct_not_valid")

}

var Status = []int64{-1, 0, 10}

func containsStatus(status []int64, val int64) bool {
	for _, v := range status {
		if v == val {
			return true
		}
	}
	return false
}
