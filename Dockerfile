FROM golang:latest AS builder

ADD . /go/api

WORKDIR /go/api

RUN go get -u github.com/swaggo/swag/cmd/swag

RUN rm -rf deploy
RUN mkdir deploy
RUN /go/bin/swag init -g adapter/http_service/main.go -o adapter/http_service/docs
RUN go mod tidy

RUN CGO_ENABLED=0 go build -o goapp adapter/http_service/main.go
RUN mv goapp ./deploy/goapp
RUN mv adapter/http_service/docs ./deploy/docs
RUN mv adapter/http_service/config.json ./deploy/config.json


FROM alpine:3.7 AS production

COPY --from=builder /go/api/deploy /api/

WORKDIR /api

ENTRYPOINT  ./goapp

