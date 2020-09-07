package term

// #include "termlinux.h"
import "C"

import (
	"fmt"
)

// linux terminal

type linux struct {
}

func newLinux() (Term, error) {
	rv := C.init()
	if rv != 0 {
		return nil, fmt.Errorf("Unabled to initialize terminal")
	}

	t := new(linux)
	return t, nil
}

func (t *linux) Restore() error {
	rv := C.restore()
	if rv != 0 {
		return fmt.Errorf("Unable to restore terminal settings")
	}
	return nil
}
