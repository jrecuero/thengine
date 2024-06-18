// traps.go module contains all data, methods and logic for any trap in the
// application.
package assets

import (
	"github.com/gdamore/tcell/v2"
	"github.com/jrecuero/thengine/app/game/dad/constants"
	"github.com/jrecuero/thengine/app/game/dad/rules"
	"github.com/jrecuero/thengine/pkg/api"
	"github.com/jrecuero/thengine/pkg/engine"
	"github.com/jrecuero/thengine/pkg/widgets"
)

// -----------------------------------------------------------------------------
// Module internal constants
// -----------------------------------------------------------------------------

const (
	detectIndex = 0
	disarmIndex = 1
)

// -----------------------------------------------------------------------------
//
// ITrap
//
// -----------------------------------------------------------------------------

type ITrap interface {
	engine.IEntity
	CanDetect(rules.IUnit, ...any) bool
	CanDisarm(rules.IUnit, ...any) bool
}

// -----------------------------------------------------------------------------
//
// Trap
//
// -----------------------------------------------------------------------------

type Trap struct {
	*widgets.Widget
	*rules.Damage
}

// NewTrap function creates a new Trap instance.
// Trap has a fixed score for the saving throws as wisdom perception for any
// default trap instance.
// In the same way by default any trap is not visible and a solid object.
func NewTrap(name string, position *api.Point, size *api.Size, style *tcell.Style,
	detectDC int, disarmDC int, diceThrow rules.IDiceThrow, damageType rules.DamageType) *Trap {
	t := &Trap{
		Widget: widgets.NewWidget(name, position, size, style),
		Damage: rules.NewDamage(diceThrow, damageType),
	}
	detect := &rules.SavingThrowDamage{
		SavingThrow: rules.NewSavingThrow(constants.Perception, detectDC),
	}
	disarm := &rules.SavingThrowDamage{
		SavingThrow: rules.NewSavingThrow(constants.Sleight, disarmDC),
	}
	t.Damage.SetSavingThrowsDamage([]*rules.SavingThrowDamage{detect, disarm})
	t.SetVisible(false)
	t.SetSolid(true)
	return t
}

// -----------------------------------------------------------------------------
// Trap private methods
// -----------------------------------------------------------------------------

// -----------------------------------------------------------------------------
// Trap public methods
// -----------------------------------------------------------------------------

// CanDetect method checks if the unit can detect the trap. It is based in a
// perception check over the detect trap DC.
// opts can contain some tools or any other condition that can help to improve
// trap detection or decrease it.
func (t *Trap) CanDetect(unit rules.IUnit, opts ...any) bool {
	if t.GetSavingThrowsDamage() == nil {
		return true
	}
	if len(t.GetSavingThrowsDamage()) == 0 {
		return true
	}
	return t.GetSavingThrowsDamage()[detectIndex].Pass(unit)
}

// CanDisarm method checks if the unit can disarm the trap. It is bases in a
// sleight of hand check over the disarm trap DC.
// opts can contain some tools or any other condition that can help to improve
// trap disarm or decrease it.
func (t *Trap) CanDisarm(unit rules.IUnit, opts ...any) bool {
	if t.GetSavingThrowsDamage() == nil {
		return true
	}
	if len(t.GetSavingThrowsDamage()) <= 1 {
		return true
	}
	return t.GetSavingThrowsDamage()[disarmIndex].Pass(unit)

}

//func (t *Trap) Update(event tcell.Event, scene engine.IScene) {
//    t.Widget.Update(event, scene)
//}

var _ ITrap = (*Trap)(nil)
var _ rules.IDamage = (*Trap)(nil)
var _ engine.IEntity = (*Trap)(nil)
