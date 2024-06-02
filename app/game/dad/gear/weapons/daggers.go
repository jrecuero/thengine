// daggers.go contains all daggers to be used.
package weapons

import "github.com/jrecuero/thengine/app/game/dad/rules"

const (
	DaggerName = "weapon/dagger/dagger"
)

func init() {
	sections := []string{rules.DbSectionGear, rules.DbSectionWeapon}
	rules.DBase.Add(sections, rules.NewDatabaseEntry(DaggerName, NewDagger))
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
