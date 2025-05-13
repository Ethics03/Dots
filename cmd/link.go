/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
)

// linkCmd represents the link command
var linkCmd = &cobra.Command{
	Use:   "link",
	Short: "Create symlinks for tracked dotfiles.",
	Long:  `Creates symbolic links from your dotfiles repo to their original paths.`,
	Run: func(cmd *cobra.Command, args []string) {

		dotfiles := map[string]string{
			"bashrc": ".bashrc",
			"zshrc":  ".zshrc",
		}

		home, err := os.UserHomeDir()
		if err != nil {
			log.Fatal("Cannot find home directory", err)
		}

		dotsDir := filepath.Join(home, ".config", ".dots")

		for name, target := range dotfiles {
			src := filepath.Join(dotsDir, name)
			desti := filepath.Join(home, target)

			err := os.Symlink(src, desti) //this is creating the symlink (./zshrc symlink pointing to zshrc in ./.config/.dots dir)
			if err != nil {
				fmt.Printf("Failed to link : %s -> %s: %v", src, desti, err)
			} else {
				fmt.Printf("Linked %s -> %s", src, desti) // this is printing the output
			}
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
