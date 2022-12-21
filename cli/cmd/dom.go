/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// domCmd represents the dom command
var domCmd = &cobra.Command{
	Use:   "dom",
	Short: "Greet dom",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Dom, you're a legend!")
	},
}

func init() {
	rootCmd.AddCommand(domCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// domCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// domCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
