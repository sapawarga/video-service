package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/sapawarga/video-service/cmd/database"
	"github.com/sapawarga/video-service/config"
	"github.com/sapawarga/video-service/repository/mysql"
	transportGRPC "github.com/sapawarga/video-service/transport/grpc"
	transportHTTP "github.com/sapawarga/video-service/transport/http"
	"github.com/sapawarga/video-service/usecase"

	kitlog "github.com/go-kit/kit/log"
	"github.com/sapawarga/proto-file/video"
	"google.golang.org/grpc"
)

var (
	filename = "cmd/server/main.go"
	method   = "main"
)

func main() {
	config, err := config.NewConfig()
	errorCheck(err)

	ctx := context.Background()
	db := database.NewConnection(config.DB)
	errChan := make(chan error)
	repo := mysql.NewVideoRepository(db)
	logger := kitlog.NewLogfmtLogger(kitlog.NewSyncWriter(os.Stderr))
	logger = kitlog.With(logger, "ts", kitlog.DefaultTimestampUTC)
	uc := usecase.NewVideo(repo, logger)

	grpcAdd := flag.String("grpc", fmt.Sprintf(":%d", config.AppGRPCPort), "gRPC listening address")
	go func() {
		logger.Log("transport", "grpc", "address", *grpcAdd, "msg", "listening")
		listener, err := net.Listen("tcp", *grpcAdd)
		if err != nil {
			errChan <- err
			return
		}
		handler := transportGRPC.MakeHandler(ctx, uc)
		grpcServer := grpc.NewServer()
		video.RegisterVideoHandlerServer(grpcServer, handler)
		logger.Log("filename", filename, "method", method, "note", "running video service grpc")
		errChan <- grpcServer.Serve(listener)
	}()

	httpAdd := flag.String("http", fmt.Sprintf(":%d", config.AppHTTPPort), "HTTP listening address")
	go func() {
		logger.Log("transport", "http", "address", *httpAdd, "msg", "listening")
		mux := http.NewServeMux()
		ctx := context.Background()
		mux.Handle("/videos/health", transportHTTP.MakeHealthyCheckHandler(ctx, uc, logger))
		mux.Handle("/videos/", transportHTTP.MakeHTTPHandler(ctx, uc, logger))
		logger.Log("filename", filename, "method", method, "note", "running video service http")
		errChan <- http.ListenAndServe(*httpAdd, accessControl(mux))
	}()

	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
		errChan <- fmt.Errorf("%s", <-c)
		logger.Log(
			"filename", filename,
			"method", method,
			"note", "Gracefully Stop Trading Account GRPC",
		)
	}()
	logger.Log(
		"filename", filename,
		"method", method,
		"note", <-errChan,
	)
}

func accessControl(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		r.Header.Set("Access-Control-Allow-Origin", "*")
		r.Header.Set("Access-Control-Allow-Methods", "GET, PUT, PATCH, POST, OPTIONS")
		r.Header.Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, Access-Control-Allow-Origin, Access-Control-Allow-Headers, scope, state, hd, code")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, PUT, PATCH, POST, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, Access-Control-Allow-Origin, Access-Control-Allow-Headers, scope, state, hd, code")

		if r.Method == "OPTIONS" {
			return
		}

		h.ServeHTTP(w, r)
	})
}

func errorCheck(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
