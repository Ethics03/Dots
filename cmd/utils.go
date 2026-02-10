package cmd

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
)

// error to stop walk early
var errFound = fmt.Errorf("found")

// findDotfile finds a dotfile in the dots directory and returns its path
// along with the corresponding home path where the symlink should be
func findDotfile(filename string) (dotsPath string, homePath string, err error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", "", fmt.Errorf("cannot find home directory: %w", err)
	}

	dotsDir := filepath.Join(home, ".config", "dots")
	filename = filepath.Clean(filename)
	baseName := filepath.Base(filename)

	// First, try with just the basename (for files in home root)
	tryPath := filepath.Join(dotsDir, baseName)

	if _, err := os.Stat(tryPath); err == nil {
		// Found it with basename
		return tryPath, filepath.Join(home, baseName), nil
	}

	// Not found with basename, walk dots directory to find it
	var foundPath string
	walkErr := filepath.WalkDir(dotsDir, func(path string, d os.DirEntry, err error) error {
		if err != nil {
			// Return error to stop walk on failures
			return err
		}

		// Skip .git directory entirely
		if d.IsDir() && d.Name() == ".git" {
			return filepath.SkipDir
		}

		// Skip directories, only check files
		if d.IsDir() {
			return nil
		}

		// Check if this file matches our target
		if d.Name() == baseName {
			foundPath = path
			return errFound // Stop walk immediately
		}

		return nil
	})

	// Check if we found the file
	if walkErr == errFound {
		// Get the relative path from dots directory
		relPath, err := filepath.Rel(dotsDir, foundPath)
		if err != nil {
			return "", "", fmt.Errorf("failed to determine relative path: %w", err)
		}
		return foundPath, filepath.Join(home, relPath), nil
	}

	// Check for actual errors
	if walkErr != nil {
		return "", "", fmt.Errorf("error walking dots directory: %w", walkErr)
	}

	// File not found
	return "", "", fmt.Errorf("'%s' is not tracked by dots", filename)
}

func copyFile(src, dst string) error {
	sourceFile, err := os.Open(src)
	if err != nil {
		return err
	}
	defer sourceFile.Close()

	destFile, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer destFile.Close()

	if _, err := io.Copy(destFile, sourceFile); err != nil {
		return err
	}

	// Preserve permissions
	srcInfo, err := os.Stat(src)
	if err != nil {
		return err
	}
	return os.Chmod(dst, srcInfo.Mode())
}

// copyDir recursively copies a directory
func copyDir(src, dst string) error {
	srcInfo, err := os.Stat(src)
	if err != nil {
		return err
	}

	if err := os.MkdirAll(dst, srcInfo.Mode()); err != nil {
		return err
	}

	entries, err := os.ReadDir(src)
	if err != nil {
		return err
	}

	for _, entry := range entries {
		srcPath := filepath.Join(src, entry.Name())
		dstPath := filepath.Join(dst, entry.Name())

		if entry.IsDir() {
			if err := copyDir(srcPath, dstPath); err != nil {
				return err
			}
		} else {
			if err := copyFile(srcPath, dstPath); err != nil {
				return err
			}
		}
	}

	return nil
}
