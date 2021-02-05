FROM golang:latest AS builder

ADD . /go/api

WORKDIR /go/api

RUN go get -u github.com/swaggo/swag/cmd/swag

RUN rm -rf deploy
RUN mkdir deploy
RUN go mod download
RUN /go/bin/swag init
RUN go test ./...

RUN CGO_ENABLED=0 go build -o goapp
RUN mv goapp ./deploy/goapp
RUN mv docs ./deploy/docs
RUN mv config.json ./deploy/config.json




FROM alpine:3.7 AS production

COPY --from=builder /go/api/deploy /api/

WORKDIR /api

ENTRYPOINT  ./goapp

