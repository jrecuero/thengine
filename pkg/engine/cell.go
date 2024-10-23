// cell.go contains required structures and method for handling a cell in the
// application. A cell contains information about the rune and the color to be
// displayed in any position in the canvas or screen.
// Cell is the base class to represent information for a cell based only in the
// rune and the color.
// CellPos contains the cell information plus a position.
// CellGroup contains a groups of CellPos that can be used for any widget or
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

// -----------------------------------------------------------------------------
//
// CellPos
//
// -----------------------------------------------------------------------------

// CellPos struct defines a cell at a given position, which is used to create
// sprite widgets.
type CellPos struct {
	position *api.Point
	cell     *Cell
}

// NewCellPos function creates a new CellPos instance.
func NewCellPos(position *api.Point, cell *Cell) *CellPos {
	return &CellPos{
		position: position,
		cell:     cell,
	}
}

// -----------------------------------------------------------------------------
// CellPos public methods
// -----------------------------------------------------------------------------

// GetCell method returns the cell from the CellPos instance.
func (s *CellPos) GetCell() *Cell {
	return s.cell
}

// GetPosition method returs the position from the CellPos instance.
func (s *CellPos) GetPosition() *api.Point {
	return s.position
}

// SetCell method sets the cell in a CellPos instance.
func (s *CellPos) SetCell(cell *Cell) {
	s.cell = cell
}

// SetPosition method sets the position in a CellPos instance.
func (s *CellPos) SetPosition(position *api.Point) {
	s.position = position
}

// -----------------------------------------------------------------------------
//
// CellGroup
//
// -----------------------------------------------------------------------------

// CellGroup type groups an array of CellPos, which is used to create sprite
// widgets.
type CellGroup []*CellPos

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

// GetSpriteCells method returns the canvas instance number.
func (f *CellFrame) GetSpriteCells() CellGroup {
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
