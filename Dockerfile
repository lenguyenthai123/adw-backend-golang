# Use the official Golang image as the base image
FROM golang:latest as builder

# Set the Current Working Directory inside the container
WORKDIR /app

# Set the Go proxy to direct
ENV GOPROXY=direct

# Copy go.mod and go.sum files
COPY go.mod go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Copy the source code into the container
COPY . .

# Build the Go app with static linking
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o main .

# Start a new stage from scratch
FROM alpine:latest

# Set the Current Working Directory inside the container
WORKDIR /root/

# Install bash (optional, for debugging)
RUN apk add --no-cache bash

# Copy the Pre-built binary file from the previous stage
COPY --from=builder /app/main .

# Ensure the binary has execute permission
RUN chmod +x ./main

# Expose port 8081 to the outside world
EXPOSE 8081

# Command to run the executable
CMD ["./main"]