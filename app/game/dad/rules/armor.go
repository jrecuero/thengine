// armor.go contains all data and methods related with any armament to be used
// by any unit.
package rules

// -----------------------------------------------------------------------------
//
// IArmor
//
// -----------------------------------------------------------------------------

// IArmor interface defines all methods any Armor should implement.
type IArmor interface {
	IBattleGear
	GetBodyPart() string
	SetBodyPart(string)
}

// -----------------------------------------------------------------------------
//
// Armor
//
// -----------------------------------------------------------------------------

// Armor structure defines all methods for any piece of Armor gear.
type Armor struct {
	*BattleGear
	bodypart string
}

func NewArmor(name string, uname string, cost int, weight int, bodypart string) *Armor {
	return &Armor{
		BattleGear: NewBattleGear(name, uname, cost, weight),
		bodypart:   bodypart,
	}
}

// -----------------------------------------------------------------------------
// Armor public methods
// -----------------------------------------------------------------------------

func (a *Armor) GetBodyPart() string {
	return a.bodypart
}

func (a *Armor) SetBodyPart(bodypart string) {
	a.bodypart = bodypart
}

var _ IDamage = (*Armor)(nil)
var _ IRollBonus = (*Armor)(nil)
var _ IBattleGear = (*Armor)(nil)
var _ IArmor = (*Armor)(nil)
