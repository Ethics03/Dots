package cmd

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/spf13/cobra"
)

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize the dots directory and git repository",
	Long: `Initialize your dotfiles management system by creating the necessary
directories, configuration files, and git repository.

This command will:
  - Create ~/.config/dots directory
  - Initialize a git repository
  - Create .gitignore file
  - Create README.md
  - Make initial commit

Example:
  dots init`,
	Run: func(cmd *cobra.Command, args []string) {
		if err := initializeDots(); err != nil {
			fmt.Printf("Error: %v\n", err)
			os.Exit(1)
		}
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
}

func initializeDots() error {
	home, err := os.UserHomeDir()
	if err != nil {
		return fmt.Errorf("cannot find home directory: %w", err)
	}

	dotsDir := filepath.Join(home, ".config", "dots")

	// Check if dots directory already exists
	if _, err := os.Stat(dotsDir); err == nil {
		return fmt.Errorf("dots directory already exists at %s\nUse 'dots status' to check your dotfiles", dotsDir)
	}

	// Print beautiful header
	fmt.Println()
	fmt.Println("╔═══════════════════════════════════════════════════════════════╗")
	fmt.Println("║                                                               ║")
	fmt.Println("║                      Initializing Dots                        ║")
	fmt.Println("║                                                               ║")
	fmt.Println("╚═══════════════════════════════════════════════════════════════╝")
	fmt.Println()

	// Create dots directory
	fmt.Println("Setting up dotfiles directory...")
	if err := os.MkdirAll(dotsDir, 0o755); err != nil {
		return fmt.Errorf("failed to create dots directory: %w", err)
	}
	fmt.Printf("   ✓ Created %s\n", dotsDir)
	fmt.Println()

	// Create .gitignore
	fmt.Println("Creating configuration files...")
	gitignoreContent := `# Ignore backup files
*.backup
*.bak
*.swp
*.tmp

# Ignore OS files
.DS_Store
Thumbs.db

# Ignore editor files
.vscode/
.idea/
*.sublime-*
`
	gitignorePath := filepath.Join(dotsDir, ".gitignore")
	if err := os.WriteFile(gitignorePath, []byte(gitignoreContent), 0o644); err != nil {
		return fmt.Errorf("failed to create .gitignore: %w", err)
	}
	fmt.Println("   ✓ .gitignore")

	// Create README.md
	readmeContent := `# My Dotfiles

This repository contains my personal dotfiles managed with [dots](https://github.com/Ethics03/Dots).

## Setup

To set up these dotfiles on a new machine:

` + "```bash" + `
# Clone this repository
git clone <your-repo-url> ~/.config/dots

# Link the dotfiles
dots link bashrc
dots link nvim
` + "```" + `

## Tracked Files

All files in this directory (except .git, .gitignore, and README.md) are tracked dotfiles.
`
	readmePath := filepath.Join(dotsDir, "README.md")
	if err := os.WriteFile(readmePath, []byte(readmeContent), 0o644); err != nil {
		return fmt.Errorf("failed to create README.md: %w", err)
	}
	fmt.Println("   ✓ README.md")
	fmt.Println()

	// init git repo
	fmt.Println("Initializing git repository...")
	gitInitCmd := exec.Command("git", "init")
	gitInitCmd.Dir = dotsDir
	if output, err := gitInitCmd.CombinedOutput(); err != nil {
		return fmt.Errorf("failed to initialize git: %w\n%s", err, output)
	}
	fmt.Println("   ✓ Repository initialized")
	fmt.Println()

	// initial commit
	fmt.Println("Creating initial commit...")

	// Stage all files
	gitAddCmd := exec.Command("git", "add", ".")
	gitAddCmd.Dir = dotsDir
	if output, err := gitAddCmd.CombinedOutput(); err != nil {
		return fmt.Errorf("failed to stage files: %w\n%s", err, output)
	}

	// Commit
	gitCommitCmd := exec.Command("git", "commit", "-m", "Initial commit: dots setup")
	gitCommitCmd.Dir = dotsDir
	if output, err := gitCommitCmd.CombinedOutput(); err != nil {
		return fmt.Errorf("failed to create initial commit: %w\n%s", err, output)
	}
	fmt.Println("   ✓ Committed initial files")
	fmt.Println()

	// Success message
	fmt.Println("╔═══════════════════════════════════════════════════════════════╗")
	fmt.Println("║                                                               ║")
	fmt.Println("║                     Setup Complete!                           ║")
	fmt.Println("║                                                               ║")
	fmt.Println("╚═══════════════════════════════════════════════════════════════╝")
	fmt.Println()
	fmt.Printf("Dotfiles directory: %s\n", dotsDir)
	fmt.Println()
	fmt.Println("┌─────────────────────────────────────────────────────────────┐")
	fmt.Println("│  Next Steps                                                 │")
	fmt.Println("├─────────────────────────────────────────────────────────────┤")
	fmt.Println("│                                                             │")
	fmt.Println("│  1. Add your first dotfile                                  │")
	fmt.Println("│     $ dots add ~/.bashrc                                    │")
	fmt.Println("│                                                             │")
	fmt.Println("│  2. Check status                                            │")
	fmt.Println("│     $ dots status                                           │")
	fmt.Println("│                                                             │")
	fmt.Println("│  3. Edit a dotfile                                          │")
	fmt.Println("│     $ dots edit bashrc                                      │")
	fmt.Println("│                                                             │")
	fmt.Println("│  4. Add a remote repository                                 │")
	fmt.Println("│     $ cd ~/.config/dots                                     │")
	fmt.Println("│     $ git remote add origin <your-repo-url>                 │")
	fmt.Println("│                                                             │")
	fmt.Println("│  5. Sync to remote                                          │")
	fmt.Println("│     $ dots sync                                             │")
	fmt.Println("│                                                             │")
	fmt.Println("└─────────────────────────────────────────────────────────────┘")
	fmt.Println()

	return nil
}
