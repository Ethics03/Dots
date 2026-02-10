package cmd

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/spf13/cobra"
)

// cloneCmd represents the clone command
var cloneCmd = &cobra.Command{
	Use:   "clone <repository-url>",
	Short: "Clone an existing dotfiles repository",
	Long: `Clone an existing dotfiles repository to ~/.config/dots.

This command will:
  - Clone the repository to ~/.config/dots
  - Show you the files available for linking

After cloning, use 'dots link <file>' to create symlinks for your dotfiles.

Example:
  dots clone https://github.com/username/dotfiles.git
  dots clone git@github.com:username/dotfiles.git`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		repoURL := args[0]

		if err := cloneDotfiles(repoURL); err != nil {
			fmt.Printf("Error: %v\n", err)
			os.Exit(1)
		}
	},
}

func init() {
	rootCmd.AddCommand(cloneCmd)
}

func cloneDotfiles(repoURL string) error {
	home, err := os.UserHomeDir()
	if err != nil {
		return fmt.Errorf("cannot find home directory: %w", err)
	}

	dotsDir := filepath.Join(home, ".config", "dots")

	// Check if dots directory already exists
	if _, err := os.Stat(dotsDir); err == nil {
		return fmt.Errorf("dots directory already exists at %s\nRemove it first or use 'dots pull' to update", dotsDir)
	}

	fmt.Printf("Cloning dotfiles from %s...\n", repoURL)

	// Clone the repository
	cloneCmd := exec.Command("git", "clone", repoURL, dotsDir)
	cloneCmd.Stdout = os.Stdout
	cloneCmd.Stderr = os.Stderr
	if err := cloneCmd.Run(); err != nil {
		return fmt.Errorf("failed to clone repository: %w", err)
	}

	fmt.Println("\n✓ Repository cloned successfully!")

	// List available dotfiles
	fmt.Println("\nAvailable dotfiles:")
	files, err := os.ReadDir(dotsDir)
	if err != nil {
		return fmt.Errorf("failed to read dots directory: %w", err)
	}

	dotfileCount := 0
	for _, file := range files {
		name := file.Name()
		// Skip git directory, README, and other meta files
		if name == ".git" || name == ".gitignore" || name == "README.md" {
			continue
		}

		fileType := "file"
		if file.IsDir() {
			fileType = "directory"
		}
		fmt.Printf("  - %s (%s)\n", name, fileType)
		dotfileCount++
	}

	if dotfileCount == 0 {
		fmt.Println("  (No dotfiles found)")
	}

	fmt.Println("\nNext steps:")
	fmt.Println("  1. Link individual files:     dots link <filename>")
	fmt.Println("  2. Check status:              dots status")
	fmt.Println("  3. Edit a dotfile:            dots edit <filename>")

	return nil
}
