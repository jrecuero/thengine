// group.go contains the group widget that allows to create a widget with a
// list of any other widget.
package widgets

import "github.com/jrecuero/thengine/pkg/engine"

// -----------------------------------------------------------------------------
//
// Group
//
// -----------------------------------------------------------------------------

type Group struct {
	*Widget
	widgets []IWidget
}

// -----------------------------------------------------------------------------
// New Group functions
// -----------------------------------------------------------------------------

func NewGroup(name string, widgets ...IWidget) *Group {
	group := &Group{
		Widget:  NewNamedWidget(name),
		widgets: make([]IWidget, len(widgets)),
	}
	for i, w := range widgets {
		group.widgets[i] = w
	}
	return group
}

// -----------------------------------------------------------------------------
// Group public methods
// -----------------------------------------------------------------------------

// Draw method renders the entity in the screen.
func (g *Group) Draw(scene engine.IScene) {
	for _, w := range g.widgets {
		w.Draw(scene)
	}
}

// GetWidgets methods returns all widgets in the Group.
func (g *Group) GetWidgets() []IWidget {
	return g.widgets
}

var _ engine.IObject = (*Group)(nil)
var _ engine.IFocus = (*Group)(nil)
var _ engine.IEntity = (*Group)(nil)
var _ IWidget = (*Group)(nil)
