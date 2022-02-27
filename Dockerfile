FROM golang:1.17 as base

# Create another stage called "dev" that is based off of our "base" stage (so we have golang available to us)
FROM base as dev

# Install the air binary so we get live code-reloading when we save files
RUN curl -sSfL https://raw.githubusercontent.com/cosmtrek/air/master/install.sh | sh -s -- -b $(go env GOPATH)/bin

# Run the air command in the directory where our code will live
WORKDIR /opt/app/mensatt
CMD ["air"]

FROM base as built

WORKDIR /go/app/mensatt
COPY . .

ENV CGO_ENABLED=0

RUN go get -d -v ./...
RUN go build -o /tmp/mensatt ./cmd/mensatt/main.go

FROM busybox

COPY --from=built /tmp/mensatt /usr/bin/mensatt
CMD ["mensatt"]