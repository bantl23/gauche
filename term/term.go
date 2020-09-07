package term

import (
	"fmt"
	"runtime"
)

// Term provide unified interface for each os
type Term interface {
	// Restore restores terminal to original settings
	Restore() error
}

// NewTerm returns a new raw terminal
func NewTerm() (Term, error) {
	if runtime.GOOS == "windows" {
		return newWindows()
	} else if runtime.GOOS == "linux" {
		return newLinux()
	} else {
		return nil, fmt.Errorf("Unsupported os: %s", runtime.GOOS)
	}
}
