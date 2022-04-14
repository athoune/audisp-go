package syscall

import (
	"bytes"
	_ "embed"
)

//go:embed linux_x86_64/unistd_64.h
var s []byte

func init() {
	Syscalls, _ = syscallReader(bytes.NewBuffer(s))
}
