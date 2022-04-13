package pid

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"regexp"
	"strings"

	"log"
)

var Numbers *regexp.Regexp

func init() {
	Numbers = regexp.MustCompile(`\d+`)
}

func SonOf(pid string) ([]string, error) {
	sons := make([]string, 0)

	files, err := ioutil.ReadDir("/proc")
	if err != nil {
		return nil, err
	}
	for _, file := range files {
		if file.IsDir() && Numbers.Match([]byte(file.Name())) {
			p, err := ppid(file.Name())
			if err != nil {
				// yolo
				log.Println(err)
				continue
			}
			if p == pid {
				sons = append(sons, file.Name())
			}
		}
	}
	return sons, nil
}

func ppid(pid string) (string, error) {
	f, err := os.Open(fmt.Sprintf("/proc/%s/status", pid))
	if err != nil {
		return "", err
	}
	defer f.Close()
	reader := bufio.NewReader(f)
	return statusPpid(reader)
}

func statusPpid(reader *bufio.Reader) (string, error) {
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				break
			}
			return "", err
		}
		if strings.HasPrefix(line, "PPid") {
			return strings.TrimLeft(strings.Split(line[:len(line)-1], ":")[1], "\t"), nil
		}
	}
	return "", fmt.Errorf("can't find ppid")
}
