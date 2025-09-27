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
}

func (g *Game) CheckBounds() {
	width, height := g.screen.Size()
	x, y := g.ball.NextPosition()
	if x <= 0 || x >= width || y <= 0 || y >= height-1 {
		g.status.Score(x >= width, x <= 0)
		g.ball.Bounce()
		return
	}
	if g.rracket.OnRacket(x, y) || g.lracket.OnRacket(x, y) {
		g.ball.Bounce()
		return
	}
}

func (g *Game) _controller() {
	ticker := time.NewTicker(50 * time.Millisecond)
	for range ticker.C {
		g.CheckBounds()
		g.ball.Move()
		g.screen.Show()
	}
}
