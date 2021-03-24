package testcases

import (
	"database/sql"
	"errors"

	"github.com/sapawarga/video-service/helper"
	"github.com/sapawarga/video-service/model"
)

var videoUpdate = &model.UpdateVideoRequest{
	ID:         1,
	Title:      helper.GenerateRandomString(10),
	Source:     "youtube",
	CategoryID: 1,
	RegencyID:  1,
	VideoURL:   helper.GenerateRandomString(10),
	Status:     10,
}

type UpdateVideo struct {
	Description         string
	UsecaseRequest      *model.UpdateVideoRequest
	GetLocationName     int64
	GetCategoryName     int64
	RepositoryRequest   *model.UpdateVideoRequest
	MockGetLocationName ResponseGetLocationName
	MockGetCategoryName ResponseGetCategoryName
	MockRepository      error
	MockUsecase         error
}

var UpdateVideoData = []UpdateVideo{
	{
		Description:       "success_update_video",
		UsecaseRequest:    videoUpdate,
		GetLocationName:   1,
		GetCategoryName:   1,
		RepositoryRequest: videoUpdate,
		MockGetLocationName: ResponseGetLocationName{
			Result: location,
			Error:  nil,
		},
		MockGetCategoryName: ResponseGetCategoryName{
			Result: category,
			Error:  nil,
		},
		MockRepository: nil,
		MockUsecase:    nil,
	}, {
		Description:       "failed_get_category",
		UsecaseRequest:    videoUpdate,
		GetLocationName:   1,
		GetCategoryName:   1,
		RepositoryRequest: videoUpdate,
		MockGetLocationName: ResponseGetLocationName{
			Result: location,
			Error:  nil,
		},
		MockGetCategoryName: ResponseGetCategoryName{
			Result: nil,
			Error:  sql.ErrNoRows,
		},
		MockRepository: nil,
		MockUsecase:    sql.ErrNoRows,
	}, {
		Description:       "failed_get_location",
		UsecaseRequest:    videoUpdate,
		GetLocationName:   1,
		GetCategoryName:   1,
		RepositoryRequest: videoUpdate,
		MockGetLocationName: ResponseGetLocationName{
			Result: nil,
			Error:  sql.ErrNoRows,
		},
		MockGetCategoryName: ResponseGetCategoryName{
			Result: category,
			Error:  nil,
		},
		MockRepository: nil,
		MockUsecase:    sql.ErrNoRows,
	}, {
		Description:       "failed_update_video",
		UsecaseRequest:    videoUpdate,
		GetLocationName:   1,
		GetCategoryName:   1,
		RepositoryRequest: videoUpdate,
		MockGetLocationName: ResponseGetLocationName{
			Result: location,
			Error:  nil,
		},
		MockGetCategoryName: ResponseGetCategoryName{
			Result: category,
			Error:  nil,
		},
		MockRepository: errors.New("something_went_wrong"),
		MockUsecase:    errors.New("something_went_wrong"),
	},
}

func UpdateVideoDescription() []string {
	var arr = []string{}
	for _, data := range UpdateVideoData {
		arr = append(arr, data.Description)
	}
	return arr
}
