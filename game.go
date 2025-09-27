package main

import (
	"time"

	"github.com/gdamore/tcell/v2"
	"github.com/realtime74/gopong/controls"
	"github.com/realtime74/gopong/scr"
)

type Game struct {
	title controls.TitleBar

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
	return game
}

func (g *Game) Start() {
	go g._controller()
}

func (g *Game) CheckBounds() {
	width, height := g.screen.Size()
	x, y := g.ball.NextPosition()
	if x <= 0 || x >= width || y <= 0 || y >= height-1 {
		scr.Flash(g.screen)
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
