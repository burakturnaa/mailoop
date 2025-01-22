# Base image
FROM golang:1.21 AS builder

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy the Go Modules manifests
COPY go.mod go.sum ./

# Download Go modules
RUN go mod download

# Copy the source code into the container
COPY . .

# Build the Go app
RUN go build -o main main.go

# Start a new stage from a smaller image
FROM debian:bullseye-slim

# Set the working directory in the container
WORKDIR /root/

# Copy the compiled binary from the builder stage
COPY --from=builder /app/main .

# Expose the port
EXPOSE 3000

# Command to run the executable
CMD ["./main"]