package controls

import (
	"fmt"
	"time"

	"github.com/gdamore/tcell/v2"
	"github.com/realtime74/gopong/scr"
)

type StatusLine struct {
	lScore, rScore int

	screen tcell.Screen
	ticker int
	level  int

	lScored bool
	rScored bool
	levelUp bool
}

func NewStatusLine(screen tcell.Screen) *StatusLine {
	return &StatusLine{
		level:  1,
		screen: screen,
	}
}

func (tb *StatusLine) Score(left, right int) (x, y int) {
	tb.lScore += left
	tb.rScore += right
	tb.lScored = left > 0
	tb.rScored = right > 0
	go tb._resetflash()
	tb.Draw()
	return tb.lScore, tb.rScore
}

func (tb *StatusLine) SetTicker(ticker int) {
	tb.ticker = ticker
	tb.Draw()
}

func (tb *StatusLine) GetScore() (x, y int) {
	return tb.lScore, tb.rScore
}

func (tb *StatusLine) SetLevel(level int, flash bool) {
	tb.level = level
	tb.levelUp = flash
	tb.Draw()
	go tb._resetflash()
}

func (tb *StatusLine) _drawScore(x, y int, score int, flashing bool) {
	bgc := bgColor
	if flashing && score > 0 {
		bgc = tcell.ColorGreen
	}
	style := tcell.StyleDefault.
		Background(bgc).
		Foreground(fgColor)

	scorestr := fmt.Sprintf(" %d ", score)
	if x < 0 {
		x = -x
		x = x - len(scorestr) + 1
	}
	scr.DrawText(tb.screen, x, y, scorestr, style)
}

func (tb *StatusLine) Draw() {
	width, height := tb.screen.Size()

	y := height - 1
	style := tcell.StyleDefault.
		Background(bgColor).
		Foreground(fgColor)
	scr.Fill(tb.screen, 0, y, width, ' ', style)

	// level
	if tb.levelUp {
		style = tcell.StyleDefault.
			Background(tcell.ColorBlue).
			Foreground(fgColor)
	}

	level := fmt.Sprintf("[ %d ]", tb.level)
	scr.DrawText(tb.screen,
		(width-len(level))/2, y, level, style)

	// player scores
	tb._drawScore(1, height-1, tb.lScore, tb.lScored)
	tb._drawScore(-(width - 1), height-1, tb.rScore, tb.rScored)
}

func (tb *StatusLine) _resetflash() {
	time.Sleep(750 * time.Millisecond)
	tb.lScored = false
	tb.rScored = false
	tb.levelUp = false
	tb.Draw()
}
