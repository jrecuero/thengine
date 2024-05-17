package main

import (
	"github.com/gdamore/tcell/v2"
	"github.com/jrecuero/thengine/pkg/api"
	"github.com/jrecuero/thengine/pkg/engine"
	"github.com/jrecuero/thengine/pkg/widgets"
)

type Player struct {
	*widgets.Widget
}

func NewPlayer(name string, position *api.Point, style *tcell.Style) *Player {
	cell := engine.NewCell(style, 'o')
	player := &Player{
		Widget: widgets.NewWidget(name, position, nil, style),
	}
	player.GetCanvas().SetCellAt(nil, cell)
	return player
}

func (p *Player) Update(event tcell.Event, scene engine.IScene) {
	if !p.HasFocus() {
		return
	}
}
