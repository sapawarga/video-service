package model

type GetListVideoRequest struct {
	RegencyID *int64
	Page      *int64
}

type CreateVideoRequest struct {
	Title      string
	Source     string
	CategoryID int64
	RegencyID  int64
	VideoURL   string
	Status     int64
}
