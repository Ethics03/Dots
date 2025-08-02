/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>

*/
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
		usr,_ := user.Current()
		dotDr := filepath.Join(usr.HomeDir,".config","dots")

		files, err := os.ReadDir(dotDr)
		if err != nil { 
				fmt.Println("Failed to read ./config/dots: ",err)
				return
		}

		fmt.Println("Dotfiles status: ")
    fmt.Printf("%-40s  ->  %s\n", "Dotfile (home)", "Target (./.config/dots folder)")
		for _,f := range files {
			
			filename := f.Name()
			if(filename == "dots.yaml"){
				continue
			} else {
			homePath := filepath.Join(usr.HomeDir,filename)
			dotPath := filepath.Join(dotDr,filename)

			link,err := os.Readlink(homePath)
			if err != nil {
				if os.IsNotExist(err){
				fmt.Printf("%-40s  ->  %s\n", "Missing symlink: " + homePath,dotPath)
				} else {
				fmt.Printf("%-40s  ->  %s\n", "Not a symlink or unreadable: "+homePath,"")
			}
				continue
		}
		absTarget,_ := filepath.Abs(dotPath)
		absLink,_ := filepath.Abs(link)

		if absTarget == absLink {
			fmt.Printf("%-40s  ->  %s\n", "Status ok: "+homePath, link)
		} else {
			fmt.Printf("Wrong target: %s -> %s (expected %s)\n",homePath,link,dotPath)
		}
	}
	}
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
