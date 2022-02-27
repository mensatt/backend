FROM golang:1.17 as base

# Create another stage called "dev" that is based off of our "base" stage (so we have golang available to us)
FROM base as dev

# Install the air binary so we get live code-reloading when we save files
RUN go install github.com/cosmtrek/air@latest

# Install dbmate for datbase migrations
RUN curl -fsSL -o /usr/local/bin/dbmate https://github.com/amacneil/dbmate/releases/latest/download/dbmate-linux-amd64
RUN chmod +x /usr/local/bin/dbmate

# Run the air command in the directory where our code will live
WORKDIR /opt/app/mensatt

# Download go module dependencies
RUN go mod download

# Start air hot-reload server
CMD ["air"]

FROM base as built

WORKDIR /go/app/mensatt
COPY . .

ENV CGO_ENABLED=0

RUN go mod download
RUN go build -o /tmp/mensatt ./cmd/mensatt/main.go

FROM busybox

COPY --from=built /tmp/mensatt /usr/bin/mensatt
CMD ["mensatt"]
