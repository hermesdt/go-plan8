package chip8

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
)

func (c *Chip8) LoadRom(filename string) {
	bs, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}

	rom := c.Memory[0x200:]
	write := bytes.NewBuffer(rom)
	write.Reset()
	n, err := write.Write(bs)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("copied %d bytes\n", n)
}
