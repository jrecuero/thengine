package main

import (
	"github.com/gdamore/tcell/v2"
	"github.com/jrecuero/thengine/app/game/dad/constants"
	"github.com/jrecuero/thengine/app/game/dad/feats"
	"github.com/jrecuero/thengine/app/game/dad/gear/body"
	"github.com/jrecuero/thengine/app/game/dad/gear/shields"
	"github.com/jrecuero/thengine/app/game/dad/gear/weapons"
	"github.com/jrecuero/thengine/app/game/dad/inventory/consumables"
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

func (a *PowerAttack) DieRoll(unit rules.IUnit) int {
	tools.Logger.WithField("module", "player").
		WithField("struct", "PowerAttack").
		WithField("method", "DieRoll").
		Debugf("power attack")
	for _, p := range unit.GetFeats() {
		if p.GetName() == "great weapon master" {
			tools.Logger.WithField("module", "player").
				WithField("struct", "PowerAttack").
				WithField("method", "DieRoll").
				Debugf("activate feat %s", p.GetName())
			p.Activate()
			unit.AddActivable(p)
		}
	}
	return 0
}

func (a *PowerAttack) Roll(rules.IUnit) int {
	tools.Logger.WithField("module", "player").
		WithField("struct", "PowerAttack").
		WithField("method", "Roll").
		Debugf("power attack")
	return a.gear.GetMainHand().RollDamage()
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
	perception := rules.GetSkillByName(player.GetSkills(), constants.Perception)
	perception.SetProficienty(2)
	sleight := rules.GetSkillByName(player.GetSkills(), constants.Sleight)
	sleight.SetProficienty(1)
	feats := []rules.IFeat{feats.NewGreatWeaponMaster()}
	player.SetFeats(feats)

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

	sections := []string{rules.DbSectionInventory, rules.DbSectionConsumables}
	//tools.Logger.WithField("module", "potion").
	//    WithField("function", "init").
	//    Debugf("DBase: %+#v", rules.DBase.GetSections()[rules.DbSectionInventory].GetSections()[rules.DbSectionConsumables])
	tmp := rules.DBase.GetCreator(sections, consumables.PotionName)
	potionCreator := tmp.(func(int) *rules.Consumable)
	potion := potionCreator(5)
	player.GetInventory().AddConsumables(potion)
	tools.Logger.WithField("module", "potion").
		WithField("function", "init").
		Debugf("Inventory: %+#v", player.GetInventory())

	strengthModifier := player.GetAbilities().GetStrength().GetModifier()
	tools.Logger.WithField("module", "player").
		WithField("function", "NewPlayer").
		Debugf("strength modifier %d", strengthModifier)
	return player
}

func (p *Player) EndTick(engine.IScene) {
	p.EndTurn()
}
