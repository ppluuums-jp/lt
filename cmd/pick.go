package cmd

import (
	"github.com/ppluuums-jp/lt/internal"
	"github.com/spf13/cobra"
)

func PickCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "pick",
		Short: "pick the specified number of people from the applicants",
		RunE:  internal.Pick,
	}
	flags := cmd.Flags()
	flags.IntP("number", "n", 1, "picking number")
	flags.StringP("txt", "t", "", "the file to get names from")
	return cmd
}
