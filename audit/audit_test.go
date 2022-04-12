package audit

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// audit(1364481363.243:24287):

func TestAudit(t *testing.T) {
	a, err := Parse("audit(1364481363.243:24287):")
	assert.NoError(t, err)
	fmt.Println(a)
	assert.Equal(t, 2013, a.TimeStamp.Year())
}
