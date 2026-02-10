package cmd

import (
	"fmt"
	"os"
	"os/user"
	"path/filepath"

	"github.com/spf13/cobra"
)

// statusCmd represents the status command
var statusCmd = &cobra.Command{
	Use:   "status",
	Short: "A command to check the status of the dots folder",
	Long: `Helps you in checking the current status of all the symlinks and the files
	connected through those symlinks to your dotfiles`,
	Run: func(cmd *cobra.Command, args []string) {
		usr, _ := user.Current()
		dotDr := filepath.Join(usr.HomeDir, ".config", "dots")

		// Check if dots directory exists
		if _, err := os.Stat(dotDr); os.IsNotExist(err) {
			fmt.Println("Dots directory not found. Run 'dots init' first.")
			return
		}

		fmt.Println("Dotfiles status: ")
		fmt.Printf("%-40s  ->  %s\n", "Dotfile (home)", "Target (./.config/dots folder)")

		// Walk the dots directory to find all dotfiles (including nested ones)
		filepath.Walk(dotDr, func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return nil
			}

			// Skip directories
			if info.IsDir() {
				return nil
			}

			filename := info.Name()
			// Skip git directory and meta files
			if filename == ".git" || filename == ".gitignore" || filename == "README.md" {
				return nil
			}

			// Get relative path from dots directory
			relPath, err := filepath.Rel(dotDr, path)
			if err != nil {
				return nil
			}

			// Skip if inside .git directory
			if len(relPath) > 4 && relPath[:4] == ".git" {
				return nil
			}

			dotPath := path
			homePath := filepath.Join(usr.HomeDir, relPath)

			link, err := os.Readlink(homePath)
			if err != nil {
				if os.IsNotExist(err) {
					fmt.Printf("%-40s  ->  %s\n", "Missing symlink: "+homePath, dotPath)
				} else {
					fmt.Printf("%-40s  ->  %s\n", "Not a symlink or unreadable: "+homePath, "")
				}
				return nil
			}
			absTarget, _ := filepath.Abs(dotPath)
			absLink, _ := filepath.Abs(link)

			if absTarget == absLink {
				fmt.Printf("%-40s  ->  %s\n", "Status ok: "+homePath, link)
			} else {
				fmt.Printf("Wrong target: %s -> %s (expected %s)\n", homePath, link, dotPath)
			}

			return nil
		})
	},
}

func init() {
	rootCmd.AddCommand(statusCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// statusCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// statusCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
