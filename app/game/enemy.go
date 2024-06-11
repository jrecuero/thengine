package main

import (
	"github.com/gdamore/tcell/v2"
	"github.com/jrecuero/thengine/app/game/dad/gear/weapons"
	"github.com/jrecuero/thengine/app/game/dad/rules"
	"github.com/jrecuero/thengine/pkg/api"
	"github.com/jrecuero/thengine/pkg/engine"
	"github.com/jrecuero/thengine/pkg/tools"
	"github.com/jrecuero/thengine/pkg/widgets"
)

type Enemy struct {
	*widgets.Widget
	*rules.Unit
}

func NewEmptyEnemy() *Enemy {
	enemy := &Enemy{
		Widget: widgets.NewEmptyWidget(),
		Unit:   rules.NewUnit("enemy"),
	}
	return enemy
}

func NewEnemy(name string, position *api.Point, style *tcell.Style) *Enemy {
	enemy := &Enemy{
		Widget: widgets.NewWidget(name, position, nil, style),
		Unit:   rules.NewUnit("enemy"),
	}
	cell := engine.NewCell(enemy.GetStyle(), 'X')
	enemy.GetCanvas().SetCellAt(nil, cell)
	enemy.populate(nil)
	return enemy
}

func (e *Enemy) populate(content map[string]any) {
	tools.Logger.WithField("module", "enemy").
		WithField("method", "populate").
		Debugf("%+v", content)
	e.SetSolid(true)
	defaults := map[string]any{
		"hp":           50,
		"strength":     10,
		"dexterity":    10,
		"constitution": 10,
		"intelligence": 10,
		"wisdom":       10,
		"charisma":     10,
	}
	e.Populate(defaults, content)
	if _, ok := content["gear"]; ok {
		e.GetGear().UnmarshalMap(content)
	} else {
		//e.GetGear().SetMainHand(weapons.NewDagger())
		e.GetGear().SetMainHand(weapons.NewPoisonDagger())
	}
	attack := rules.NewWeaponMeleeAttack("attack/weapon/melee", e.GetGear())
	e.GetAttacks().AddAttack(attack)
}

func (e *Enemy) UnmarshalMap(content map[string]any, origin *api.Point) error {
	if err := e.Entity.UnmarshalMap(content, origin); err != nil {
		return err
	}
	if uname, ok := content["uname"].(string); ok {
		e.SetUName(uname)
	}
	tools.Logger.WithField("module", "enemy").
		WithField("method", "UnmarshalJSON").
		Debugf("%s|%s is being unmarshal", e.GetName(), e.GetUName())

	e.populate(content)
	return nil
}
