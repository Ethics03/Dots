package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// linkCmd represents the link command
var linkCmd = &cobra.Command{
	Use:   "link",
	Short: "Create symlinks for tracked dotfiles.",
	Long:  `Creates symbolic links from your dotfiles repo to their original paths.`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 1 {
			fmt.Println("Usage: dots link <dotfile>")
			fmt.Println("Example: dots link bashrc")
			return
		}

		name := args[0]

		// Find the dotfile in dots directory
		src, desti, err := findDotfile(name)
		if err != nil {
			fmt.Printf("Error: %v\n", err)
			return
		}

		if _, err := os.Lstat(desti); err == nil {
			fmt.Printf("Destination already exists: %s (skipping)\n", desti)
			return
		}

		err = os.Symlink(src, desti)
		if err != nil {
			fmt.Printf("Failed to link: %s -> %s: %v\n", src, desti, err)
		} else {
			fmt.Printf("Linked %s -> %s\n", src, desti)
		}
	},
}

func init() {
	rootCmd.AddCommand(linkCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// linkCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// linkCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
