/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add <file>",
	Short: "Add a dotfile to tracking and create a symlink",
	Long: `Add a dotfile to the dots directory and create a symlink from the original location.

Example:
  dots add ~/.bashrc        # Add bashrc to tracking
  dots add ~/.config/nvim   # Add entire nvim config directory
  dots add .zshrc           # Add from current directory`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		filePath := args[0]

		if err := addDotfile(filePath); err != nil {
			fmt.Printf("Error: %v\n", err)
			os.Exit(1)
		}
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}

func addDotfile(filePath string) error {
	// Get home directory
	home, err := os.UserHomeDir()
	if err != nil {
		return fmt.Errorf("cannot find home directory: %w", err)
	}

	// Resolve to absolute path
	absPath, err := filepath.Abs(filePath)
	if err != nil {
		return fmt.Errorf("cannot resolve path: %w", err)
	}

	// Check if source file/directory exists
	srcInfo, err := os.Lstat(absPath)
	if err != nil {
		return fmt.Errorf("source does not exist: %s", absPath)
	}

	// Check if it's already a symlink
	if srcInfo.Mode()&os.ModeSymlink != 0 {
		return fmt.Errorf("source is already a symlink: %s", absPath)
	}

	// Get dots directory path
	dotsDir := filepath.Join(home, ".config", "dots")

	// Check if file is already inside dots directory (prevent recursive symlinks)
	if strings.HasPrefix(absPath, dotsDir+string(filepath.Separator)) || absPath == dotsDir {
		return fmt.Errorf("cannot add files from within dots directory: %s", absPath)
	}

	// check base name
	baseName := filepath.Base(absPath)

	// preserving structure of sub-dirs
	relPath, err := filepath.Rel(home, absPath)
	var dotsPath string

	if err == nil && !filepath.IsAbs(relPath) && relPath != ".." && !strings.HasPrefix(relPath, ".."+string(filepath.Separator)) {
		// preserve structure
		dotsPath = filepath.Join(dotsDir, relPath)
	} else {
		// file outside home dir, preserve base name
		dotsPath = filepath.Join(dotsDir, baseName)
	}

	// Create parent directory if needed
	dotsParent := filepath.Dir(dotsPath)
	if err := os.MkdirAll(dotsParent, 0o755); err != nil {
		return fmt.Errorf("failed to create parent directory: %w", err)
	}

	// Check if destination already exists
	if _, err := os.Lstat(dotsPath); err == nil {
		return fmt.Errorf("dotfile already exists in dots directory: %s", dotsPath)
	}

	// Copy file or directory to dots directory
	if srcInfo.IsDir() {
		if err := copyDir(absPath, dotsPath); err != nil {
			return fmt.Errorf("failed to copy directory: %w", err)
		}
		fmt.Printf("Copied directory: %s -> %s\n", absPath, dotsPath)
	} else {
		if err := copyFile(absPath, dotsPath); err != nil {
			return fmt.Errorf("failed to copy file: %w", err)
		}
		fmt.Printf("Copied file: %s -> %s\n", absPath, dotsPath)
	}

	// Remove original file/directory
	if err := os.RemoveAll(absPath); err != nil {
		return fmt.Errorf("failed to remove original: %w", err)
	}

	// creating symlink
	if err := os.Symlink(dotsPath, absPath); err != nil {
		return fmt.Errorf("failed to create symlink: %w", err)
	}

	fmt.Printf("Created symlink: %s -> %s\n", absPath, dotsPath)
	fmt.Println("✓ Dotfile added successfully!")

	return nil
}
