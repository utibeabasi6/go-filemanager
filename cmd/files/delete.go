/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package files

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// deleteCmd represents the delete command
var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete a file",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		if _, err := os.Stat(deleteFileName); !os.IsNotExist(err) {
			err := os.Remove(deleteFileName)
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println(deleteFileName, "deleted")
			}
		}
	},
}

func init() {
	deleteCmd.Flags().StringVarP(&deleteFileName, "name", "n", "", "The name of the file to delete")
	if err := deleteCmd.MarkFlagRequired("name"); err != nil {
		fmt.Println(err)
	}
	FilesCmd.AddCommand(deleteCmd)
}
