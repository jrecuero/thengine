// light.go contains all light body armor.
package body

import "github.com/jrecuero/thengine/app/game/dad/rules"

const (
	PaddedArmorName = "armor/body/light/padded"
)

func init() {
	sections := []string{rules.DbSectionGear, rules.DbSectionBody}
	rules.DBase.Add(sections, rules.NewDatabaseEntry(PaddedArmorName, NewPaddedBodyArmor))
}

// -----------------------------------------------------------------------------
//
// PaddedBodyArmor
//
// -----------------------------------------------------------------------------

func NewPaddedBodyArmor() *rules.BodyGear {
	armor := rules.NewBodyGear(PaddedArmorName, "paddedarmor", 5, 8)
	armor.SetAC(1)
	return armor
}
