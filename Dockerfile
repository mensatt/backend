FROM golang:1.22-alpine as base

# Install git for VCS build information
RUN apk add --no-cache git

# Create another stage called "dev" that is based off of our "base" stage (so we have golang available to us)
FROM base as dev

# Install the air binary so we get live code-reloading when we save files
RUN go install github.com/cosmtrek/air@latest

# Run the air command in the directory where our code will live
WORKDIR /opt/mensatt
RUN git config --global --add safe.directory /opt/mensatt

# Start air hot-reload server
CMD ["air"]

FROM base as built

WORKDIR /opt/mensatt
#RUN git config --global --add safe.directory /opt/mensatt

# Install required dependencies here for better cache utilization
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the code and build the binary
COPY . .
RUN go build -o ./build/mensatt ./cmd/mensatt

FROM alpine:3.19 as prod

COPY --from=built /opt/mensatt/build/mensatt /usr/bin/mensatt

USER nobody:nobody

EXPOSE 4000
CMD ["mensatt"]
