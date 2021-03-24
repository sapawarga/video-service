package endpoint

import (
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

func (cvr *CreateVideoRequest) ValidateInputs() error {
	return validation.ValidateStruct(cvr,
		validation.Field(cvr.Title, validation.Required, validation.Length(10, 0)),
		validation.Field(cvr.Source, validation.Required, validation.In("youtube")),
		validation.Field(cvr.CategoryID, validation.Required),
		validation.Field(cvr.VideoURL, validation.Required, validation.Match(regexp.MustCompile("/^(https://www.youtube.com)/.+$/"))),
		validation.Field(cvr.Status, validation.Required, validation.In(-1, 0, 10)),
	)
}
