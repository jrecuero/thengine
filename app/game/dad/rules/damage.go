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
	Poison      DamageType = "poison"
)

// -----------------------------------------------------------------------------
//
// SavingThrowDamage
//
// -----------------------------------------------------------------------------

type SavingThrowDamage struct {
	*SavingThrow
	*Damage
}

// -----------------------------------------------------------------------------
//
// IDamage
//
// -----------------------------------------------------------------------------

// IDamage interface defines all methods any object that produces damage has to
// implement.
type IDamage interface {
	GetApplyStatus() []any
	GetDiceThrow() IDiceThrow
	GetDamageType() DamageType
	GetSavingThrow() []SavingThrowDamage
	RollDamageValue() int
	SetApplyStatus([]any)
	SetDiceThrow(IDiceThrow)
	SetDamageType(DamageType)
	SetSavingThrow([]SavingThrowDamage)
}

// -----------------------------------------------------------------------------
//
// Damage
//
// -----------------------------------------------------------------------------

// Damage structure represents any object that produces damage.
type Damage struct {
	diceThrow   IDiceThrow
	damageType  DamageType
	applyStatus []any
	savingThrow []SavingThrowDamage
}

func NewDamage(diceThrow IDiceThrow, damageType DamageType) *Damage {
	return &Damage{
		diceThrow:   diceThrow,
		damageType:  damageType,
		applyStatus: nil,
		savingThrow: nil,
	}
}

func NewNoDamage() *Damage {
	return &Damage{
		diceThrow:   nil,
		damageType:  NullDamage,
		applyStatus: nil,
		savingThrow: nil,
	}
}

// -----------------------------------------------------------------------------
// Damage public methods
// -----------------------------------------------------------------------------

func (d *Damage) GetApplyStatus() []any {
	return d.applyStatus
}

func (d *Damage) GetDiceThrow() IDiceThrow {
	return d.diceThrow
}

func (d *Damage) GetDamageType() DamageType {
	return d.damageType
}

func (d *Damage) GetSavingThrow() []SavingThrowDamage {
	return d.savingThrow
}

func (d *Damage) RollDamageValue() int {
	if d.diceThrow != nil {
		return d.diceThrow.Roll()
	}
	return 0
}

func (d *Damage) SetApplyStatus(status []any) {
	d.applyStatus = status
}

func (d *Damage) SetDiceThrow(diceThrow IDiceThrow) {
	d.diceThrow = diceThrow
}

func (d *Damage) SetDamageType(damageType DamageType) {
	d.damageType = damageType
}

func (d *Damage) SetSavingThrow(savingThrow []SavingThrowDamage) {
	d.savingThrow = savingThrow
}

var _ IDamage = (*Damage)(nil)
