package internal

import (
	"errors"
	"fmt"
	"time"

	"github.com/spf13/cobra"
)

func Shuffle(cmd *cobra.Command, args []string) error {
	// Parse the inputs
	parser := Parser{}
	presenters, err := parser.Parse(cmd, args)
	if err != nil {
		return fmt.Errorf("failed to parse applicants: %w", err)
	}
	if len(presenters) < 1 {
		return errors.New("not provided the applicants")
	}

	// Shuffle the presenters
	presenters.Shuffle()

	// Show the progress bar
	progress := NewProgress("The Order is ...", 100)
	progress.Start()
	progress.Increment()
	progress.Stop()

	// Output the presenters
	for i, p := range presenters {
		fmt.Printf("No. %d %s\n", i+1, p.Name)
		time.Sleep(750 * time.Millisecond)

	}

	return nil
}
