package pkg

import (
	"fmt"
)

type Opcode struct {
	Value uint16
	Chip8 *Chip8
}

func (o *Opcode) execute() {
	switch {
	case o.Value == 0x00E0:
		o.DispClr()
	case o.Value == 0x00EE:
		o.Return()
	case o.Value>>12 == 0x0:
		o.Call()
	case o.Value>>12 == 0x1:
		o.Jump()
	case o.Value>>12 == 0x2:
		o.CallSub()
	case o.Value>>12 == 0x3:
		o.SkipEq()
	case o.Value>>12 == 0x4:
		o.SkipNeq()
	case o.Value>>12 == 0x5:
		o.SkipEqVY()
	case o.Value>>12 == 0x6:
		o.Set()
	case o.Value>>12 == 0x7:
		o.Add()
	case o.Value&0xF00F == 0x8000:
		o.SetVY()
	case o.Value&0xF00F == 0x8001:
		o.OrVY()
	case o.Value&0xF00F == 0x8002:
		o.AndVY()
	case o.Value&0xF00F == 0x8003:
		o.XorVY()
	case o.Value&0xF00F == 0x8004:
		o.AddVY()
	case o.Value&0xF00F == 0x8005:
		o.SubVY()
	case o.Value&0xF00F == 0x8006:
		o.ShiftRight()
	case o.Value&0xF00F == 0x8007:
		o.VYSub()
	case o.Value&0xF00F == 0x800E:
		o.ShiftLeft()
	case o.Value>>12 == 0x9:
		o.SkipNeqVY()
	case o.Value>>12 == 0xA:
		o.SetI()
	case o.Value>>12 == 0xB:
		o.JumpPlusV0()
	case o.Value>>12 == 0xC:
		o.SetRandomMask()
	case o.Value>>12 == 0xD:
		o.Draw()
	case o.Value&0xF0FF == 0xE09E:
		o.SkipKeyPressed()
	case o.Value&0xF0FF == 0xE0A1:
		o.SkipNotKeyPressed()
	case o.Value&0xF0FF == 0xF007:
		o.SetFromDelay()
	case o.Value&0xF0FF == 0xF00A:
		o.ReadKey()
	case o.Value&0xF0FF == 0xF015:
		o.SetDelay()
	case o.Value&0xF0FF == 0xF018:
		o.SetSound()
	case o.Value&0xF0FF == 0xF01E:
		o.AddI()
	case o.Value&0xF0FF == 0xF029:
		o.SetISprite()
	case o.Value&0xF0FF == 0xF033:
		o.SetBCD()
	case o.Value&0xF0FF == 0xF055:
		o.RegDump()
	case o.Value&0xF0FF == 0xF065:
		o.RegLoad()
	default:
		panic(fmt.Sprintf("unkown opcode %+v", o.Value))
	}
}

func (o *Opcode) Call() {
	panic("not implemented yet")
}
func (o *Opcode) DispClr() {
	for i := range o.Chip8.Screen {
		o.Chip8.Screen[i] = 0
	}
	o.Chip8.PC += 2
}
func (o *Opcode) Return() {
	o.Chip8.SP--
	o.Chip8.PC = o.Chip8.Stack[o.Chip8.SP]
}
func (o *Opcode) Jump() {
	address := o.Value & 0x0FFF
	o.Chip8.PC = address
}
func (o *Opcode) CallSub() {
	o.Chip8.Stack[o.Chip8.SP] = o.Chip8.PC
	o.Chip8.SP++
	o.Chip8.PC = o.Value & 0x0FFF
}
func (o *Opcode) SkipEq() {
	n := uint8(o.Value & 0x00FF)
	x := (o.Value & 0x0F00) >> 8
	if o.Chip8.V[x] == n {
		o.Chip8.PC += 4
		return
	}
	o.Chip8.PC += 2
}
func (o *Opcode) SkipNeq() {
	n := uint8(o.Value & 0x00FF)
	x := (o.Value & 0x0F00) >> 8
	if o.Chip8.V[x] != n {
		o.Chip8.PC += 4
		return
	}
	o.Chip8.PC += 2
}
func (o *Opcode) SkipEqVY() {
	x := (o.Value & 0x0F00) >> 8
	y := (o.Value & 0x00F0) >> 4
	if o.Chip8.V[x] == o.Chip8.V[y] {
		o.Chip8.PC += 4
		return
	}
	o.Chip8.PC += 2
}
func (o *Opcode) Set() {
	n := uint8(o.Value & 0x00FF)
	x := (o.Value & 0x0F00) >> 8
	o.Chip8.V[x] = n
	o.Chip8.PC += 2
}
func (o *Opcode) Add() {
	n := uint8(o.Value & 0x00FF)
	x := (o.Value & 0x0F00) >> 8
	o.Chip8.V[x] += n
	o.Chip8.PC += 2
}
func (o *Opcode) SetVY() {
	x := (o.Value & 0x0F00) >> 8
	y := (o.Value & 0x00F0) >> 4
	o.Chip8.V[x] = o.Chip8.V[y]
	o.Chip8.PC += 2
}
func (o *Opcode) OrVY() {
	x := (o.Value & 0x0F00) >> 8
	y := (o.Value & 0x00F0) >> 4
	o.Chip8.V[x] |= o.Chip8.V[y]
	o.Chip8.PC += 2
}
func (o *Opcode) AndVY() {
	x := (o.Value & 0x0F00) >> 8
	y := (o.Value & 0x00F0) >> 4
	o.Chip8.V[x] &= o.Chip8.V[y]
	o.Chip8.PC += 2
}
func (o *Opcode) XorVY() {
	x := (o.Value & 0x0F00) >> 8
	y := (o.Value & 0x00F0) >> 4
	o.Chip8.V[x] ^= o.Chip8.V[y]
	o.Chip8.PC += 2
}
func (o *Opcode) AddVY() {
	x := (o.Value & 0x0F00) >> 8
	y := (o.Value & 0x00F0) >> 4
	o.Chip8.V[0xF] = 0x0
	hasCarray := (0xFF - o.Chip8.V[x]) < o.Chip8.V[y]
	if hasCarray {
		o.Chip8.V[0xF] = 0x1
	}
	o.Chip8.V[x] += o.Chip8.V[y]
	o.Chip8.PC += 2
}
func (o *Opcode) SubVY() {
	panic("not implemented yet")
}
func (o *Opcode) ShiftRight() {
	panic("not implemented yet")
}
func (o *Opcode) VYSub() {
	panic("not implemented yet")
}
func (o *Opcode) ShiftLeft() {
	panic("not implemented yet")
}
func (o *Opcode) SkipNeqVY() {
	panic("not implemented yet")
}
func (o *Opcode) SetI() {
	panic("not implemented yet")
}
func (o *Opcode) JumpPlusV0() {
	panic("not implemented yet")
}
func (o *Opcode) SetRandomMask() {
	panic("not implemented yet")
}
func (o *Opcode) Draw() {
	panic("not implemented yet")
}
func (o *Opcode) SkipKeyPressed() {
	panic("not implemented yet")
}
func (o *Opcode) SkipNotKeyPressed() {
	panic("not implemented yet")
}
func (o *Opcode) SetFromDelay() {
	panic("not implemented yet")
}
func (o *Opcode) ReadKey() {
	panic("not implemented yet")
}
func (o *Opcode) SetDelay() {
	panic("not implemented yet")
}
func (o *Opcode) SetSound() {
	panic("not implemented yet")
}
func (o *Opcode) AddI() {
	panic("not implemented yet")
}
func (o *Opcode) SetISprite() {
	panic("not implemented yet")
}
func (o *Opcode) SetBCD() {
	panic("not implemented yet")
}
func (o *Opcode) RegDump() {
	panic("not implemented yet")
}
func (o *Opcode) RegLoad() {
	panic("not implemented yet")
}
