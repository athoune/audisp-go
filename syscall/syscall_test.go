package syscall

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSyscall(t *testing.T) {
	s, err := syscall("../contrib/unistd_64.h")
	assert.NoError(t, err)
	assert.Len(t, s, 436)
}
