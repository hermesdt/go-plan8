package chip8

import (
	"fmt"
	"strings"
)

type Chip8 struct {
	Screen     [64 * 32]uint8
	Memory     [4096]uint8
	V          [16]uint8
	I          uint16
	PC         uint16
	SP         uint16
	Stack      [16]uint16
	Key        [16]uint8
	DelayTimer uint8
	SoundTimer uint8
}

func (c *Chip8) FetchOpcode() Opcode {
	return Opcode{
		Value: uint16(c.Memory[c.PC])<<8 | uint16(c.Memory[c.PC+1]),
		Chip8: c,
	}
}

func (c *Chip8) Draw() string {
	lines := []string{}
	for row := 0; row < 32; row++ {
		line := []string{}
		for col := 0; col < 8; col++ {
			line = append(line, fmt.Sprintf("%08b", c.Screen[row*8+col]))
		}
		joinedLine := strings.Join(line, "")
		lines = append(lines, strings.ReplaceAll(joinedLine, "0", " "))
	}
	return strings.Join(lines, "\n")
}

func NewChip8() *Chip8 {
	c := &Chip8{
		PC: 0x200,
	}
	c.LoadFontSet()
	return c
}
