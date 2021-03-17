package endpoint

import "github.com/sapawarga/video-service/model"

type VideoResponse struct {
	Data     []*model.VideoResponse `json:"data"`
	Metadata *model.Metadata        `json:"metadata"`
}
