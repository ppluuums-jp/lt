package cmd

import (
	"github.com/spf13/cobra"
)

func LtCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:                   "lt COMMAND [OPTIONS] <ARGS...>",
		Short:                 "A tiny tool for managing lighting talks",
		TraverseChildren:      true,
		DisableFlagsInUseLine: true,
		Version:               "v1.0.0",
	}
	return cmd
}
