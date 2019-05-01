package chip8

import (
	"fmt"
	"strings"
)

type Chip8 struct {
	Screen [64 * 32]uint8
	Memory [4096]uint8
	V      [16]uint8
	I      uint16
	PC     uint16
	SP     uint16
	Stack  [16]uint16
}

func (c *Chip8) fetchOpcode() Opcode {
	return Opcode{
		Value: uint16(c.Memory[c.PC])<<8 | uint16(c.Memory[c.PC+1]),
		Chip8: c,
	}
}

func (c *Chip8) draw() string {
	lines := []string{}
	for row := 0; row < 32; row++ {
		line := []string{}
		for col := 0; col < 8; col++ {
			line = append(line, fmt.Sprintf("%08b", c.Screen[row*8+col]))
		}
		lines = append(lines, strings.Join(line, ""))
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
