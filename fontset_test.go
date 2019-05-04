package chip8

import "fmt"
import "testing"
import "github.com/stretchr/testify/assert"

func TestFontSet_0(t *testing.T) {
	chip8 := Chip8{}
	chip8.LoadFontSet()
	character := 0
	digit_offset := character*5 + 0x050

	digit := fmt.Sprintf("%08b\n%08b\n%08b\n%08b\n%08b",
		chip8.Memory[digit_offset+0],
		chip8.Memory[digit_offset+1],
		chip8.Memory[digit_offset+2],
		chip8.Memory[digit_offset+3],
		chip8.Memory[digit_offset+4])

	assert.Equal(t, digit,
		"11110000\n"+
			"10010000\n"+
			"10010000\n"+
			"10010000\n"+
			"11110000")
}

func TestFontSet_1(t *testing.T) {
	chip8 := Chip8{}
	chip8.LoadFontSet()
	character := 1
	digit_offset := character*5 + 0x050

	digit := fmt.Sprintf("%08b\n%08b\n%08b\n%08b\n%08b",
		chip8.Memory[digit_offset+0],
		chip8.Memory[digit_offset+1],
		chip8.Memory[digit_offset+2],
		chip8.Memory[digit_offset+3],
		chip8.Memory[digit_offset+4])

	assert.Equal(t, digit,
		"00100000\n"+
			"01100000\n"+
			"00100000\n"+
			"00100000\n"+
			"01110000")
}
