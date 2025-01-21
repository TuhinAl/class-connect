#!/bin/bash

# Navigate to project root directory
cd "$(dirname "$0")/.."

# Create bin directory if it doesn't exist
mkdir -p bin

# Clean previous build
rm -f bin/golang-api

# Build for Linux
echo "Building for Linux..."
GOOS=linux GOARCH=amd64 go build -o bin/golang-api main.go

echo "Build complete! Binary is located at bin/golang-api"

# Run the application
echo "Starting the server..."
./bin/golang-api
