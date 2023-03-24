package internal_test

import (
	"testing"

	"github.com/ppluuums-jp/lt/internal"
)

func TestProgress(t *testing.T) {
	progress := internal.NewProgress("Progress: ", 10)
	progress.Start()
	for i := 0; i < 10; i++ {
		progress.Increment()
	}
	progress.Stop()
}
