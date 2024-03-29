package testcases

import (
	"database/sql"
	"errors"

	"github.com/sapawarga/video-service/lib/converter"
	"github.com/sapawarga/video-service/lib/generator"
	"github.com/sapawarga/video-service/model"
)

var _, currentTime = generator.GetCurrentTimeUTC()
var category = &model.Category{
	ID:   1,
	Name: "category",
}
var videoResponses = []*model.VideoResponse{
	{
		ID:         1,
		Title:      "Test Video 1",
		CategoryID: 1,
		Source:     "youtube",
		VideoURL:   "https://youtube.com/UDOHE",
		RegencyID:  sql.NullInt64{Int64: 1, Valid: true},
		Status:     10,
		CreatedAt:  currentTime,
		UpdatedAt:  currentTime,
		CreatedBy:  1,
		UpdatedBy:  1,
	}, {
		ID:         2,
		Title:      "Test Video 2",
		CategoryID: 1,
		Source:     "youtube",
		VideoURL:   "https://youtube.com/UDOHE",
		RegencyID:  sql.NullInt64{Int64: 1, Valid: true},
		Status:     10,
		CreatedAt:  currentTime,
		UpdatedAt:  currentTime,
		CreatedBy:  1,
		UpdatedBy:  1,
	},
}

var videoUsecase = []*model.Video{
	{
		ID:        1,
		Title:     "Test Video 1",
		Category:  category,
		Source:    "youtube",
		VideoURL:  "https://youtube.com/UDOHE",
		Regency:   location,
		Status:    10,
		CreatedAt: currentTime,
		UpdatedAt: currentTime,
		CreatedBy: 1,
		UpdatedBy: 1,
	}, {
		ID:        2,
		Title:     "Test Video 2",
		Category:  category,
		Source:    "youtube",
		VideoURL:  "https://youtube.com/UDOHE",
		Regency:   location,
		Status:    10,
		CreatedAt: currentTime,
		UpdatedAt: currentTime,
		CreatedBy: 1,
		UpdatedBy: 1,
	},
}

type ResponseGetListVideo struct {
	Result []*model.VideoResponse
	Error  error
}

type ResponseMetadata struct {
	Result *int64
	Error  error
}

type ResponseUsecase struct {
	Result *model.VideoWithMetadata
	Error  error
}
type GetListVideo struct {
	Description             string
	UsecaseRequest          model.GetListVideoRequest
	GetListVideoRepoRequest model.GetListVideoRepoRequest
	GetLocationByID         int64
	MockGetLocationByID     ResponseGetLocation
	MockGetListVideoRepo    ResponseGetListVideo
	MockGetMetadata         ResponseMetadata
	MockGetCategoryName     ResponseGetCategoryName
	MockUsecaseResponse     ResponseUsecase
}

var regencyID = converter.SetPointerInt64(1)
var page = converter.SetPointerInt64(1)
var limit = converter.SetPointerInt64(10)
var offset = converter.SetPointerInt64(0)

var GetListVideoData = []GetListVideo{
	{
		Description: "success_get_list_video",
		UsecaseRequest: model.GetListVideoRequest{
			RegencyID: regencyID,
			Page:      page,
		},
		GetListVideoRepoRequest: model.GetListVideoRepoRequest{
			RegencyID: regencyID,
			Limit:     limit,
			Offset:    offset,
		},
		GetLocationByID: 1,
		MockGetLocationByID: ResponseGetLocation{
			Result: location,
			Error:  nil,
		},
		MockGetListVideoRepo: ResponseGetListVideo{
			Result: videoResponses,
			Error:  nil,
		},
		MockGetMetadata: ResponseMetadata{
			Result: converter.SetPointerInt64(int64(len(videoResponses))),
			Error:  nil,
		},
		MockGetCategoryName: ResponseGetCategoryName{
			Result: categoryName,
			Error:  nil,
		},
		MockUsecaseResponse: ResponseUsecase{
			Result: &model.VideoWithMetadata{
				Data: videoUsecase,
				Metadata: &model.Metadata{
					Page:      1,
					TotalPage: 1,
					Total:     int64(len(videoResponses)),
				},
			},
		},
	}, {
		Description: "failed_get_list_video",
		UsecaseRequest: model.GetListVideoRequest{
			RegencyID: regencyID,
			Page:      page,
		},
		MockGetCategoryName: ResponseGetCategoryName{
			Result: categoryName,
			Error:  nil,
		},
		GetLocationByID: 1,
		MockGetLocationByID: ResponseGetLocation{
			Result: location,
			Error:  nil,
		},
		GetListVideoRepoRequest: model.GetListVideoRepoRequest{
			RegencyID: regencyID,
			Limit:     limit,
			Offset:    offset,
		},
		MockGetListVideoRepo: ResponseGetListVideo{
			Result: nil,
			Error:  errors.New("failed_get_video"),
		},
		MockGetMetadata: ResponseMetadata{
			Result: nil,
			Error:  errors.New("failed_get_metadata"),
		},
		MockUsecaseResponse: ResponseUsecase{
			Result: nil,
			Error:  errors.New("failed_get_video"),
		},
	}, {
		Description: "failed_get_metadata",
		UsecaseRequest: model.GetListVideoRequest{
			RegencyID: regencyID,
			Page:      page,
		},
		GetListVideoRepoRequest: model.GetListVideoRepoRequest{
			RegencyID: regencyID,
			Limit:     limit,
			Offset:    offset,
		},
		GetLocationByID: 1,
		MockGetLocationByID: ResponseGetLocation{
			Result: location,
			Error:  nil,
		},
		MockGetCategoryName: ResponseGetCategoryName{
			Result: categoryName,
			Error:  nil,
		},
		MockGetListVideoRepo: ResponseGetListVideo{
			Result: videoResponses,
			Error:  nil,
		},
		MockGetMetadata: ResponseMetadata{
			Result: nil,
			Error:  errors.New("failed_get_metadata"),
		},
		MockUsecaseResponse: ResponseUsecase{
			Result: nil,
			Error:  errors.New("failed_get_metadata"),
		},
	}, {
		Description: "failed_get_category",
		UsecaseRequest: model.GetListVideoRequest{
			RegencyID: regencyID,
			Page:      page,
		},
		GetLocationByID: 1,
		MockGetLocationByID: ResponseGetLocation{
			Result: location,
			Error:  nil,
		},
		GetListVideoRepoRequest: model.GetListVideoRepoRequest{
			RegencyID: regencyID,
			Limit:     limit,
			Offset:    offset,
		},
		MockGetCategoryName: ResponseGetCategoryName{
			Result: nil,
			Error:  errors.New("something went wrong"),
		},
		MockGetListVideoRepo: ResponseGetListVideo{
			Result: videoResponses,
			Error:  nil,
		},
		MockGetMetadata: ResponseMetadata{
			Result: nil,
			Error:  nil,
		},
		MockUsecaseResponse: ResponseUsecase{
			Result: nil,
			Error:  errors.New("something went wrong"),
		},
	}, {
		Description: "failed_get_locatioan",
		UsecaseRequest: model.GetListVideoRequest{
			RegencyID: regencyID,
			Page:      page,
		},
		GetLocationByID: 1,
		MockGetLocationByID: ResponseGetLocation{
			Result: nil,
			Error:  errors.New("something_went_wrong"),
		},
		GetListVideoRepoRequest: model.GetListVideoRepoRequest{
			RegencyID: regencyID,
			Limit:     limit,
			Offset:    offset,
		},
		MockGetCategoryName: ResponseGetCategoryName{
			Result: nil,
			Error:  nil,
		},
		MockGetListVideoRepo: ResponseGetListVideo{
			Result: videoResponses,
			Error:  nil,
		},
		MockGetMetadata: ResponseMetadata{
			Result: nil,
			Error:  nil,
		},
		MockUsecaseResponse: ResponseUsecase{
			Result: nil,
			Error:  errors.New("something went wrong"),
		},
	},
}

func ListVideoDescription() []string {
	var arr = []string{}
	for _, data := range GetListVideoData {
		arr = append(arr, data.Description)
	}
	return arr
}
