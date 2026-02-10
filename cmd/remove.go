package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
)

// removeCmd represents the remove command
var removeCmd = &cobra.Command{
	Use:   "remove <file>",
	Short: "Remove a dotfile from tracking and restore the original",
	Long: `Remove a dotfile from the dots directory, delete the symlink,
and restore the original file to its location.

This command will:
  - Remove the symlink from the original location
  - Copy the file back from ~/.config/dots
  - Delete the file from ~/.config/dots

Example:
  dots remove bashrc
  dots remove .zshrc
  dots remove .config/nvim`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		filename := args[0]

		if err := removeDotfile(filename); err != nil {
			fmt.Printf("Error: %v\n", err)
			os.Exit(1)
		}
	},
}

func init() {
	rootCmd.AddCommand(removeCmd)
}

func removeDotfile(filename string) error {
	home, err := os.UserHomeDir()
	if err != nil {
		return fmt.Errorf("cannot find home directory: %w", err)
	}

	dotsDir := filepath.Join(home, ".config", "dots")

	// Clean the filename (remove leading ./ or ~/)
	filename = filepath.Clean(filename)
	filename = filepath.Base(filename)

	// Path in dots directory
	dotsPath := filepath.Join(dotsDir, filename)

	// Check if file exists in dots directory
	dotsInfo, err := os.Stat(dotsPath)
	if err != nil {
		if os.IsNotExist(err) {
			return fmt.Errorf("'%s' is not tracked by dots (not found in %s)", filename, dotsDir)
		}
		return fmt.Errorf("failed to access dots file: %w", err)
	}

	// Determine the home path (where symlink should be)
	homePath := filepath.Join(home, filename)

	// Check if symlink exists
	linkInfo, err := os.Lstat(homePath)
	if err != nil {
		if os.IsNotExist(err) {
			// Symlink doesn't exist, just remove from dots directory
			fmt.Printf("⚠ Warning: No symlink found at %s\n", homePath)
			fmt.Println("Removing from dots directory only...")
		} else {
			return fmt.Errorf("failed to check symlink: %w", err)
		}
	} else {
		// Check if it's actually a symlink
		if linkInfo.Mode()&os.ModeSymlink != 0 {
			// Verify it points to our dots file
			linkTarget, err := os.Readlink(homePath)
			if err != nil {
				return fmt.Errorf("failed to read symlink: %w", err)
			}

			absLinkTarget, _ := filepath.Abs(linkTarget)
			absDotsPath, _ := filepath.Abs(dotsPath)

			if absLinkTarget != absDotsPath {
				return fmt.Errorf("symlink at %s points to %s, not %s\nManual intervention required",
					homePath, linkTarget, dotsPath)
			}

			// Remove the symlink
			if err := os.Remove(homePath); err != nil {
				return fmt.Errorf("failed to remove symlink: %w", err)
			}
			fmt.Printf("✓ Removed symlink: %s\n", homePath)

			// Copy file/directory back from dots to home
			if dotsInfo.IsDir() {
				if err := copyDir(dotsPath, homePath); err != nil {
					return fmt.Errorf("failed to restore directory: %w", err)
				}
				fmt.Printf("✓ Restored directory: %s\n", homePath)
			} else {
				if err := copyFile(dotsPath, homePath); err != nil {
					return fmt.Errorf("failed to restore file: %w", err)
				}
				fmt.Printf("✓ Restored file: %s\n", homePath)
			}
		} else {
			// Not a symlink, something else exists there
			return fmt.Errorf("file at %s exists but is not a symlink\nManual intervention required", homePath)
		}
	}

	// Remove from dots directory
	if err := os.RemoveAll(dotsPath); err != nil {
		return fmt.Errorf("failed to remove from dots directory: %w", err)
	}
	fmt.Printf("✓ Removed from dots directory: %s\n", dotsPath)

	fmt.Println("\n✓ Dotfile removed successfully!")
	return nil
}
