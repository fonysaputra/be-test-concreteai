FROM golang:1.22-bullseye

ENV TZ Asia/Jakarta

WORKDIR /go/src/ms-go-gin
COPY . .

RUN go get -d -v ./... \
    && go get -u github.com/air-verse/air \
    && go install github.com/air-verse/air \
    && go build -o  /usr/local/go/bin/app cmd/main.go

CMD app
