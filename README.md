# dots

**dots** is a minimal, fast, and developer-friendly dotfile manager built in Go. It helps you effortlessly manage, version, and sync your dotfiles using symlinks and Git, without the complexity of bloated tools or manual setup.

Designed with simplicity in mind, dots lets you keep your config files organized in a single directory, backed by version control, and easily portable across machines. Whether you're setting up a new dev environment or just want clean dotfile hygiene, dots makes it painless.

## Features

- Automatically symlink your dotfiles from a central location.
- Simple `add`, `link`, `init`, `edit` and `sync` commands.
- Built with Go — fast, easy to maintain.
- Git integration for remote backup and sharing.

## Getting Started

### Clone the Repository and Build the Binary

```bash
git clone https://github.com/Ethics03/Dots.git
cd dots
go build -o dots
```

# Optional: Move the Binary to a Directory in Your $PATH

```bash
sudo mv dots /usr/local/bin
```

# Directory Structure

```bash
~/.config/.dots/
├── bashrc
├── zshrc
└── vimrc
```

# Commands 


Initialize your dotfiles repository. This sets up the .dots directory and prepares it for managing your dotfiles.
```bash
dots init
```

link makes a symlink with the mentioned dotfile in your .config/.dots/ directory to the main configuration in $HOME
```bash
dots link <dotfile to symlink> eg: bashrc 
```


Lets you edit your dotfile in the .dots directory
```bash
dots edit
```














