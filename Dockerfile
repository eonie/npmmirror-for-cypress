## Build
FROM golang:1.16-alpine AS build

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY *.go ./

RUN go build -o /npmmirror-for-cypress

## Deploy
FROM alpine:latest

RUN addgroup -S mfc && adduser -S mfc -G mfc
USER mfc

WORKDIR /app

COPY --from=build /npmmirror-for-cypress /app/

EXPOSE 8090

ENTRYPOINT ["/app/npmmirror-for-cypress"] 