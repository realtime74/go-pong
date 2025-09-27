package scr

import (
	"time"
	"github.com/gdamore/tcell/v2"
)

func Fill(
	screen tcell.Screen,
	x, y, width int,
	ch rune,
	style tcell.Style) {
	for i := 0; i < width; i++ {
		screen.SetContent(x+i, y, ch, nil, style)
	}
}

func HFill(
	screen tcell.Screen,
	x, y, height int,
	ch rune,
	style tcell.Style) {
	for i := 0; i < height; i++ {
		screen.SetContent(x, y+i, ch, nil, style)
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

func Flash(screen tcell.Screen) {
	width, height := screen.Size()
	style := tcell.StyleDefault.
		Background(tcell.ColorRed).
		Foreground(tcell.ColorRed)
	Fill(screen, 0, height-1, width, '!', style)
	screen.Show()

	go func() {
		time.Sleep(300 * time.Millisecond)
		Fill(screen, 0, height-1, width, ' ', tcell.StyleDefault)
		screen.Show()
	}()
}
