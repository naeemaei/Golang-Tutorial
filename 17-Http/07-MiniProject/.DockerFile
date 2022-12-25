FROM golang:1.18-alpine as builder

WORKDIR /app

COPY go.* ./
RUN go mod download

COPY . ./

RUN go build -v -o server

COPY ./ ./app/server
EXPOSE 80
CMD ["/app/server"]