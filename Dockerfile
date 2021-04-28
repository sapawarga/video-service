FROM golang:1.15 AS compile-image

# RUN apk --no-cache add gcc g++ make ca-certificates git
# Set necessary environmet variables needed for our image
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

WORKDIR /build

COPY go.mod .
COPY go.sum .
COPY .env.example ./.env
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 go test -v ./...
RUN make build

FROM alpine:3.11.3 as run-image

LABEL maintainer="GoSapawarga <setiadi.yon3@gmail.com>"

ENV PROJECT_PATH=/build

WORKDIR /app

COPY --from=compile-image ${PROJECT_PATH}/video-service-grpc /app/video-service-grpc
COPY --from=compile-image ${PROJECT_PATH}/.env /app/.env

RUN apk --update add tzdata ca-certificates && \
    update-ca-certificates 2>/dev/null || true 

EXPOSE 3001

ENTRYPOINT [ "/app/video-service-grpc" ]
