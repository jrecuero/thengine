// light.go contains all light body armor.
package body

import "github.com/jrecuero/thengine/app/game/dad/rules"

// -----------------------------------------------------------------------------
//
// PaddedBodyArmor
//
// -----------------------------------------------------------------------------

func NewPaddedBodyArmor() *rules.BodyGear {
	armor := rules.NewBodyGear("armor/body/light/padded", 5, 8)
	armor.SetAC(1)
	return armor
}
