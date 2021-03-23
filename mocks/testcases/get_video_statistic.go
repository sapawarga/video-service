package testcases

import (
	"database/sql"
	"errors"

	"github.com/sapawarga/video-service/helper"
	"github.com/sapawarga/video-service/model"
)

var randomString = []string{
	helper.GenerateRandomString(5), helper.GenerateRandomString(5),
}

var videoStatistic = []*model.VideoStatistic{
	{
		ID:    1,
		Name:  sql.NullString{String: randomString[0], Valid: true},
		Count: 2,
	}, {
		ID:    2,
		Name:  sql.NullString{String: randomString[1], Valid: true},
		Count: 4,
	},
}

var videoStatisticUC = []*model.VideoStatisticUC{
	{
		ID:    1,
		Name:  randomString[0],
		Count: 2,
	}, {
		ID:    2,
		Name:  randomString[1],
		Count: 4,
	},
}

type ResponseGetVideoStatisticRepo struct {
	Result []*model.VideoStatistic
	Error  error
}

type ResponseGetVideoStatisticUC struct {
	Result []*model.VideoStatisticUC
	Error  error
}

type GetVideoStatistic struct {
	Description         string
	MockResponseRepo    ResponseGetVideoStatisticRepo
	MockResponseUsecase ResponseGetVideoStatisticUC
}

var GetVideoStatisticData = []GetVideoStatistic{
	{
		Description: "success_get_video_statistic",
		MockResponseRepo: ResponseGetVideoStatisticRepo{
			Result: videoStatistic,
			Error:  nil,
		},
		MockResponseUsecase: ResponseGetVideoStatisticUC{
			Result: videoStatisticUC,
			Error:  nil,
		},
	}, {
		Description: "failed_get_video_statistic",
		MockResponseRepo: ResponseGetVideoStatisticRepo{
			Result: nil,
			Error:  errors.New("something_went_wrong"),
		},
		MockResponseUsecase: ResponseGetVideoStatisticUC{
			Result: nil,
			Error:  errors.New("something_went_wrong"),
		},
	},
}

func ListVideoStatistic() []string {
	var arr = []string{}
	for _, data := range GetVideoStatisticData {
		arr = append(arr, data.Description)
	}
	return arr
}
