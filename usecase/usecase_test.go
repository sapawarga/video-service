package usecase_test

import (
	"context"
	"fmt"
	"os"
	"reflect"

	"github.com/golang/mock/gomock"
	"github.com/sapawarga/video-service/mocks"
	"github.com/sapawarga/video-service/mocks/testcases"
	"github.com/sapawarga/video-service/usecase"

	kitlog "github.com/go-kit/kit/log"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/ginkgo/extensions/table"
	. "github.com/onsi/gomega"
)

var _ = Describe("Usecase", func() {
	var (
		mockVideoRepo *mocks.MockDatabaseI
		video         usecase.UsecaseI
	)

	BeforeEach(func() {
		logger := kitlog.NewLogfmtLogger(kitlog.NewSyncWriter(os.Stderr))
		mockSvc := gomock.NewController(GinkgoT())
		mockSvc.Finish()
		mockVideoRepo = mocks.NewMockDatabaseI(mockSvc)
		video = usecase.NewVideo(mockVideoRepo, logger)
	})

	// DECLARE UNIT TEST FUNCTION

	var GetListVideoLogic = func(idx int) {
		ctx := context.Background()
		data := testcases.GetListVideoData[idx]
		mockVideoRepo.EXPECT().GetListVideo(ctx, gomock.Any()).Return(data.MockGetListVideoRepo.Result, data.MockGetListVideoRepo.Error).Times(1)
		mockVideoRepo.EXPECT().GetMetadataVideo(ctx, gomock.Any()).Return(data.MockGetMetadata.Result, data.MockGetMetadata.Error).Times(1)
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
		mockVideoRepo.EXPECT().GetLocationNameByID(ctx, data.GetLocationName).Return(data.MockGetLocationName.Result, data.MockGetLocationName.Error).Times(1)
		resp, err := video.GetDetailVideo(ctx, data.UsecaseRequest)
		if err != nil {
			Expect(err).NotTo(BeNil())
			Expect(resp).To(BeNil())
		} else {
			Expect(err).To(BeNil())
		}
	}

	var unitTestLogic = map[string]map[string]interface{}{
		"GetListVideo":   {"func": GetListVideoLogic, "test_case_count": len(testcases.GetListVideoData), "desc": testcases.ListVideoDescription()},
		"GetDetailVideo": {"func": GetDetailVideoLogic, "test_case_count": len(testcases.GetDetailVideoData), "desc": testcases.DetailVideoDescription()},
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
