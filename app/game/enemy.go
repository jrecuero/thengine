package main

import (
	"github.com/gdamore/tcell/v2"
	"github.com/jrecuero/thengine/pkg/api"
	"github.com/jrecuero/thengine/pkg/engine"
	"github.com/jrecuero/thengine/pkg/widgets"
)

type Enemy struct {
	*widgets.Widget
}

func NewEnemy(name string, position *api.Point, style *tcell.Style) *Enemy {
	cell := engine.NewCell(style, 'X')
	enemy := &Enemy{
		Widget: widgets.NewWidget(name, position, nil, style),
	}
	enemy.GetCanvas().SetCellAt(nil, cell)
	enemy.SetSolid(true)
	return enemy
}
