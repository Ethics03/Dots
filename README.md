<div align="center">

# 📦 dots

### A minimal, fast, and developer-friendly dotfile manager

**Effortlessly manage, version, and sync your dotfiles using symlinks and Git**

[![Go Version](https://img.shields.io/badge/Go-1.23.1-00ADD8?style=flat&logo=go)](https://golang.org)
[![License](https://img.shields.io/badge/License-MIT-blue.svg)](LICENSE)
[![Platform](https://img.shields.io/badge/Platform-Linux%20%7C%20macOS-lightgrey)](https://github.com/Ethics03/Dots)

[Features](#-features) • [Installation](#-installation) • [Quick Start](#-quick-start) • [Commands](#-commands) • [Examples](#-examples)

</div>

---

## ✨ Features

- **🚀 Simple & Fast** - Minimal overhead, instant operations, no bloat
- **🔄 Git Integration** - Full version control with sync, push, pull commands
- **🔗 Symlink Management** - Automatic symlink creation and tracking
- **☁️ Cross-Machine Sync** - Clone and deploy dotfiles to new machines instantly
- **🛡️ Safe Operations** - Built-in checks to prevent data loss and recursive symlinks
- **📝 No Config Files** - Filesystem-based tracking, no YAML complexity
- **⚡ Written in Go** - Single binary, no dependencies, cross-platform compilation

---

## 📥 Installation

### Prerequisites

- **Go** 1.20+ (for building from source)
- **Git** (for version control features)

### Quick Install (Linux/macOS)

```bash
git clone https://github.com/Ethics03/Dots.git
cd Dots
chmod +x Install.sh
./Install.sh
```

The install script will:
- ✅ Check for Go and Git
- ✅ Build the binary
- ✅ Install to `/usr/local/bin`
- ✅ Verify installation

### Manual Installation

```bash
git clone https://github.com/Ethics03/Dots.git
cd Dots
go build -o dots
sudo mv dots /usr/local/bin
```

### Platform-Specific Notes

<details>
<summary><b>🐧 Linux</b></summary>

All features fully supported. Symlinks work natively.

```bash
./Install.sh
```
</details>

<details>
<summary><b>🍎 macOS</b></summary>

All features fully supported. May require `sudo` for installation.

```bash
./Install.sh
```

**Note:** If you get a security warning, go to System Preferences → Security & Privacy and allow the binary.
</details>

<details>
<summary><b>🪟 Windows</b></summary>

⚠️ **Limited Support** - Requires Developer Mode or Administrator privileges for symlinks.

**Option 1: Using WSL (Recommended)**
```bash
# Inside WSL
./Install.sh
```

**Option 2: Native Windows**
```powershell
git clone https://github.com/Ethics03/Dots.git
cd Dots
go build -o dots.exe
move dots.exe $HOME\bin\  # Add to PATH
```

**Enable Developer Mode for Symlinks:**
1. Settings → Update & Security → For Developers
2. Enable "Developer Mode"
3. Restart terminal
</details>

---

## 🚀 Quick Start

### Initialize Your Dotfiles

```bash
# Initialize dots (creates ~/.config/dots and git repo)
dots init
```

Output:
```
╔═══════════════════════════════════════════════════════════════╗
║                                                               ║
║                      Initializing Dots                        ║
║                                                               ║
╚═══════════════════════════════════════════════════════════════╝

Setting up dotfiles directory...
   ✓ Created /home/user/.config/dots

Creating configuration files...
   ✓ .gitignore
   ✓ README.md

Initializing git repository...
   ✓ Repository initialized

Creating initial commit...
   ✓ Committed initial files

╔═══════════════════════════════════════════════════════════════╗
║                                                               ║
║                     Setup Complete!                           ║
║                                                               ║
╚═══════════════════════════════════════════════════════════════╝
```

### Add Your First Dotfile

```bash
# Add a dotfile to tracking
dots add ~/.bashrc

# Add a directory
dots add ~/.config/nvim
```

### Check Status

```bash
dots status
```

### Sync to Remote

```bash
# Set up remote repository
cd ~/.config/dots
git remote add origin https://github.com/username/dotfiles.git

# Sync your dotfiles
dots sync
```

---

## 📚 Commands

### Core Commands

| Command | Description | Example |
|---------|-------------|---------|
| `dots init` | Initialize dotfiles directory and git repo | `dots init` |
| `dots add <file>` | Add a dotfile to tracking | `dots add ~/.bashrc` |
| `dots remove <file>` | Remove a dotfile from tracking | `dots remove bashrc` |
| `dots link <file>` | Create symlink for a dotfile | `dots link bashrc` |
| `dots status` | Check status of all dotfiles | `dots status` |
| `dots edit <file>` | Edit a dotfile using `$EDITOR` | `dots edit bashrc` |

### Git Commands

| Command | Description | Example |
|---------|-------------|---------|
| `dots sync` | Commit and push changes | `dots sync -m "Update config"` |
| `dots push` | Push committed changes | `dots push` |
| `dots pull` | Pull changes from remote | `dots pull` |
| `dots clone <url>` | Clone existing dotfiles repo | `dots clone git@github.com:user/dots.git` |

### Utility Commands

| Command | Description | Example |
|---------|-------------|---------|
| `dots create <file>` | Create a new dotfile | `dots create .zshrc` |
| `dots setup <path>` | Create directory structure | `dots setup nvim/lua/plugins` |

---

## 💡 Examples

### Complete Workflow

```bash
# 1. Initialize
dots init

# 2. Add your dotfiles
dots add ~/.bashrc
dots add ~/.zshrc
dots add ~/.config/nvim
dots add ~/.gitconfig

# 3. Check what's tracked
dots status

# 4. Set up remote backup
cd ~/.config/dots
git remote add origin git@github.com:yourusername/dotfiles.git

# 5. Sync to GitHub
dots sync -m "Initial dotfiles commit"
```

### Setting Up a New Machine

```bash
# Clone your dotfiles
dots clone git@github.com:yourusername/dotfiles.git

# Link what you need
dots link bashrc
dots link zshrc
dots link nvim

# Check everything is working
dots status
```

### Editing and Syncing

```bash
# Edit a dotfile
dots edit bashrc

# Sync changes to remote
dots sync -m "Add new aliases"
```

### Removing a Dotfile

```bash
# Remove from tracking (restores original file)
dots remove bashrc
```

---

## 📂 Directory Structure

### After Initialization

```
~/.config/dots/
├── .git/              # Git repository
├── .gitignore         # Ignore patterns
├── README.md          # Auto-generated README
├── bashrc             # Your dotfiles
├── zshrc
├── gitconfig
└── nvim/              # Nested configs supported
    ├── init.lua
    └── lua/
        └── plugins/
```

### Symlinks in Home Directory

```
~/
├── .bashrc -> ~/.config/dots/bashrc
├── .zshrc -> ~/.config/dots/zshrc
├── .gitconfig -> ~/.config/dots/gitconfig
└── .config/
    └── nvim -> ~/.config/dots/nvim
```

---

## 🛡️ Safety Features

- ✅ **Prevents recursive symlinks** - Won't add files from within `~/.config/dots`
- ✅ **Validates symlink targets** - Ensures symlinks point to correct locations
- ✅ **Backup on remove** - Restores original files when removing from tracking
- ✅ **Git stash on pull** - Automatically stashes uncommitted changes before pulling
- ✅ **Path validation** - Checks if files exist before operations
- ✅ **Permission preservation** - Maintains file permissions when copying

---

## 🌍 Platform Support

| Platform | Status | Notes |
|----------|--------|-------|
| **Linux** | ✅ Full Support | All features work perfectly |
| **macOS** | ✅ Full Support | All features work perfectly |
| **Windows** | ⚠️ Limited | Requires Developer Mode or Admin for symlinks |

### Windows Users

Symlinks on Windows require either:
- **Developer Mode** enabled (Windows 10+), OR
- Running as Administrator

**Recommended:** Use WSL (Windows Subsystem for Linux) for full compatibility.

---

## 🤝 Contributing

Contributions are welcome! Here's how you can help:

1. **Fork the repository**
2. **Create a feature branch** (`git checkout -b feature/amazing-feature`)
3. **Commit your changes** (`git commit -m 'Add amazing feature'`)
4. **Push to the branch** (`git push origin feature/amazing-feature`)
5. **Open a Pull Request**

### Development Setup

```bash
git clone https://github.com/Ethics03/Dots.git
cd Dots
go build -o dots
./dots --help
```

---

## 🐛 Troubleshooting

<details>
<summary><b>Symlink creation fails</b></summary>

**Issue:** `failed to create symlink: permission denied`

**Solution:**
- Linux/Mac: Check file permissions
- Windows: Enable Developer Mode or run as Administrator
</details>

<details>
<summary><b>Git commands fail</b></summary>

**Issue:** `git: command not found`

**Solution:** Install Git:
- Ubuntu/Debian: `sudo apt install git`
- macOS: `brew install git` or Xcode Command Line Tools
- Windows: Download from [git-scm.com](https://git-scm.com)
</details>

<details>
<summary><b>Remote push fails</b></summary>

**Issue:** `failed to push: authentication failed`

**Solution:**
- Set up SSH keys: `ssh-keygen -t ed25519 -C "your@email.com"`
- Add to GitHub: Settings → SSH Keys → Add new
- Or use HTTPS with personal access token
</details>

<details>
<summary><b>Status shows wrong symlinks</b></summary>

**Issue:** Symlinks point to wrong locations

**Solution:**
```bash
# Remove incorrect symlink
dots remove <filename>

# Re-add the dotfile
dots add ~/.<filename>
```
</details>

---

## 📄 License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

---

## 🙏 Acknowledgments

- Built with [Cobra](https://github.com/spf13/cobra) - CLI framework
- Inspired by [GNU Stow](https://www.gnu.org/software/stow/), [yadm](https://yadm.io/), and [chezmoi](https://www.chezmoi.io/)

---

## 📞 Support

- 🐛 **Bug Reports:** [GitHub Issues](https://github.com/Ethics03/Dots/issues)
- 💡 **Feature Requests:** [GitHub Issues](https://github.com/Ethics03/Dots/issues)
- 📖 **Documentation:** [GitHub Wiki](https://github.com/Ethics03/Dots/wiki)

---

<div align="center">

**Made with ❤️ by [Ethics03](https://github.com/Ethics03)**

⭐ Star this repo if you find it useful!

</div>
