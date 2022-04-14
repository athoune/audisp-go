package sshd

import (
	"encoding/hex"
	"fmt"

	"github.com/athoune/audisp-go/audisp"
	"github.com/athoune/audisp-go/message"
)

type Sshd struct {
	ID uint
}

type Sshds map[uint]*Sshd

func New() Sshds {
	return make(Sshds)
}

func (s Sshds) Snitch(a audisp.LineReader) error {
	messages := message.New(a)
	for messages.Next() {
		err := messages.Error()
		if err != nil {
			return err
		}
		m := messages.Message()
		if m.Type == "EXECVE" {
			c, ok := m.Get("a0")
			if ok && c == "/usr/sbin/sshd" {
				s[m.ID] = &Sshd{
					ID: m.ID,
				}
			}
		}
		_, ok := s[m.ID]
		if ok {
			if m.Type == "PROCTITLE" {
				pt, _ := m.Get("proctitle")
				s, err := hex.DecodeString(pt)
				if err != nil {
					return err
				}
				fmt.Println("proc title", string(s))
			}
			fmt.Println(m.Raw())
		}

	}
	return nil
}
