FROM golang:latest

ADD . /go/api

WORKDIR /go/api

RUN go mod download

CMD ["go", "run", "main.go"]
