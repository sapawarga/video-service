package testcases

import (
	"database/sql"
	"errors"

	"github.com/sapawarga/video-service/helper"
	"github.com/sapawarga/video-service/model"
)

var newVideoRequest = &model.CreateVideoRequest{
	Title:      helper.GenerateRandomString(12),
	Source:     "youtube",
	CategoryID: 1,
	RegencyID:  1,
	VideoURL:   helper.GenerateRandomString(10),
	Status:     10,
}

type CreateNewVideo struct {
	Description         string
	UsecaseRequest      *model.CreateVideoRequest
	GetLocationName     int64
	GetCategoryName     int64
	RepositoryRequest   *model.CreateVideoRequest
	MockGetLocationName ResponseGetLocationName
	MockGetCategoryName ResponseGetCategoryName
	MockRepository      error
	MockUsecase         error
}

var CreateNewVideoData = []CreateNewVideo{
	{
		Description:       "success_insert_video",
		UsecaseRequest:    newVideoRequest,
		GetLocationName:   1,
		GetCategoryName:   1,
		RepositoryRequest: newVideoRequest,
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
		UsecaseRequest:    newVideoRequest,
		GetLocationName:   1,
		GetCategoryName:   1,
		RepositoryRequest: newVideoRequest,
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
		UsecaseRequest:    newVideoRequest,
		GetLocationName:   1,
		GetCategoryName:   1,
		RepositoryRequest: newVideoRequest,
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
		Description:       "failed_insert_video",
		UsecaseRequest:    newVideoRequest,
		GetLocationName:   1,
		GetCategoryName:   1,
		RepositoryRequest: newVideoRequest,
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

func CreateNewVideoDescription() []string {
	var arr = []string{}
	for _, data := range CreateNewVideoData {
		arr = append(arr, data.Description)
	}
	return arr
}