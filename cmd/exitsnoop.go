/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"context"
	"github.com/ahubaoan/emage/pkg/bpf"
	"github.com/ahubaoan/emage/pkg/bpf/controller/exitsnoop"
	"github.com/ahubaoan/emage/pkg/bpf/module"
	"github.com/ahubaoan/emage/pkg/logger"
	"github.com/spf13/cobra"
)

var exitSnoopConfig = module.ExitSnoop{}

var exitsnoopCmd = &cobra.Command{
	Use:   "exitsnoop",
	Short: "listen process exit",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		EnvInit()
		logger.ComLog.Info("exitsnoop called")
		ctx, _ := context.WithCancel(context.TODO())
		exitsnoop.Start(ctx, exitSnoopConfig.ExitSnoopKern)
	},
}

func init() {
	exitsnoopCmd.Flags().BoolVarP(&exitSnoopConfig.FilterCroup, "filter_cgroup", "c",
		bpf.GetDefaultVal(exitSnoopConfig, "FilterCroup").(bool), "filter cgoup")
	exitsnoopCmd.Flags().Uint64VarP(&exitSnoopConfig.TargetPid, "target_pid", "t",
		bpf.GetDefaultVal(exitSnoopConfig, "TargetPid").(uint64), "target pid")
	exitsnoopCmd.Flags().BoolVarP(&exitSnoopConfig.TraceFailedOnly, "trace_failed_only", "o",
		bpf.GetDefaultVal(exitSnoopConfig, "TraceFailedOnly").(bool), "trace failed exit only")
	rootCmd.AddCommand(exitsnoopCmd)
}
