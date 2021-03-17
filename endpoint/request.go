package endpoint

type GetVideoRequest struct {
	RegencyID *int64 `json:"regency_id"`
	Page      *int64 `json:"page"`
}
