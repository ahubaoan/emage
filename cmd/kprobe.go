/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// kprobeCmd represents the kprobe command
var kprobeCmd = &cobra.Command{
	Use:   "kprobe",
	Short: "A brief description of your command",
	Long:  `.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("kprobe called")
	},
}

func init() {
	rootCmd.AddCommand(kprobeCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// kprobeCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// kprobeCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
