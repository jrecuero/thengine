// cell.go contains required structures and method for handling a cell in the
// application. A cell contains information about the rune and the color to be
// displayed in any position in the canvas or screen.
package engine

import (
	"fmt"

	"github.com/gdamore/tcell/v2"
)

// -----------------------------------------------------------------------------
//
// Cell
//
// -----------------------------------------------------------------------------

// Cell structure defines a cell that is drawn in the canvas.
// Style instance identifies color and other attributes.
// Ch Rune identifies the character to be displayed in the cell.
type Cell struct {
	Style *tcell.Style
	Rune  rune
}

// NewCell function creates a new Cell instance with the given color and rune.
func NewCell(style *tcell.Style, ch rune) *Cell {
	if style == nil {
		style = &tcell.StyleDefault
	}
	return &Cell{
		Style: style,
		Rune:  ch,
	}
}

// NewEmptyCell function creates a new Cell instance without any color or rune.
func NewEmptyCell() *Cell {
	return &Cell{}
}

// CloneCell function creates a new Cell instance with same attributes as the
// given Cell instance.
func CloneCell(cell *Cell) *Cell {
	return &Cell{
		Style: cell.Style,
		Rune:  cell.Rune,
	}
}

// -----------------------------------------------------------------------------
// Cell public methods
// -----------------------------------------------------------------------------

// Clone method clones all attributes from the given Cell in to the instance.
func (c *Cell) Clone(cell *Cell) {
	c.Style = cell.Style
	c.Rune = cell.Rune
}

// IsEqual method checks if the given Cell is equal to the instance, where Color
// and Rune should be the same.
func (c *Cell) IsEqual(cell *Cell) bool {
	return CompareStyle(c.Style, cell.Style) && (c.Rune == cell.Rune)
}

// ToString method returns the cell instance information as a string.
func (c *Cell) ToString() string {
	return fmt.Sprintf("[%c]%s", c.Rune, StyleToString(c.Style))
}

// SaveToDict method saves the instance information as a map.
// TODO: to be revisited.
func (c *Cell) SaveToDict() map[string]any {
	result := map[string]any{}
	result["style"] = c.Style
	result["rune"] = c.Rune
	return result
}
