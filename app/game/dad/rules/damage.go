// damage.go contains all information related with damage
package rules

// -----------------------------------------------------------------------------
//
// DamageType
//
// -----------------------------------------------------------------------------

// DamageType type for any weapon.
type DamageType string

const (
	NullDamage  DamageType = "null"
	Bludgeoning DamageType = "bludgeoning"
	Piercing    DamageType = "piercing"
	Slashing    DamageType = "slashing"
	Magical     DamageType = "magical"
)
