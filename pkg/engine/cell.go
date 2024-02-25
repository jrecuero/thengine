// cell.go contains required structures and method for handling a cell in the
// application. A cell contains information about the rune and the color to be
// displayed in any position in the canvas or screen.
package engine

import (
	"fmt"

	"github.com/jrecuero/thengine/pkg/api"
)

// -----------------------------------------------------------------------------
//
// Cell
//
// -----------------------------------------------------------------------------

// Cell structure defines a cell that is drawn in the canvas.
// Color *Color instance identifies the color (foreground and background
// colors) of the character in the cell.
// Ch Rune identifies the character to be displayed in the cell.
type Cell struct {
	Color *api.Color
	Rune  rune
}

// NewCell function creates a new Cell instance with the given color and rune.
func NewCell(color *api.Color, ch rune) *Cell {
	return &Cell{
		Color: color,
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
		Color: api.CloneColor(cell.Color),
		Rune:  cell.Rune,
	}
}

// -----------------------------------------------------------------------------
// Cell public methods
// -----------------------------------------------------------------------------

// Clone method clones all attributes from the given Cell in to the instance.
func (c *Cell) Clone(cell *Cell) {
	c.Color = api.CloneColor(cell.Color)
	c.Rune = cell.Rune
}

// IsEqual method checks if the given Cell is equal to the instance, where Color
// and Rune should be the same.
func (c *Cell) IsEqual(cell *Cell) bool {
	return c.Color.IsEqual(cell.Color) && (c.Rune == cell.Rune)
}

// ToString method returns the cell instance information as a string.
func (c *Cell) ToString() string {
	return fmt.Sprintf("[%c]%s", c.Rune, c.Color.ToString())
}

// SaveToDict method saves the instance information as a map.
func (c *Cell) SaveToDict() map[string]any {
	result := map[string]any{}
	result["color"] = c.Color.SaveToDict()
	result["rune"] = c.Rune
	return result
}
