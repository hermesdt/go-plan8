package chip8

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDispClr(t *testing.T) {
	screen := [2048]uint8{}
	for i := 0; i < 3; i++ {
		screen[i*10] = 1
	}
	o := Opcode{
		Value: 0x00E0,
		Chip8: &Chip8{
			Screen: screen,
		},
	}

	o.execute()

	assert.Equal(t, o.Chip8.Screen[0], uint8(0x0))
	assert.Equal(t, o.Chip8.Screen[10], uint8(0x0))
	assert.Equal(t, o.Chip8.Screen[20], uint8(0x0))
}

func TestReturn(t *testing.T) {
	var stack [16]uint16
	var SP uint16 = 1
	var prevPC uint16 = 0xABCD
	stack[0] = prevPC

	o := Opcode{
		Value: 0x00EE,
		Chip8: &Chip8{
			PC:    0x1010,
			SP:    SP,
			Stack: stack,
		},
	}

	o.execute()

	assert.Equal(t, o.Chip8.PC, prevPC)
	assert.Equal(t, o.Chip8.SP, uint16(0x0))
}

func TestJump(t *testing.T) {
	o := Opcode{
		Value: 0x1E04,
		Chip8: &Chip8{
			PC: 0x1010,
		},
	}

	o.execute()

	assert.Equal(t, o.Chip8.PC, uint16(0x0E04))
}

func TestCallSub(t *testing.T) {
	var stack [16]uint16
	var SP uint16 = 0x2
	stack[SP] = 0x0111
	o := Opcode{
		Value: 0x2028,
		Chip8: &Chip8{
			PC:    0x1010,
			SP:    SP,
			Stack: stack,
		},
	}

	o.execute()

	assert.Equal(t, o.Chip8.PC, uint16(0x028))
	assert.Equal(t, o.Chip8.SP, uint16(0x3))
	assert.Equal(t, o.Chip8.Stack[o.Chip8.SP-1], uint16(0x1010))
}

func TestSkipEq_true(t *testing.T) {
	var v [16]uint8
	v[7] = uint8(0x42)
	var pc uint16 = 0x0010
	o := Opcode{
		Value: 0x3742,
		Chip8: &Chip8{
			V:  v,
			PC: pc,
		},
	}

	o.execute()

	assert.Equal(t, o.Chip8.PC, pc+4)
}

func TestSkipEq_false(t *testing.T) {
	var v [16]uint8
	v[7] = uint8(0x42)
	var pc uint16 = 0x0010
	o := Opcode{
		Value: 0x3750,
		Chip8: &Chip8{
			V:  v,
			PC: pc,
		},
	}

	o.execute()

	assert.Equal(t, o.Chip8.PC, pc+2)
}

func TestSkipNeq_true(t *testing.T) {
	var v [16]uint8
	v[7] = uint8(0x42)
	var pc uint16 = 0x0010
	o := Opcode{
		Value: 0x4750,
		Chip8: &Chip8{
			V:  v,
			PC: pc,
		},
	}

	o.execute()

	assert.Equal(t, o.Chip8.PC, pc+4)
}

func TestSkipNeq_false(t *testing.T) {
	var v [16]uint8
	v[7] = uint8(0x42)
	var pc uint16 = 0x0010
	o := Opcode{
		Value: 0x4742,
		Chip8: &Chip8{
			V:  v,
			PC: pc,
		},
	}

	o.execute()

	assert.Equal(t, o.Chip8.PC, pc+2)
}

func TestSkipEqVY_true(t *testing.T) {
	var v [16]uint8
	v[4] = uint8(0x42)
	v[7] = uint8(0x42)
	var pc uint16 = 0x0010
	o := Opcode{
		Value: 0x5740,
		Chip8: &Chip8{
			V:  v,
			PC: pc,
		},
	}

	o.execute()

	assert.Equal(t, o.Chip8.PC, pc+4)
}

func TestSkipEqVY_false(t *testing.T) {
	var v [16]uint8
	v[4] = uint8(0x42)
	v[7] = uint8(0x43)
	var pc uint16 = 0x0010
	o := Opcode{
		Value: 0x5740,
		Chip8: &Chip8{
			V:  v,
			PC: pc,
		},
	}

	o.execute()

	assert.Equal(t, o.Chip8.PC, pc+2)
}

func TestSet(t *testing.T) {
	var v [16]uint8
	var pc uint16 = 0x0010
	o := Opcode{
		Value: 0x6A11,
		Chip8: &Chip8{
			V:  v,
			PC: pc,
		},
	}

	o.execute()

	assert.Equal(t, o.Chip8.V[0xA], uint8(0x11))
	assert.Equal(t, o.Chip8.PC, pc+2)
}

func TestAdd_nocarry(t *testing.T) {
	var v [16]uint8
	var pc uint16 = 0x0010
	v[0] = 0x21
	o := Opcode{
		Value: 0x7032,
		Chip8: &Chip8{
			V:  v,
			PC: pc,
		},
	}

	o.execute()

	assert.Equal(t, o.Chip8.V[0x0], uint8(0x53))
	assert.Equal(t, o.Chip8.PC, pc+2)
}

func TestAdd_carry(t *testing.T) {
	var v [16]uint8
	var pc uint16 = 0x0010
	v[0] = 0xFA
	o := Opcode{
		Value: 0x7032,
		Chip8: &Chip8{
			V:  v,
			PC: pc,
		},
	}

	o.execute()

	assert.Equal(t, o.Chip8.V[0x0], uint8(0x2C))
	assert.Equal(t, o.Chip8.V[0xF], uint8(0x0))
	assert.Equal(t, o.Chip8.PC, pc+2)
}

func TestSetVY(t *testing.T) {
	var v [16]uint8
	var pc uint16 = 0x0010
	v[0x1] = 0x42
	v[0xD] = 0x91
	o := Opcode{
		Value: 0x81D0,
		Chip8: &Chip8{
			V:  v,
			PC: pc,
		},
	}

	o.execute()

	assert.Equal(t, o.Chip8.V[0x1], uint8(0x91))
	assert.Equal(t, o.Chip8.PC, pc+2)
}

func TestOrVY(t *testing.T) {
	var v [16]uint8
	var pc uint16 = 0x0010
	v[0x1] = 0x42
	v[0xD] = 0x91
	o := Opcode{
		Value: 0x81D1,
		Chip8: &Chip8{
			V:  v,
			PC: pc,
		},
	}

	o.execute()

	assert.Equal(t, o.Chip8.V[0x1], uint8((0x42)|(0x91)))
	assert.Equal(t, o.Chip8.PC, pc+2)
}

func TestAndVY(t *testing.T) {
	var v [16]uint8
	var pc uint16 = 0x0010
	v[0x1] = 0x42
	v[0xD] = 0x92
	o := Opcode{
		Value: 0x81D2,
		Chip8: &Chip8{
			V:  v,
			PC: pc,
		},
	}

	o.execute()

	assert.Equal(t, o.Chip8.V[0x1], uint8((0x42)&(0x92)))
	assert.Equal(t, o.Chip8.PC, pc+2)
}

func TestXorVY(t *testing.T) {
	var v [16]uint8
	var pc uint16 = 0x0010
	v[0x1] = 0x42
	v[0xD] = 0x91
	o := Opcode{
		Value: 0x81D3,
		Chip8: &Chip8{
			V:  v,
			PC: pc,
		},
	}

	o.execute()

	assert.Equal(t, o.Chip8.V[0x1], uint8((0x42)^(0x91)))
	assert.Equal(t, o.Chip8.PC, pc+2)
}

func TestAddVY_carry(t *testing.T) {
	var v [16]uint8
	var pc uint16 = 0x0010
	v[0x1] = 0xF2
	v[0xD] = 0x91
	o := Opcode{
		Value: 0x81D4,
		Chip8: &Chip8{
			V:  v,
			PC: pc,
		},
	}

	o.execute()

	assert.Equal(t, o.Chip8.V[0x1], uint8(0x83))
	assert.Equal(t, o.Chip8.V[0xF], uint8(0x1))
	assert.Equal(t, o.Chip8.PC, pc+2)
}

func TestAddVY_nocarry(t *testing.T) {
	var v [16]uint8
	var pc uint16 = 0x0010
	v[0x1] = 0x22
	v[0xD] = 0x91
	o := Opcode{
		Value: 0x81D4,
		Chip8: &Chip8{
			V:  v,
			PC: pc,
		},
	}

	o.execute()

	assert.Equal(t, o.Chip8.V[0x1], uint8(0xB3))
	assert.Equal(t, o.Chip8.V[0xF], uint8(0x0))
	assert.Equal(t, o.Chip8.PC, pc+2)
}

func TestSubVY_borrow(t *testing.T) {
	var v [16]uint8
	var pc uint16 = 0x0010
	v[0x1] = 0x42
	v[0xD] = 0x91
	o := Opcode{
		Value: 0x81D5,
		Chip8: &Chip8{
			V:  v,
			PC: pc,
		},
	}

	o.execute()

	assert.Equal(t, o.Chip8.V[0x1], uint8(0xb1))
	assert.Equal(t, o.Chip8.V[0xF], uint8(0x1))
	assert.Equal(t, o.Chip8.PC, pc+2)
}

func TestSubVY_noborrow(t *testing.T) {
	var v [16]uint8
	var pc uint16 = 0x0010
	v[0x1] = 0x42
	v[0xD] = 0x33
	o := Opcode{
		Value: 0x81D5,
		Chip8: &Chip8{
			V:  v,
			PC: pc,
		},
	}

	o.execute()

	assert.Equal(t, o.Chip8.V[0x1], uint8(0x0F))
	assert.Equal(t, o.Chip8.V[0xF], uint8(0x0))
	assert.Equal(t, o.Chip8.PC, pc+2)
}

func TestShiftRight(t *testing.T) {
	var v [16]uint8
	var pc uint16 = 0x0010
	v[0x3] = 0x42
	o := Opcode{
		Value: 0x8306,
		Chip8: &Chip8{
			V:  v,
			PC: pc,
		},
	}

	o.execute()

	assert.Equal(t, o.Chip8.V[0x3], uint8(0x21))
	assert.Equal(t, o.Chip8.V[0xF], uint8(0x0))
	assert.Equal(t, o.Chip8.PC, pc+2)
}

func TestShiftRight_second(t *testing.T) {
	var v [16]uint8
	var pc uint16 = 0x0010
	v[0x3] = 0x43
	o := Opcode{
		Value: 0x8306,
		Chip8: &Chip8{
			V:  v,
			PC: pc,
		},
	}

	o.execute()

	assert.Equal(t, o.Chip8.V[0x3], uint8(0x21))
	assert.Equal(t, o.Chip8.V[0xF], uint8(0x1))
	assert.Equal(t, o.Chip8.PC, pc+2)
}

func TestVYSub_borrow(t *testing.T) {
	var v [16]uint8
	var pc uint16 = 0x0010
	v[0x1] = 0x42
	v[0xD] = 0x33
	o := Opcode{
		Value: 0x81D7,
		Chip8: &Chip8{
			V:  v,
			PC: pc,
		},
	}

	o.execute()

	assert.Equal(t, o.Chip8.V[0x1], uint8(0xF1))
	assert.Equal(t, o.Chip8.V[0xF], uint8(0x1))
	assert.Equal(t, o.Chip8.PC, pc+2)
}

func TestVYSub_noborrow(t *testing.T) {
	var v [16]uint8
	var pc uint16 = 0x0010
	v[0x1] = 0x42
	v[0xD] = 0x50
	o := Opcode{
		Value: 0x81D7,
		Chip8: &Chip8{
			V:  v,
			PC: pc,
		},
	}

	o.execute()

	assert.Equal(t, o.Chip8.V[0x1], uint8(0x0e))
	assert.Equal(t, o.Chip8.V[0xF], uint8(0x0))
	assert.Equal(t, o.Chip8.PC, pc+2)
}

func TestShiftLeft(t *testing.T) {
	var v [16]uint8
	var pc uint16 = 0x0010
	v[0x3] = 0x43
	o := Opcode{
		Value: 0x830E,
		Chip8: &Chip8{
			V:  v,
			PC: pc,
		},
	}

	o.execute()

	assert.Equal(t, o.Chip8.V[0x3], uint8(0x86))
	assert.Equal(t, o.Chip8.V[0xF], uint8(0x0))
	assert.Equal(t, o.Chip8.PC, pc+2)
}

func TestShiftLeft_second(t *testing.T) {
	var v [16]uint8
	var pc uint16 = 0x0010
	v[0x3] = 0xA2
	o := Opcode{
		Value: 0x830E,
		Chip8: &Chip8{
			V:  v,
			PC: pc,
		},
	}

	o.execute()

	assert.Equal(t, o.Chip8.V[0x3], uint8(0x44))
	assert.Equal(t, o.Chip8.V[0xF], uint8(0x80))
	assert.Equal(t, o.Chip8.PC, pc+2)
}

func TestSkipNeqVY_true(t *testing.T) {
	var v [16]uint8
	v[2] = uint8(0x42)
	v[5] = uint8(0x41)
	var pc uint16 = 0x0010
	o := Opcode{
		Value: 0x9250,
		Chip8: &Chip8{
			V:  v,
			PC: pc,
		},
	}

	o.execute()

	assert.Equal(t, o.Chip8.PC, pc+4)
}

func TestSkipNeqVY_false(t *testing.T) {
	var v [16]uint8
	v[2] = uint8(0x42)
	v[5] = uint8(0x42)
	var pc uint16 = 0x0010
	o := Opcode{
		Value: 0x9250,
		Chip8: &Chip8{
			V:  v,
			PC: pc,
		},
	}

	o.execute()

	assert.Equal(t, o.Chip8.PC, pc+2)
}

func TestSetI(t *testing.T) {
	var pc uint16 = 0x0010
	o := Opcode{
		Value: 0xA911,
		Chip8: &Chip8{
			PC: pc,
		},
	}

	o.execute()

	assert.Equal(t, o.Chip8.I, uint16(0x0911))
	assert.Equal(t, o.Chip8.PC, pc+2)
}

func TestJumpPlusV0(t *testing.T) {
	var v [16]uint8
	v[0] = uint8(0x12)
	var pc uint16 = 0x0010
	o := Opcode{
		Value: 0xBF50,
		Chip8: &Chip8{
			V:  v,
			PC: pc,
		},
	}

	o.execute()

	assert.Equal(t, o.Chip8.PC, uint16(0x0F62))
}
