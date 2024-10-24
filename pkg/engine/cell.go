// cell.go contains required structures and method for handling a cell in the
// application. A cell contains information about the rune and the color to be
// displayed in any position in the canvas or screen.
// Cell is the base class to represent information for a cell based only in the
// rune and the color.
// sprite.
package engine

import (
	"fmt"

	"github.com/gdamore/tcell/v2"
	"github.com/jrecuero/thengine/pkg/api"
)

// -----------------------------------------------------------------------------
//
// Cell
//
// -----------------------------------------------------------------------------

// Cell structure defines a cell that is drawn in the canvas.
// Style instance identifies color and other attributes.
// Ch Rune identifies the character to be displayed in the cell.
// position keep the position if provided, but with a default empty value.
type Cell struct {
	position *api.Point
	Style    *tcell.Style
	Rune     rune
}

// NewCell function creates a new Cell instance with the given color and rune.
func NewCell(style *tcell.Style, ch rune) *Cell {
	if style == nil {
		style = &tcell.StyleDefault
	}
	return &Cell{
		position: nil,
		Style:    style,
		Rune:     ch,
	}
}

// NewCellAt function creates a new Cell instance with the given color, rune
// and at the given position.
func NewCellAt(style *tcell.Style, ch rune, position *api.Point) *Cell {
	cell := NewCell(style, ch)
	cell.SetPosition(position)
	return cell
}

// NewEmptyCell function creates a new Cell instance without any color or rune.
func NewEmptyCell() *Cell {
	return &Cell{}
}

// CloneCell function creates a new Cell instance with same attributes as the
// given Cell instance.
func CloneCell(cell *Cell) *Cell {
	return &Cell{
		position: cell.GetPosition(),
		Style:    cell.Style,
		Rune:     cell.Rune,
	}
}

// -----------------------------------------------------------------------------
// Cell public methods
// -----------------------------------------------------------------------------

// Clone method clones all attributes from the given Cell in to the instance.
func (c *Cell) Clone(cell *Cell) {
	c.Style = cell.Style
	c.Rune = cell.Rune
	c.position = cell.GetPosition()
}

// GetPosition method returs the position from the Cell instance.
func (c *Cell) GetPosition() *api.Point {
	return c.position
}

// IsEqual method checks if the given Cell is equal to the instance, where Color
// and Rune should be the same.
func (c *Cell) IsEqual(cell *Cell) bool {
	return CompareStyle(c.Style, cell.Style) &&
		(c.Rune == cell.Rune) &&
		(c.GetPosition() == cell.GetPosition())

}

// ToString method returns the cell instance information as a string.
func (c *Cell) ToString() string {
	if c.GetPosition() != nil {
		return fmt.Sprintf("[%c]%s %s", c.Rune, StyleToString(c.Style), c.GetPosition().ToString())
	}
	return fmt.Sprintf("[%c]%s", c.Rune, StyleToString(c.Style))
}

// SaveToDict method saves the instance information as a map.
// TODO: to be revisited.
func (c *Cell) SaveToDict() map[string]any {
	result := map[string]any{}
	result["style"] = c.Style
	result["rune"] = c.Rune
	result["position"] = c.GetPosition()
	return result
}

// SetPosition method sets the position in a Cell instance.
func (c *Cell) SetPosition(position *api.Point) {
	c.position = position
}

// -----------------------------------------------------------------------------
//
// CellGroup
//
// -----------------------------------------------------------------------------

// CellGroup type groups an array of Cell, which is used to create sprite
// widgets.
type CellGroup []*Cell

// -----------------------------------------------------------------------------
//
// CellFrame
//
// -----------------------------------------------------------------------------

// CellFrame structure defines the cell frame to be used and how many ticks
// has to be mantained. This is used to create sprite widget animations.
type CellFrame struct {
	cells    CellGroup
	maxTicks int
	ticks    int
}

// NewCellFrame function creates a new SpriteFrame instance.
func NewCellFrame(cells CellGroup, maxTicks int) *CellFrame {
	return &CellFrame{
		cells:    cells,
		maxTicks: maxTicks,
		ticks:    0,
	}
}

// -----------------------------------------------------------------------------
// CellFrame public methods
// -----------------------------------------------------------------------------

// GetCells method returns the canvas instance number.
func (f *CellFrame) GetCells() CellGroup {
	return f.cells
}

// Inc method increase the actual counter for the sprite frame instance and
// returns if that counter has reached the maxTicks.
func (f *CellFrame) Inc() bool {
	f.ticks++
	return f.ticks >= f.maxTicks
}

// Reset method resets the sprite frame counter value.
func (f *CellFrame) Reset() {
	f.ticks = 0
}

// -----------------------------------------------------------------------------
//
// CellFrames
//
// -----------------------------------------------------------------------------

// CellFrames type groups an array of CellFrame, which is used to create sprite
// widgets animations.
type CellFrames []*CellFrame
