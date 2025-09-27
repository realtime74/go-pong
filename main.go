package main

import (
	"fmt"
	"time"

	"github.com/gdamore/tcell/v2"

	"github.com/realtime74/gopong/controls"
	"github.com/realtime74/gopong/scr"
)

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
	scr.DrawText(screen, width/2-len(text)/2, height/2, text, style)

	titlebar := controls.NewTitleBar(screen, "go-pong")
	titlebar.Draw()

	screen.Show()
	time.Sleep(10 * time.Second)
	fmt.Println("Hello, World!")
}
