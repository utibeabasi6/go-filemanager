/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package files

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// createCmd represents the create command
var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Creates a file",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		if _, err := os.Stat(createFileName); os.IsNotExist(err) {
			new, err := os.Create(createFileName)
			if err != nil {
				fmt.Println(err)
			} else {
				if createFilePerm != 0666 {
					err := os.Chmod(createFileName, os.FileMode(createFilePerm))
					if err != nil {
						fmt.Println(err)
					}
				}
			}
			defer new.Close()
		} else {
			fmt.Println(createFileName, "already exists")
		}

	},
}

func init() {
	createCmd.Flags().StringVarP(&createFileName, "name", "n", "", "The name of the file to create")
	createCmd.Flags().IntVarP(&createFilePerm, "mode", "m", 0666, "The permissions to create the file with.")
	if err := createCmd.MarkFlagRequired("name"); err != nil {
		fmt.Println(err)
	}
	FilesCmd.AddCommand(createCmd)
}
