package core

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetFiber(t *testing.T) {
	assert.NotNil(t, GetFiber())
}
