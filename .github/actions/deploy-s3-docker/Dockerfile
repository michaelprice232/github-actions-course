# Ideally this should use a multi-stage build from scatch and the image uploaded to a Doceker registry, but including here for easier reference
FROM golang:1.23

WORKDIR /usr/src/app

# pre-copy/cache go.mod for pre-downloading dependencies and only redownloading them in subsequent builds if they change
COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .
RUN go build -v -o /usr/local/bin/app main.go

ENTRYPOINT ["app"]