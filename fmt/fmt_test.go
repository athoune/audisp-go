package fmt

import (
	"fmt"
	"io"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNaked(t *testing.T) {
	f := New(`pid=147491`)
	k, v, err := f.Next()
	assert.NoError(t, err)
	assert.Equal(t, "pid", k)
	assert.Equal(t, "147491", v)
	_, _, err = f.Next()
	assert.Equal(t, io.EOF, err)
}

func TestDouble(t *testing.T) {
	f := New(`name="John Doe"`)
	k, v, err := f.Next()
	assert.NoError(t, err)
	assert.Equal(t, "name", k)
	assert.Equal(t, "John Doe", v)
	_, _, err = f.Next()
	assert.Equal(t, io.EOF, err)
}

func TestSimple(t *testing.T) {
	f := New(`name='Mister X'`)
	k, v, err := f.Next()
	assert.NoError(t, err)
	assert.Equal(t, "name", k)
	assert.Equal(t, "Mister X", v)
	_, _, err = f.Next()
	assert.Equal(t, io.EOF, err)
}

func TestFmt(t *testing.T) {
	f := New(`node=sd-127470 type=USER_START msg=audit(1648302613.336:67497): pid=147491 uid=0 auid=1000 ses=117 msg='op=PAM:session_open grantors=pam_selinux,pam_loginuid,pam_keyinit,pam_permit,pam_umask,pam_unix,pam_systemd,pam_mail,pam_limits,pam_env,pam_env,pam_selinux acct="mlecarme" exe="/usr/sbin/sshd" hostname=88.123.196.115 addr=88.123.196.115 terminal=ssh res=success'`)
	for {
		k, v, err := f.Next()
		if err == io.EOF {
			break
		}
		fmt.Println(k, v)
	}
	assert.False(t, true)
}
