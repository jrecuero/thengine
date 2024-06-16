// traps.go module contains all data, methods and logic for any trap in the
// application.
package assets

import (
	"github.com/gdamore/tcell/v2"
	"github.com/jrecuero/thengine/app/game/dad/rules"
	"github.com/jrecuero/thengine/pkg/api"
	"github.com/jrecuero/thengine/pkg/engine"
	"github.com/jrecuero/thengine/pkg/widgets"
)

// -----------------------------------------------------------------------------
//
// ITrap
//
// -----------------------------------------------------------------------------

type ITrap interface {
	engine.IEntity
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

func NewTrap(name string, position *api.Point, size *api.Size, style *tcell.Style) *Trap {
	t := &Trap{
		Widget: widgets.NewWidget(name, position, size, style),
		Damage: nil,
	}
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

//func (t *Trap) Update(event tcell.Event, scene engine.IScene) {
//    t.Widget.Update(event, scene)
//}
