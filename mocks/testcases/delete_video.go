package testcases

import (
	"database/sql"
	"errors"
)

type DeleteVideo struct {
	Description             string
	UsecaseRequest          int64
	GetDetailVideoRequest   int64
	DeleteVideoRequest      int64
	MockVideoDetailResponse ResponseGetDetailVideo
	MockDeleteVideo         error
	MockUsecase             error
}

var DeleteVideoData = []DeleteVideo{
	{
		Description:           "delete_video_successful",
		UsecaseRequest:        1,
		GetDetailVideoRequest: 1,
		DeleteVideoRequest:    1,
		MockVideoDetailResponse: ResponseGetDetailVideo{
			Result: videoDetail,
			Error:  nil,
		},
		MockDeleteVideo: nil,
		MockUsecase:     nil,
	}, {
		Description:           "failed_get_data_detail",
		UsecaseRequest:        1,
		GetDetailVideoRequest: 1,
		DeleteVideoRequest:    1,
		MockVideoDetailResponse: ResponseGetDetailVideo{
			Result: nil,
			Error:  sql.ErrNoRows,
		},
		MockDeleteVideo: nil,
		MockUsecase:     sql.ErrNoRows,
	}, {
		Description:           "failed_delete_video",
		UsecaseRequest:        1,
		GetDetailVideoRequest: 1,
		DeleteVideoRequest:    1,
		MockVideoDetailResponse: ResponseGetDetailVideo{
			Result: videoDetail,
			Error:  nil,
		},
		MockDeleteVideo: errors.New("something_went_wrong"),
		MockUsecase:     errors.New("something_went_wrong"),
	},
}

func DeleteVideoDescription() []string {
	var arr = []string{}
	for _, data := range DeleteVideoData {
		arr = append(arr, data.Description)
	}
	return arr
}
