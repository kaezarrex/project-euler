package problems

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEuler1(t *testing.T) {
	assert.Equal(t, 23, Euler1(10))
	assert.Equal(t, 233168, Euler1(1000))
}

func TestEuler2(t *testing.T) {
	assert.Equal(t, 4613732, Euler2(4000000))
}

func TestEuler3(t *testing.T) {
	assert.Equal(t, 29, Euler3(13195))
	assert.Equal(t, 6857, Euler3(600851475143))
}

func TestEuler4(t *testing.T) {
	assert.Equal(t, 9009, Euler4(2))
	assert.Equal(t, 906609, Euler4(3))
}

func TestEuler5(t *testing.T) {
	assert.Equal(t, 2520, Euler5(10))
	assert.Equal(t, 232792560, Euler5(20))
}

func TestEuler6(t *testing.T) {
	assert.Equal(t, 2640, Euler6(10))
	assert.Equal(t, 25164150, Euler6(100))
}

func TestEuler7(t *testing.T) {
	assert.Equal(t, 13, Euler7(6))
	assert.Equal(t, 104743, Euler7(10001))
}

func TestEuler8(t *testing.T) {
	assert.Equal(t, 5832, Euler8(4))
	assert.Equal(t, 23514624000, Euler8(13))
}

func TestEuler9(t *testing.T) {
	assert.Equal(t, 31875000, Euler9(1000))
}

func TestEuler10(t *testing.T) {
	assert.Equal(t, 17, Euler10(10))
	assert.Equal(t, 142913828922, Euler10(2000000))
}

func TestEuler11(t *testing.T) {
	assert.Equal(t, 70600674, Euler11())
}
