// weapons.go contains all weapons to be used
package weapons

import (
	"github.com/jrecuero/thengine/app/game/dad/rules"
)

// -----------------------------------------------------------------------------
//
// Sword
//
// -----------------------------------------------------------------------------

type Shortsword struct {
	*rules.Weapon
}

func NewSwordsword() *Shortsword {
	htype := rules.NewHandheldType(1)
	return &Shortsword{
		Weapon: rules.NewWeapon("shortsword", 10, 2, htype, rules.DiceThrow1d6, rules.Piercing),
	}
}

var _ rules.IWeapon = (*Shortsword)(nil)
