package audit

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

type Audit struct {
	TimeStamp time.Time
	ID        uint
}

func Parse(txt string) (*Audit, error) {
	if !strings.HasPrefix(txt, "audit(") {
		return nil, fmt.Errorf("not an audit msg : [%s]", txt)
	}
	blobs := strings.Split(txt[6:len(txt)-2], ":")
	id, err := strconv.ParseInt(blobs[1], 10, 64)
	if err != nil {
		return nil, err
	}
	ts, err := strconv.ParseFloat(blobs[0], 64)
	if err != nil {
		return nil, err
	}
	return &Audit{
		TimeStamp: time.UnixMicro(int64(ts * 1000000)),
		ID:        uint(id),
	}, nil

}
