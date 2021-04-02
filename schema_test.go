package res_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/suisrc/res.zgo"
)

func TestIsResponseError(t *testing.T) {
	res := &res.Success{
		Success: true,
	}
	assert.NotNil(t, res)

	// b := result.FixResponseError(nil, res)
	// assert.True(t, b)
	// err := errors.New("hello")
	// b = result.FixResponseError(nil, err)
	// assert.False(t, b)
	// b = result.FixResponseError(nil, result.Err403Forbidden)
	// assert.False(t, b)
	// b = result.FixResponseError(nil, result.ResError(nil, result.Err403Forbidden))
	// assert.True(t, b)
}
