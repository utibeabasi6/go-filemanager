/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package files

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all files in the specified directory",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		err := filepath.Walk(listDirPath, func(path string, info os.FileInfo, err error) error {
			if err != nil {
				fmt.Println(err)
				return err
			}
			if !info.IsDir() {
				fmt.Println(path)
			}
			return nil
		})
		if err != nil {
			fmt.Println(err)
		}
	},
}

func init() {
	dir, _ := os.Getwd()

	listCmd.Flags().StringVarP(&listDirPath, "path", "p", dir, "The path to directory to list files in")
	FilesCmd.AddCommand(listCmd)
}
