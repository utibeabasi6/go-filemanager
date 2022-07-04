/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package directories

import (
	"github.com/spf13/cobra"
)

// DirectoriesCmd represents the directories command
var DirectoriesCmd = &cobra.Command{
	Use:   "directories",
	Short: "A brief description of your command",
	Long:  ``,
	// Run: func(cmd *cobra.Command, args []string) {
	// 	fmt.Println("directories called")
	// },
}

func init() {
	DirectoriesCmd.AddCommand(createCmd)
	DirectoriesCmd.AddCommand(deleteCmd)
	DirectoriesCmd.AddCommand(listCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// directoriesCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// directoriesCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
