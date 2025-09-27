package main

import (
	"math/rand"
	"time"

	"github.com/realtime74/gopong/controls"
)

func ComputerMove(g *Game, r *controls.Racket) {
	ticker := time.NewTicker(250 * time.Millisecond)

	defer ticker.Stop()

	for range ticker.C {
		// introduce some randomness in player
		// time
		delay := rand.Intn(250)
		time.Sleep(
			time.Duration(delay) * time.Millisecond)

		w, _ := g.screen.Size()
		bx, by := g.ball.Position()
		rx, ry := r.Position()

		// if ball is behind the bar, move away
		avoidDirection := 1
		if by > ry {
			avoidDirection = -1
		}
		if rx < w/2 && bx <= rx {
			r.Move(g.ticker, avoidDirection)
			continue
		}
		if rx > w/2 && bx >= rx {
			r.Move(g.ticker, avoidDirection)
			continue
		}

		// move towards the ball
		if by < ry+1 {
			r.Move(g.ticker, -1)
			continue
		}
		if by > ry-1 {
			r.Move(g.ticker, 1)
			continue
		}
	}
}
