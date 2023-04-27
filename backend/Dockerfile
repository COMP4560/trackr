# syntax=docker/dockerfile:1

FROM golang:1.18-alpine

RUN apk add --no-cache --update gcc g++

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
COPY .env ./
RUN go mod download

COPY src ./src

RUN go build -o /trackr src/main.go

EXPOSE 8080

CMD [ "/trackr" ]
