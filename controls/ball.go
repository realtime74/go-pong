package controls

import (
	"github.com/gdamore/tcell/v2"
)

type Ball struct {
	x, y   int
	dx, dy int
	bx, by int
	t0     int
	Level  int
	Yboost int

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
		Level:  3,
		screen: screen,
	}

	return b
}

func (b *Ball) ResetAngle() {
	b.dx = (b.dx / b.dx) * 10
	b.dy = b.dy / b.dy
}

// would be easier if ball bounces itself

// otherwise we need to check outside for collision, then

// fix the next position based on:
//    - the position on the bounce
//    - calculate the y position for this x position

// at the end:
//    - we need to find the collision point generically
//      in the game zone, more reliable for future extensions
// howto:
//    walk from the current position with tick increments until
//    - you hit a wall
//    - you hit a racket
//    - later: obstacles
//    return coordinates of the bounce point

// return the next step (+1) movement of the ball
// from the last position up to the current position

func (b *Ball) NextPosition(tick int) (x, y int) {

	// usually 5
	delta_tick := tick - b.t0

	//          5            /   100 (at level 10) => 0.05
	//          5            /   20 (at level 50)  => 0.2
	//
	t := float64(delta_tick) / (1000.0 / float64(b.Level))

	// dx = 1, level = 10, t = 0.05
	// => 1 * 0.05 = 0.05 each call
	// dx = 1, level = 50, t = 0.2
	// => 1 * 0.2 = 0.2 each call => up to level 50 still 5 scans per cell
	// should then never jump more than 1 cell per call

	dx := int(float64(b.dx) * t)

	dy := int(float64(b.dy*b.Yboost) * t)

	return b.bx + dx, b.by + dy
}

func (b *Ball) Move(ticker int) {
	b.Clear()
	b.x, b.y = b.NextPosition(ticker)
	b.Draw()
}

func (b *Ball) Bounce(tick int, dx, dy int) {
	b.bx = b.x
	b.by = b.y
	b.dx *= dx
	b.dy *= dy
	if b.dy > 100 {
		b.dy = 100
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
