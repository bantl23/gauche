package term

import (
	"fmt"
)

// Windows terminal

type windows struct {
}

func newWindows() (Term, error) {
	t := new(windows)
	fmt.Println("New")
	return t, nil
}

func (t *windows) Restore() error {
	fmt.Println("Restore")
	return nil
}
