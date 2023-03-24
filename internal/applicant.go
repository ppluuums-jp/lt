package internal

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

type Applicant struct {
	Order    int
	Name     string
	Title    string
	Division string
}

func (a Applicant) String() string {
	return fmt.Sprintf("%d. %s %s %s", a.Order, a.Name, a.Division, a.Title)
}

func (a Applicant) Validate() error {
	if a.Name == "" {
		return errors.New("applicant name cannot be empty")
	}
	return nil
}

type Applicants []Applicant

func (a Applicants) Shuffle() {
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(a), func(i, j int) {
		a[i], a[j] = a[j], a[i]
	})
}
