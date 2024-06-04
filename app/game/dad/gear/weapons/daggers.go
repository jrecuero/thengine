// daggers.go contains all daggers to be used.
package weapons

import "github.com/jrecuero/thengine/app/game/dad/rules"

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
	return rules.NewWeapon(DaggerName, "dagger", 2, 1, htype, rules.DiceThrow1d4, rules.Piercing)
}

// -----------------------------------------------------------------------------
//
// PoisonDagger
//
// -----------------------------------------------------------------------------

func NewPoisonDagger() *rules.Weapon {
	htype := rules.NewHandheldType(1)
	dagger := rules.NewWeapon(DaggerName, "poison dagger", 2, 1, htype, rules.DiceThrow1d4, rules.Piercing)
	poisonDamage := &rules.SavingThrowDamage{
		SavingThrow: rules.NewSavingThrow(rules.DexterityAS, 12),
		Damage:      rules.NewDamage(rules.DiceThrow1d3, rules.Poison),
	}
	dagger.Damage.SetSavingThrows([]*rules.SavingThrowDamage{poisonDamage})
	return dagger
}
