FROM golang:1.22-alpine as base

# Create another stage called "dev" that is based off of our "base" stage (so we have golang available to us)
FROM base as dev

# Install the air binary so we get live code-reloading when we save files
RUN go install github.com/cosmtrek/air@latest

# Run the air command in the directory where our code will live
WORKDIR /opt/mensatt

# Start air hot-reload server
CMD ["air"]

FROM base as built

WORKDIR /go/app/mensatt
COPY . .

RUN go mod download
RUN go build -o /tmp/mensatt ./cmd/mensatt/main.go

FROM alpine:3.19 as prod

COPY --from=built /tmp/mensatt /usr/bin/mensatt

USER nobody:nobody

EXPOSE 4000
CMD ["mensatt"]
