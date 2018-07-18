package generator

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewsMatrix(t *testing.T) {
	m := NewMatrix(2, 3)
	assert.Equal(t, 2*3, len(m.Vals))
	assert.Equal(t, 2*3, cap(m.Vals))
	assert.Equal(t, 2, m.Stride)
}

func TestWidth(t *testing.T) {
	m := NewMatrix(2, 3)
	assert.Equal(t, 2, m.Width())
}

func TestHeight(t *testing.T) {
	m := NewMatrix(2, 3)
	assert.Equal(t, 3, m.Height())
}

func TestValOffset(t *testing.T) {
	w, h := 2, 3
	tests := []struct {
		x        int
		y        int
		expected int
	}{
		{0, 0, 0},
		{1, 1, 3},
		{-1, 1, -1},
		{9, 9, -1},
		{4, 0, -1},
	}

	for _, test := range tests {
		m := NewMatrix(w, h)
		actual := m.ValOffset(test.x, test.y)
		assert.Equal(t, test.expected, actual)
	}
}
