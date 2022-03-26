package audisp

import (
	"bufio"
	"net"

	"github.com/athoune/audisp-go/fmt"
)

type Audisp struct {
	conn   net.Conn
	reader *bufio.Reader
}

func New(path string) (*Audisp, error) {
	conn, err := net.Dial("unix", path)
	if err != nil {
		return nil, err
	}
	return &Audisp{
		conn:   conn,
		reader: bufio.NewReader(conn),
	}, nil
}

func (a *Audisp) Line() (*fmt.Fmt, error) {
	text, err := a.reader.ReadString('\n')
	if err != nil {
		return nil, err
	}
	return fmt.New(text), nil
}

func (a *Audisp) Close() error {
	return a.conn.Close()
}
