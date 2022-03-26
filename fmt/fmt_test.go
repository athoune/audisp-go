package fmt

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNaked(t *testing.T) {
	f := New(`pid=147491`)
	assert.True(t, f.Next())
	assert.NoError(t, f.Error())
	k, v := f.KeyValue()
	assert.Equal(t, "pid", k)
	assert.Equal(t, "147491", v)
	assert.False(t, f.Next())
}

func TestChariotReturn(t *testing.T) {
	f := New("pid=147491\n")
	assert.True(t, f.Next())
	assert.NoError(t, f.Error())
	k, v := f.KeyValue()
	assert.Equal(t, "pid", k)
	assert.Equal(t, "147491", v)
	assert.False(t, f.Next())
}
func TestSpaces(t *testing.T) {
	f := New(`pid=147491 name=auditd`)
	assert.True(t, f.Next())
	assert.NoError(t, f.Error())
	k, v := f.KeyValue()
	assert.Equal(t, "pid", k)
	assert.Equal(t, "147491", v)

	assert.True(t, f.Next())
	assert.NoError(t, f.Error())
	k, v = f.KeyValue()
	assert.Equal(t, "name", k)
	assert.Equal(t, "auditd", v)
	assert.False(t, f.Next())
}
func TestDouble(t *testing.T) {
	f := New(`name="John Doe"`)
	assert.True(t, f.Next())
	assert.NoError(t, f.Error())
	k, v := f.KeyValue()
	assert.Equal(t, "name", k)
	assert.Equal(t, "John Doe", v)
	assert.False(t, f.Next())
}

func TestSimple(t *testing.T) {
	f := New(`name='Mister X'`)
	assert.True(t, f.Next())
	assert.NoError(t, f.Error())
	k, v := f.KeyValue()
	assert.Equal(t, "name", k)
	assert.Equal(t, "Mister X", v)
	assert.False(t, f.Next())
}

func TestFmt(t *testing.T) {
	f := New(`type=USER_LOGIN msg=audit(1648319421.985:67697): pid=164731 uid=0 auid=4294967295 ses=4294967295 msg='op=login acct="root" exe="/usr/sbin/sshd" hostname=? addr=104.194.75.112 terminal=sshd res=failed'`)
	cpt := 0
	for f.Next() {
		assert.NoError(t, f.Error())
		k, v := f.KeyValue()
		fmt.Printf("[%s] : [%s]\n", k, v)
		cpt++
	}
	assert.Equal(t, 7, cpt)
}
