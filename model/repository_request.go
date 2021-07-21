package model

type GetListVideoRepoRequest struct {
	RegencyID  *int64
	Limit      *int64
	Offset     *int64
	CategoryID *int64
	Title      *string
}
