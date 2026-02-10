#!/bin/bash

# Colors for output
GREEN="\033[1;32m"
YELLOW="\033[1;33m"
RED="\033[1;31m"
BLUE="\033[1;34m"
RESET="\033[0m"

set -e  # Exit on error

echo -e "${BLUE}╔═══════════════════════════════════════════════════════════╗${RESET}"
echo -e "${BLUE}║                                                           ║${RESET}"
echo -e "${BLUE}║                Installing Dots                            ║${RESET}"
echo -e "${BLUE}║                                                           ║${RESET}"
echo -e "${BLUE}╚═══════════════════════════════════════════════════════════╝${RESET}"
echo

# Check if Go is installed
if ! command -v go &> /dev/null; then
    echo -e "${RED}Error: Go is not installed${RESET}"
    echo "Please install Go from https://golang.org/dl/"
    exit 1
fi

# Check if git is installed
if ! command -v git &> /dev/null; then
    echo -e "${RED}Error: Git is not installed${RESET}"
    echo "Please install Git first"
    exit 1
fi

# Build the binary
echo -e "${YELLOW}Building dots binary...${RESET}"
if go build -o dots; then
    echo -e "${GREEN}✓ Binary built successfully${RESET}"
else
    echo -e "${RED}✗ Failed to build binary${RESET}"
    exit 1
fi

# Install to /usr/local/bin
INSTALL_DIR="/usr/local/bin"
echo
echo -e "${YELLOW}Installing to $INSTALL_DIR...${RESET}"

if [ -w "$INSTALL_DIR" ]; then
    mv dots "$INSTALL_DIR/dots"
    echo -e "${GREEN}✓ Installed to $INSTALL_DIR/dots${RESET}"
else
    echo -e "${YELLOW}Need sudo privileges to install to $INSTALL_DIR${RESET}"
    sudo mv dots "$INSTALL_DIR/dots"
    echo -e "${GREEN}✓ Installed to $INSTALL_DIR/dots${RESET}"
fi

# Verify installation
echo
if command -v dots &> /dev/null; then
    echo -e "${GREEN}✓ Installation successful!${RESET}"
    echo
    echo -e "${BLUE}═══════════════════════════════════════════════════════════${RESET}"
    echo -e "${GREEN}Dots has been installed successfully!${RESET}"
    echo -e "${BLUE}═══════════════════════════════════════════════════════════${RESET}"
    echo
    echo "Next steps:"
    echo "  1. Initialize dots:       dots init"
    echo "  2. Add a dotfile:         dots add ~/.bashrc"
    echo "  3. Check status:          dots status"
    echo "  4. Get help:              dots --help"
    echo
else
    echo -e "${RED}✗ Installation verification failed${RESET}"
    exit 1
fi
