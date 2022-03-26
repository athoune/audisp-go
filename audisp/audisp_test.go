package audisp

import (
	"net"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.uber.org/goleak"
)

func TestMain(m *testing.M) {
	goleak.VerifyTestMain(m)
}

func TestAudisp(t *testing.T) {
	folder := t.TempDir()
	l, err := net.Listen("unix", folder+"/socket")
	assert.NoError(t, err)
	msg := make(chan string, 10)
	stop := make(chan interface{})
	go func() {
		conn, err := l.Accept()
		assert.NoError(t, err)
		for {
			select {
			case m := <-msg:
				conn.Write([]byte(m + "\n"))
			case <-stop:
				return // stop the goroutine
			}
		}
	}()
	a, err := New(folder + "/socket")
	assert.NoError(t, err)
	defer a.Close()
	msg <- "app=boo"
	msg <- "name='Bon Sinclar"
	line, err := a.Line()
	assert.NoError(t, err)
	cpt := 0
	keys := []string{"app", "name"}
	for line.Next() {
		assert.NoError(t, line.Error())
		k, _ := line.KeyValue()
		assert.Equal(t, keys[cpt], k)
		cpt++
	}
	stop <- nil

}
