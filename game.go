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
	lWall   controls.Wall
	rWall   controls.Wall

	ball *controls.Ball

	screen tcell.Screen
	ticker int

	options GameOptions
}

type GameStatus struct {
	scoreLeft    int
	scoreRight   int
	ballX        int
	ballY        int
	screenHeight int
	screenWidth  int
}

type GameOptions struct {
	computerPlayers int
	startLevel      int
	levelCap        int
	scoreCap        int
}

func NewGame(screen tcell.Screen, opts GameOptions) Game {
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
	game.ball.Level = opts.startLevel
	game.ball.Draw()
	game.status = *controls.NewStatusLine(screen)
	game.status.SetLevel(game.ball.Level, false)
	game.status.Draw()
	game.lWall = controls.NewWall(screen, 0)
	game.rWall = controls.NewWall(screen, width-1)
	game.lWall.Draw()
	game.rWall.Draw()

	game.ticker = 0
	game.options = opts

	return game
}

func (g *Game) Start() {
	go g._controller()
	if g.options.computerPlayers > 0 {
		go ComputerMove(g, g.lracket)
	}
	if g.options.computerPlayers > 1 {
		go ComputerMove(g, g.rracket)
	}
}

func (g *Game) Status() GameStatus {
	gs := GameStatus{}

	gs.scoreLeft, gs.scoreRight = g.status.GetScore()
	gs.ballX, gs.ballY = g.ball.Position()
	gs.screenWidth, gs.screenHeight = g.screen.Size()
	return gs
}

func (g *Game) CheckBounds(tick int) {
	width, height := g.screen.Size()
	x, y := g.ball.NextPosition(g.ticker)

	// wall bounce
	leftBounce := x <= 0
	rightBounce := x >= width-1
	topBounce := y <= 0
	bottomBounce := y >= height-1

	if leftBounce {
		g.status.Score(0, 1)
	}
	if rightBounce {
		g.status.Score(1, 0)
	}
	if leftBounce || rightBounce {
		g.ball.Bounce(tick, -1, 1)
		g.lWall.Flash(leftBounce)
		g.rWall.Flash(rightBounce)
	}
	if topBounce || bottomBounce {
		g.ball.Bounce(tick, 1, -1)
	}
	if leftBounce || rightBounce || topBounce || bottomBounce {
		if g.ball.Yboost > 1 {
			g.ball.Yboost -= 1
		}
	}

	// racket bounce
	dx, dy := -1, 1
	rackets := []*controls.Racket{g.lracket, g.rracket}
	for _, racket := range rackets {
		if racket.OnRacket(x, y) {
			if g.ticker-racket.LastMove < 100 {
				g.ball.Yboost += 5
			} else {
				g.ball.Yboost = 1
			}
			g.ball.Bounce(tick, dx, dy)
			return
		}
	}
}

func (g *Game) _controller() {
	ticker := time.NewTicker(10 * time.Millisecond)
	for range ticker.C {
		g.ticker++

		if g.ball.Level > g.options.levelCap {
			g.screen.PostEvent(tcell.NewEventKey(tcell.KeyEscape, 0, tcell.ModNone))
		}
		ls, rs := g.status.GetScore()
		if ls >= g.options.scoreCap || rs >= g.options.scoreCap {
			g.screen.PostEvent(tcell.NewEventKey(tcell.KeyEscape, 0, tcell.ModNone))
		}

		// level up every 3000 ticks
		if g.ticker%3000 == 0 {
			g.ball.Level += 1
			g.status.SetLevel(g.ball.Level, true)
		}
		// Move the ball faster
		g.CheckBounds(g.ticker)
		g.ball.Move(g.ticker)
		g.screen.Show()
		g.ticker++
	}
}
