PHONY: build

build:
	@echo "Building the project..."
	go build -o email-service
	@echo "Build complete."

dev:
	@echo "Running in development mode..."
	go run main.go --workers=3 --queue-size=100


