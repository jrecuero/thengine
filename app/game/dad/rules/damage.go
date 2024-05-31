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

// -----------------------------------------------------------------------------
//
// IDamage
//
// -----------------------------------------------------------------------------

// IDamage interface defines all methods any object that produces damage has to
// implement.
type IDamage interface {
	GetDiceThrow() IDiceThrow
	GetDamageType() DamageType
	RollDamageValue() int
	SetDiceThrow(IDiceThrow)
	SetDamageType(DamageType)
}

// -----------------------------------------------------------------------------
//
// Damage
//
// -----------------------------------------------------------------------------

// Damage structure represents any object that produces damage.
type Damage struct {
	diceThrow  IDiceThrow
	damageType DamageType
}

func NewDamage(diceThrow IDiceThrow, damageType DamageType) *Damage {
	return &Damage{
		diceThrow:  diceThrow,
		damageType: damageType,
	}
}

// -----------------------------------------------------------------------------
// Damage public methods
// -----------------------------------------------------------------------------

func (d *Damage) GetDiceThrow() IDiceThrow {
	return d.diceThrow
}

func (d *Damage) GetDamageType() DamageType {
	return d.damageType
}

func (d *Damage) RollDamageValue() int {
	if d.diceThrow != nil {
		return d.diceThrow.Roll()
	}
	return 0
}

func (d *Damage) SetDiceThrow(diceThrow IDiceThrow) {
	d.diceThrow = diceThrow
}

func (d *Damage) SetDamageType(damageType DamageType) {
	d.damageType = damageType
}

var _ IDamage = (*Damage)(nil)
