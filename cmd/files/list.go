/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package files

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all files in the specified directory",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		files, err := ioutil.ReadDir(listDirPath)
		if err != nil {
			fmt.Println(err)
		}
		for _, file := range files {
			if !file.IsDir() {
				fmt.Println(file.Name())
			}
		}
	},
}

func init() {
	dir, _ := os.Getwd()

	listCmd.Flags().StringVarP(&listDirPath, "path", "p", dir, "The path to directory to list files in")
	FilesCmd.AddCommand(listCmd)
}
