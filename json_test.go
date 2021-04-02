package res_test

import (
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
	res "github.com/suisrc/res.zgo"
)

func TestJson(t *testing.T) {
	re0 := &res.Success{
		Success: true,
	}
	buf, _ := res.JSONMarshal(re0)
	log.Println(buf)
	assert.NotNil(t, nil)
}
