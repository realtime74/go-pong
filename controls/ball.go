package controls

import (
	"github.com/gdamore/tcell/v2"
)

type Ball struct {
	x, y   int
	dx, dy int
	bx, by int
	t0     int

	screen tcell.Screen
}

func NewBall(screen tcell.Screen, x, y int) *Ball {
	b := &Ball{
		x:      x,
		y:      y,
		bx:     x,
		by:     y,
		dx:     10,
		dy:     1,
		screen: screen,
	}

	return b
}

func (b *Ball) NextPosition(tick int) (x, y int) {
	t := float64(tick-b.t0) / 10.0
	dx := int(float64(b.dx) * t)
	dy := int(float64(b.dy) * t)

	return b.bx + dx, b.by + dy
}

func (b *Ball) Move(ticker int) {
	b.Clear()
	b.x, b.y = b.NextPosition(ticker)
	b.Draw()
}

func (b *Ball) Bounce(tick int, x, y bool) {
	b.bx = b.x
	b.by = b.y
	if x {
		b.dx = -b.dx
	}
	if y {
		b.dy = -b.dy
	}
	b.t0 = tick
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
