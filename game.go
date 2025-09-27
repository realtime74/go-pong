package main

import (
	"time"

	"github.com/gdamore/tcell/v2"
	"github.com/realtime74/gopong/controls"
)

type Game struct {
	title  controls.TitleBar
	status controls.StatusLine

	lracket *controls.Racket
	rracket *controls.Racket

	ball *controls.Ball

	screen tcell.Screen
	ticker int
}

func NewGame(screen tcell.Screen) Game {
	game := Game{}
	game.screen = screen

	width, height := screen.Size()

	game.title = *controls.NewTitleBar(screen, "go-pong")
	game.title.Draw()

	game.rracket = controls.NewRacket(screen, width-2, height/2)
	game.rracket.Draw()
	game.lracket = controls.NewRacket(screen, 1, height/2)
	game.lracket.Draw()
	game.ball = controls.NewBall(screen, width/2, height/2)
	game.ball.Draw()
	game.status = *controls.NewStatusLine(screen)
	game.status.Draw()

	return game
}

func (g *Game) Start() {
	go g._controller()
	go ComputerMove(g)
}

func (g *Game) CheckBounds(tick int) {
	width, height := g.screen.Size()
	x, y := g.ball.NextPosition(g.ticker)

	if x <= 0 {
		g.status.Score(0, 1)
		//g.ball.ResetAngle()
		g.ball.Bounce(tick, -1, 1)
		return
	}
	if x >= width {
		g.status.Score(1, 0)
		//g.ball.ResetAngle()
		g.ball.Bounce(tick, -1, 1)
		return
	}
	if y <= 0 || y >= height-1 {
		g.ball.Bounce(tick, 1, -1)
		return
	}

	if g.rracket.OnRacket(x, y) {
		dx, dy := -1, 1
		if g.ticker-g.rracket.LastMove < 500 {
			dx, dy = -1, 2
		} else {
			//g.ball.ResetAngle()
		}
		g.ball.Bounce(tick, dx, dy)
		return
	}
	if g.lracket.OnRacket(x, y) {
		dx, dy := -1, 1
		if g.ticker-g.lracket.LastMove < 500 {
			dx, dy = -1, 2
		} else {
			//g.ball.ResetAngle()
		}
		g.ball.Bounce(tick, dx, dy)
		return
	}

}

func (g *Game) _controller() {
	ticker := time.NewTicker(10 * time.Millisecond)
	for range ticker.C {
		g.ticker++
		if g.ticker%5 == 0 {
			g.CheckBounds(g.ticker)
			g.ball.Move(g.ticker)
			g.screen.Show()
			g.ticker++
		}
	}
}
