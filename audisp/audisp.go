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

// New return an Audisp instance, connected to a UNIX socket path to an audisp af_unix
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

// Line read one line
func (a *Audisp) Line() (*fmt.Fmt, error) {
	line, err := a.reader.ReadString('\n')
	if err != nil {
		return nil, err
	}
	return fmt.New(line), nil
}

// Close the UNIX socket connection
func (a *Audisp) Close() error {
	return a.conn.Close()
}
