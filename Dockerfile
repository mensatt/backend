FROM golang:1.19-alpine as base

# Set environment variables for the user
ARG USER_ID
ARG GROUP_ID

# Install libvips-dev for image processing
RUN apk add --no-cache vips-dev build-base git

# Create a group and user
RUN addgroup -g $GROUP_ID mensatt && \
    adduser -D -s /bin/sh -u $USER_ID -G mensatt mensatt

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

RUN ls -asl
RUN pwd
#RUN git log

RUN git config --global --add safe.directory /go/app/mensatt
RUN git log

RUN go mod download
RUN go build -o /tmp/mensatt ./cmd/mensatt/main.go

FROM alpine:3.18 as prod

COPY --from=built /tmp/mensatt /usr/bin/mensatt

# Install libvips-dev for image processing
RUN apk add --no-cache vips

EXPOSE 4000
CMD ["mensatt"]
