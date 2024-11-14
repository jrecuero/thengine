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
// ICell
//
// -----------------------------------------------------------------------------

// ICell interface define the interface to be used for any cell implementation
// so we can create cells that contains some additional information.
type ICell interface {
	Clone(ICell)
	GetPayload() any
	GetPosition() *api.Point
	GetRune() rune
	GetStyle() *tcell.Style
	IsEqual(ICell) bool
	ToString() string
	SaveToDict() map[string]any
	SetPayload(any)
	SetPosition(*api.Point)
	SetRune(rune)
	SetStyle(*tcell.Style)
}

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
	payload  any
	position *api.Point
	rune     rune
	style    *tcell.Style
}

// NewCell function creates a new Cell instance with the given color and rune.
func NewCell(style *tcell.Style, ch rune) *Cell {
	if style == nil {
		style = &tcell.StyleDefault
	}
	return &Cell{
		payload:  nil,
		position: nil,
		rune:     ch,
		style:    CloneStyle(style),
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
func CloneCell(cell ICell) *Cell {
	return &Cell{
		payload:  nil,
		position: api.ClonePoint(cell.GetPosition()),
		rune:     cell.GetRune(),
		style:    CloneStyle(cell.GetStyle()),
	}
}

// -----------------------------------------------------------------------------
// Cell public methods
// -----------------------------------------------------------------------------

// Clone method clones all attributes from the given Cell in to the instance.
func (c *Cell) Clone(cell ICell) {
	c.position = api.ClonePoint(cell.GetPosition())
	c.rune = cell.GetRune()
	c.style = CloneStyle(cell.GetStyle())
}

func (c *Cell) GetPayload() any {
	return c.payload
}

// GetPosition method returs the position from the Cell instance.
func (c *Cell) GetPosition() *api.Point {
	return c.position
}

func (c *Cell) GetRune() rune {
	return c.rune
}

func (c *Cell) GetStyle() *tcell.Style {
	return c.style
}

// IsEqual method checks if the given Cell is equal to the instance, where Color
// and Rune should be the same.
func (c *Cell) IsEqual(cell ICell) bool {
	return CompareStyle(c.GetStyle(), cell.GetStyle()) &&
		(c.GetRune() == cell.GetRune()) &&
		(c.GetPosition() == cell.GetPosition())

}

// ToString method returns the cell instance information as a string.
func (c *Cell) ToString() string {
	if c.GetPosition() != nil {
		return fmt.Sprintf("[%c]%s %s", c.GetRune(), StyleToString(c.GetStyle()), c.GetPosition().ToString())
	}
	return fmt.Sprintf("[%c]%s", c.GetRune(), StyleToString(c.GetStyle()))
}

// SaveToDict method saves the instance information as a map.
// TODO: to be revisited.
func (c *Cell) SaveToDict() map[string]any {
	result := map[string]any{}
	result["style"] = c.GetStyle()
	result["rune"] = c.GetRune()
	result["position"] = c.GetPosition()
	return result
}

func (c *Cell) SetPayload(payload any) {
	c.payload = payload
}

// SetPosition method sets the position in a Cell instance.
func (c *Cell) SetPosition(position *api.Point) {
	c.position = api.ClonePoint(position)
}

func (c *Cell) SetRune(r rune) {
	c.rune = r
}

func (c *Cell) SetStyle(style *tcell.Style) {
	c.style = CloneStyle(style)
}

// -----------------------------------------------------------------------------
//
// CellGroup
//
// -----------------------------------------------------------------------------

// CellGroup type groups an array of Cell, which is used to create sprite
// widgets.
type CellGroup []ICell

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

var _ ICell = (*Cell)(nil)
