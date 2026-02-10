# dots

**dots** is a minimal, fast, and developer-friendly dotfile manager built in Go. It helps you effortlessly manage, version, and sync your dotfiles using symlinks and Git, without the complexity of bloated tools or manual setup.

Designed with simplicity in mind, dots lets you keep your config files organized in a single directory, backed by version control, and easily portable across machines. Whether you're setting up a new dev environment or just want clean dotfile hygiene, dots makes it painless.

## Features

- **Simple & Fast** - Minimal overhead, instant operations
- **Git Integration** - Full version control with sync, push, pull commands
- **Symlink Management** - Automatic symlink creation and tracking
- **Cross-Machine Sync** - Clone and deploy dotfiles to new machines
- **Safe Operations** - Built-in checks to prevent data loss
- **No Config Files** - Filesystem-based tracking, no YAML complexity

## Getting Started

### Quick Install

```bash
git clone https://github.com/Ethics03/Dots.git
cd Dots
chmod +x Install.sh
./Install.sh
```

The install script will:
- Build the binary
- Install to `/usr/local/bin`
- Verify installation

### Manual Installation

```bash
git clone https://github.com/Ethics03/Dots.git
cd Dots
go build -o dots
sudo mv dots /usr/local/bin
```

## Quick Start

```bash
# Initialize dots
dots init

# Add your first dotfile
dots add ~/.bashrc

# Check status
dots status

# Sync to remote (after setting up git remote)
cd ~/.config/dots
git remote add origin <your-repo-url>
dots sync
```

## Commands

### Core Commands

**`dots init`** - Initialize dotfiles directory and git repository
```bash
dots init
```

**`dots add <file>`** - Add a dotfile to tracking
```bash
dots add ~/.bashrc
dots add ~/.config/nvim
```

**`dots remove <file>`** - Remove a dotfile from tracking
```bash
dots remove bashrc
```

**`dots link <file>`** - Create symlink for a dotfile
```bash
dots link bashrc
```

**`dots status`** - Check status of all dotfiles
```bash
dots status
```

**`dots edit <file>`** - Edit a dotfile
```bash
dots edit bashrc
```

### Git Commands

**`dots sync`** - Commit and push changes
```bash
dots sync
dots sync -m "Update vim config"
```

**`dots push`** - Push committed changes
```bash
dots push
```

**`dots pull`** - Pull changes from remote
```bash
dots pull
```

**`dots clone <url>`** - Clone existing dotfiles repository
```bash
dots clone https://github.com/username/dotfiles.git
```

## Directory Structure

After initialization:
```
~/.config/dots/
├── .git/
├── .gitignore
├── README.md
├── bashrc          # Your dotfiles
├── zshrc
└── nvim/
```

Your home directory:
```
~/
├── .bashrc -> ~/.config/dots/bashrc  # Symlinks
├── .zshrc -> ~/.config/dots/zshrc
└── .config/
    └── nvim -> ~/.config/dots/nvim
```














