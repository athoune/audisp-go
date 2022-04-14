package socket

import (
	"net"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParse4(t *testing.T) {
	saddr, err := ParseSaddr("020000357F0000359006002C8A7F0000")
	assert.NoError(t, err)
	assert.Equal(t, int32(53), saddr.Port)
	assert.Equal(t, net.IP{127, 0, 0, 53}, saddr.IP)
}

func TestParse6(t *testing.T) {
	saddr, err := ParseSaddr("0A000050000000002A010E0C00010000000000000000000100000000")
	assert.NoError(t, err)
	assert.Equal(t, int32(80), saddr.Port)
	assert.Equal(t, "inet6", saddr.Family)
	assert.Equal(t, net.ParseIP("2a01:e0c:1::1"), saddr.IP) // it's free.fr IP
}
