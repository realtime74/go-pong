package controls

import (
	"github.com/gdamore/tcell/v2"
)

type Ball struct {
	x, y   int
	dx, dy int

	screen tcell.Screen
}

func NewBall(screen tcell.Screen, x, y int) *Ball {
	b := &Ball{
		x:      x,
		y:      y,
		dx:     1,
		dy:     0,
		screen: screen,
	}

	return b
}

func (b *Ball) NextPosition() (x, y int) {
	return b.x + b.dx, b.y + b.dy
}

func (b *Ball) Move() {
	b.Clear()
	b.x, b.y = b.NextPosition()
	b.Draw()
}

func (b *Ball) Bounce() {
	b.dx = -b.dx
	b.dy = -b.dy
}

func (b *Ball) Position() (dx, dy int) {
	return b.x, b.y
}

func (b *Ball) Clear() {
	b.screen.SetContent(b.x, b.y, ' ', nil, tcell.StyleDefault)
}

func (b *Ball) Draw() {
	b.screen.SetContent(b.x, b.y, '@', nil, tcell.StyleDefault)
}
