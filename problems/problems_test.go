package problems

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEuler1(t *testing.T) {
	assert.Equal(t, 233168, Euler1(1000))
}

func TestEuler2(t *testing.T) {
	assert.Equal(t, 4613732, Euler2(4000000))
}
