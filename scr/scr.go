package scr

import "github.com/gdamore/tcell/v2"

func Fill(
	screen tcell.Screen,
	x, y, width int,
	ch rune,
	style tcell.Style) {
	for i := 0; i < width; i++ {
		screen.SetContent(x+i, y, ch, nil, style)
	}
}

func DrawText(
	screen tcell.Screen,
	x, y int, text string,
	style tcell.Style) {
	for i, r := range text {
		screen.SetContent(x+i, y, r, nil, style)
	}
}
