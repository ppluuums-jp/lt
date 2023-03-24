package internal

import (
	"bufio"
	"errors"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

type Parser struct{}

func (p Parser) Parse(cmd *cobra.Command, args []string) (Applicants, error) {
	// Check if specified both -t and args
	file, err := cmd.Flags().GetString("txt")
	if file != "" && len(args) > 0 {
		return nil, errors.New("cannot specify both a file and command line arguments")
	}

	// Parse applicants depends on the flags
	var applicants Applicants
	if file != "" {
		applicants, err = parseFromFile(file)
		if err != nil {
			return nil, fmt.Errorf("failed to parse from file: %w", err)
		}
	} else {
		applicants, err = parseFromArgs(args)
		if err != nil {
			return nil, fmt.Errorf("failed to parse from args: %w", err)
		}
	}

	return applicants, nil
}

func parseFromFile(file string) (Applicants, error) {
	f, err := os.Open(file)
	if err != nil {
		return nil, fmt.Errorf("failed to open file: %w", err)
	}
	defer f.Close()

	// Check if specified file is a directory
	fileInfo, err := f.Stat()
	if err != nil {
		return nil, fmt.Errorf("failed to get file info: %w", err)
	}
	if fileInfo.IsDir() {
		return nil, fmt.Errorf("file %s is a directory", file)
	}

	// Actual parsing
	scanner := bufio.NewScanner(f)
	var applicants Applicants
	for scanner.Scan() {
		name := scanner.Text()
		applicant := Applicant{Name: name}
		if err := applicant.Validate(); err != nil {
			return nil, fmt.Errorf("invalid applicant: %w", err)
		}
		applicants = append(applicants, applicant)
	}
	if len(applicants) == 0 {
		return nil, errors.New("empty list of applicants")
	}

	return applicants, nil
}

func parseFromArgs(args []string) (Applicants, error) {
	if len(args) == 0 {
		return nil, errors.New("no applicants specified")
	}

	// Actual parsing
	var applicants Applicants
	for _, arg := range args {
		applicants = append(applicants, Applicant{Name: arg})
	}

	return applicants, nil
}
