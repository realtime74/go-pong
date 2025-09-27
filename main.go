package main

import (
	"fmt"

	"github.com/gdamore/tcell/v2"

	"github.com/realtime74/gopong/controls"
	"github.com/realtime74/gopong/scr"
)

type Game struct {
	screen  tcell.Screen
	lracket *controls.Racket
	rracket *controls.Racket
	title   controls.TitleBar
}

func _loop(game Game) bool {
	game.screen.Show()
	ev := game.screen.PollEvent()
	switch ev := ev.(type) {
	case *tcell.EventKey:
		switch ev.Key() {
		case tcell.KeyEscape, tcell.KeyCtrlC:
			return false
		case tcell.KeyUp:
			game.rracket.Move(-1)
		}
		switch ev.Rune() {
		case 'q':
			return false
		case 'k':
			game.rracket.Move(-1)
		case 'j':
			game.rracket.Move(1)
		}
	}

	return true
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
	scr.DrawText(screen, width/2-len(text)/2, height/2, text, style)

	game := Game{}
	game.screen = screen

	game.title = *controls.NewTitleBar(screen, "go-pong")
	game.title.Draw()

	game.rracket = controls.NewRacket(screen, width-2, height/2)
	game.rracket.Draw()
	game.lracket = controls.NewRacket(screen, 1, height/2)
	game.lracket.Draw()

	for {
		if !_loop(game) {
			return
		}
	}

	fmt.Println("Hello, World!")
}
