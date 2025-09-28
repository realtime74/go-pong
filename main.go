package main

import (
	"fmt"
	"os"

	"github.com/gdamore/tcell/v2"
)

func _loop(game *Game) bool {
	game.screen.Show()
	ev := game.screen.PollEvent()
	switch ev := ev.(type) {
	case *tcell.EventKey:
		switch ev.Key() {
		case tcell.KeyEscape, tcell.KeyCtrlC:
			return false
		}
		switch ev.Rune() {
		case 'q':
			return false
		case 'k':
			game.rracket.Move(game.ticker, -1)
		case 'j':
			game.rracket.Move(game.ticker, 1)
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

	options := GameOptions{
		computerPlayers: 1,
		startLevel:      3,
	}

	for ii, value := range os.Args[1:] {
		switch value {
		case "-1":
			options.computerPlayers = 1
		case "-2":
			options.computerPlayers = 2
		case "--level":
			if len(os.Args) > ii+2 {
				fmt.Sscanf(os.Args[ii+2], "%d", &options.startLevel)
			}
		}
	}

	game := NewGame(screen, options)
	game.Start()
	server := NewRESTServer(&game)
	server.Start()

	for {
		if !_loop(&game) {
			return
		}
	}
}
