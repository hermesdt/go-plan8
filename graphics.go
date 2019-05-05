package chip8

import (
	"strings"
)

const (
	ScreenWidth  int = 64
	ScreenHeight int = 32
)

var b2i = map[bool]uint8{false: 0, true: 1}

type Screen struct {
	px [ScreenWidth * ScreenHeight]bool
}

func (s *Screen) Set(y, x int, value bool) {
	s.px[y*ScreenWidth+x] = value
}

func (s *Screen) Get(y, x int) bool {
	return s.px[y*ScreenWidth+x]
}

func (s *Screen) GetByte(y, x int) byte {
	var b byte = 0
	for i := 0; i < 8; i++ {
		row := y
		col := (x + i) % ScreenWidth

		bit := b2i[s.Get(row, col)]
		b |= bit << (7 - uint(i))
	}
	return b
}

func (s *Screen) DrawByte(y, x int, b byte) bool {
	collision := false
	for i := 0; i < 8; i++ {
		row := y
		col := (x + i) % ScreenWidth
		mask := byte(0x80 >> uint(i))
		value := b2i[(b&mask)>>(7-uint(i)) == 1]
		oldValue := b2i[s.Get(row, col)]

		s.Set(row, col, oldValue^value == 1)
		if oldValue^value == 0 && oldValue == 1 {
			collision = collision || true
		}
	}
	return collision
}

func (s *Screen) Clear() {
	for y := 0; y < ScreenHeight; y++ {
		for x := 0; x < ScreenWidth; x++ {
			s.Set(y, x, false)
		}
	}
}

func (s *Screen) Render() string {
	lines := []string{}
	for row := 0; row < ScreenHeight; row++ {
		line := []string{}
		for col := 0; col < ScreenWidth; col++ {
			if s.Get(row, col) {
				line = append(line, "1")
			} else {
				line = append(line, "0")
			}

		}
		lines = append(lines, strings.Join(line, ""))
	}
	return strings.Join(lines, "\n")
}
