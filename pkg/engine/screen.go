// screen.go contains all structures and methods required to handle the screen
// that will be displayed. Screen is mostly composed from a canvas with the
// size of the screen.
package engine

import (
	"github.com/jrecuero/thengine/pkg/api"
	"github.com/nsf/termbox-go"
)

// -----------------------------------------------------------------------------
//
// IScreen
//
// -----------------------------------------------------------------------------

// IScreen interface defines all functions a Screen has to implement.
type IScreen interface {
	Draw(bool)
	RenderCellAt(*api.Point, *Cell) bool
}

// -----------------------------------------------------------------------------
//
// Screen
//
// -----------------------------------------------------------------------------

// Screen struct contains all required information to display all application
// data in the display.
// oldCanvas Canvas instance contains the last canvas being flushed.
// Canvas Canvas instance contains the latest canvas to be flushed.
// DryRun bool flag is set true for testing where termbox is not called.
type Screen struct {
	OldCanvas *Canvas
	Canvas    *Canvas
	DryRun    bool
}

// NewScreen function creates a new screen with the given width and height.
func NewScreen(size *api.Size) *Screen {
	return &Screen{
		OldCanvas: NewCanvas(size),
		Canvas:    NewCanvas(size),
	}
}

// -----------------------------------------------------------------------------
// Package private methods
// -----------------------------------------------------------------------------

// renderCell function updates a cell in the canvas based on the previous value
// for that cell in the canvas.
func renderCell(oldCell *Cell, newCell *Cell) {
	if newCell.Rune != 0 {
		oldCell.Rune = newCell.Rune
	}
	if newCell.Color.Fg != api.ColorDefault {
		oldCell.Color.Fg = newCell.Color.Fg
	}
	if newCell.Color.Bg != api.ColorDefault {
		oldCell.Color.Bg = newCell.Color.Bg
	}
}

// -----------------------------------------------------------------------------
// Screen private methods
// -----------------------------------------------------------------------------

// drawCanvasInDisplay function draws the canvas content into the displays using
// termbox API.
func (s *Screen) drawCanvasInDisplay() {
	for r, rows := range s.Canvas.Rows {
		for c, cell := range rows.Cols {
			if cell == nil {
				continue
			}
			// skip termbox call
			if !s.DryRun {
				termbox.SetCell(c, r, cell.Rune, termbox.Attribute(cell.Color.Fg), termbox.Attribute(cell.Color.Bg))
			}
		}
	}
}

// -----------------------------------------------------------------------------
// Screen public methods
// -----------------------------------------------------------------------------

// Draw method draws the canvas content in the display.
func (s *Screen) Draw(flush bool) {
	if flush || !s.OldCanvas.IsEqual(s.Canvas) {
		s.drawCanvasInDisplay()
		if !s.DryRun {
			termbox.Flush()
		}
		s.OldCanvas = CloneCanvas(s.Canvas)
	}
}

// GetRect method returns the rectangule for the screen.
func (s *Screen) GetRect() *api.Rect {
	return s.Canvas.GetRect()
}

// RenderCellAt method renders the cell in the screen canvas.
func (s *Screen) RenderCellAt(point *api.Point, cell *Cell) bool {
	if canvasCell := s.Canvas.GetCellAt(point); canvasCell != nil {
		renderCell(canvasCell, cell)
		return true
	}
	// if the cell in the screen was nil, create a new one.
	s.Canvas.SetCellAt(point, CloneCell(cell))
	return true
}

// Size method returns the screen size as width and height.
func (s *Screen) Size() *api.Size {
	return s.Canvas.Size()
}

var _ IScreen = (*Screen)(nil)
