/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"context"
	"fmt"
	"github.com/ahubaoan/emage/bpf/exitsnoop"
	"github.com/ahubaoan/emage/config/module"
	"github.com/spf13/cobra"
)

var exitSnoopConfig = module.ExitSnoop{}

var exitsnoopCmd = &cobra.Command{
	Use:   "exitsnoop",
	Short: "listen process exit",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("exitsnoop called")
		ctx, _ := context.WithCancel(context.TODO())
		exitsnoop.Start(ctx, exitSnoopConfig.ExitSnoopKern)
	},
}

func init() {
	exitsnoopCmd.Flags().BoolVarP(&exitSnoopConfig.FilterCroup, "filter_cgroup", "c", false, "filter cgoup")
	exitsnoopCmd.Flags().Uint32VarP(&exitSnoopConfig.TargetPid, "target_pid", "t", 0, "target pid")
	exitsnoopCmd.Flags().BoolVarP(&exitSnoopConfig.TraceFailedOnly, "trace_filed_only", "f", false, "trace filed exit only")
	rootCmd.AddCommand(exitsnoopCmd)
}
