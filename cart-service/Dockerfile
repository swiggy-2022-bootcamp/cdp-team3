# syntax=docker/dockerfile:1

# Base Image for Golang project. Alpine has smaller footprint than Ubuntu.
FROM golang:1.16-alpine

# Create a directly called /app and change directory into it
WORKDIR /app

# Copy Go dependency definitions into the image
COPY go.mod ./
COPY go.sum ./

# Get Dependencies
RUN go mod download

# Copy entire Golang Source Code into the image
COPY *.go ./

# Build the Golang Application  
RUN go build -o /main

# Expose the application as a port and run it
EXPOSE 8080
CMD ["/main"]
