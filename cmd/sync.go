/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"time"

	"github.com/spf13/cobra"
)

var syncMessage string

// syncCmd represents the sync command
var syncCmd = &cobra.Command{
	Use:   "sync",
	Short: "Commit and push your dotfiles to the remote repository",
	Long: `Sync your dotfiles by committing all changes and pushing to the remote repository.

This command will:
  - Stage all changes in ~/.config/dots
  - Commit with a message (auto-generated or custom)
  - Push to the remote repository

Example:
  dots sync                           # Auto-generated commit message
  dots sync -m "Update vim config"    # Custom commit message`,
	Run: func(cmd *cobra.Command, args []string) {
		if err := syncDotfiles(); err != nil {
			fmt.Printf("Error: %v\n", err)
			os.Exit(1)
		}
	},
}

func init() {
	rootCmd.AddCommand(syncCmd)
	syncCmd.Flags().StringVarP(&syncMessage, "message", "m", "", "Commit message")
}

func syncDotfiles() error {
	home, err := os.UserHomeDir()
	if err != nil {
		return fmt.Errorf("cannot find home directory: %w", err)
	}

	dotsDir := filepath.Join(home, ".config", "dots")

	// Check if dots directory exists
	if _, err := os.Stat(dotsDir); os.IsNotExist(err) {
		return fmt.Errorf("dots directory not found. Run 'dots init' first")
	}

	// Check if it's a git repository
	gitDir := filepath.Join(dotsDir, ".git")
	if _, err := os.Stat(gitDir); os.IsNotExist(err) {
		return fmt.Errorf("not a git repository. Run 'dots init' to initialize")
	}

	fmt.Println("Checking git status...")

	// Check if there are any changes
	statusCmd := exec.Command("git", "status", "--porcelain")
	statusCmd.Dir = dotsDir
	output, err := statusCmd.Output()
	if err != nil {
		return fmt.Errorf("failed to check git status: %w", err)
	}

	if len(output) == 0 {
		fmt.Println("✓ No changes to sync")
		return nil
	}

	fmt.Println("Changes detected. Staging files...")

	// Stage all changes
	addCmd := exec.Command("git", "add", "-A")
	addCmd.Dir = dotsDir
	if output, err := addCmd.CombinedOutput(); err != nil {
		return fmt.Errorf("failed to stage files: %w\n%s", err, output)
	}
	fmt.Println("✓ Files staged")

	// Generate commit message if not provided
	if syncMessage == "" {
		syncMessage = fmt.Sprintf("Update dotfiles - %s", time.Now().Format("2006-01-02 15:04:05"))
	}

	fmt.Printf("Committing with message: \"%s\"\n", syncMessage)

	// Commit changes
	commitCmd := exec.Command("git", "commit", "-m", syncMessage)
	commitCmd.Dir = dotsDir
	if output, err := commitCmd.CombinedOutput(); err != nil {
		return fmt.Errorf("failed to commit: %w\n%s", err, output)
	}
	fmt.Println("✓ Changes committed")

	// Check if remote is configured
	remoteCmd := exec.Command("git", "remote", "get-url", "origin")
	remoteCmd.Dir = dotsDir
	remoteOutput, err := remoteCmd.Output()

	if err != nil || len(remoteOutput) == 0 {
		fmt.Println("\n⚠ No remote repository configured")
		fmt.Println("To add a remote:")
		fmt.Printf("  cd %s\n", dotsDir)
		fmt.Println("  git remote add origin <repository-url>")
		fmt.Println("  git push -u origin main")
		return nil
	}

	fmt.Println("Pushing to remote...")

	// Push to remote
	pushCmd := exec.Command("git", "push")
	pushCmd.Dir = dotsDir
	pushCmd.Stdout = os.Stdout
	pushCmd.Stderr = os.Stderr
	if err := pushCmd.Run(); err != nil {
		return fmt.Errorf("failed to push: %w", err)
	}

	fmt.Println("\n✓ Dotfiles synced successfully!")
	return nil
}
