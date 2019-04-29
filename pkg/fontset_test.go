package chip8

import "fmt"
import "testing"
import "github.com/stretchr/testify/assert"

func TestFontSet_0(t *testing.T) {
	chip8 := Chip8{}
	chip8.LoadFontSet()
	offset := 0x050

	digit := fmt.Sprintf("%b\n%b\n%b\n%b\n%b",
		chip8.Memory[offset+0],
		chip8.Memory[offset+1],
		chip8.Memory[offset+2],
		chip8.Memory[offset+3],
		chip8.Memory[offset+4])

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

	digit := fmt.Sprintf("%b\n%b\n%b\n%b\n%b",
		chip8.Memory[0x050],
		chip8.Memory[0x051],
		chip8.Memory[0x052],
		chip8.Memory[0x053],
		chip8.Memory[0x054])

	assert.Equal(t, digit,
		"11110000"+
			"10010000"+
			"10010000"+
			"10010000"+
			"11110000")
}
