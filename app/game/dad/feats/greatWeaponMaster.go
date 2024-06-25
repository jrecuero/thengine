// greatWeaponMaster.go package implement Great Weapon Maste feat for any
// character.
package feats

import (
	"github.com/jrecuero/thengine/app/game/dad/constants"
	"github.com/jrecuero/thengine/app/game/dad/rules"
)

// -----------------------------------------------------------------------------
//
// GreatWeaponMaster
//
// -----------------------------------------------------------------------------

func NewGreatWeaponMaster() *rules.Feat {
	f := rules.NewFeat("great weapon master", false, false, false)
	f.SetDescription(`On your turn, when you score a critical hit with a melee
weapon or reduce a creature to 0 hit points with one, you can make one melee
weapon attack as a bonus action. Before you make a melee attack with a heavy
weapon that you are proficient with, you can choose to take a -5 penalty to the
attack roll. If the attack hits, you add +10 to the attack’s damage.`)
	effects := map[string]any{
		constants.DieRoll:    -5,
		constants.DamageRoll: 10,
	}
	f.SetEffects(effects)
	return f
}
