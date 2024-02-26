// canvas.go contains required structures and method for handling any canvas
// in the application. A canvas represents a sequence of characters to be
// displayed in the screen in a given sequence of coordinates.
package engine

import (
	"os"
	"strings"

	"github.com/jrecuero/thengine/pkg/api"
	"github.com/jrecuero/thengine/pkg/tools"
)

// -----------------------------------------------------------------------------
// Row
// -----------------------------------------------------------------------------

// Row struct defines a sequence of cells for a row in the canvas. It defines
// all columns in the canvas.
type Row struct {
	Cols []*Cell
}

// NewRow function creates a new Row instance with the give size.
func NewRow(size int) *Row {
	return &Row{
		Cols: make([]*Cell, size),
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
// displayed in the screen.
type Canvas struct {
	Rows []*Row
}

// NewCanvas function creates a new Canvas instance with the given number of
// columns and rows.
func NewCanvas(size *api.Size) *Canvas {
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
func NewCanvasFromString(str string, color *api.Color) *Canvas {
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
			// Use black/white as default color if not provided.
			if color == nil {
				color = api.NewColor(api.ColorBlack, api.ColorWhite)
			}
			cell := NewCell(color, ch)
			canvas.SetCellAt(api.NewPoint(col, row), cell)
		}
	}
	return canvas
}

// NewCanvasFromFile function creates a new canvas with the content of the
// file.
func NewCanvasFromFile(filename string, color *api.Color) *Canvas {
	content, err := os.ReadFile(filename)
	if err != nil {
		tools.Logger.WithField("module", "canvas").Errorf("Error opening %s Err=%+v", filename, err)
		return nil
	}
	return NewCanvasFromString(string(content), color)
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
func (c *Canvas) FillWithCell(cell *Cell) {
	for row, rows := range c.Rows {
		for col := range rows.Cols {
			// Create a new instance for every cell position.
			c.SetCellAt(api.NewPoint(col, row), CloneCell(cell))
		}
	}
}

// GetCellAt method returns the cell in the canvas at the given row and column.
func (c *Canvas) GetCellAt(point *api.Point) *Cell {
	if c.IsInside(point) {
		return c.Rows[point.Y].Cols[point.X]
	}
	return nil
}

// GetColorAt method returns the Color in the canvas at the given row and
// column.
func (c *Canvas) GetColorAt(point *api.Point) *api.Color {
	if cell := c.GetCellAt(point); cell != nil {
		return cell.Color
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
		ch := cell.Rune
		return ch
	}
	return 0
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

// Reder method render the canvas into the screen.
func (c *Canvas) Render(screen IScreen) {
	for r, rows := range c.Rows {
		for c, cell := range rows.Cols {
			position := api.NewPoint(c, r)
			screen.RenderCellAt(position, cell)
		}
	}
}

// SaveToDict method saves the instance information as a map.
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
func (c *Canvas) SetCellAt(point *api.Point, cell *Cell) bool {
	if c.IsInside(point) {
		c.Rows[point.Y].Cols[point.X] = cell
		return true
	}
	return false
}

// SetColorAt method sets the given Color to the cell at the given row and
// column.
func (c *Canvas) SetColorAt(point *api.Point, color *api.Color) bool {
	if cell := c.GetCellAt(point); cell != nil {
		cell.Color = color
		return true
	}
	return false
}

// SetRineAt method sets the given Rune to the cell at the given row and
// column.
func (c *Canvas) SetRuneAt(point *api.Point, ch rune) bool {
	if cell := c.GetCellAt(point); cell != nil {
		cell.Rune = ch
		return true
	}
	return false
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
