package main

import (
	"bufio"
	"fmt"
	"io"
	"os"

	"github.com/bantl23/gauche/term"
)

func mainWithReturn() int {
	t, err := term.NewTerm()
	if err != nil {
		fmt.Println("unable to initialize term")
		return 1
	}
	defer t.Restore()

	done := false
	reader := bufio.NewReader(os.Stdin)
	prompt := true
	for !done {
		if prompt == true {
			fmt.Printf("gauche> ")
		}
		item, err := reader.ReadByte()
		if err != nil {
			if err == io.EOF {
				prompt = false
				continue
			} else {
				fmt.Printf("Error: %#v\r\n", err)
				done = true
				continue
			}
		}

		prompt = true
		if item >= 33 && item <= 126 {
			fmt.Printf("%03d 0x%02x %s\r\n", item, item, string(item))
		} else {
			fmt.Printf("%03d 0x%02x\r\n", item, item)
		}

		if item == 'q' || item == 0x03 {
			done = true
			continue
		}
	}

	return 0
}

func main() {
	os.Exit(mainWithReturn())
}
