package model

// GetListVideoRequest ...
type GetListVideoRequest struct {
	Search     *string
	RegencyID  *int64
	Page       *int64
	Limit      *int64
	CategoryID *int64
	Title      *string
	SortBy     string
	SortOrder  string
}

// CreateVideoRequest ...
type CreateVideoRequest struct {
	Title              string
	Source             string
	CategoryID         int64
	RegencyID          *int64
	VideoURL           string
	Status             int64
	Sequence           int64
	IsPushNotification bool
}

// UpdateVideoRequest ...
type UpdateVideoRequest struct {
	ID                 *int64
	Title              *string
	Source             *string
	CategoryID         *int64
	RegencyID          *int64
	VideoURL           *string
	Status             *int64
	Sequence           *int64
	IsPushNotification *bool
}
