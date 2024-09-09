#!/bin/bash

# Check if Go is installed
if ! command -v go &> /dev/null; then
    echo "Go is not installed. Please install Go and try again."
    exit 1
fi


# Create a temporary directory
TEMP_DIR=$(mktemp -d)
cd "$TEMP_DIR"

git clone git@github.com:mnsdojo/xfetch.git

go build -o xfetch main.go

# Move the binary to /usr/local/bin
sudo mv xfetch /usr/local/bin/

# Clean up
cd ..
rm -rf "$TEMP_DIR"


echo "xfetch has been installed successfully. You can now run it by typing 'xfetch' in the terminal."
