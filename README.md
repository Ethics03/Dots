# dots

**dots** is a minimal, fast, and developer-friendly dotfile manager built in Go. It helps you effortlessly manage, version, and sync your dotfiles using symlinks and Git, without the complexity of bloated tools or manual setup.

Designed with simplicity in mind, dots lets you keep your config files organized in a single directory, backed by version control, and easily portable across machines. Whether you're setting up a new dev environment or just want clean dotfile hygiene, dots makes it painless.

##  Features

-  Automatically symlink your dotfiles from a central location.
-  Simple `add`, `link`, `init`, and `sync` commands.
-  Built with Go — fast, easy to maintain.
-  Git integration for remote backup and sharing


##  Getting Started

```bash
git clone https://github.com/Ethics03/Dots.git
cd dots
go build -o dots


Optional Move the binary to a directory in your $PATH:

For easier access, you can move the compiled binary to /usr/local/bin:

sudo mv dots /usr/local/bin


Directory Structure: 

~/.config/.dots/
├── bashrc
├── zshrc
└── vimrc

Important: These filenames should not start with a dot (e.g., bashrc instead of .bashrc). The symlink will handle the dot for you.










