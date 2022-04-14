package message

/*

See https://access.redhat.com/articles/4409591#audit-event-fields-1

*/

import (
	"time"

	"github.com/athoune/audisp-go/audisp"
	"github.com/athoune/audisp-go/audit"
	"github.com/athoune/audisp-go/fmt"
)

type Messages struct {
	audisp       audisp.LineReader
	currentLine  *fmt.Fmt
	currentError error
}

func New(a audisp.LineReader) *Messages {
	return &Messages{
		audisp: a,
	}
}

type Message struct {
	Type      string
	TimeStamp time.Time
	ID        uint
	line      *fmt.Fmt
	values    map[string]string
}

// Get is lazy
func (m *Message) Get(key string) (string, bool) {
	v, ok := m.values[key]
	if ok {
		return v, true
	}
	for m.line.Next() {
		k, v := m.line.KeyValue()
		m.values[k] = v
		if k == key {
			return v, true
		}
	}
	return "", false
}

func (m *Messages) Next() bool {
	m.currentLine, m.currentError = m.audisp.Line()
	return m.currentError == nil
}

func (m *Messages) Error() error {
	return m.currentError
}

func newMessage(line *fmt.Fmt) (*Message, error) {
	m := &Message{
		values: make(map[string]string),
	}
	line.Next()
	k, v := line.KeyValue()
	// assert k == Type
	m.values[k] = v
	m.Type = v
	line.Next()
	k, v = line.KeyValue()
	// assert k = msg
	a, err := audit.Parse(v)
	if err != nil {
		return nil, err
	}
	m.TimeStamp = a.TimeStamp
	m.ID = a.ID
	return m, nil
}

// Message return next Message
func (m *Messages) Message() *Message {
	mm, err := newMessage(m.currentLine)
	if err != nil {
		m.currentError = err
		return nil
	}
	return mm

}
