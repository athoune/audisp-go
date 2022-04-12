package sshd

import (
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

func (s Sshds) Snitch(a *audisp.Audisp) error {
	messages := message.New(a)
	for messages.Next() {
		err := messages.Error()
		if err != nil {
			return err
		}
		m := messages.Message()
		if m.Values["type"] == "EXECVE" {
			c, ok := m.Values["a0"]
			if ok && c == "/usr/sbin/sshd" {
				s[m.ID] = &Sshd{
					ID: m.ID,
				}
			}
		}
		_, ok := s[m.ID]
		if ok {
			fmt.Println(m.ID, m.Values)
		}

	}
	return nil
}
