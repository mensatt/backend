FROM golang:1.19-bullseye as base

# Set environment variables for the user
ARG USER_ID
ARG GROUP_ID

# Create a group and user
RUN addgroup --gid $GROUP_ID mensatt && \
    adduser --disabled-password --gecos '' --uid $USER_ID --gid $GROUP_ID mensatt

# This is required to set proper permissions for the assets directory when using a volume
RUN mkdir -p /opt/app/mensatt/assets && chown -R mensatt:mensatt /opt/app/mensatt/assets

# Use the new user
USER mensatt

# Create another stage called "dev" that is based off of our "base" stage (so we have golang available to us)
FROM base as dev

# Install the air binary so we get live code-reloading when we save files
RUN go install github.com/cosmtrek/air@latest

# Run the air command in the directory where our code will live
WORKDIR /opt/app/mensatt

# Start air hot-reload server
CMD ["air"]

FROM base as built

WORKDIR /go/app/mensatt
COPY . .

#ENV CGO_ENABLED=1

RUN go mod download
#RUN go build -tags netgo -ldflags="-extldflags=-static" -o /tmp/mensatt ./cmd/mensatt/main.go
RUN go build -o /tmp/mensatt ./cmd/mensatt/main.go

FROM busybox:glibc as prod

COPY --from=built /tmp/mensatt /usr/bin/mensatt

EXPOSE 4000
CMD ["mensatt"]
