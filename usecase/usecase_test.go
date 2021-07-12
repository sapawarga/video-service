package usecase_test

import (
	"context"
	"fmt"
	"os"
	"reflect"

	"github.com/golang/mock/gomock"
	mock_repository "github.com/sapawarga/video-service/mocks"
	"github.com/sapawarga/video-service/mocks/testcases"
	"github.com/sapawarga/video-service/usecase"

	kitlog "github.com/go-kit/kit/log"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/ginkgo/extensions/table"
	. "github.com/onsi/gomega"
)

var _ = Describe("Usecase", func() {
	var (
		mockVideoRepo *mock_repository.MockDatabaseI
		video         usecase.UsecaseI
	)

	BeforeEach(func() {
		logger := kitlog.NewLogfmtLogger(kitlog.NewSyncWriter(os.Stderr))
		mockSvc := gomock.NewController(GinkgoT())
		mockSvc.Finish()
		mockVideoRepo = mock_repository.NewMockDatabaseI(mockSvc)
		video = usecase.NewVideo(mockVideoRepo, logger)
	})

	// DECLARE UNIT TEST FUNCTION

	var GetListVideoLogic = func(idx int) {
		ctx := context.Background()
		data := testcases.GetListVideoData[idx]
		mockVideoRepo.EXPECT().GetListVideo(ctx, gomock.Any()).Return(data.MockGetListVideoRepo.Result, data.MockGetListVideoRepo.Error).Times(1)
		mockVideoRepo.EXPECT().GetMetadataVideo(ctx, gomock.Any()).Return(data.MockGetMetadata.Result, data.MockGetMetadata.Error).Times(1)
		mockVideoRepo.EXPECT().GetCategoryNameByID(ctx, gomock.Any()).Return(data.MockGetCategoryName.Result, data.MockGetCategoryName.Error).Times(len(data.MockGetListVideoRepo.Result))
		resp, err := video.GetListVideo(ctx, &data.UsecaseRequest)
		if err != nil {
			Expect(err).NotTo(BeNil())
			Expect(resp).To(BeNil())
		} else {
			Expect(err).To(BeNil())
		}
	}

	var GetDetailVideoLogic = func(idx int) {
		ctx := context.Background()
		data := testcases.GetDetailVideoData[idx]
		mockVideoRepo.EXPECT().GetDetailVideo(ctx, data.GetVideoDetailRequest).Return(data.MockGetDetailRepo.Result, data.MockGetDetailRepo.Error).Times(1)
		mockVideoRepo.EXPECT().GetCategoryNameByID(ctx, data.GetCategoryName).Return(data.MockGetCategoryName.Result, data.MockGetCategoryName.Error).Times(1)
		mockVideoRepo.EXPECT().GetLocationByID(ctx, data.GetLocationName).Return(data.MockGetLocation.Result, data.MockGetLocation.Error).Times(1)
		resp, err := video.GetDetailVideo(ctx, data.UsecaseRequest)
		if err != nil {
			Expect(err).NotTo(BeNil())
			Expect(resp).To(BeNil())
		} else {
			Expect(err).To(BeNil())
		}
	}

	var GetVideoStatisticLogic = func(idx int) {
		ctx := context.Background()
		data := testcases.GetVideoStatisticData[idx]
		mockVideoRepo.EXPECT().GetVideoStatistic(ctx).Return(data.MockResponseRepo.Result, data.MockResponseRepo.Error).Times(1)
		resp, err := video.GetStatisticVideo(ctx)
		if err != nil {
			Expect(err).NotTo(BeNil())
			Expect(resp).To(BeNil())
		} else {
			Expect(err).To(BeNil())
		}
	}

	var CreateNewVideoLogic = func(idx int) {
		ctx := context.Background()
		data := testcases.CreateNewVideoData[idx]
		mockVideoRepo.EXPECT().GetCategoryNameByID(ctx, data.GetCategoryName).Return(data.MockGetCategoryName.Result, data.MockGetCategoryName.Error).Times(1)
		mockVideoRepo.EXPECT().GetLocationByID(ctx, data.GetLocationName).Return(data.MockGetLocation.Result, data.MockGetLocation.Error).Times(1)
		mockVideoRepo.EXPECT().Insert(ctx, data.RepositoryRequest).Return(data.MockRepository).Times(1)
		if err := video.CreateNewVideo(ctx, data.UsecaseRequest); err != nil {
			Expect(err).NotTo(BeNil())
		} else {
			Expect(err).To(BeNil())
		}
	}

	var UpdateVideoLogic = func(idx int) {
		ctx := context.Background()
		data := testcases.UpdateVideoData[idx]
		mockVideoRepo.EXPECT().GetCategoryNameByID(ctx, data.GetCategoryName).Return(data.MockGetCategoryName.Result, data.MockGetCategoryName.Error).Times(1)
		mockVideoRepo.EXPECT().GetLocationByID(ctx, data.GetLocationName).Return(data.MockGetLocation.Result, data.MockGetLocation.Error).Times(1)
		mockVideoRepo.EXPECT().GetDetailVideo(ctx, data.GetDetailVideoRequest).Return(data.MockVideoDetail.Result, data.MockVideoDetail.Error).Times(1)
		mockVideoRepo.EXPECT().Update(ctx, data.RepositoryRequest).Return(data.MockRepository).Times(1)
		if err := video.UpdateVideo(ctx, data.UsecaseRequest); err != nil {
			Expect(err).NotTo(BeNil())
		} else {
			Expect(err).To(BeNil())
		}
	}

	var DeleteVideoLogic = func(idx int) {
		ctx := context.Background()
		data := testcases.DeleteVideoData[idx]
		mockVideoRepo.EXPECT().GetDetailVideo(ctx, data.GetDetailVideoRequest).Return(data.MockVideoDetailResponse.Result, data.MockVideoDetailResponse.Error).Times(1)
		mockVideoRepo.EXPECT().Delete(ctx, data.DeleteVideoRequest).Return(data.MockDeleteVideo).Times(1)
		if err := video.DeleteVideo(ctx, data.UsecaseRequest); err != nil {
			Expect(err).NotTo(BeNil())
		} else {
			Expect(err).To(BeNil())
		}
	}

	var CheckReadinessLogic = func(idx int) {
		ctx := context.Background()
		data := testcases.CheckReadinessData[idx]
		mockVideoRepo.EXPECT().HealthCheckReadiness(ctx).Return(data.MockCheckReadiness).Times(1)
		if err := video.CheckHealthReadiness(ctx); err != nil {
			Expect(err).NotTo(BeNil())
		} else {
			Expect(err).To(BeNil())
		}
	}

	var unitTestLogic = map[string]map[string]interface{}{
		"GetListVideo":      {"func": GetListVideoLogic, "test_case_count": len(testcases.GetListVideoData), "desc": testcases.ListVideoDescription()},
		"GetDetailVideo":    {"func": GetDetailVideoLogic, "test_case_count": len(testcases.GetDetailVideoData), "desc": testcases.DetailVideoDescription()},
		"GetStatisticVideo": {"func": GetVideoStatisticLogic, "test_case_count": len(testcases.GetVideoStatisticData), "desc": testcases.ListVideoStatistic()},
		"CreateNewVideo":    {"func": CreateNewVideoLogic, "test_case_count": len(testcases.CreateNewVideoData), "desc": testcases.CreateNewVideoDescription()},
		"UpdateVideo":       {"func": UpdateVideoLogic, "test_case_count": len(testcases.CreateNewVideoData), "desc": testcases.UpdateVideoDescription()},
		"DeleteVideo":       {"func": DeleteVideoLogic, "test_case_count": len(testcases.DeleteVideoData), "desc": testcases.DeleteVideoDescription()},
		"CheckReadiness":    {"func": CheckReadinessLogic, "test_case_count": len(testcases.CheckReadinessData), "desc": testcases.CheckReadinessDescription()},
	}

	for _, val := range unitTestLogic {
		s := reflect.ValueOf(val["desc"])
		var arr []TableEntry
		for i := 0; i < val["test_case_count"].(int); i++ {
			fmt.Println(s.Index(i).String())
			arr = append(arr, Entry(s.Index(i).String(), i))
		}
		DescribeTable("Function ", val["func"], arr...)
	}
})
