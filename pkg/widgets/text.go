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

// Text structure defines all attributes and method for any basic and common
// text widget.
type Text struct {
	//*engine.Entity
	*Widget
	str string
}

// NewText function creates a new Text instance widget.
func NewText(name string, position *api.Point, size *api.Size, style *tcell.Style, str string) *Text {
	text := &Text{
		//Entity: engine.NewEntity(name, position, size, style),
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
	str := string(t.str)
	lines := strings.Split(str, "\n")
	maxW := 0
	for _, line := range lines {
		if len(line) > maxW {
			maxW = len(line)
		}
	}
	newSize := api.NewSize(maxW, len(lines))
	t.SetSize(newSize)
	canvas := engine.NewCanvas(newSize)
	for line := 0; line < newSize.H; line++ {
		for col := 0; col < newSize.W; col++ {
			if col >= len(lines[line]) {
				continue
			}
			point := api.NewPoint(col, line)
			cell := engine.NewCell(t.GetStyle(), rune(lines[line][col]))
			canvas.SetCellAt(point, cell)
		}
	}
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
