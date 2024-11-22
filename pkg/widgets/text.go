// text.go contains all data methods requried for handling any text widget.
package widgets

import (
	"strings"

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
	// Widget provides all widget attributes and methods for Text.
	*Widget

	// anchor is used to provide the end position for the label string.
	anchor *api.Point

	// label is the string to be output
	label string
}

// -----------------------------------------------------------------------------
// New Text functions
// -----------------------------------------------------------------------------

// NewText function creates a new Text instance widget.
func NewText(name string, position *api.Point, size *api.Size, style *tcell.Style, label string) *Text {
	text := &Text{
		Widget: NewWidget(name, position, size, style),
		anchor: nil,
		label:  label,
	}
	text.updateCanvas()
	return text
}

// NewAnchorText function create a new Text instace widget that has to be
// anchored.
func NewAnchorText(name string, pos *api.Point, size *api.Size, style *tcell.Style, label string) *Text {
	text := NewText(name, pos, size, style, label)
	text.SetAnchor()
	return text
}

// -----------------------------------------------------------------------------
// Text private methods
// -----------------------------------------------------------------------------

// updateCanvas method updates the text widget canvas with the string
// information.
func (t *Text) updateCanvas() {
	canvas := engine.NewCanvasFromString(t.label, t.GetStyle())
	t.SetCanvas(canvas)
	if t.IsAnchor() {
		t.SetAnchor()
	}
}

// -----------------------------------------------------------------------------
// Text public methods
// -----------------------------------------------------------------------------

// GetAnchot method returns the Text widget anchor.
func (t *Text) GetAnchor() *api.Point {
	return t.anchor
}

// GetText method returns the Text instance string.
func (t *Text) GetText() string {
	return t.label
}

// IsAnchor method returns if the Text widget is anchored or not.
func (t *Text) IsAnchor() bool {
	return t.anchor != nil
}

// Refresh method refreshes the Text widget with latest attribute values.
func (t *Text) Refresh() {
	t.updateCanvas()
}

// SetAnchor method sets the anchor attribute based on the string.
func (t *Text) SetAnchor() *api.Point {
	split := strings.Split(t.label, "\n")
	lines := len(split)
	cols := len(split[lines-1])
	anchor := api.ClonePoint(t.GetPosition())
	anchor.AddScale(cols, lines-1)
	t.anchor = anchor
	return t.anchor
}

// SetText method sets the Text instance string.
func (t *Text) SetText(label string) {
	t.label = label
	t.updateCanvas()
}

var _ engine.IObject = (*Text)(nil)
var _ engine.IFocus = (*Text)(nil)
var _ engine.IEntity = (*Text)(nil)
