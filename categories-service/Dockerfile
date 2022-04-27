# syntax=docker/dockerfile:1

# Base Image for Golang project. Alpine has smaller footprint than Ubuntu.
FROM golang:alpine as builder

# Set necessary environment variables needed for our image
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

# Create a directly called /app and change directory into it
WORKDIR /app

# Copy project files into the image
COPY ./ ./

# Get Dependencies
RUN go mod download

# Run tests
RUN go test ./...

# Build the Golang Application  
RUN go build -o /main .

FROM alpine
COPY --from=build /app /.

# Expose the application as a port and run it
EXPOSE ${CATEGORIES_SERVICE_PORT}


CMD ["/main"]
