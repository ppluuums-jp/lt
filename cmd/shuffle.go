package cmd

import (
	"github.com/ppluuums-jp/lt/internal"
	"github.com/spf13/cobra"
)

func ShuffleCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "shuffle",
		Short: "shuffles the order of the applicants",
		RunE:  internal.Shuffle,
	}
	flags := cmd.Flags()
	flags.StringP("txt", "t", "", "the file to get names from")
	return cmd
}
