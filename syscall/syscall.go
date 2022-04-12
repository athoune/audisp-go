package syscall

import (
	"bufio"
	"io"
	"os"
	"strconv"
	"strings"
)

// New return syscall list
func New() []string {
	s, _ := syscall("/usr/include/x86_64-linux-gnu/asm/unistd_64.h")
	return s
}

type line struct {
	name string
	rank int64
}

func syscall(path string) ([]string, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	start := "#define __NR_"
	size := len(start)
	reader := bufio.NewReader(f)
	lines := make([]line, 0)
	for {
		l, err := reader.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				break
			}
			return nil, err
		}
		if !strings.HasPrefix(l, start) {
			continue
		}
		l = strings.Trim(l, "\n")
		blobs := strings.Split(l[size:], " ")
		rank, err := strconv.ParseInt(blobs[1], 10, 64)
		if err != nil {
			return nil, err
		}
		lines = append(lines, line{blobs[0], rank})
	}

	m := lines[len(lines)-1].rank
	s := make([]string, m+1)
	for _, line := range lines {
		s[line.rank] = line.name
	}

	return s, nil
}
