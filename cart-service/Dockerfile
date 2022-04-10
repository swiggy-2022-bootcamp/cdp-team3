# syntax=docker/dockerfile:1

# Base Image for Golang project. Alpine has smaller footprint than Ubuntu.
FROM golang:1.18-alpine3.15

# Set necessary environment variables needed for our image
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

# Setup Arguments to be used throughout the Dockerfile
ARG SERVER_PORT=3010

# Create a directly called /app and change directory into it
WORKDIR /app

# Copy project files into the image
COPY ./ ./

# Set up environetment variables
ENV CART_SERVICE_PORT=${SERVER_PORT}

# Get Dependencies
RUN go mod download

# Run tests
RUN go test ./...

# Build the Golang Application  
RUN go build -o /main .

# Expose the application as a port and run it
EXPOSE ${SERVER_PORT}

CMD ["/main"]
