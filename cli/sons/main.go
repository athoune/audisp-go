package main

import (
	"fmt"
	"os"

	"github.com/athoune/audisp-go/pid"
)

func main() {
	sons, err := pid.SonOf(os.Args[1])
	if err != nil {
		panic(err)
	}
	fmt.Println(sons)
}
