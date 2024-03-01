package engine_test

import (
	"github.com/gdamore/tcell/v2"
	"github.com/jrecuero/thengine/pkg/engine"
)

var cells []*engine.Cell

func createCells() {
	cells = []*engine.Cell{
		{
			Style: engine.NewStyle(tcell.ColorBlack, tcell.ColorWhite, 0),
			Rune:  '0',
		},
		{
			Style: engine.NewStyle(tcell.ColorBlue, tcell.ColorRed, 0),
			Rune:  '1',
		},
		{
			Style: engine.NewStyle(tcell.ColorLightBlue, tcell.ColorDefault, 0),
			Rune:  '2',
		},
		{
			Style: engine.NewStyle(tcell.ColorDefault, tcell.ColorDarkMagenta, 0),
			Rune:  '3',
		},
		{
			Style: engine.NewStyle(tcell.ColorDarkCyan, tcell.ColorYellow, 0),
			Rune:  '4',
		},
		{
			Style: engine.NewStyle(tcell.ColorDarkGray, tcell.ColorLightBlue, 0),
			Rune:  '5',
		},
	}
}
