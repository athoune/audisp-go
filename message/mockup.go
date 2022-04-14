package message

import (
	"io"
	"strings"

	"github.com/athoune/audisp-go/fmt"
)

type MessagesMockup struct {
	msgs []string
	poz  int
}

func NewMessagesMockup(txt string) *MessagesMockup {
	return &MessagesMockup{
		strings.Split(txt, "\n"),
		0,
	}
}

func (m *MessagesMockup) Close() error {
	return nil
}

func (m *MessagesMockup) Line() (*fmt.Fmt, error) {
	m.poz++
	if m.poz >= len(m.msgs) {
		return nil, io.EOF
	}
	return fmt.New(m.msgs[m.poz-1]), nil
}
