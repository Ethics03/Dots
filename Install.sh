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
