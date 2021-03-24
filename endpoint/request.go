package endpoint

import (
	"errors"
	"regexp"

	validation "github.com/go-ozzo/ozzo-validation/v4"
)

type GetVideoRequest struct {
	RegencyID *int64 `json:"regency_id"`
	Page      *int64 `json:"page"`
}

type RequestID struct {
	ID int64 `httpquery:"id" json:"id"`
}

type CreateVideoRequest struct {
	Title      *string `json:"title"`
	Source     *string `json:"source"`
	CategoryID *int64  `json:"category_id"`
	RegencyID  *int64  `json:"regency_id"`
	VideoURL   *string `json:"video_url"`
	Status     *int64  `json:"status"`
}

type UpdateVideoRequest struct {
	ID         *int64  `json:"id"`
	Title      *string `json:"title"`
	Source     *string `json:"source"`
	CategoryID *int64  `json:"category_id"`
	RegencyID  *int64  `json:"regency_id"`
	VideoURL   *string `json:"video_url"`
	Status     *int64  `json:"status"`
}

func ValidateInputs(in interface{}) error {
	if obj, ok := in.(*CreateVideoRequest); ok {
		return validation.ValidateStruct(obj,
			validation.Field(obj.Title, validation.Required, validation.Length(10, 0)),
			validation.Field(obj.Source, validation.Required, validation.In("youtube")),
			validation.Field(obj.CategoryID, validation.Required),
			validation.Field(obj.VideoURL, validation.Required, validation.Match(regexp.MustCompile("/^(https://www.youtube.com)/.+$/"))),
			validation.Field(obj.Status, validation.Required, validation.In(-1, 0, 10)),
		)
	} else if obj, ok := in.(*UpdateVideoRequest); ok {
		return validation.ValidateStruct(obj,
			validation.Field(obj.ID, validation.Required),
			validation.Field(obj.Title, validation.Required, validation.Length(10, 0)),
			validation.Field(obj.Source, validation.Required, validation.In("youtube")),
			validation.Field(obj.CategoryID, validation.Required),
			validation.Field(obj.VideoURL, validation.Required, validation.Match(regexp.MustCompile("/^(https://www.youtube.com)/.+$/"))),
			validation.Field(obj.Status, validation.Required, validation.In(-1, 0, 10)),
		)
	}
	return errors.New("format_struct_not_valid")

}
