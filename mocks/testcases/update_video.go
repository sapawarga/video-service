package testcases

import (
	"database/sql"
	"errors"

	"github.com/sapawarga/video-service/lib/converter"
	"github.com/sapawarga/video-service/lib/generator"
	"github.com/sapawarga/video-service/model"
)

var videoUpdate = &model.UpdateVideoRequest{
	ID:         converter.SetPointerInt64(1),
	Title:      converter.SetPointerString(generator.GenerateRandomString(10)),
	Source:     converter.SetPointerString("youtube"),
	CategoryID: converter.SetPointerInt64(1),
	RegencyID:  converter.SetPointerInt64(1),
	VideoURL:   converter.SetPointerString(generator.GenerateRandomString(10)),
	Status:     converter.SetPointerInt64(10),
}

type UpdateVideo struct {
	Description           string
	UsecaseRequest        *model.UpdateVideoRequest
	GetDetailVideoRequest int64
	GetLocationName       int64
	GetCategoryName       int64
	RepositoryRequest     *model.UpdateVideoRequest
	MockGetLocation       ResponseGetLocation
	MockGetCategoryName   ResponseGetCategoryName
	MockVideoDetail       ResponseGetDetailVideo
	MockRepository        error
	MockUsecase           error
}

var UpdateVideoData = []UpdateVideo{
	{
		Description:           "success_update_video",
		UsecaseRequest:        videoUpdate,
		GetLocationName:       1,
		GetCategoryName:       1,
		GetDetailVideoRequest: 1,
		RepositoryRequest:     videoUpdate,
		MockGetLocation: ResponseGetLocation{
			Result: location,
			Error:  nil,
		},
		MockGetCategoryName: ResponseGetCategoryName{
			Result: categoryName,
			Error:  nil,
		},
		MockVideoDetail: ResponseGetDetailVideo{
			Result: videoDetail,
			Error:  nil,
		},
		MockRepository: nil,
		MockUsecase:    nil,
	}, {
		Description:           "failed_get_category",
		UsecaseRequest:        videoUpdate,
		GetLocationName:       1,
		GetCategoryName:       1,
		GetDetailVideoRequest: 1,
		RepositoryRequest:     videoUpdate,
		MockGetLocation: ResponseGetLocation{
			Result: location,
			Error:  nil,
		},
		MockGetCategoryName: ResponseGetCategoryName{
			Result: nil,
			Error:  sql.ErrNoRows,
		},
		MockVideoDetail: ResponseGetDetailVideo{
			Result: videoDetail,
			Error:  nil,
		},
		MockRepository: nil,
		MockUsecase:    sql.ErrNoRows,
	}, {
		Description:           "failed_get_location",
		UsecaseRequest:        videoUpdate,
		GetLocationName:       1,
		GetCategoryName:       1,
		GetDetailVideoRequest: 1,
		RepositoryRequest:     videoUpdate,
		MockGetLocation: ResponseGetLocation{
			Result: nil,
			Error:  sql.ErrNoRows,
		},
		MockGetCategoryName: ResponseGetCategoryName{
			Result: categoryName,
			Error:  nil,
		},
		MockVideoDetail: ResponseGetDetailVideo{
			Result: videoDetail,
			Error:  nil,
		},
		MockRepository: nil,
		MockUsecase:    sql.ErrNoRows,
	}, {
		Description:           "failed_update_video",
		UsecaseRequest:        videoUpdate,
		GetLocationName:       1,
		GetCategoryName:       1,
		GetDetailVideoRequest: 1,
		RepositoryRequest:     videoUpdate,
		MockGetLocation: ResponseGetLocation{
			Result: location,
			Error:  nil,
		},
		MockGetCategoryName: ResponseGetCategoryName{
			Result: categoryName,
			Error:  nil,
		},
		MockVideoDetail: ResponseGetDetailVideo{
			Result: videoDetail,
			Error:  nil,
		},
		MockRepository: errors.New("something_went_wrong"),
		MockUsecase:    errors.New("something_went_wrong"),
	}, {
		Description:           "failed_get_video_detail",
		UsecaseRequest:        videoUpdate,
		GetLocationName:       1,
		GetCategoryName:       1,
		GetDetailVideoRequest: 1,
		RepositoryRequest:     videoUpdate,
		MockGetLocation: ResponseGetLocation{
			Result: location,
			Error:  nil,
		},
		MockGetCategoryName: ResponseGetCategoryName{
			Result: categoryName,
			Error:  nil,
		},
		MockVideoDetail: ResponseGetDetailVideo{
			Result: nil,
			Error:  sql.ErrNoRows,
		},
		MockRepository: nil,
		MockUsecase:    sql.ErrNoRows,
	},
}

func UpdateVideoDescription() []string {
	var arr = []string{}
	for _, data := range UpdateVideoData {
		arr = append(arr, data.Description)
	}
	return arr
}
