package main

import (
	"bufio"
	_fmt "fmt"
	"log"
	"net"

	"github.com/athoune/audisp-go/fmt"
)

func main() {
	conn, err := net.Dial("unix", "/var/run/audispd_events")
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	reader := bufio.NewReader(conn)
	for {
		text, err := reader.ReadString('\n')
		if err != nil {
			log.Println(err)
			return
		}
		line := fmt.New(text)
		_fmt.Println()
		_fmt.Println(text)
		for line.Next() {
			err := line.Error()
			if err != nil {
				_fmt.Println("Error :", err)
			}
			k, v := line.KeyValue()
			_fmt.Printf("%s => %s\n", k, v)
		}
	}

}
