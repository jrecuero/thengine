package engine_test

import (
	"github.com/gdamore/tcell/v2"
	"github.com/jrecuero/thengine/pkg/engine"
)

var cells []*engine.Cell

func createCells() {
	cells = []*engine.Cell{
		engine.NewCell(engine.NewStyle(tcell.ColorBlack, tcell.ColorWhite, 0), '0'),
		engine.NewCell(engine.NewStyle(tcell.ColorBlue, tcell.ColorRed, 0), '1'),
		engine.NewCell(engine.NewStyle(tcell.ColorLightBlue, tcell.ColorDefault, 0), '2'),
		engine.NewCell(engine.NewStyle(tcell.ColorDefault, tcell.ColorDarkMagenta, 0), '3'),
		engine.NewCell(engine.NewStyle(tcell.ColorDarkCyan, tcell.ColorYellow, 0), '4'),
		engine.NewCell(engine.NewStyle(tcell.ColorDarkGray, tcell.ColorLightBlue, 0), '5'),
	}
}
