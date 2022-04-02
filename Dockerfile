FROM golang:latest

WORKDIR /tribun-news-unofficial-api

COPY . .

RUN go mod download

RUN go build

CMD ["./tribun-news-unofficial-api"]