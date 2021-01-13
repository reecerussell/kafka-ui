#!/bin/bash

build_ui() {
    echo "Starting UI build"

    cd ./ui

    echo "Installing Node depedencies..."
    npm install

    echo "Rebuilding node-sass..."
    npm rebuild node-sass

    echo "Building UI project"
    npm run build

    cd ..
}

build_server() {
    echo "Starting server build"

    echo "Downloading Go modules..."
    go mod download

    echo "Building server..."
    go build -tags dynamic -o server ./cmd/main.go
}

main() {
    build_ui
    build_server

    echo "UI Assets: ./ui/build"
    echo "Server: ./server"
}

main