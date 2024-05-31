// weapon.go contains all data and methods common to any weapon to be used.
package rules

// -----------------------------------------------------------------------------
//
// IWeapon
//
// -----------------------------------------------------------------------------

// IWeapon interface defines all methods any Weapon structure should implement.
type IWeapon interface {
	IHandheld
}

// -----------------------------------------------------------------------------
//
// Weapon
//
// -----------------------------------------------------------------------------

// Weapon structure defines all attributes and methods for the basic weapon.
type Weapon struct {
	*Handheld
}

func NewWeapon(name string, cost int, weight int, htype *HandheldType, diceThrow IDiceThrow, damageType DamageType) *Weapon {
	weapon := &Weapon{
		Handheld: NewHandheld(name, cost, weight, htype),
	}
	weapon.Damage = NewDamage(diceThrow, damageType)
	return weapon
}

// -----------------------------------------------------------------------------
// Weapon public methods
// -----------------------------------------------------------------------------

var _ IHandheld = (*Weapon)(nil)
var _ IWeapon = (*Weapon)(nil)
