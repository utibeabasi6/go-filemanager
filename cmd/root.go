/*
Copyright © 2022 Utibeabasi Umanah utibeabasiumanah6@gmail.com

*/
package cmd

import (
	"os"

	"github.com/Utibeabasi6/go-filemanager/cmd/directories"
	"github.com/Utibeabasi6/go-filemanager/cmd/files"
	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "go-filemanager",
	Short: "A commandline tool for working with files and directories",
	Long:  ``,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func addSubCommands() {
	rootCmd.AddCommand(files.FilesCmd)
	rootCmd.AddCommand(directories.DirectoriesCmd)
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.go-filemanager.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	addSubCommands()
}
