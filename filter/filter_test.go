package filter

import (
	"fmt"
	"testing"

	"github.com/athoune/audisp-go/message"
	"github.com/stretchr/testify/assert"
)

func TestPipe(t *testing.T) {
	messages := message.NewMessagesMockup(
		`type=SYSCALL msg=audit(1649877826.389:602570): arch=c000003e syscall=59 success=yes exit=0 a0=564757dfee20 a1=564757cfd320 a2=564757e1ecf0 a3=8 items=2 ppid=2100871 pid=2109512 auid=1000 uid=1000 gid=1000 euid=1000 suid=1000 fsuid=1000 egid=1000 sgid=1000 fsgid=1000 tty=pts2 ses=710 comm="curl" exe="/usr/bin/curl" key="susp_activity"
type=EXECVE msg=audit(1649877826.389:602570): argc=2 a0="curl" a1="free.fr"
type=PATH msg=audit(1649877826.389:602570): item=0 name="/usr/bin/curl" inode=18877655 dev=08:02 mode=0100755 ouid=0 ogid=0 rdev=00:00 nametype=NORMAL cap_fp=0 cap_fi=0 cap_fe=0 cap_fver=0 cap_frootid=0
type=PATH msg=audit(1649877826.389:602570): item=1 name="/lib64/ld-linux-x86-64.so.2" inode=33816588 dev=08:02 mode=0100755 ouid=0 ogid=0 rdev=00:00 nametype=NORMAL cap_fp=0 cap_fi=0 cap_fe=0 cap_fver=0 cap_frootid=0
type=PROCTITLE msg=audit(1649877826.389:602570): proctitle=6375726C00667265652E6672
type=SYSCALL msg=audit(1649877826.417:602571): arch=c000003e syscall=42 success=yes exit=0 a0=7 a1=7fc6e7469dcc a2=10 a3=ffffffffffffff0a items=0 ppid=2100871 pid=2109512 auid=1000 uid=1000 gid=1000 euid=1000 suid=1000 fsuid=1000 egid=1000 sgid=1000 fsgid=1000 tty=pts2 ses=710 comm="curl" exe="/usr/bin/curl" key="network_connect_4"
`)
	reader, err := New(`line.syscall == "59"`, true, message.New(messages))
	assert.NoError(t, err)
	cpt := 0
	for reader.Next() {
		err = reader.Error()
		assert.NoError(t, err)
		msg := reader.Message()
		fmt.Println(msg)
		cpt++
	}
	assert.Equal(t, 5, cpt)
}
