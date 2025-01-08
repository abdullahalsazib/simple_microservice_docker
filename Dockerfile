# Use the official Golang image from the Docker Hub as a builder image
FROM golang:1.20-alpine AS builder

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy the Go Modules manifests to the container
COPY go.mod go.sum ./

# Download all the dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod tidy

# Copy the source code into the container
COPY . .

# Build the Go app
RUN go build -o main .

# Start a new stage from scratch (smaller image size)
FROM alpine:latest  

# Install sqlite3 if you need it for your app
RUN apk --no-cache add sqlite

# Set the Current Working Directory inside the container
WORKDIR /root/

# Copy the Pre-built binary file from the builder image
COPY --from=builder /app/main .

# Expose the port your app runs on (8080 in this case)
EXPOSE 8080

# Run the Go binary (your app)
CMD ["./main"]
