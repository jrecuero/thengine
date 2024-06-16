package main

import (
	"github.com/gdamore/tcell/v2"
	"github.com/jrecuero/thengine/app/game/assets"
	"github.com/jrecuero/thengine/app/game/dad/gear/body"
	"github.com/jrecuero/thengine/app/game/dad/gear/shields"
	"github.com/jrecuero/thengine/app/game/dad/gear/weapons"
	"github.com/jrecuero/thengine/app/game/dad/rules"
	"github.com/jrecuero/thengine/app/game/dad/spells"
	"github.com/jrecuero/thengine/pkg/api"
	"github.com/jrecuero/thengine/pkg/engine"
	"github.com/jrecuero/thengine/pkg/tools"
	"github.com/jrecuero/thengine/pkg/widgets"
)

type PowerAttack struct {
	rules.IAttackRoll
	gear rules.IGear
}

func NewPowerAttack(gear rules.IGear) *PowerAttack {
	return &PowerAttack{
		gear: gear,
	}
}

func (a *PowerAttack) Roll() int {
	return 2 * a.gear.GetMainHand().RollDamage()
}

func (a *PowerAttack) RollSavingThrows(rules.IUnit) int {
	return 0
}

type Player struct {
	*widgets.Widget
	*rules.Unit
}

func NewPlayer(name string, position *api.Point, style *tcell.Style) *Player {
	cell := engine.NewCell(style, 'H')
	//cell := engine.NewCell(style, '🧝')
	//cell := engine.NewCell(style, '🐱')
	player := &Player{
		Widget: widgets.NewWidget(name, position, nil, style),
		Unit:   rules.NewUnit("player"),
	}
	player.GetCanvas().SetCellAt(nil, cell)
	player.SetZLevel(1)
	player.GetHitPoints().SetMaxScore(100)
	player.GetHitPoints().SetScore(100)
	player.GetAbilities().GetStrength().SetScore(14)
	player.GetAbilities().GetDexterity().SetScore(12)
	player.GetAbilities().GetConstitution().SetScore(10)
	player.GetAbilities().GetIntelligence().SetScore(10)
	player.GetAbilities().GetWisdom().SetScore(10)
	player.GetAbilities().GetCharisma().SetScore(10)

	weaponEntry := rules.DBase.GetSections()[rules.DbSectionGear].GetSections()[rules.DbSectionWeapon].GetEntries()[weapons.ShortswordName]
	weaponCreator := weaponEntry.GetCreator().(func() rules.IHandheld)
	player.GetGear().SetMainHand(weaponCreator())
	//player.GetGear().SetMainHand(weapons.NewShortsword())

	shieldCreator := rules.DBase.GetCreator([]string{rules.DbSectionGear, rules.DbSectionShield}, shields.ShieldName).(func() *rules.Shield)
	player.GetGear().SetOffHand(shieldCreator())

	player.GetGear().SetBody(body.NewPaddedBodyArmor())

	weaponAttack := rules.NewWeaponMeleeAttack("attack/weapon/melee", player.GetGear())
	player.GetAttacks().AddAttack(weaponAttack)

	powerAttack := rules.NewSpecialAttack("attack/special/power", NewPowerAttack(player.GetGear()))
	player.GetAttacks().AddAttack(powerAttack)

	magicalAttack := rules.NewSpellMeleeAttack("attack/spell/melee/guiding-bolt", spells.NewGuidingBolt())
	player.GetAttacks().AddAttack(magicalAttack)

	strengthModifier := player.GetAbilities().GetStrength().GetModifier()
	tools.Logger.WithField("module", "player").
		WithField("function", "NewPlayer").
		Debugf("strength modifier %d", strengthModifier)
	return player
}

func (p *Player) Update(event tcell.Event, scene engine.IScene) {
	if !p.IsActive() {
		return
	}
	traps := getTrapsInScene(scene)
	if t := isAnyTrapAdjacent(p, traps); t != nil {
		t.SetVisible(true)
		if trap, ok := t.(*assets.Trap); ok {
			if trap.GetSavingThrows() != nil && len(trap.GetSavingThrows()) != 0 {
				trapScore := trap.GetSavingThrows()[0].GetScore()
				trapDC := trap.GetSavingThrows()[0].GetDC()
				_, _ = trapScore, trapDC
			}
		}
	}
}
