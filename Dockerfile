FROM golang:latest

ADD . /go/api

WORKDIR /go/api

RUN ls
RUN go get github.com/gorilla/handlers
RUN go get github.com/gorilla/mux
RUN go get github.com/lib/pq
RUN go get github.com/dgrijalva/jwt-go
RUN go get github.com/bitly/go-simplejson
RUN go get github.com/mitchellh/mapstructure

CMD ["go", "run", "main.go"]
