/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/spf13/cobra"
)

// editCmd represents the edit command
var editCmd = &cobra.Command{
	Use:   "edit",
	Short: "edit command lets you edit the dotfile mentioned",
	Long: `Provide the filepath/filename of the dotfile you want to edit, it won't 
	apply any changes until you do "dots save"`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		fileEdit := args[0]

		home, err := os.UserHomeDir()
		if err != nil {
			fmt.Println("Could not find home directory")
			return
		}

		dotPath := filepath.Join(home, ".config", ".dots", fileEdit)

		editor := os.Getenv("EDITOR")
		if editor == "" {
			editor = "nano" //fallback EDITOR
		}

		editcmd := exec.Command(editor, dotPath)
		editcmd.Stdin = os.Stdin
		editcmd.Stdout = os.Stdout
		editcmd.Stderr = os.Stderr

		if err := editcmd.Run(); err != nil {
			fmt.Println("failed to open editor: ", err)
		}

	},
}

func init() {
	rootCmd.AddCommand(editCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// editCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// editCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
