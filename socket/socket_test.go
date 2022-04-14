package socket

import (
	"net"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParse4(t *testing.T) {
	saddr, err := Parse4("020000357F0000359006002C8A7F0000")
	assert.NoError(t, err)
	assert.Equal(t, int32(53), saddr.Port)
	assert.Equal(t, net.IP{127, 0, 0, 53}, saddr.IP)
}
