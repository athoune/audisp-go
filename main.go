package main

import (
	_fmt "fmt"

	"github.com/athoune/audisp-go/audisp"
)

func main() {
	a, err := audisp.New("/var/run/audispd_events")
	if err != nil {
		panic(err)
	}
	defer a.Close()
	for {
		line, err := a.Line()
		if err != nil {
			_fmt.Println("Error :", err)
			break
		}
		_fmt.Println()
		for line.Next() {
			err := line.Error()
			if err != nil {
				_fmt.Println("Error :", err)
				continue
			}
			k, v := line.KeyValue()
			_fmt.Printf("%s => %s\n", k, v)
		}
	}
}
