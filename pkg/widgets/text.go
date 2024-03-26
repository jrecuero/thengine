// text.go contains all data methods requried for handling any text widget.
package widgets

import (
	"github.com/gdamore/tcell/v2"
	"github.com/jrecuero/thengine/pkg/api"
	"github.com/jrecuero/thengine/pkg/engine"
)

// -----------------------------------------------------------------------------
// Text private types
// -----------------------------------------------------------------------------
type Runes []rune

// -----------------------------------------------------------------------------
//
// Text
//
// -----------------------------------------------------------------------------

// Text structure defines all attributes and methods for any basic and common
// text widget.
type Text struct {
	*Widget
	str string
}

// NewText function creates a new Text instance widget.
func NewText(name string, position *api.Point, size *api.Size, style *tcell.Style, str string) *Text {
	text := &Text{
		Widget: NewWidget(name, position, size, style),
		str:    str,
	}
	text.updateCanvas()
	return text
}

// -----------------------------------------------------------------------------
// Text private methods
// -----------------------------------------------------------------------------

// updateCanvas method updates the text widget canvas with the string
// information.
func (t *Text) updateCanvas() {
	canvas := engine.NewCanvasFromString(t.str, t.GetStyle())
	t.SetCanvas(canvas)
}

// -----------------------------------------------------------------------------
// Text public methods
// -----------------------------------------------------------------------------

// GetText method returns the Text instance string.
func (t *Text) GetText() string {
	return t.str
}

// SetText method sets the Text instance string.
func (t *Text) SetText(str string) {
	t.str = str
	t.updateCanvas()
}

var _ engine.IEntity = (*Text)(nil)
