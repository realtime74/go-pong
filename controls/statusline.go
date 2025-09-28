package controls

import (
	"fmt"
	"time"

	"github.com/gdamore/tcell/v2"
	"github.com/realtime74/gopong/scr"
)

type StatusLine struct {
	lScore, rScore int

	flashing bool
	screen   tcell.Screen
	ticker   int
	level    int
}

func NewStatusLine(screen tcell.Screen) *StatusLine {
	return &StatusLine{
		screen: screen,
	}
}

func (tb *StatusLine) Score(left, right int) (x, y int) {
	tb.lScore += left
	tb.rScore += right
	tb._flash()
	return tb.lScore, tb.rScore
}

func (tb *StatusLine) SetTicker(ticker int) {
	tb.ticker = ticker
	tb.Draw()
}

func (tb *StatusLine) GetScore() (x, y int) {
	return tb.lScore, tb.rScore
}

func (tb *StatusLine) SetLevel(level int) {
	tb.level = level
	tb._flash()
	tb.Draw()
}

func (tb *StatusLine) Draw() {
	width, height := tb.screen.Size()

	bgc := bgColor
	if tb.flashing {
		bgc = tcell.ColorRed
	}

	style := tcell.StyleDefault.
		Background(bgc).
		Foreground(fgColor)

	lscore := fmt.Sprintf("Score: %d", tb.lScore)
	rscore := fmt.Sprintf("Score: %d", tb.rScore)

	y := height - 1
	scr.Fill(tb.screen, 0, y, width, ' ', style)

	// level
	level := fmt.Sprintf("<<%d>> [%d]", tb.level, tb.ticker)
	scr.DrawText(tb.screen,
		(width-len(level))/2, y, level, style)

	// player scores
	scr.DrawText(tb.screen, 1, y, lscore, style)
	scr.DrawText(tb.screen,
		width-len(rscore)-1, y, rscore, style)
}

func (tb *StatusLine) _flash() {
	tb.flashing = true
	tb.Draw()

	go func() {
		time.Sleep(300 * time.Millisecond)
		tb.flashing = false
		tb.Draw()
	}()
}
