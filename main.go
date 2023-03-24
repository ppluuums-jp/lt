package main

import (
	"log"

	"github.com/ppluuums-jp/lt/cmd"
)

func main() {
	rootCmd := cmd.LtCommand()
	rootCmd.AddCommand(cmd.PickCommand())
	rootCmd.AddCommand(cmd.ShuffleCommand())
	err := rootCmd.Execute()
	if err != nil {
		log.Panicf("oops: %v", err)
	}
}
