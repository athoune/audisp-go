package filter

import (
	"fmt"
	"time"

	"github.com/antonmedv/expr"
	"github.com/antonmedv/expr/vm"
	"github.com/athoune/audisp-go/message"
)

type FilteredReader struct {
	ids     map[uint]time.Time
	program *vm.Program
	source  message.MessagesReader
	current *message.Message
	vm      *vm.VM
	err     error
	env     map[string]interface{}
	sonsToo bool
}

func New(code string, sonsToo bool, source message.MessagesReader) (message.MessagesReader, error) {
	f := &FilteredReader{
		sonsToo: sonsToo,
		source:  source,
		ids:     make(map[uint]time.Time),
		vm:      &vm.VM{},
		env: map[string]interface{}{
			"sprintf": fmt.Sprintf,
		},
	}
	var err error
	f.program, err = expr.Compile(code,
		expr.Env(f.env),
		expr.AllowUndefinedVariables(), // Allow to use undefined variables.
	)
	if err != nil {
		return nil, err
	}
	return f, nil
}

func (f *FilteredReader) Next() bool {
	n := f.source.Next()
	if !n {
		return false
	}
	f.current = f.source.Message()
	if f.source.Error() != nil {
		f.err = f.source.Error()
		return false
	}
	if f.sonsToo {
		_, ok := f.ids[f.current.ID] // auditd is already exist
		if ok {                      // it's a son
			return true
		}
	}
	f.env["line"] = f.current
	out, err := f.vm.Run(f.program, f.env)
	if err != nil {
		f.err = err
		return false
	}
	resp, ok := out.(bool)
	if !ok {
		f.err = fmt.Errorf("not a boolean : %v", resp)
		return false
	}
	if resp {
		_, ok := f.ids[f.current.ID]
		if !ok {
			f.ids[f.current.ID] = time.Now()
		}
	}
	return resp
}

func (f *FilteredReader) Error() error {
	return f.source.Error()
}

func (f *FilteredReader) Message() *message.Message {
	return f.current
}
