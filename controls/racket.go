package controls

import (
	"github.com/gdamore/tcell/v2"
	"github.com/realtime74/gopong/scr"
)

type Racket struct {
	x, y int

	screen tcell.Screen
}

const RacketHeight = 3

func NewRacket(screen tcell.Screen, x, y int) *Racket {
	return &Racket{
		x:      x,
		y:      y,
		screen: screen,
	}
}

func (r *Racket) Position() (dx, dy int) {
	return r.x, r.y
}

func (r *Racket) Move(dy int) {
	r.Clear()

	_, height := r.screen.Size()
	r.y += dy
	miny := 1 + RacketHeight/2
	maxy := height - 2 - RacketHeight/2
	if r.y <= miny {
		r.y = miny
	}
	if r.y >= maxy {
		r.y = maxy
	}
	r.Draw()
}

func (r *Racket) OnRacket(x, y int) bool {
	if x != r.x {
		return false
	}
	return y >= r.y-RacketHeight/2 && y <= r.y+RacketHeight/2
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
