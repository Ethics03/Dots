#!/bin/bash

# Colors for output
GREEN="\033[1;32m"
YELLOW="\033[1;33m"
RED="\033[1;31m"
BLUE="\033[1;34m"
RESET="\033[0m"

#set variables
CONFIG_DIR="$HOME/.config/dots"
SRC_DIR="./dots"
DOTS_YAML="./dots.yaml"

echo -e "${BLUE} Configuring and installing Dots..${RESET}"

#creating the config dir if missing
if [ ! -d "$CONFIG_DIR" ]; then
  echo -e "${YELLOW}Creating config directory at $CONFIG_DIR${RESET}"
  mkdir -p "$CONFIG_DIR"

else
  echo -e "${GREEN}Config directory already exists: $CONFIG_DIR${RESET}"

fi

# Copy dots folder contents
echo -e "${BLUE} Copying dotfiles to $CONFIG_DIR${RESET}"
cp -r "$SRC_DIR/"* "$CONFIG_DIR/" || {
  echo -e "${RED}Failed to copy dots folder!${RESET}"
  exit 1
}

# Copy dots.yaml
echo -e "${BLUE}📄 Copying dots.yaml to $CONFIG_DIR${RESET}"
cp "$DOTS_YAML" "$CONFIG_DIR/" || {
  echo -e "${RED}Failed to copy dots.yaml!${RESET}"
  exit 1
}
