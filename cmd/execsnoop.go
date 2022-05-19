/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"github.com/ahubaoan/emage/bpf/execsnoop"
	"github.com/spf13/cobra"
)

// execsnoopCmd represents the execsnoop command
var execsnoopCmd = &cobra.Command{
	Use:   "execsnoop",
	Short: "exec snoop",
	Long:  `Monitor program exec`,
	Run: func(cmd *cobra.Command, args []string) {
		e := execsnoop.ExecSnoop{}
		e.Start()
	},
}

func init() {
	rootCmd.AddCommand(execsnoopCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// execsnoopCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// execsnoopCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
