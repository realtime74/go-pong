package controls

import (
	"time"

	"github.com/gdamore/tcell/v2"
	"github.com/realtime74/gopong/scr"
)

type Wall struct {
	X int

	screen tcell.Screen
}

func NewWall(screen tcell.Screen, x int) Wall {
	return Wall{
		X:      x,
		screen: screen,
	}
}

func (w *Wall) Draw() {
	_, height := w.screen.Size()
	style := tcell.StyleDefault.
		Background(_Color).
		Foreground(_Color)

	scr.HFill(w.screen, w.X, 1, height-2, ' ', style)
}

func (w *Wall) Flash(gate bool) {
	if !gate {
		return
	}
	_, height := w.screen.Size()
	style := tcell.StyleDefault.
		Background(tcell.ColorRed).
		Foreground(tcell.ColorRed)
	scr.HFill(w.screen, w.X, 1, height-2, ' ', style)
	w.screen.Show()
	go w._resetFlash()
}

func (w *Wall) _resetFlash() {
	time.Sleep(250 * time.Millisecond)
	w.Draw()
}
