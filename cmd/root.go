/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"github.com/ahubaoan/emage/config"
	"github.com/ahubaoan/emage/pkg/logger"
	"github.com/cilium/ebpf/rlimit"
	"go.uber.org/zap"
	"os"

	"github.com/spf13/cobra"
)

var confFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "emage",
	Short: "A ebpf app can be used for monitoring, security, auditing, etc.",
	Long:  `A ebpf app can be used for monitoring, security, auditing, etc.. Customized configuration can realize these gons.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("1111")
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().StringVarP(&confFile, "conf_file", "f", "", "choose a config file to start ")
}

func EnvInit() {
	if confFile != "" {
		fmt.Println("init config from file ", confFile)
		config.InitConfig(confFile)
	} else {
		fmt.Println("init config by default set")
		config.InitConfigDefault()
	}

	// Allow the current process to lock memory for eBPF resources.
	if err := rlimit.RemoveMemlock(); err != nil {
		logger.ComLog.Fatal("execsnoopCmd RemoveMemlock error", zap.Error(err))
	}
}
