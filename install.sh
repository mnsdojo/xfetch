#!/bin/bash

# Check if Go is installed
if ! command -v go &> /dev/null; then
    echo "Go is not installed. Please install Go and try again."
    exit 1
fi

# Create a temporary directory
TEMP_DIR=$(mktemp -d)
cd "$TEMP_DIR" || exit

# Clone the repository
git clone https://github.com/mnsdojo/xfetch.git
cd xfetch || exit

# Initialize Go module if go.mod doesn't exist
if [ ! -f go.mod ]; then
    go mod init github.com/mnsdojo/xfetch
    go mod tidy
fi

# Build the program
go build -o xfetch

# Set the correct permissions
chmod +x xfetch

# Move the binary to /usr/local/bin (or other specified directory)
INSTALL_DIR=${1:-/usr/local/bin}
sudo mv xfetch "$INSTALL_DIR/"

# Ensure correct permissions after moving
sudo chmod 755 "$INSTALL_DIR/xfetch"

# Clean up temporary directory
cd ..
rm -rf "$TEMP_DIR"

# Add xfetch to both .bashrc and .zshrc for universal support
if [[ ":$PATH:" != *":$INSTALL_DIR:"* ]]; then
    echo 'export PATH=$PATH:'"$INSTALL_DIR" >> ~/.bashrc
    echo 'export PATH=$PATH:'"$INSTALL_DIR" >> ~/.zshrc
    echo "xfetch has been added to both your .bashrc and .zshrc files."
    echo "Please run 'source ~/.bashrc' or 'source ~/.zshrc' or start a new terminal session to use xfetch."
else
    echo "xfetch is already in your PATH."
fi

echo "xfetch has been installed successfully. You can now run it by typing 'xfetch' in the terminal."
