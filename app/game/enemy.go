package main

import (
	"github.com/gdamore/tcell/v2"
	"github.com/jrecuero/thengine/app/game/dad/gear/weapons"
	"github.com/jrecuero/thengine/app/game/dad/rules"
	"github.com/jrecuero/thengine/pkg/api"
	"github.com/jrecuero/thengine/pkg/engine"
	"github.com/jrecuero/thengine/pkg/widgets"
)

type Enemy struct {
	*widgets.Widget
	*rules.Unit
}

func NewEnemy(name string, position *api.Point, style *tcell.Style) *Enemy {
	cell := engine.NewCell(style, 'X')
	enemy := &Enemy{
		Widget: widgets.NewWidget(name, position, nil, style),
		Unit:   rules.NewUnit("enemy"),
	}
	enemy.GetCanvas().SetCellAt(nil, cell)
	enemy.SetSolid(true)
	enemy.GetHitPoints().SetMaxScore(50)
	enemy.GetHitPoints().SetScore(50)
	enemy.GetAbilities().GetStrength().SetScore(10)
	enemy.GetAbilities().GetDexterity().SetScore(10)
	enemy.GetAbilities().GetConstitution().SetScore(10)
	enemy.GetAbilities().GetIntelligence().SetScore(10)
	enemy.GetAbilities().GetWisdom().SetScore(10)
	enemy.GetAbilities().GetCharisma().SetScore(10)
	enemy.GetGear().SetMainHand(weapons.NewSwordsword())
	//attack := rules.NewDefaultAttack(6)
	attack := rules.NewWeaponAttack(enemy.GetGear())
	enemy.GetAttacks().AddAttack(attack)
	return enemy
}
