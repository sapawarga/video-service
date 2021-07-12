package testcases

import (
	"database/sql"
	"errors"

	"github.com/sapawarga/video-service/helper"
	"github.com/sapawarga/video-service/model"
)

var _, currentTime = helper.GetCurrentTimeUTC()
var category = &model.Category{
	ID:   1,
	Name: "category",
}
var videoResponses = []*model.VideoResponse{
	{
		ID:         1,
		Title:      sql.NullString{String: "Test Video 1", Valid: true},
		CategoryID: sql.NullInt64{Int64: 1, Valid: true},
		Source:     sql.NullString{String: "youtube", Valid: true},
		VideoURL:   sql.NullString{String: "https://youtube.com/UDOHE", Valid: true},
		RegencyID:  sql.NullInt64{Int64: 1, Valid: true},
		Status:     sql.NullInt64{Int64: 10, Valid: true},
		CreatedAt:  sql.NullInt64{Int64: currentTime, Valid: true},
		UpdatedAt:  sql.NullInt64{Int64: currentTime, Valid: true},
		CreatedBy:  sql.NullInt64{Int64: 1, Valid: true},
		UpdatedBy:  sql.NullInt64{Int64: 1, Valid: true},
	}, {
		ID:         2,
		Title:      sql.NullString{String: "Test Video 2", Valid: true},
		CategoryID: sql.NullInt64{Int64: 1, Valid: true},
		Source:     sql.NullString{String: "youtube", Valid: true},
		VideoURL:   sql.NullString{String: "https://youtube.com/UDOHE", Valid: true},
		RegencyID:  sql.NullInt64{Int64: 1, Valid: true},
		Status:     sql.NullInt64{Int64: 10, Valid: true},
		CreatedAt:  sql.NullInt64{Int64: currentTime, Valid: true},
		UpdatedAt:  sql.NullInt64{Int64: currentTime, Valid: true},
		CreatedBy:  sql.NullInt64{Int64: 1, Valid: true},
		UpdatedBy:  sql.NullInt64{Int64: 1, Valid: true},
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

var regencyID = helper.SetPointerInt64(1)
var page = helper.SetPointerInt64(1)
var limit = helper.SetPointerInt64(10)
var offset = helper.SetPointerInt64(0)

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
			Result: helper.SetPointerInt64(int64(len(videoResponses))),
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
