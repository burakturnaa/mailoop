# Base image
FROM golang:1.21 AS builder

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy the Go Modules manifests
COPY go.mod go.sum ./

# Download Go modules
RUN go mod download

# Copy the source code into the container
COPY ./ ./

# Build the Go app
RUN CGO_ENABLED=0 GOOS=linux go build -o /mailoop-api-go

# Expose the port
EXPOSE 3000

# Command to run the executable
CMD ["/mailoop-api-go"]