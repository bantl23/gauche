package prompt

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/bantl23/gauche/term"
)

const (
	keyEndOfText      = 0x03
	keyCarriageReturn = 0x0d
	keyEscape         = 0x1b
	keySpace          = 0x20
	keyLeftBracket    = 0x5b
	keyTilde          = 0x7e
	keyDelete         = 0x7f
)

// Prompt struct defines the prompt options
type Prompt struct {
	prompt   string
	exitText string
	reader   *bufio.Reader
	term     term.Term
}

// Config is the config options
type Config struct {
	Prompt   string
	ExitText string
}

// NewPrompt creates a new prompt
func NewPrompt(cfg *Config) (*Prompt, error) {
	p := new(Prompt)
	p.reader = bufio.NewReader(os.Stdin)
	t, err := term.NewTerm()
	if err != nil {
		return nil, err
	}
	p.term = t
	p.prompt = cfg.Prompt
	p.exitText = cfg.ExitText
	return p, nil
}

// Readline gets the next line
func (p *Prompt) Readline() (string, error) {
	fmt.Printf("%s", p.prompt)
	data := ""
	for {
		b, err := p.reader.ReadByte()
		if err != nil {
			if err == io.EOF {
				continue
			} else {
				return "", err
			}
		}
		switch {
		case b == keyDelete:
			if len(data) > 0 {
				data = data[:len(data)-1]
			}
		case b == keyEscape:
			var d0 byte
			var d1 byte
			for {
				d, err := p.reader.ReadByte()
				if err != nil {
					if err == io.EOF {
						continue
					} else {
						return "", err
					}
				}
				d0 = d
				break
			}
			for {
				d, err := p.reader.ReadByte()
				if err != nil {
					if err == io.EOF {
						continue
					} else {
						return "", err
					}
				}
				d1 = d
				break
			}
			if d0 == keyLeftBracket {
				if d1 == 'A' {
					fmt.Printf("up arrow\r\n")
				} else if d1 == 'B' {
					fmt.Printf("dn arrow\r\n")
				} else if d1 == 'C' {
					fmt.Printf("rt arrow\r\n")
				} else if d1 == 'D' {
					fmt.Printf("lt arrow\r\n")
				} else {
					fmt.Printf("none arrow\r\n")
				}
			} else {
				fmt.Printf("none\r\n")
			}
		case b >= keySpace && b <= keyTilde:
			data = data + string(b)
		case b == keyCarriageReturn:
			if data == strings.TrimSpace(p.exitText) {
				fmt.Printf("\r\n")
				return "", io.EOF
			}
			return data, nil
		case b == keyEndOfText:
			fmt.Printf("\r\n")
			return "", io.EOF
		}
		fmt.Printf("\r%s%s", p.prompt, data)
	}
}

// Close closes terminal and restores
func (p *Prompt) Close() error {
	return p.term.Restore()
}
