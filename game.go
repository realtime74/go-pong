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
	game.status.SetLevel(game.ball.Level)
	game.status.Draw()

	return game
}

func (g *Game) Start() {
	go g._controller()
	go ComputerMove(g, g.lracket)
	go ComputerMove(g, g.rracket)
}

func (g *Game) CheckBounds(tick int) {
	width, height := g.screen.Size()
	x, y := g.ball.NextPosition(g.ticker)

	if x <= 0 {
		g.status.Score(0, 1)
		if g.ball.Yboost > 1 {
			g.ball.Yboost -= 1
		}
		g.ball.Bounce(tick, -1, 1)
		return
	}
	if x >= width {
		g.status.Score(1, 0)
		if g.ball.Yboost > 1 {
			g.ball.Yboost -= 1
		}
		g.ball.Bounce(tick, -1, 1)
		return
	}
	if y <= 0 || y >= height-1 {
		g.ball.Bounce(tick, 1, -1)
		if g.ball.Yboost > 1 {
			g.ball.Yboost -= 1
		}
		return
	}

	dx, dy := -1, 1
	if g.rracket.OnRacket(x, y) {
		if g.ticker-g.rracket.LastMove < 100 {
			g.ball.Yboost += 5
		} else {
			g.ball.Yboost = 1
		}
		g.ball.Bounce(tick, dx, dy)
		return
	}
	if g.lracket.OnRacket(x, y) {
		dx, dy := -1, 1
		if g.ticker-g.lracket.LastMove < 100 {
			g.ball.Yboost += 5
		} else {
			g.ball.Yboost = 1
		}
		g.ball.Bounce(tick, dx, dy)
		return
	}
}

func (g *Game) _controller() {
	ticker := time.NewTicker(10 * time.Millisecond)
	for range ticker.C {
		g.ticker++
		g.status.SetTicker(g.ticker / 100)

		// level up every 3000 ticks
		if g.ticker%3000 == 0 {
			g.ball.Level += 1
			g.status.SetLevel(g.ball.Level)
		}
		// Move the ball every 5 ticks (50ms)
		if g.ticker%5 == 0 {
			g.CheckBounds(g.ticker)
			g.ball.Move(g.ticker)
			g.screen.Show()
			g.ticker++
		}
	}
}
