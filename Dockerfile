FROM golang:1.15-alpine3.12 AS compile-image

RUN apk --no-cache add gcc g++ make ca-certificates git
# Set necessary environmet variables needed for our image
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

WORKDIR /build

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 go test -v ./...
RUN make build

FROM gcr.io/distroless/static as run-image

LABEL maintainer="GoSapawarga <setiadi.yon3@gmail.com>"

ENV PROJECT_PATH=/build

RUN apk --update add tzdata ca-certificates && \
    update-ca-certificates 2>/dev/null || true 

COPY --from=compile-image ${PROJECT_PATH}/video-service /

ENTRYPOINT [ "/video-service" ]
