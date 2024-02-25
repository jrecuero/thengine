package engine_test

import (
	"github.com/jrecuero/thengine/pkg/api"
	"github.com/jrecuero/thengine/pkg/engine"
)

var cells []*engine.Cell

func createCells() {
	cells = []*engine.Cell{
		{
			Color: api.NewColor(api.ColorBlack, api.ColorWhite),
			Rune:  '0',
		},
		{
			Color: api.NewColor(api.ColorBlue, api.ColorRed),
			Rune:  '1',
		},
		{
			Color: api.NewColor(api.ColorLightBlue, api.ColorDefault),
			Rune:  '2',
		},
		{
			Color: api.NewColor(api.ColorDefault, api.ColorMagenta),
			Rune:  '3',
		},
		{
			Color: api.NewColor(api.ColorCyan, api.ColorYellow),
			Rune:  '4',
		},
		{
			Color: api.NewColor(api.ColorDarkGray, api.ColorLightBlue),
			Rune:  '5',
		},
	}
}
