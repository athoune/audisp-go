package fmt

import (
	"errors"
	"fmt"
	"strings"
)

type Fmt struct {
	txt   string
	poz   int
	key   string
	value string
	err   error
}

func New(txt string) *Fmt {
	if txt[len(txt)-1] == '\n' {
		txt = txt[:len(txt)-1]
	}
	return &Fmt{
		txt: txt,
	}
}

func (f *Fmt) Next() bool {
	if f.poz >= len(f.txt) {
		return false
	}
	poz := strings.Index(f.txt[f.poz:], "=")
	if poz == -1 {
		f.err = errors.New("can't find =")
		return false
	}
	f.key = f.txt[f.poz : f.poz+poz]
	f.poz += poz
	first := f.txt[f.poz+1]
	if first == '"' || first == '\'' {
		poz = strings.Index(f.txt[f.poz+2:], string(first))
		if poz == -1 {
			f.err = fmt.Errorf("can't find %v at %d", first, f.poz+2)
		}
		f.value = f.txt[f.poz+2 : f.poz+poz+2]
		poz += 2
	} else {
		poz = strings.Index(f.txt[f.poz:], " ")
		if poz == -1 {
			poz = len(f.txt) - f.poz // until the end
		}
		f.value = f.txt[f.poz+1 : f.poz+poz]
	}
	f.poz += poz + 1
	return true
}

func (f *Fmt) KeyValue() (string, string) {
	return f.key, f.value
}

func (f *Fmt) Error() error {
	return f.err
}
