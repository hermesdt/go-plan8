package main

import (
	"fmt"
	"os"

	chip8 "github.com/hermesdt/go-plan8"
)

func main() {
	// app := app.New()

	// w := app.NewWindow("Hello")
	// w.SetContent(widget.NewVBox(
	// 	widget.NewLabel("Hello Fyne!"),
	// 	widget.NewButton("Quit", func() {
	// 		app.Quit()
	// 	}),
	// ))

	// w.Canvas().SetOnTypedRune(func(r rune) {
	// 	fmt.Println("read rune", r)
	// })

	// w.ShowAndRun()

	c := chip8.NewChip8()
	c.LoadRom(os.Args[1])

	for {
		o := c.FetchOpcode()
		o.Execute()
		fmt.Println(c.Draw())
		// time.Sleep(10 * time.Millisecond)
	}
}
