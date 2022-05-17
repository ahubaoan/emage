/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package main

import (
	"github.com/ahubaoan/emage/cmd"
	"github.com/cilium/ebpf/rlimit"
	"log"
)

func init() {
	// Allow the current process to lock memory for eBPF resources.
	if err := rlimit.RemoveMemlock(); err != nil {
		log.Fatal(err)
	}
}

func main() {
	cmd.Execute()
}
