package main

import (
	"fmt"
	"io"
	"os"

	"github.com/bantl23/gauche/prompt"
)

func mainWithReturn() int {
	p, err := prompt.NewPrompt(&prompt.Config{
		Prompt:   "gauche> ",
		ExitText: "quit",
	})
	if err != nil {
		fmt.Println("unable to get new prompt")
	}
	defer p.Close()

	done := false
	for !done {
		line, err := p.Readline()
		if err != nil {
			if err == io.EOF {
				done = true
				continue
			}
			fmt.Println("Error getting line")
			done = true
			continue
		}

		fmt.Printf("\r\nmain: %s\r\n", line)
	}

	return 0
}

func main() {
	os.Exit(mainWithReturn())
}
