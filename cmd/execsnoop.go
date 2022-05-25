/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"context"
	"github.com/ahubaoan/emage/pkg/bpf"
	"github.com/ahubaoan/emage/pkg/bpf/controller/execsnoop"
	"github.com/ahubaoan/emage/pkg/bpf/module"
	"github.com/ahubaoan/emage/pkg/logger"
	"github.com/spf13/cobra"
)

var execsnoopConfig = module.ExecSnoop{}

var execsnoopCmd = &cobra.Command{
	Use:   "execsnoop",
	Short: "exec snoop",
	Long:  `Monitor program exec`,
	Run: func(cmd *cobra.Command, args []string) {
		EnvInit()
		logger.ComLog.Info("execsnoop called")
		ctx, _ := context.WithCancel(context.TODO())
		execsnoop.Start(ctx, execsnoopConfig.ExecSnoopKern)
	},
}

func init() {
	execsnoopCmd.Flags().BoolVarP(&execsnoopConfig.FilterCroup, "filter_cgroup", "c",
		bpf.GetDefaultVal(execsnoopConfig, "FilterCroup").(bool), "filter cgoup")
	rootCmd.AddCommand(execsnoopCmd)
}
