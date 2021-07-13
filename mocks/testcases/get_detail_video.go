package testcases

import (
	"database/sql"

	"github.com/sapawarga/video-service/lib/converter"
	"github.com/sapawarga/video-service/model"
)

var videoDetail = &model.VideoResponse{
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
	BPSCode: sql.NullString{String: "code", Valid: true},
	Name:    sql.NullString{String: "location", Valid: true},
}
var categoryName = converter.SetPointerString("category")

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
				Title:     videoDetail.Title,
				Category:  category,
				Source:    videoDetail.Source,
				VideoURL:  videoDetail.VideoURL,
				Regency:   location,
				Status:    videoDetail.Status,
				CreatedAt: converter.SetPointerInt64(videoDetail.CreatedAt),
				UpdatedAt: converter.SetPointerInt64(videoDetail.UpdatedAt),
				CreatedBy: converter.SetPointerInt64(videoDetail.CreatedBy),
				UpdatedBy: converter.SetPointerInt64(videoDetail.UpdatedBy),
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
