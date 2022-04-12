package message

/*

See https://access.redhat.com/articles/4409591#audit-event-fields-1

*/

import (
	"strings"
	"time"

	"github.com/athoune/audisp-go/audisp"
	"github.com/athoune/audisp-go/audit"
	"github.com/athoune/audisp-go/fmt"
)

type Messages struct {
	audisp       *audisp.Audisp
	currentLine  *fmt.Fmt
	currentError error
}

func New(a *audisp.Audisp) *Messages {
	return &Messages{
		audisp: a,
	}
}

type Message struct {
	ID        uint
	TimeStamp time.Time
	Values    map[string]string
}

func (m *Messages) Next() bool {
	m.currentLine, m.currentError = m.audisp.Line()
	return m.currentError == nil
}

func (m *Messages) Error() error {
	return m.currentError
}

func (m *Messages) Message() *Message {
	mm := &Message{
		Values: make(map[string]string),
	}
	for m.currentLine.Next() {
		err := m.currentLine.Error()
		if err != nil {
			m.currentError = err
			return nil
		}
		k, v := m.currentLine.KeyValue()
		mm.Values[k] = v
		if k == "msg" && strings.HasPrefix(v, "audit(") {
			a, err := audit.Parse(v)
			if err != nil {
				m.currentError = err
				return nil
			}
			mm.ID = a.ID
			mm.TimeStamp = a.TimeStamp
		}
	}
	return mm
}
