package http

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_IsStatusCode2XX(t *testing.T) {
	assert.True(t, IsStatusCode2XX(200))
	assert.True(t, IsStatusCode2XX(299))
	assert.True(t, IsStatusCode2XX(250))

	assert.False(t, IsStatusCode2XX(199))
	assert.False(t, IsStatusCode2XX(300))
	assert.False(t, IsStatusCode2XX(400))
	assert.False(t, IsStatusCode2XX(500))
	assert.False(t, IsStatusCode2XX(95))
}
