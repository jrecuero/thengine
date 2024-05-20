package main

import (
	"github.com/gdamore/tcell/v2"
	"github.com/jrecuero/thengine/app/game/dad/rules"
	"github.com/jrecuero/thengine/pkg/api"
	"github.com/jrecuero/thengine/pkg/engine"
	"github.com/jrecuero/thengine/pkg/widgets"
)

type Player struct {
	*widgets.Widget
	*rules.Unit
}

func NewPlayer(name string, position *api.Point, style *tcell.Style) *Player {
	cell := engine.NewCell(style, 'o')
	player := &Player{
		Widget: widgets.NewWidget(name, position, nil, style),
		Unit:   rules.NewUnit(),
	}
	player.GetCanvas().SetCellAt(nil, cell)
	player.GetHitPoints().SetMaxScore(100)
	player.GetHitPoints().SetScore(100)
	player.GetAbilities().GetStrength().SetScore(10)
	attack := rules.NewDefaultAttack(6)
	player.GetAttacks().AddAttack(attack)
	return player
}

func (p *Player) Update(event tcell.Event, scene engine.IScene) {
	if !p.HasFocus() {
		return
	}
}
