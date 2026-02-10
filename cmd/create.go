package cmd

import (
	"fmt"
	"os"
	"github.com/spf13/cobra"
)

// createCmd represents the create command
var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Used to create new dotfiles for syncing.",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {

	  if len(args) == 0 {
			fmt.Println("Error: Please provide a filename.")
			return
		}

		editFile := args[0]
		homedir,err := os.UserHomeDir()

		if err != nil {
			fmt.Println("Error creating file:", err)
			return
		}

 		dirPath := fmt.Sprintf("%s/.config/dots", homedir)
	if err := os.MkdirAll(dirPath, 0755); err != nil {
    fmt.Println("Error creating directory:", err)
   		 return
				}

		dotpath := fmt.Sprintf("%s/%s", dirPath, editFile)
		created, err := os.Create(dotpath)
			if err != nil {
    		fmt.Println("Error creating file:", err)
    		return
	}
			defer created.Close()

		fmt.Printf("Created file: %s\n", dotpath)



	},
}

func init() {
	rootCmd.AddCommand(createCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// createCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// createCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
