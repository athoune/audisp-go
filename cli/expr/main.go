package main

import (
	"fmt"
	"os"

	"github.com/athoune/audisp-go/audisp"
	"github.com/athoune/audisp-go/filter"
	"github.com/athoune/audisp-go/message"
)

func main() {
	a, err := audisp.New("/var/run/audispd_events")
	if err != nil {
		panic(err)
	}
	defer a.Close()
	reader, err := filter.New(os.Args[1], true, message.New(a))
	for reader.Next() {
		msg := reader.Message()
		fmt.Println(msg.Raw())
	}
	err = reader.Error()
	if err != nil {
		panic(err)
	}
}
