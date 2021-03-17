package endpoint

type GetVideoRequest struct {
	RegencyID *int64 `json:"regency_id"`
	Page      *int64 `json:"page"`
}

type RequestID struct {
	ID int64 `httpquery:"id" json:"id"`
}
