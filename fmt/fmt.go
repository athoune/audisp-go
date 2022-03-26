package fmt

import (
	"errors"
	"fmt"
	"io"
	"strings"
)

type Fmt struct {
	txt string
	poz int
}

func New(txt string) *Fmt {
	return &Fmt{
		txt: txt,
	}
}

func (f *Fmt) Next() (string, string, error) {
	if f.poz >= len(f.txt) {
		return "", "", io.EOF
	}
	poz := strings.Index(f.txt[f.poz:], "=")
	if poz == -1 {
		return "", "", errors.New("can't find =")
	}
	key := f.txt[f.poz : f.poz+poz]
	f.poz += poz
	var value string
	first := f.txt[f.poz+1]
	if first == '"' || first == '\'' {
		poz = strings.Index(f.txt[f.poz+2:], string(first))
		if poz == -1 {
			return "", "", fmt.Errorf("can't find %v at %d", first, f.poz+2)
		}
		value = f.txt[f.poz+2 : f.poz+poz+2]
		poz += 2
	} else {
		poz = strings.Index(f.txt[f.poz:], " ")
		if poz == -1 {
			poz = len(f.txt) - f.poz - 1 // until the end
		}
		value = f.txt[f.poz+1 : f.poz+poz+1]
	}
	f.poz += poz + 1
	return key, value, nil
}
