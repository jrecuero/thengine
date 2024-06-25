// activatable.go package contains the interface for any structure that can be
// activated and has a passive or active attribute.
package rules

import (
	"fmt"

	"github.com/jrecuero/thengine/pkg/tools"
)

// -----------------------------------------------------------------------------
//
// IActivable
//
// -----------------------------------------------------------------------------

type IActivable interface {
	Activate() error
	Clean()
	Deactivate() error
	GetEffects() map[string]any
	IsActivated() bool
	IsPassive() bool
	IsSustained() bool
	SetEffects(map[string]any)
}

// -----------------------------------------------------------------------------
//
// Activable
//
// -----------------------------------------------------------------------------

type Activable struct {
	effects     map[string]any
	isactivated bool
	ispassive   bool
	issustained bool
}

func NewActivable(ispassive bool, issustained bool, isactivated bool) *Activable {
	return &Activable{
		effects:     make(map[string]any),
		isactivated: isactivated,
		ispassive:   ispassive,
		issustained: issustained,
	}
}

// -----------------------------------------------------------------------------
// Activable public methods
// -----------------------------------------------------------------------------

func (a *Activable) Activate() error {
	if a.ispassive {
		return fmt.Errorf("cannot activate passive instance")
	}
	a.isactivated = true
	return nil
}

func (a *Activable) Clean() {
	if a.isactivated {
		if a.ispassive {
			return
		}
		if a.issustained {
			return
		}
		tools.Logger.WithField("module", "activable").
			WithField("method", "Clean").
			Debugf("deactivate %+#v", a)
		a.Deactivate()
	}
}

func (a *Activable) Deactivate() error {
	if a.ispassive {
		return fmt.Errorf("cannot deactivate passive instance")
	}
	a.isactivated = false
	return nil
}

func (a *Activable) GetEffects() map[string]any {
	return a.effects
}

func (a *Activable) IsActivated() bool {
	return a.isactivated
}

func (a *Activable) IsPassive() bool {
	return a.ispassive
}

func (a *Activable) IsSustained() bool {
	return a.issustained
}

func (a *Activable) SetEffects(effects map[string]any) {
	a.effects = effects
}

var _ IActivable = (*Activable)(nil)
