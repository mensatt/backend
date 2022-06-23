FROM golang:1.18 as base

# Create another stage called "dev" that is based off of our "base" stage (so we have golang available to us)
FROM base as dev

# Install the air binary so we get live code-reloading when we save files
RUN go install github.com/cosmtrek/air@latest

# Install dbmate for datbase migrations
RUN go install github.com/amacneil/dbmate@latest

# Run the air command in the directory where our code will live
WORKDIR /opt/app/mensatt

# Start air hot-reload server
CMD ["air"]

FROM base as built

WORKDIR /go/app/mensatt
COPY . .

ENV CGO_ENABLED=1

RUN go mod download
RUN go build -ldflags="-extldflags=-static" -o /tmp/mensatt ./cmd/mensatt/main.go

FROM busybox as prod

COPY --from=built /tmp/mensatt /usr/bin/mensatt

EXPOSE 4000
CMD ["mensatt"]
