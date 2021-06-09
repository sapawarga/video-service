package http

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/sapawarga/video-service/endpoint"
	"github.com/sapawarga/video-service/helper"
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
	r.Handle("/videos/health/live", kithttp.NewServer(endpoint.MakeCheckHealthy(ctx), decodeNoRequest, encodeResponse, opts...)).Methods(helper.HTTP_GET)
	r.Handle("/videos/health/ready", kithttp.NewServer(endpoint.MakeCheckReadiness(ctx, fs), decodeNoRequest, encodeResponse, opts...)).Methods(helper.HTTP_GET)
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
	r.Handle("/videos/", processVideoGetList).Methods(helper.HTTP_GET)
	r.Handle("/videos/statistic", processGetVideoStatistic).Methods(helper.HTTP_GET)
	r.Handle("/videos/{id}", processGetDetailVideo).Methods(helper.HTTP_GET)
	r.Handle("/videos/", processCreateVideo).Methods(helper.HTTP_POST)
	r.Handle("/videos/{id}", processUpdateVideo).Methods(helper.HTTP_PUT)
	r.Handle("/videos/{id}", processDeleteVideo).Methods(helper.HTTP_DELETE)

	return r
}

func decodeGetListVideo(ctx context.Context, r *http.Request) (interface{}, error) {
	regIDString := r.URL.Query().Get("regency_id")
	pageString := r.URL.Query().Get("page")

	if pageString == "0" || pageString == "" {
		pageString = "1"
	}
	regID, _ := helper.ConvertFromStringToInt64(regIDString)
	pageInt, _ := helper.ConvertFromStringToInt64(pageString)
	request := &endpoint.GetVideoRequest{
		RegencyID: regID,
		Page:      pageInt,
	}

	return request, nil
}

func decodeGetByID(ctx context.Context, r *http.Request) (interface{}, error) {
	params := mux.Vars(r)
	_, id := helper.ConvertFromStringToInt64(params["id"])
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
	id, _ := helper.ConvertFromStringToInt64(params["id"])

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
		if status.Code == helper.STATUS_CREATED {
			w.WriteHeader(http.StatusCreated)
		} else if status.Code == helper.STATUS_UPDATED || status.Code == helper.STATUS_DELETED {
			w.WriteHeader(http.StatusNoContent)
			return json.NewEncoder(w).Encode(nil)
		}
	} else {
		w.WriteHeader(http.StatusOK)
	}
	return json.NewEncoder(w).Encode(response)
}

func encodeError(_ context.Context, err error, w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusBadRequest)

	json.NewEncoder(w).Encode(map[string]interface{}{
		"error": err.Error(),
	})

}
