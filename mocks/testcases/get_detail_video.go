package testcases

import (
	"database/sql"

	"github.com/sapawarga/video-service/helper"
	"github.com/sapawarga/video-service/model"
)

var videoDetail = &model.VideoResponse{
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
}

type ResponseGetDetailVideo struct {
	Result *model.VideoResponse
	Error  error
}

type ResponseGetCategoryName struct {
	Result *string
	Error  error
}

type ResponseGetLocation struct {
	Result *model.Location
	Error  error
}

type ResponseUsecaseGetDetail struct {
	Result *model.VideoDetail
	Error  error
}

type GetDetailVideo struct {
	Description           string
	UsecaseRequest        int64
	GetVideoDetailRequest int64
	GetLocationName       int64
	GetCategoryName       int64
	MockGetDetailRepo     ResponseGetDetailVideo
	MockGetLocation       ResponseGetLocation
	MockGetCategoryName   ResponseGetCategoryName
	MockUsecaseResponse   ResponseUsecaseGetDetail
}

var location = &model.Location{
	ID:      1,
	BPSCode: "code",
	Name:    "location",
}
var categoryName = helper.SetPointerString("category")

var GetDetailVideoData = []GetDetailVideo{
	{
		Description:           "success_get_video_detail",
		UsecaseRequest:        1,
		GetVideoDetailRequest: 1,
		GetLocationName:       1,
		GetCategoryName:       1,
		MockGetDetailRepo: ResponseGetDetailVideo{
			Result: videoDetail,
			Error:  nil,
		},
		MockGetLocation: ResponseGetLocation{
			Result: location,
			Error:  nil,
		},
		MockGetCategoryName: ResponseGetCategoryName{
			Result: categoryName,
			Error:  nil,
		},
		MockUsecaseResponse: ResponseUsecaseGetDetail{
			Result: &model.VideoDetail{
				ID:        videoDetail.ID,
				Title:     videoDetail.Title.String,
				Category:  category,
				Source:    videoDetail.Source.String,
				VideoURL:  videoDetail.VideoURL.String,
				Regency:   location,
				Status:    videoDetail.Status.Int64,
				CreatedAt: helper.SetPointerInt64(videoDetail.CreatedAt.Int64),
				UpdatedAt: helper.SetPointerInt64(videoDetail.UpdatedAt.Int64),
				CreatedBy: helper.SetPointerInt64(videoDetail.CreatedBy.Int64),
				UpdatedBy: helper.SetPointerInt64(videoDetail.UpdatedBy.Int64),
			},
			Error: nil,
		},
	}, {
		Description:           "failed_get_video_detail",
		UsecaseRequest:        1,
		GetVideoDetailRequest: 1,
		GetLocationName:       1,
		GetCategoryName:       1,
		MockGetDetailRepo: ResponseGetDetailVideo{
			Result: nil,
			Error:  sql.ErrNoRows,
		},
		MockGetLocation: ResponseGetLocation{
			Result: location,
			Error:  nil,
		},
		MockGetCategoryName: ResponseGetCategoryName{
			Result: categoryName,
			Error:  nil,
		},
		MockUsecaseResponse: ResponseUsecaseGetDetail{
			Result: nil,
			Error:  sql.ErrNoRows,
		},
	}, {
		Description:           "failed_get_location_name",
		UsecaseRequest:        1,
		GetVideoDetailRequest: 1,
		GetLocationName:       1,
		GetCategoryName:       1,
		MockGetDetailRepo: ResponseGetDetailVideo{
			Result: videoDetail,
			Error:  nil,
		},
		MockGetLocation: ResponseGetLocation{
			Result: nil,
			Error:  sql.ErrNoRows,
		},
		MockGetCategoryName: ResponseGetCategoryName{
			Result: categoryName,
			Error:  nil,
		},
		MockUsecaseResponse: ResponseUsecaseGetDetail{
			Result: nil,
			Error:  sql.ErrNoRows,
		},
	}, {
		Description:           "failed_get_category_name",
		UsecaseRequest:        1,
		GetVideoDetailRequest: 1,
		GetLocationName:       1,
		GetCategoryName:       1,
		MockGetDetailRepo: ResponseGetDetailVideo{
			Result: videoDetail,
			Error:  nil,
		},
		MockGetLocation: ResponseGetLocation{
			Result: location,
			Error:  nil,
		},
		MockGetCategoryName: ResponseGetCategoryName{
			Result: nil,
			Error:  sql.ErrNoRows,
		},
		MockUsecaseResponse: ResponseUsecaseGetDetail{
			Result: nil,
			Error:  sql.ErrNoRows,
		},
	},
}

func DetailVideoDescription() []string {
	var arr = []string{}
	for _, data := range GetDetailVideoData {
		arr = append(arr, data.Description)
	}
	return arr
}
