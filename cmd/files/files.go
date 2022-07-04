/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package files

import (
	"github.com/spf13/cobra"
)

var (
	createFileName string
	deleteFileName string
	listDirPath    string
	createFilePerm int
)

// FilesCmd represents the files command
var FilesCmd = &cobra.Command{
	Use:   "files",
	Short: "Common file operations",
	Long:  ``,
}

func init() {
}
