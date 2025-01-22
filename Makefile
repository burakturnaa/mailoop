# Simple Makefile for a Go project

# Build and test the application
all: build test

build:
	@echo "Building..."
	
	
	@go build -o main.exe main.go

# Run the application
run:
	@go run main.go

# Build and test the docker application
# docker-build:
# 	@echo "Building..."
# 	@docker build -t mailoop-app .

# Run the docker application
docker-run:
	@docker run -it --rm -p 3000:3000 --env-file .env mailoop-app

# Create and run Docker Compose services
docker-build:
	@echo "Building..."
	@docker compose up --build -d

# Shutdown Docker Compose services
docker-down:
	@docker compose down

# Test the application
test:
	@echo "Testing..."
	@go test ./... -v

# Integration tests
itest:
	@echo "Running integration tests..."
	@go test ./configs/setup -v

# Clean build artifacts
clean:
	@echo "Cleaning..."
	@rm -f main.exe

# Watch for file changes and reload
watch:
	@powershell -ExecutionPolicy Bypass -Command "if (Get-Command air -ErrorAction SilentlyContinue) { \
		air; \
		Write-Output 'Watching...'; \
	} else { \
		Write-Output 'Installing air...'; \
		go install github.com/air-verse/air@latest; \
		air; \
		Write-Output 'Watching...'; \
	}"

.PHONY: all build run test clean watch docker-run docker-down itest
