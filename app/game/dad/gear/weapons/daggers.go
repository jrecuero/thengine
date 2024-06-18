// daggers.go contains all daggers to be used.
package weapons

import (
	"github.com/jrecuero/thengine/app/game/dad/constants"
	"github.com/jrecuero/thengine/app/game/dad/rules"
)

const (
	DaggerName       = "weapon/dagger/dagger"
	PoisonDaggerName = "weapon/dagger/poison-dager"
)

func init() {
	sections := []string{rules.DbSectionGear, rules.DbSectionWeapon}
	rules.DBase.Add(sections, rules.NewDatabaseEntry(DaggerName, NewDagger))
	rules.DBase.Add(sections, rules.NewDatabaseEntry(PoisonDaggerName, NewPoisonDagger))
}

// -----------------------------------------------------------------------------
//
// Dagger
//
// -----------------------------------------------------------------------------

func NewDagger() *rules.Weapon {
	htype := rules.NewHandheldType(1)
	return rules.NewWeapon(DaggerName, "dagger", 2, 1, htype, rules.DiceThrow1d4, constants.Piercing)
}

// -----------------------------------------------------------------------------
//
// PoisonDagger
//
// -----------------------------------------------------------------------------

func NewPoisonDagger() *rules.Weapon {
	htype := rules.NewHandheldType(1)
	dagger := rules.NewWeapon(DaggerName, "poison dagger", 2, 1, htype, rules.DiceThrow1d4, constants.Piercing)
	poisonDamage := &rules.SavingThrowDamage{
		SavingThrow: rules.NewSavingThrow(constants.Dexterity, 12),
		Damage:      rules.NewDamage(rules.DiceThrow1d3, constants.Poison),
	}
	dagger.Damage.SetSavingThrowsDamage([]*rules.SavingThrowDamage{poisonDamage})
	return dagger
}
