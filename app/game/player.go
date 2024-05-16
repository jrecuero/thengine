package main

import (
	"github.com/gdamore/tcell/v2"
	"github.com/jrecuero/thengine/pkg/api"
	"github.com/jrecuero/thengine/pkg/engine"
	"github.com/jrecuero/thengine/pkg/tools"
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
	player.SetFocusType(engine.SingleFocus)
	player.SetFocusEnable(true)
	return player
}

func (p *Player) Move(args ...any) {
	direction := args[0].(string)
	scene := args[1].(engine.IScene)
	tools.Logger.WithField("module", "Player").WithField("method", "Move").Infof(direction)
	x, y := p.GetPosition().Get()
	var newPosition *api.Point
	switch direction {
	case "up":
		newPosition = api.NewPoint(x, y-1)
	case "down":
		newPosition = api.NewPoint(x, y+1)
	case "left":
		newPosition = api.NewPoint(x-1, y)
	case "right":
		newPosition = api.NewPoint(x+1, y)
	}
	p.SetPosition(newPosition)
	collisions := scene.CheckCollisionWith(p)
	for _, ent := range collisions {
		if _, ok := ent.(*Wall); ok {
			p.SetPosition(api.NewPoint(x, y))
		}
	}
}

func (p *Player) Update(event tcell.Event, scene engine.IScene) {
	if !p.HasFocus() {
		return
	}
	actions := []*widgets.KeyboardAction{
		{Key: tcell.KeyUp, Callback: p.Move, Args: []any{"up", scene}},
		{Key: tcell.KeyDown, Callback: p.Move, Args: []any{"down", scene}},
		{Key: tcell.KeyRight, Callback: p.Move, Args: []any{"right", scene}},
		{Key: tcell.KeyLeft, Callback: p.Move, Args: []any{"left", scene}},
	}
	p.HandleKeyboardForActions(event, actions)
}
