package http

import (
	"context"
	"encoding/json"
	"net/http"
	"strings"

	"github.com/sapawarga/video-service/endpoint"
	"github.com/sapawarga/video-service/lib/constants"
	"github.com/sapawarga/video-service/lib/converter"
	"github.com/sapawarga/video-service/usecase"

	kitlog "github.com/go-kit/kit/log"
	"github.com/go-kit/kit/transport"
	kithttp "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
)

type err interface {
	error() error
}

func MakeHealthyCheckHandler(ctx context.Context, fs usecase.UsecaseI, logger kitlog.Logger) http.Handler {
	opts := []kithttp.ServerOption{
		kithttp.ServerErrorHandler(transport.NewLogErrorHandler(logger)),
		kithttp.ServerErrorEncoder(encodeError),
	}

	r := mux.NewRouter()
	r.Handle("/health/live", kithttp.NewServer(endpoint.MakeCheckHealthy(ctx), decodeNoRequest, encodeResponse, opts...)).Methods(constants.HTTP_GET)
	r.Handle("/health/ready", kithttp.NewServer(endpoint.MakeCheckReadiness(ctx, fs), decodeNoRequest, encodeResponse, opts...)).Methods(constants.HTTP_GET)
	return r
}

func MakeHTTPHandler(ctx context.Context, fs usecase.UsecaseI, logger kitlog.Logger) http.Handler {
	opts := []kithttp.ServerOption{
		kithttp.ServerErrorHandler(transport.NewLogErrorHandler(logger)),
		kithttp.ServerErrorEncoder(encodeError),
	}

	processVideoGetList := kithttp.NewServer(endpoint.MakeGetListVideo(ctx, fs), decodeGetListVideo, encodeResponse, opts...)
	processGetDetailVideo := kithttp.NewServer(endpoint.MakeGetDetailVideo(ctx, fs), decodeGetByID, encodeResponse, opts...)
	processGetVideoStatistic := kithttp.NewServer(endpoint.MakeGetVideoStatistic(ctx, fs), decodeNoRequest, encodeResponse, opts...)
	processCreateVideo := kithttp.NewServer(endpoint.MakeCreateNewVideo(ctx, fs), decodeCreateVideo, encodeResponse, opts...)
	processUpdateVideo := kithttp.NewServer(endpoint.MakeUpdateVideo(ctx, fs), decodeUpdateVideo, encodeResponse, opts...)
	processDeleteVideo := kithttp.NewServer(endpoint.MakeDeleteVideo(ctx, fs), decodeGetByID, encodeResponse, opts...)

	r := mux.NewRouter()

	// TODO: handle token middleware
	r.Handle("/videos/", processVideoGetList).Methods(constants.HTTP_GET)
	r.Handle("/videos/", processCreateVideo).Methods(constants.HTTP_POST)
	r.Handle("/videos/statistics", processGetVideoStatistic).Methods(constants.HTTP_GET)
	r.Handle("/videos/{id}", processGetDetailVideo).Methods(constants.HTTP_GET)
	r.Handle("/videos/{id}", processUpdateVideo).Methods(constants.HTTP_PUT)
	r.Handle("/videos/{id}", processDeleteVideo).Methods(constants.HTTP_DELETE)

	return r
}

func decodeGetListVideo(ctx context.Context, r *http.Request) (interface{}, error) {
	regIDString := r.URL.Query().Get("kabkota_id")
	pageString := r.URL.Query().Get("page")
	limitString := r.URL.Query().Get("limit")
	categoryIDString := r.URL.Query().Get("category_id")
	title := r.URL.Query().Get("title")
	search := r.URL.Query().Get("search")
	sortBy := r.URL.Query().Get("sort_by")
	sortOrder := "DESC"
	if r.URL.Query().Get("sort_order") != "" {
		sortOrder = constants.AscOrDesc[r.URL.Query().Get("sort_order")]
	}
	status, _ := converter.ConvertFromStringToInt64(r.URL.Query().Get("status"))
	var statusDef int64 = 10
	if status != nil {
		statusDef = *status
	}
	if pageString == "0" || pageString == "" {
		pageString = "1"
	}
	if limitString == "" || limitString == "0" {
		limitString = "10"
	}
	if sortBy == "" {
		sortBy = "created_at"
	}
	_, regID := converter.ConvertFromStringToInt64(regIDString)
	var pointerRegID *int64 = nil
	if regID != 0 {
		pointerRegID = &regID
	}
	pageInt, _ := converter.ConvertFromStringToInt64(pageString)
	limit, _ := converter.ConvertFromStringToInt64(limitString)
	_, categoryID := converter.ConvertFromStringToInt64(categoryIDString)
	var pointerCatID *int64 = nil
	if categoryID != 0 {
		pointerCatID = &categoryID
	}
	return &endpoint.GetVideoRequest{
		Search:     converter.SetPointerString(search),
		RegencyID:  pointerRegID,
		Page:       pageInt,
		Limit:      limit,
		CategoryID: pointerCatID,
		Title:      converter.SetPointerString(title),
		SortBy:     sortBy,
		SortOrder:  sortOrder,
		Status:     statusDef,
	}, nil
}

func decodeGetByID(ctx context.Context, r *http.Request) (interface{}, error) {
	params := mux.Vars(r)
	_, id := converter.ConvertFromStringToInt64(params["id"])
	request := &endpoint.RequestID{
		ID: id,
	}
	return request, nil
}

func decodeNoRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	return r, nil
}

func decodeCreateVideo(ctx context.Context, r *http.Request) (interface{}, error) {
	reqBody := &endpoint.CreateVideoRequest{}
	if err := json.NewDecoder(r.Body).Decode(reqBody); err != nil {
		return nil, err
	}

	return reqBody, nil
}

func decodeUpdateVideo(ctx context.Context, r *http.Request) (interface{}, error) {
	params := mux.Vars(r)
	id, _ := converter.ConvertFromStringToInt64(params["id"])

	reqBody := &endpoint.UpdateVideoRequest{}
	if err := json.NewDecoder(r.Body).Decode(reqBody); err != nil {
		return nil, err
	}

	reqBody.ID = id
	return reqBody, nil
}

func encodeResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	if e, ok := response.(err); ok && e.error() != nil {
		encodeError(ctx, e.error(), w)
		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	status, ok := response.(*endpoint.StatusResponse)
	if ok {
		if status.Code == constants.STATUS_CREATED {
			w.WriteHeader(http.StatusCreated)
		} else if status.Code == constants.STATUS_UPDATED || status.Code == constants.STATUS_DELETED {
			w.WriteHeader(http.StatusNoContent)
			_ = json.NewEncoder(w).Encode(nil)
			return nil
		}
	} else {
		w.WriteHeader(http.StatusOK)
	}
	return json.NewEncoder(w).Encode(response)
}

func encodeError(_ context.Context, err error, w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	if strings.ContainsAny(err.Error(), "error") {
		w.WriteHeader(http.StatusUnprocessableEntity)
	} else {
		w.WriteHeader(http.StatusInternalServerError)
	}

	json.NewEncoder(w).Encode(map[string]interface{}{
		"error": err.Error(),
	})
}
