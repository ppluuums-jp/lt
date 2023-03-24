package internal

import (
	"errors"
	"fmt"
	"time"

	"github.com/spf13/cobra"
)

func Pick(cmd *cobra.Command, args []string) error {
	// Get the number if number flag is true
	number, err := cmd.Flags().GetInt("number")
	if err != nil {
		return fmt.Errorf("failed to get number flag: %w", err)
	}
	if number <= 0 {
		return errors.New("number must be greater than 0")
	}

	// Parse the inputs
	parser := Parser{}
	applicants, err := parser.Parse(cmd, args)
	if err != nil {
		return fmt.Errorf("failed to parse applicants: %w", err)
	}
	if len(applicants) < number {
		return errors.New("not enough applicants to choose from")
	}

	// Shuffle the applicants
	applicants.Shuffle()
	presenters := append([]Applicant{}, applicants[:number]...)

	// Show the progress bar
	progress := NewProgress("Choosing ...", 100)
	progress.Start()
	progress.Increment()
	progress.Stop()

	// Output the presenters
	for _, p := range presenters {
		fmt.Printf("- %s\n", p.Name)
		<-time.Tick(750 * time.Millisecond)
	}

	return nil
}
