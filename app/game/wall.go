package main

import (
	"github.com/gdamore/tcell/v2"
	"github.com/jrecuero/thengine/pkg/api"
	"github.com/jrecuero/thengine/pkg/engine"
	"github.com/jrecuero/thengine/pkg/widgets"
)

type Wall struct {
	*widgets.Widget
}

func NewWall(name string, position *api.Point, size *api.Size, style *tcell.Style) *Wall {
	wall := &Wall{
		Widget: widgets.NewWidget(name, position, size, style),
	}
	cell := engine.NewCell(style, '#')
	wall.GetCanvas().FillWithCell(cell)
	wall.SetSolid(true)
	return wall
}

func NewEmptyWall() *Wall {
	wall := &Wall{
		Widget: widgets.NewWidget("", nil, nil, nil),
	}
	wall.SetSolid(true)
	return wall
}
