// damage.go contains all information related with damage
package rules

import (
	"fmt"

	"github.com/jrecuero/thengine/pkg/tools"
)

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

func (s *SavingThrowDamage) ToString() string {
	return fmt.Sprintf("%s %s", s.SavingThrow.ToString(), s.Damage.ToString())
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
	GetSavingThrows() []*SavingThrowDamage
	RollDamageValue() int
	RollSavingThrowsDamage(IUnit) int
	SetApplyStatus([]any)
	SetDiceThrow(IDiceThrow)
	SetDamageType(DamageType)
	SetSavingThrows([]*SavingThrowDamage)
	ToString() string
}

// -----------------------------------------------------------------------------
//
// Damage
//
// -----------------------------------------------------------------------------

// Damage structure represents any object that produces damage.
type Damage struct {
	diceThrow    IDiceThrow
	damageType   DamageType
	applyStatus  []any
	savingThrows []*SavingThrowDamage
}

func NewDamage(diceThrow IDiceThrow, damageType DamageType) *Damage {
	return &Damage{
		diceThrow:    diceThrow,
		damageType:   damageType,
		applyStatus:  nil,
		savingThrows: nil,
	}
}

func NewNoDamage() *Damage {
	return &Damage{
		diceThrow:    nil,
		damageType:   NullDamage,
		applyStatus:  nil,
		savingThrows: nil,
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

func (d *Damage) GetSavingThrows() []*SavingThrowDamage {
	return d.savingThrows
}

func (d *Damage) RollDamageValue() int {
	if d.diceThrow != nil {
		return d.diceThrow.Roll()
	}
	return 0
}

func (d *Damage) RollSavingThrowsDamage(unit IUnit) int {
	result := 0
	if d.savingThrows != nil {
		for _, stDamage := range d.savingThrows {
			if pass := stDamage.Pass(unit); pass {
				damage := stDamage.RollDamageValue()
				result += damage
				tools.Logger.WithField("module", "damage").
					WithField("method", "RollSavingThrowsDamage").
					Debugf("saving throw damage  %d->%d", damage, result)
			}
		}
	}
	return result
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

func (d *Damage) SetSavingThrows(savingThrows []*SavingThrowDamage) {
	d.savingThrows = savingThrows
}

func (d *Damage) ToString() string {
	if d.savingThrows != nil {
		return fmt.Sprintf("damage %s %s %s", d.diceThrow.ToString(), d.damageType, d.savingThrows[0].ToString())
	}
	return fmt.Sprintf("damage %s %s", d.diceThrow.ToString(), d.damageType)
}

var _ IDamage = (*Damage)(nil)
