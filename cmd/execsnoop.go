/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"context"
	"fmt"
	"github.com/ahubaoan/emage/bpf/execsnoop"
	"github.com/ahubaoan/emage/config/module"
	"github.com/ahubaoan/emage/pkg/bpf"
	"github.com/cilium/ebpf/rlimit"
	"github.com/spf13/cobra"
	"log"
)

var execsnoopConfig = module.ExecSnoop{}

var execsnoopCmd = &cobra.Command{
	Use:   "execsnoop",
	Short: "exec snoop",
	Long:  `Monitor program exec`,
	Run: func(cmd *cobra.Command, args []string) {
		// Allow the current process to lock memory for eBPF resources.
		if err := rlimit.RemoveMemlock(); err != nil {
			log.Fatal(err)
		}
		fmt.Println("execsnoop called")
		ctx, _ := context.WithCancel(context.TODO())
		execsnoop.Start(ctx, execsnoopConfig.ExecSnoopKern)
	},
}

func init() {
	execsnoopCmd.LocalFlags().BoolVarP(&execsnoopConfig.FilterCroup, "filter_cgroup", "",
		bpf.GetDefaultVal(execsnoopConfig, "FilterCroup").(bool), "filter cgoup")
	rootCmd.AddCommand(execsnoopCmd)
}
