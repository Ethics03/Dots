package cmd

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/spf13/cobra"
)

// pushCmd represents the push command
var pushCmd = &cobra.Command{
	Use:   "push",
	Short: "Push committed changes to the remote repository",
	Long: `Push your committed dotfiles to the remote repository.

Note: This only pushes already committed changes.
Use 'dots sync' to commit and push in one step.

Example:
  dots push`,
	Run: func(cmd *cobra.Command, args []string) {
		if err := pushDotfiles(); err != nil {
			fmt.Printf("Error: %v\n", err)
			os.Exit(1)
		}
	},
}

func init() {
	rootCmd.AddCommand(pushCmd)
}

func pushDotfiles() error {
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

	// Check if remote is configured
	remoteCmd := exec.Command("git", "remote", "get-url", "origin")
	remoteCmd.Dir = dotsDir
	output, err := remoteCmd.Output()

	if err != nil || len(output) == 0 {
		return fmt.Errorf("no remote repository configured\nAdd a remote with: cd %s && git remote add origin <url>", dotsDir)
	}

	// Check for uncommitted changes
	statusCmd := exec.Command("git", "status", "--porcelain")
	statusCmd.Dir = dotsDir
	statusOutput, err := statusCmd.Output()
	if err != nil {
		return fmt.Errorf("failed to check git status: %w", err)
	}

	if len(statusOutput) > 0 {
		fmt.Println("Warning: You have uncommitted changes")
		fmt.Println("Use 'dots sync' to commit and push, or commit manually first")
		return fmt.Errorf("uncommitted changes detected")
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

	fmt.Println("\n✓ Changes pushed successfully!")
	return nil
}
