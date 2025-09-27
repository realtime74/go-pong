package controls

import (
	"github.com/gdamore/tcell/v2"
	"github.com/realtime74/gopong/scr"
)

var bgColor = tcell.NewRGBColor(100, 116, 139)
var fgColor = tcell.NewRGBColor(226, 232, 240)

type TitleBar struct {
	Title string

	screen tcell.Screen
}

func NewTitleBar(screen tcell.Screen, title string) *TitleBar {
	return &TitleBar{
		Title:  title,
		screen: screen,
	}
}

func (tb *TitleBar) Draw() {
	width, _ := tb.screen.Size()
	style := tcell.StyleDefault.
		Background(bgColor).
		Foreground(fgColor)

	scr.Fill(tb.screen, 0, 0, width, ' ', style)
	scr.DrawText(tb.screen,
		width/2-len(tb.Title)/2, 0,
		tb.Title,
		style)
}
