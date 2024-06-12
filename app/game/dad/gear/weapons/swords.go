// swords.go contains all swords to be used
package weapons

import (
	"github.com/jrecuero/thengine/app/game/dad/constants"
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

type ShortSword struct {
	*rules.Weapon
}

func NewShortsword() rules.IHandheld {
	htype := rules.NewHandheldType(1)
	return &ShortSword{
		Weapon: rules.NewWeapon(ShortswordName, "shortsword", 10, 2, htype, rules.DiceThrow1d6, constants.Piercing),
	}
}

func (w *ShortSword) DieRollBonus(bonus string) int {
	if bonus == constants.Strength {
		return 2
	}
	return 0
}
