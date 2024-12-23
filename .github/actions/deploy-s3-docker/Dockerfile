# Ideally this image should be uploaded to a Docker registry, but building in-place for easier reference
FROM golang:1.23 AS builder

WORKDIR /usr/src/app

# pre-copy/cache go.mod for pre-downloading dependencies and only redownloading them in subsequent builds if they change
COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o /usr/local/bin/app main.go

FROM alpine:3

RUN apk add --no-cache ca-certificates

COPY --from=builder /usr/local/bin/app /app

ENTRYPOINT ["/app"]