package chip8

type Chip8 struct {
	Screen     Screen
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

func NewChip8() *Chip8 {
	c := &Chip8{
		PC: 0x200,
	}
	c.LoadFontSet()
	return c
}
