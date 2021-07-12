package model

// GetListVideoRequest ...
type GetListVideoRequest struct {
	RegencyID *int64
	Page      *int64
	Limit     *int64
}

// CreateVideoRequest ...
type CreateVideoRequest struct {
	Title      string
	Source     string
	CategoryID int64
	RegencyID  *int64
	VideoURL   string
	Status     int64
	Sequence   int64
}

// UpdateVideoRequest ...
type UpdateVideoRequest struct {
	ID         *int64
	Title      *string
	Source     *string
	CategoryID *int64
	RegencyID  *int64
	VideoURL   *string
	Status     *int64
	Sequence   *int64
}
