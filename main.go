package main

import (
	"fmt"
	"time"

	"github.com/gdamore/tcell/v2"
)

func drawText(screen tcell.Screen, x, y int, text string, style tcell.Style) {
	for i, r := range text {
		screen.SetContent(x+i, y, r, nil, style)
	}
}

func main() {
	screen, err := tcell.NewScreen()
	if err != nil {
		panic(fmt.Errorf("Error creating screen: %v", err))
	}
	defer screen.Fini()

	screen.Init()
	screen.Clear()

	width, height := screen.Size()

	style := tcell.StyleDefault.
		Background(tcell.ColorBlack).
		Foreground(tcell.ColorWhite)
	screen.SetStyle(style)

	text := "Hello, World!"
	drawText(screen, width/2-len(text)/2, height/2, text, style)

	screen.Show()
	time.Sleep(10 * time.Second)
	fmt.Println("Hello, World!")
}
