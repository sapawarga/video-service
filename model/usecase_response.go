package model

type VideoWithMetadata struct {
	Data     []*VideoResponse
	Metadata *Metadata
}

type Metadata struct {
	Page      int64
	TotalPage int64
	Total     int64
}
