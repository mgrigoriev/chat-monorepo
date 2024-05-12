# syntax=docker/dockerfile:experimental

FROM golang:1.22-alpine AS build

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o /bin/main cmd/main.go

FROM alpine:latest AS final

WORKDIR /

COPY --from=build /bin/main /main

RUN mkdir /swaggerui
COPY --from=build /app/swaggerui swaggerui/

EXPOSE 8080
ENTRYPOINT ["/main"]
