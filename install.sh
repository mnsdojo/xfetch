#!/bin/bash

# Check if Go is installed
if ! command -v go &> /dev/null; then
    echo "Go is not installed. Please install Go and try again."
    exit 1
fi

# Create a temporary directory
TEMP_DIR=$(mktemp -d)
cd "$TEMP_DIR"

# Clone the repository
git clone https://github.com/mnsdojo/xfetch.git
cd xfetch

# Initialize Go module if go.mod doesn't exist
if [ ! -f go.mod ]; then
    go mod init github.com/mnsdojo/xfetch
    go mod tidy
fi

# Build the program
go build -o xfetch 

# Move the binary to /usr/local/bin
sudo mv xfetch /usr/local/bin/

# Clean up
cd ..
rm -rf "$TEMP_DIR"

# Determine the user's shell
SHELL_NAME=$(basename "$SHELL")

# Add xfetch to the appropriate shell configuration file
if [ "$SHELL_NAME" = "zsh" ]; then
    echo 'export PATH=$PATH:/usr/local/bin' >> ~/.zshrc
    echo "xfetch has been added to your .zshrc file."
    echo "Please run 'source ~/.zshrc' or start a new terminal session to use xfetch."
elif [ "$SHELL_NAME" = "bash" ]; then
    echo 'export PATH=$PATH:/usr/local/bin' >> ~/.bashrc
    echo "xfetch has been added to your .bashrc file."
    echo "Please run 'source ~/.bashrc' or start a new terminal session to use xfetch."
else
    echo "Your shell ($SHELL_NAME) is not recognized. Please add '/usr/local/bin' to your PATH manually."
fi

echo "xfetch has been installed successfully. You can now run it by typing 'xfetch' in the terminal."
