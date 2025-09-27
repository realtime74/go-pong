package controls

import (
	"github.com/gdamore/tcell/v2"
	"github.com/realtime74/gopong/scr"
)

type Racket struct {
	x, y int

	screen tcell.Screen
}

const RacketHeight = 5

func NewRacket(screen tcell.Screen, x, y int) *Racket {
	return &Racket{
		x:      x,
		y:      y,
		screen: screen,
	}
}

func (r *Racket) Move(dy int) {
	r.Clear()

	_, height := r.screen.Size()
	r.y += dy
	miny := RacketHeight / 2
	maxy := height - RacketHeight/2
	if r.y < miny {
		r.y = miny
	}
	if r.y > maxy {
		r.y = maxy
	}
	r.Draw()
}

func (r *Racket) Clear() {
	scr.HFill(
		r.screen,
		r.x, r.y-RacketHeight/2, RacketHeight,
		' ', tcell.StyleDefault)
}

func (r *Racket) Draw() {
	scr.HFill(
		r.screen,
		r.x, r.y-RacketHeight/2, RacketHeight,
		'|', tcell.StyleDefault)
}
