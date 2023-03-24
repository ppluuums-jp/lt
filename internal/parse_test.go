package internal_test

import (
	"testing"

	"github.com/ppluuums-jp/lt/internal"
	"github.com/spf13/cobra"
)

func TestDefaultParser_Parse(t *testing.T) {
	type args struct {
		cmd  *cobra.Command
		args []string
	}
	tests := []struct {
		name    string
		p       internal.Parser
		args    args
		want    internal.Applicants
		wantErr bool
	}{
		{
			name: "Parse from file",
			p:    internal.Parser{},
			args: args{
				cmd: &cobra.Command{
					Use: "test",
					Run: func(cmd *cobra.Command, args []string) {
						// This is a dummy Run function to prevent the "no run function" error
					},
				},
				args: []string{},
			},
			want: []internal.Applicant{
				{Name: "Alice"},
				{Name: "Natasha"},
				{Name: "Bob"},
			},
			wantErr: false,
		},
		{
			name: "Parse from args",
			p:    internal.Parser{},
			args: args{
				cmd: &cobra.Command{
					Use: "test",
				},
				args: []string{"John Doe", "Jane Doe", "Bob Smith"},
			},
			want: []internal.Applicant{
				{Name: "John Doe"},
				{Name: "Jane Doe"},
				{Name: "Bob Smith"},
			},
			wantErr: false,
		},
		{
			name: "Empty list of applicants",
			p:    internal.Parser{},
			args: args{
				cmd: &cobra.Command{
					Use: "test",
					Run: func(cmd *cobra.Command, args []string) {
						// This is a dummy Run function to prevent the "no run function" error
					},
				},
				args: []string{},
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.name == "Parse from file" {
				tt.args.cmd.Flags().StringP("txt", "t", "", "the file to get names from")
				tt.args.cmd.Flags().Set("txt", "../applicants.txt")
			}
			got, err := tt.p.Parse(tt.args.cmd, tt.args.args)
			if (err != nil) != tt.wantErr {
				t.Errorf("DefaultParser.Parse() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !isEqual(got, tt.want) {
				t.Errorf("DefaultParser.Parse() = %v, want %v", got, tt.want)
			}
		})
	}
}

func isEqual(a, b internal.Applicants) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v.Name != b[i].Name {
			return false
		}
	}
	return true
}
