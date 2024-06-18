// damage.go contains all information related with damage
package rules

import (
	"fmt"

	"github.com/jrecuero/thengine/app/game/dad/constants"
	"github.com/jrecuero/thengine/pkg/tools"
)

// -----------------------------------------------------------------------------
//
// DamageType
//
// -----------------------------------------------------------------------------

// DamageType type for any weapon.
type DamageType string

//const (
//    NullDamage  DamageType = "null"
//    Bludgeoning DamageType = "bludgeoning"
//    Piercing    DamageType = "piercing"
//    Slashing    DamageType = "slashing"
//    Magical     DamageType = "magical"
//    Poison      DamageType = "poison"
//)

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
	GetSavingThrowsDamage() []*SavingThrowDamage
	RollDamageValue() int
	RollSavingThrowsDamage(IUnit) int
	SetApplyStatus([]any)
	SetDiceThrow(IDiceThrow)
	SetDamageType(DamageType)
	SetSavingThrowsDamage([]*SavingThrowDamage)
	ToString() string
}

// -----------------------------------------------------------------------------
//
// Damage
//
// -----------------------------------------------------------------------------

// Damage structure represents any object that produces damage.
type Damage struct {
	dicethrow          IDiceThrow
	damageType         DamageType
	applyStatus        []any
	savingThrowsDamage []*SavingThrowDamage
}

func NewDamage(dicethrow IDiceThrow, damageType DamageType) *Damage {
	return &Damage{
		dicethrow:          dicethrow,
		damageType:         damageType,
		applyStatus:        nil,
		savingThrowsDamage: nil,
	}
}

func NewNoDamage() *Damage {
	return &Damage{
		dicethrow:          nil,
		damageType:         constants.NullDamage,
		applyStatus:        nil,
		savingThrowsDamage: nil,
	}
}

// -----------------------------------------------------------------------------
// Damage public methods
// -----------------------------------------------------------------------------

func (d *Damage) GetApplyStatus() []any {
	return d.applyStatus
}

func (d *Damage) GetDiceThrow() IDiceThrow {
	return d.dicethrow
}

func (d *Damage) GetDamageType() DamageType {
	return d.damageType
}

func (d *Damage) GetSavingThrowsDamage() []*SavingThrowDamage {
	return d.savingThrowsDamage
}

func (d *Damage) PassSavingThrowsDamage(unit IUnit) bool {
	if d.savingThrowsDamage != nil {
		for _, savingThrowDamage := range d.savingThrowsDamage {
			if pass := savingThrowDamage.Pass(unit); pass {
				return true
			}
		}
	}
	return false
}

func (d *Damage) RollDamageValue() int {
	if d.dicethrow != nil {
		return d.dicethrow.Roll()
	}
	return 0
}

func (d *Damage) RollSavingThrowsDamage(unit IUnit) int {
	result := 0
	if d.savingThrowsDamage != nil {
		for _, savingThrowDamage := range d.savingThrowsDamage {
			if pass := savingThrowDamage.Pass(unit); pass {
				damage := savingThrowDamage.Damage.RollDamageValue()
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

func (d *Damage) SetDiceThrow(dicethrow IDiceThrow) {
	d.dicethrow = dicethrow
}

func (d *Damage) SetDamageType(damageType DamageType) {
	d.damageType = damageType
}

func (d *Damage) SetSavingThrowsDamage(savingThrowsDamage []*SavingThrowDamage) {
	d.savingThrowsDamage = savingThrowsDamage
}

func (d *Damage) ToString() string {
	if d.savingThrowsDamage != nil {
		return fmt.Sprintf("\nsaving throw damage %s %s %s",
			d.dicethrow.ToString(), d.damageType, d.savingThrowsDamage[0].ToString())
	}
	return fmt.Sprintf("damage %s %s\n", d.dicethrow.ToString(), d.damageType)
}

var _ IDamage = (*Damage)(nil)
