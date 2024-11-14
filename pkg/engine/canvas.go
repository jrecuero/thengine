// canvas.go contains required structures and method for handling any canvas
// in the application. A canvas represents a sequence of characters to be
// displayed in the camera in a given sequence of coordinates.
package engine

import (
	"os"
	"strings"

	"github.com/gdamore/tcell/v2"
	"github.com/jrecuero/thengine/pkg/api"
	"github.com/jrecuero/thengine/pkg/tools"
)

// -----------------------------------------------------------------------------
// Package public constants
// -----------------------------------------------------------------------------

var (
	CanvasRectSingleLine []rune = []rune{tcell.RuneULCorner,
		tcell.RuneURCorner,
		tcell.RuneLLCorner,
		tcell.RuneLRCorner,
		tcell.RuneHLine,
		tcell.RuneVLine}
	CanvasRectDoubleLine []rune = []rune{'╔', '╗', '╚', '╝', '═', '║'}
)

// -----------------------------------------------------------------------------
// iterCanvas
// -----------------------------------------------------------------------------

// iterCanvas structure defines index required to iterate a canvas.
type iterCanvas struct {
	Row int
	Col int
}

// newiterCanvas function creates a new iterCanvas instance.
func newiterCanvas() *iterCanvas {
	return &iterCanvas{}
}

// -----------------------------------------------------------------------------
// Row
// -----------------------------------------------------------------------------

// Row struct defines a sequence of cells for a row in the canvas. It defines
// all columns in the canvas.
type Row struct {
	Cols []ICell
}

// NewRow function creates a new Row instance with the give size.
func NewRow(size int) *Row {
	return &Row{
		Cols: make([]ICell, size),
	}
}

// SaveToDict method saves the instance information as a map.
func (r *Row) SaveToDict() map[string]any {
	result := map[string]any{}
	cols := []map[string]any{}
	for _, col := range r.Cols {
		cols = append(cols, col.SaveToDict())
	}
	result["cols"] = cols
	return result
}

// -----------------------------------------------------------------------------
//
// Canvas
//
// -----------------------------------------------------------------------------

// Canvas struct defines all rows and columns for the characters to be
// displayed in the camera.
type Canvas struct {
	Rows []*Row
	iter *iterCanvas
}

// -----------------------------------------------------------------------------
// package [Canvas] public functions
// -----------------------------------------------------------------------------

// NewCanvas function creates a new Canvas instance with the given number of
// columns and rows.
func NewCanvas(size *api.Size) *Canvas {
	if size == nil {
		return nil
	}
	cols, rows := size.Get()
	canvas := Canvas{
		Rows: make([]*Row, rows),
	}
	for i := 0; i < rows; i++ {
		canvas.Rows[i] = NewRow(cols)
	}
	return &canvas
}

// CloneCanvas function clones the given Canvas into a new Canvas instance with
// same size and all same cells.
func CloneCanvas(canvas *Canvas) *Canvas {
	size := canvas.Size()
	newCanvas := NewCanvas(size)
	newCanvas.Clone(canvas)
	return newCanvas
}

// NewCanvasFromString function creates a new canvas where the content is the
// given string (multi-line is allowed) with the given color.
func NewCanvasFromString(str string, style *tcell.Style) *Canvas {
	// Calculate width and height based on the number of lines in the string
	// and the max length for every line.
	lines := strings.Split(str, "\n")
	width := 0
	height := len(lines)
	for _, line := range lines {
		width = tools.Max(width, len(line))
	}
	canvas := NewCanvas(api.NewSize(width, height))
	for row, line := range lines {
		for col, ch := range line {
			cell := NewCell(style, ch)
			canvas.SetCellAt(api.NewPoint(col, row), cell)
		}
	}
	return canvas
}

// NewCanvasFromFile function creates a new canvas with the content of the
// file.
func NewCanvasFromFile(filename string, style *tcell.Style) *Canvas {
	content, err := os.ReadFile(filename)
	if err != nil {
		tools.Logger.WithField("module", "canvas").
			WithField("method", "NewCanvasFromFile").
			Errorf("Error opening %s Err=%+v", filename, err)
		return nil
	}
	return NewCanvasFromString(string(content), style)
}

// -----------------------------------------------------------------------------
// Canvas iterator methods.
// -----------------------------------------------------------------------------

// CreateIter method creates a new canvas iterator.
func (c *Canvas) CreateIter() {
	c.iter = newiterCanvas()
}

// IterHasNext method checks if there are still some entries to iterate.
func (c *Canvas) IterHasNext() bool {
	return (c.iter.Col < c.Width()) && (c.iter.Row < c.Height())
}

// IterGetNext method returns the next entry to iterate and increase iterator
// counters.
func (c *Canvas) IterGetNext() (int, int, ICell) {
	col := c.iter.Col
	row := c.iter.Row
	point := api.NewPoint(col, row)
	cell := c.GetCellAt(point)
	c.iter.Col++
	if c.iter.Col >= c.Width() {
		c.iter.Col = 0
		c.iter.Row++
	}
	return col, row, cell
}

// -----------------------------------------------------------------------------
// Canvas public methods.
// -----------------------------------------------------------------------------

// Clone method clones the given Canvas into the instance. Every cell will be
// copied to the Canvas instance.
func (c *Canvas) Clone(canvas *Canvas) {
	if c.Size().IsEqual(canvas.Size()) {
		for x, row := range canvas.Rows {
			for y, cell := range row.Cols {
				if cell == nil {
					continue
				}
				// create a new instance for every cell position.
				c.Rows[x].Cols[y] = CloneCell(cell)
			}
		}
	}
}

// FillWithCell method fills all cells in the canvas with the given cell.
func (c *Canvas) FillWithCell(cell ICell) {
	for row, rows := range c.Rows {
		for col := range rows.Cols {
			// Create a new instance for every cell position.
			c.SetCellAt(api.NewPoint(col, row), CloneCell(cell))
		}
	}
}

// GetCellAt method returns the cell in the canvas at the given row and column.
func (c *Canvas) GetCellAt(point *api.Point) ICell {
	if point == nil {
		point = api.NewPoint(0, 0)
	}
	if c.IsInside(point) {
		return c.Rows[point.Y].Cols[point.X]
	}
	return nil
}

// GetRect method returns the Rect instance for the Canvas. The Rect will have
// a zero origin (0, 0) and the Canvas width and height.
func (c *Canvas) GetRect() *api.Rect {
	return api.NewRect(api.NewPoint(0, 0), c.Size())
}

// GetRuneAt method returns the Rune in the canvas at the given row and
// column.
func (c *Canvas) GetRuneAt(point *api.Point) rune {
	if cell := c.GetCellAt(point); cell != nil {
		ch := cell.GetRune()
		return ch
	}
	return 0
}

// GetColorAt method returns the Color in the canvas at the given row and
// column.
func (c *Canvas) GetStyleAt(point *api.Point) *tcell.Style {
	if cell := c.GetCellAt(point); cell != nil {
		return cell.GetStyle()
	}
	return nil
}

// Height method returns the canvas number of rows.
func (c *Canvas) Height() int {
	return len(c.Rows)
}

// IsEqual method checks if the given Canvas has all same cells as the
// instance.
func (c *Canvas) IsEqual(canvas *Canvas) bool {
	// If both Canvas don't have the same size they can not be equals.
	if !c.Size().IsEqual(canvas.Size()) {
		return false
	}
	// check every cell is equal in both canvas.
	for x, rows := range canvas.Rows {
		for y, cell := range rows.Cols {
			canvasCell := c.Rows[x].Cols[y]
			// if both cells are nil, continue.
			if (canvasCell == nil) && (cell == nil) {
				continue
			}
			// if only one cell is nil, canvas are not equal.
			if (canvasCell == nil) || (cell == nil) {
				return false
			}
			// if cell is not equal, canvas are not equal.
			if !canvasCell.IsEqual(cell) {
				return false
			}
		}
	}
	return true
}

// IsInside method returns if the given cell position is inside the canvas.
func (c *Canvas) IsInside(point *api.Point) bool {
	if (point.X >= c.Width()) || (point.X < 0) {
		return false
	}
	if (point.Y >= c.Height()) || (point.Y < 0) {
		return false
	}
	return true
}

// Render method renders the canvas into the camera.
func (c *Canvas) Render(camera ICamera) {
	for r, rows := range c.Rows {
		for c, cell := range rows.Cols {
			position := api.NewPoint(c, r)
			camera.RenderCellAt(position, cell)
		}
	}
}

// RenderRectAt method renders part of the canvas defines by the given
// rectangle at the given position.
func (c *Canvas) RenderRectAt(camera ICamera, rectOffset *api.Point, rectSize *api.Size, offset *api.Point) {
	for row := 0; row < rectSize.H; row++ {
		for col := 0; col < rectSize.W; col++ {
			canvasRow := row + rectOffset.Y
			canvasCol := col + rectOffset.X
			if cell := c.GetCellAt(api.NewPoint(canvasCol, canvasRow)); cell != nil {
				camera.RenderCellAt(api.NewPoint(offset.X+col, offset.Y+row), cell)
			}
		}
	}
}

// RenderAt method renders the canvas into the camera at the given position.
func (c *Canvas) RenderAt(camera ICamera, offset *api.Point) {
	for r, rows := range c.Rows {
		for c, cell := range rows.Cols {
			if cell != nil {
				position := api.NewPoint(c, r)
				position.Add(offset)
				camera.RenderCellAt(position, cell)
			}
		}
	}
}

// SaveToDict method saves the instance information as a map.
// TODO: to be revisited.
func (c *Canvas) SaveToDict() map[string]any {
	result := map[string]any{}
	rows := []map[string]any{}
	for _, r := range c.Rows {
		rows = append(rows, r.SaveToDict())
	}
	result["rows"] = rows
	return result
}

// SetCellAt method sets the given cell to the given row and column.
func (c *Canvas) SetCellAt(point *api.Point, cell ICell) bool {
	// if no point value is being passed, set the (0, 0) point as default.
	if point == nil {
		point = api.NewPoint(0, 0)
	}
	if c.IsInside(point) {
		c.Rows[point.Y].Cols[point.X] = cell
		return true
	}
	return false
}

// SetRineAt method sets the given Rune to the cell at the given row and
// column.
// If point given is nil, it updates the rhune in all cells in the canvas.
func (c *Canvas) SetRuneAt(point *api.Point, ch rune) bool {
	if point != nil {
		if cell := c.GetCellAt(point); cell != nil {
			cell.SetRune(ch)
			return true
		} else {
			for _, rows := range c.Rows {
				for _, cell := range rows.Cols {
					if cell != nil {
						cell.SetRune(ch)
					}
				}
			}
		}
	}
	return false
}

// SetStyleAt method sets the given Style to the cell at the given row and
// column.
// If point given is nil, it updates style in all cells in the canvas.
func (c *Canvas) SetStyleAt(point *api.Point, style *tcell.Style) bool {
	if point != nil {
		if cell := c.GetCellAt(point); cell != nil {
			cell.SetStyle(style)
			return true
		}
	} else {
		for _, rows := range c.Rows {
			for _, cell := range rows.Cols {
				if cell != nil {
					cell.SetStyle(style)
				}
			}
		}
	}
	return false
}

// SetStyleForRune method sets the given style for cells with the given rhune.
// If the given run is 0 it updates all empty rhunes.
func (c *Canvas) SetStyleForRune(ch rune, style *tcell.Style) {
	for x, rows := range c.Rows {
		for y, cell := range rows.Cols {
			if cell != nil && cell.GetRune() == ch {
				cell.SetStyle(style)
			} else if cell == nil && ch == 0 {
				c.SetCellAt(api.NewPoint(x, y), NewCell(style, ' '))
			}
		}
	}
}

// Size method returns the canvas number of columns and the number of rows.
func (c *Canvas) Size() *api.Size {
	if c.Height() == 0 {
		return api.NewSize(0, 0)
	}
	return api.NewSize(c.Width(), c.Height())
}

// ToString method returns the cell instance information as a string.
func (c *Canvas) ToString() string {
	result := ""
	for row, rows := range c.Rows {
		for col := range rows.Cols {
			cell := c.GetCellAt(api.NewPoint(col, row))
			if cell != nil {
				result = result + cell.ToString() + "\n"
			}
		}
	}
	return result
}

// Width method returns the canvas number of columns.
func (c *Canvas) Width() int {
	if c.Height() == 0 {
		return 0
	}
	return len(c.Rows[0].Cols)
}

// WriteStringInCanvas method writes the given string in the canvas. Any
// character exciding the canvas size is missed.
func (c *Canvas) WriteStringInCanvas(str string, style *tcell.Style) {
	lines := strings.Split(str, "\n")
	for row, line := range lines {
		if row >= c.Height() {
			break
		}
		for col, ch := range line {
			if col >= c.Width() {
				break
			}
			cell := NewCell(style, ch)
			c.SetCellAt(api.NewPoint(col, row), cell)
		}
	}
}

// WriteStringInCanvasAt method writes the given string in the canvas at the
// give position. Any character exciding the canvas size is missed.
func (c *Canvas) WriteStringInCanvasAt(str string, style *tcell.Style, position *api.Point) {
	lines := strings.Split(str, "\n")
	for rowLine, line := range lines {
		row := rowLine + position.Y
		if row >= c.Height() {
			break
		}
		for colLine, ch := range line {
			col := colLine + position.X
			if col >= c.Width() {
				break
			}
			cell := NewCell(style, ch)
			c.SetCellAt(api.NewPoint(col, row), cell)
		}
	}
}

// WriteRectangleInCanvasAt method write the given rectangle in the canvas at
// the given position.
func (c *Canvas) WriteRectangleInCanvasAt(position *api.Point, size *api.Size, style *tcell.Style, pattern []rune) {
	// if no position is passed, set the (0, 0) position as default.
	if position == nil {
		position = api.NewPoint(0, 0)
	}
	// if no size is passed, set the canvas size as default.
	if size == nil {
		size = c.Size()
	}
	var ul, ur, ll, lr, hl, vl rune
	if len(pattern) == 1 {
		ul = pattern[0]
		ur = pattern[0]
		ll = pattern[0]
		lr = pattern[0]
		hl = pattern[0]
		vl = pattern[0]
	} else if len(pattern) == 6 {
		ul = pattern[0]
		ur = pattern[1]
		ll = pattern[2]
		lr = pattern[3]
		hl = pattern[4]
		vl = pattern[5]
	} else {
		return
	}
	x, y := position.Get()
	for col := 1; col < (size.W - 1); col++ {
		cell := NewCell(style, hl)
		c.SetCellAt(api.NewPoint(x+col, y), cell)
		c.SetCellAt(api.NewPoint(x+col, y+size.H-1), cell)
	}
	for row := 1; row < (size.H - 1); row++ {
		cell := NewCell(style, vl)
		c.SetCellAt(api.NewPoint(x, y+row), cell)
		c.SetCellAt(api.NewPoint(x+size.W-1, y+row), cell)
	}
	c.SetCellAt(api.NewPoint(x, y), NewCell(style, ul))
	c.SetCellAt(api.NewPoint(x+size.W-1, y), NewCell(style, ur))
	c.SetCellAt(api.NewPoint(x, y+size.H-1), NewCell(style, ll))
	c.SetCellAt(api.NewPoint(x+size.W-1, y+size.H-1), NewCell(style, lr))
}
