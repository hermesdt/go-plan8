package chip8

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestScreenDraw(t *testing.T) {
	var screen Screen
	screen.DrawByte(0, 8, 0xff)
	screen.DrawByte(1, 8, 0xc3)
	screen.DrawByte(2, 8, 0xc3)
	screen.DrawByte(3, 8, 0xff)

	assert.Equal(t, screen.GetByte(0, 8), byte(0xff))
	assert.Equal(t, screen.GetByte(1, 8), byte(0xc3))
	assert.Equal(t, screen.GetByte(2, 8), byte(0xc3))
	assert.Equal(t, screen.GetByte(3, 8), byte(0xff))
}
