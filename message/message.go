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

type MessagesReader interface {
	Next() bool
	Error() error
	Message() *Message
}

type Messages struct {
	audisp       audisp.LineReader
	currentLine  *fmt.Fmt
	currentError error
}

func New(a audisp.LineReader) MessagesReader {
	return &Messages{
		audisp: a,
	}
}

func (m *Messages) Next() bool {
	m.currentLine, m.currentError = m.audisp.Line()
	return m.currentError == nil
}

func (m *Messages) Error() error {
	return m.currentError
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

type Message struct {
	raw       string
	Type      string
	TimeStamp time.Time
	ID        uint
	line      *fmt.Fmt
	values    map[string]string
}

func newMessage(line *fmt.Fmt) (*Message, error) {
	m := &Message{
		line:   line,
		raw:    line.Raw(),
		values: make(map[string]string),
	}
	// assert k == Type
	m.Type, _ = m.Get("type")
	// assert k = msg
	v, _ := m.Get("msg")
	a, err := audit.Parse(v)
	if err != nil {
		return nil, err
	}
	m.TimeStamp = a.TimeStamp
	m.ID = a.ID
	return m, nil
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

func (m *Message) Fetch(k interface{}) interface{} {
	kk, ok := k.(string)
	if !ok {
		return nil
	}
	v, ok := m.Get(kk)
	if !ok {
		return nil
	}
	return v
}

func (m *Message) Raw() string {
	return m.raw
}
