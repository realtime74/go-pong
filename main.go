package main

import (
	"fmt"

	"github.com/gdamore/tcell/v2"
)

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

	style := tcell.StyleDefault.
		Background(tcell.ColorBlack).
		Foreground(tcell.ColorWhite)
	screen.SetStyle(style)

	game := NewGame(screen)
	game.Start()

	for {
		if !_loop(game) {
			return
		}
	}
}
