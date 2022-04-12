package main

import (
	"github.com/athoune/audisp-go/audisp"
	"github.com/athoune/audisp-go/sshd"
)

func main() {
	a, err := audisp.New("/var/run/audispd_events")
	if err != nil {
		panic(err)
	}
	defer a.Close()
	sshds := sshd.New()
	err = sshds.Snitch(a)
	if err != nil {
		panic(err)
	}
}
