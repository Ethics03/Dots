package cmd

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/spf13/cobra"
)

// pullCmd represents the pull command
var pullCmd = &cobra.Command{
	Use:   "pull",
	Short: "Pull the latest changes from the remote repository",
	Long: `Pull the latest dotfiles from the remote repository.

This command will:
  - Fetch changes from the remote repository
  - Merge them into your local dotfiles

Example:
  dots pull`,
	Run: func(cmd *cobra.Command, args []string) {
		if err := pullDotfiles(); err != nil {
			fmt.Printf("Error: %v\n", err)
			os.Exit(1)
		}
	},
}

func init() {
	rootCmd.AddCommand(pullCmd)
}

func pullDotfiles() error {
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
	if output, err := remoteCmd.Output(); err != nil || len(output) == 0 {
		return fmt.Errorf("no remote repository configured\nAdd a remote with: cd %s && git remote add origin <url>", dotsDir)
	}

	fmt.Println("Pulling changes from remote...")

	// Check for uncommitted changes
	statusCmd := exec.Command("git", "status", "--porcelain")
	statusCmd.Dir = dotsDir
	output, err := statusCmd.Output()
	if err != nil {
		return fmt.Errorf("failed to check git status: %w", err)
	}

	if len(output) > 0 {
		fmt.Println("Warning: You have uncommitted changes")
		fmt.Println("Stashing changes before pull...")

		// Stash changes
		stashCmd := exec.Command("git", "stash", "push", "-m", "Auto-stash before pull")
		stashCmd.Dir = dotsDir
		if output, err := stashCmd.CombinedOutput(); err != nil {
			return fmt.Errorf("failed to stash changes: %w\n%s", err, output)
		}
		fmt.Println("✓ Changes stashed")

		defer func() {
			fmt.Println("\nApplying stashed changes...")
			popCmd := exec.Command("git", "stash", "pop")
			popCmd.Dir = dotsDir
			if output, err := popCmd.CombinedOutput(); err != nil {
				fmt.Printf("Warning: Failed to apply stashed changes: %v\n%s\n", err, output)
				fmt.Println("You can manually apply them with: cd ~/.config/dots && git stash pop")
			} else {
				fmt.Println("✓ Stashed changes applied")
			}
		}()
	}

	// Pull changes
	pullCmd := exec.Command("git", "pull")
	pullCmd.Dir = dotsDir
	pullCmd.Stdout = os.Stdout
	pullCmd.Stderr = os.Stderr
	if err := pullCmd.Run(); err != nil {
		return fmt.Errorf("failed to pull: %w", err)
	}

	fmt.Println("\n✓ Dotfiles pulled successfully!")
	fmt.Println("\nNote: You may need to run 'dots status' to check symlink status")
	return nil
}
