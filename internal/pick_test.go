package internal_test

import (
	"bytes"
	"errors"
	"strconv"
	"testing"

	"github.com/ppluuums-jp/lt/internal"
	"github.com/spf13/cobra"
)

func TestPick(t *testing.T) {
	tests := []struct {
		name         string
		args         []string
		number       int
		expectedErr  error
		expectedText string
	}{
		{
			name:         "not_enough_applicants",
			args:         []string{"Alice", "Bob", "Charlie", "Dave"},
			number:       5,
			expectedErr:  errors.New("not enough applicants to choose from"),
			expectedText: "",
		},
		{
			name:         "invalid_number",
			args:         []string{"Alice", "Bob", "Charlie", "Dave"},
			number:       0,
			expectedErr:  errors.New("number must be greater than 0"),
			expectedText: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cmd := &cobra.Command{}
			cmd.Flags().Int("number", 0, "Number of presenters to choose")

			cmd.SetArgs(tt.args)
			cmd.Flags().Set("number", strconv.Itoa(tt.number))

			var buf bytes.Buffer
			cmd.SetOut(&buf)

			err := internal.Pick(cmd, tt.args)

			if tt.expectedErr != nil {
				if err == nil || err.Error() != tt.expectedErr.Error() {
					t.Errorf("expected error %q but got %q", tt.expectedErr, err)
				}
			} else {
				if err != nil {
					t.Errorf("unexpected error: %v", err)
				}
				t.Run("output", func(t *testing.T) {
					if buf.String() != tt.expectedText {
						t.Errorf("expected output %q but got %q", tt.expectedText, buf.String())
					}
				})
			}
		})
	}
}
