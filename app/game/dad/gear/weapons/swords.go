// swords.go contains all swords to be used
package weapons

import (
	"github.com/jrecuero/thengine/app/game/dad/rules"
)

const (
	ShortswordName = "weapon/sword/shortsword"
)

func init() {
	sections := []string{rules.DbSectionGear, rules.DbSectionWeapon}
	rules.DBase.Add(sections, rules.NewDatabaseEntry(ShortswordName, NewShortsword))
}

// -----------------------------------------------------------------------------
//
// Shortsword
//
// -----------------------------------------------------------------------------

func NewShortsword() *rules.Weapon {
	htype := rules.NewHandheldType(1)
	return rules.NewWeapon(ShortswordName, "shortsword", 10, 2, htype, rules.DiceThrow1d6, rules.Piercing)
}
