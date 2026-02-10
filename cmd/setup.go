package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
)

// setupCmd represents the setup command
var setupCmd = &cobra.Command{
	Use:   "setup",
	Short: "Setup your folder structure for your config files",
	Long: `This command helps you to create your folder structure for your configuration files


For example: dot setup /nvim/lua/plugins -> sets up ~/.config/dots/.config/nvim/lua/plugins

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		setupDirStructure(args[0])
	},
}

func init() {
	rootCmd.AddCommand(setupCmd)
}

func setupDirStructure(path string) error {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return fmt.Errorf("failed to get home dir: %w ", err)
	}

	fullpath := filepath.Join(homeDir, ".config", "dots", ".config", path)

	if _, err := os.Stat(fullpath); err == nil {
		fmt.Printf("Directory already exists: ~/.config/dots/.config/%s\n", path)
		return nil
	}

	if err := os.MkdirAll(fullpath, 0o755); err != nil {
		return fmt.Errorf("Failed to create directory: %w\n", path)
	}

	fmt.Printf("Setup Done: ~/.config/dots/.config/%s\n", path)
	return nil
}
