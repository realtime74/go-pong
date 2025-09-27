package main

import (
	"time"
)

func Start(g *Game) {
	go ComputerMove(g)
}

func ComputerMove(g *Game) {
	ticker := time.NewTicker(600 * time.Millisecond)
	defer ticker.Stop()

	for range ticker.C {
		_, by := g.ball.Position()
		_, ry := g.lracket.Position()

		if by < ry {
			g.lracket.Move(-1)
		} else {
			g.lracket.Move(1)
		}
	}
}
