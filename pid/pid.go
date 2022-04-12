package pid

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func SonOf(pid uint) ([]uint64, error) {
	sons := make([]uint64, 0)
	f, err := os.Open(fmt.Sprintf("/proc/%d/stats", pid))
	if err != nil {
		return nil, err
	}
	defer f.Close()
	reader := bufio.NewReader(f)
	spid := fmt.Sprintf("%d", pid)
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				break
			}
			return nil, err
		}
		blobs := strings.Split(line, ":")
		if blobs[0] == "PPid" {
			value := strings.TrimLeft(blobs[1], " ")
			if value == spid {
				ppid, err := strconv.ParseUint(value, 10, 64)
				if err != nil {
					return nil, err
				}
				sons = append(sons, ppid)
			}
		}
	}
	return sons, nil
}
