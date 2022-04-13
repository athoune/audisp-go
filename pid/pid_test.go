package pid

import (
	"bufio"
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPid(t *testing.T) {
	raw := `Name:	vfio-irqfd-clea
Umask:	0000
State:	I (idle)
Tgid:	99
Ngid:	0
Pid:	99
PPid:	2
TracerPid:	0
Uid:	0	0	0	0
Gid:	0	0	0	0
FDSize:	64
Groups:
NStgid:	99
NSpid:	99
NSpgid:	0
NSsid:	0
Threads:	1
SigQ:	0/15481
SigPnd:	0000000000000000
ShdPnd:	0000000000000000
SigBlk:	0000000000000000
SigIgn:	ffffffffffffffff
SigCgt:	0000000000000000
CapInh:	0000000000000000
CapPrm:	0000003fffffffff
CapEff:	0000003fffffffff
CapBnd:	0000003fffffffff
CapAmb:	0000000000000000
NoNewPrivs:	0
Seccomp:	0
Speculation_Store_Bypass:	not vulnerable
Cpus_allowed:	3
Cpus_allowed_list:	0-1
Mems_allowed:	00000000,00000000,00000000,00000000,00000000,00000000,00000000,00000000,00000000,00000000,00000000,00000000,00000000,00000000,00000000,00000000,00000000,00000000,00000000,00000000,00000000,00000000,00000000,00000000,00000000,00000000,00000000,00000000,00000000,00000000,00000000,00000001
Mems_allowed_list:	0
voluntary_ctxt_switches:	2
nonvoluntary_ctxt_switches:	0
`
	reader := bufio.NewReader(bytes.NewBuffer([]byte(raw)))
	ppid, err := statusPpid(reader)
	assert.NoError(t, err)
	assert.Equal(t, "2", ppid)
}
