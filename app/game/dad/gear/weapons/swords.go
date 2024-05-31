// weapons.go contains all weapons to be used
package weapons

import (
	"github.com/jrecuero/thengine/app/game/dad/rules"
)

// -----------------------------------------------------------------------------
//
// ShortSword
//
// -----------------------------------------------------------------------------

func NewSwordsword() *rules.Weapon {
	htype := rules.NewHandheldType(1)
	return rules.NewWeapon("weapon/sword/shortsword", 10, 2, htype, rules.DiceThrow1d6, rules.Piercing)
}
