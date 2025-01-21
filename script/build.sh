#!/bin/bash

# Navigate to project root directory
cd "$(dirname "$0")/.."

# Check if bin directory exists
if [ ! -d "bin" ]; then
    echo "Creating bin directory..."
    mkdir bin
else
    echo "bin directory already exists"
fi

# Clean previous build
rm -f bin/golang-api

# Build for Linux
echo "Building CLASS-CONNECT Application..."
GOOS=linux GOARCH=amd64 go build -o bin/golang-api cmd/api/main.go

echo "Build complete! Binary is located at bin/golang-api"

# Run the application
echo "Starting the server..."

# Check if the binary was actually created
if [ -f "bin/golang-api" ]; then
    ./bin/golang-api
else
    echo "Error: Build failed, binary not created"
    exit 1
fi
