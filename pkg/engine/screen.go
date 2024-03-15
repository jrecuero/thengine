// screen.go contains all structures and methods required to handle the screen
// that will be displayed. Screen is mostly composed from a canvas with the
// size of the screen.
package engine

import (
	"github.com/gdamore/tcell/v2"
	"github.com/jrecuero/thengine/pkg/api"
)

// -----------------------------------------------------------------------------
//
// IScreen
//
// -----------------------------------------------------------------------------

// IScreen interface defines all functions a Screen has to implement.
type IScreen interface {
	Draw(bool, tcell.Screen)
	GetOrigin() *api.Point
	RenderCellAt(*api.Point, *Cell) bool
	SetDryRun(bool)
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
// TODO: Screen requires an origin point to be used as offset in the engine
// display tcell.Screen.
type Screen struct {
	origin    *api.Point
	oldCanvas *Canvas
	canvas    *Canvas
	dryRun    bool
}

// NewScreen function creates a new screen with the given width and height.
func NewScreen(origin *api.Point, size *api.Size) *Screen {
	if origin == nil {
		origin = api.NewPoint(0, 0)
	}
	return &Screen{
		origin:    origin,
		oldCanvas: NewCanvas(size),
		canvas:    NewCanvas(size),
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
	if newCell.Style != nil {
		fg, bg, attrs := newCell.Style.Decompose()
		_ = oldCell.Style.Foreground(fg).Background(bg).Attributes(attrs)
	}
}

// -----------------------------------------------------------------------------
// Screen private methods
// -----------------------------------------------------------------------------

// drawCanvasInDisplay function draws the canvas content into the displays using
// termbox API.
func (s *Screen) drawCanvasInDisplay(screen tcell.Screen) {
	for r, rows := range s.canvas.Rows {
		for c, cell := range rows.Cols {
			if cell == nil {
				continue
			}
			// skip termbox call
			if !s.dryRun {
				fg, bg, attrs := cell.Style.Decompose()
				style := tcell.StyleDefault.Background(bg).Foreground(fg).Attributes(attrs)

				screen.SetContent(c+s.origin.X, r+s.origin.Y, cell.Rune, nil, style)
			}
		}
	}
}

// -----------------------------------------------------------------------------
// Screen public methods
// -----------------------------------------------------------------------------

// Draw method draws the canvas content in the display.
func (s *Screen) Draw(flush bool, screen tcell.Screen) {
	if flush || !s.oldCanvas.IsEqual(s.canvas) {
		s.drawCanvasInDisplay(screen)
		if !s.dryRun {
			screen.Show()
		}
		s.oldCanvas = CloneCanvas(s.canvas)
	}
}

// GetCanvas method returns the screen canvas.
func (s *Screen) GetCanvas() *Canvas {
	return s.canvas
}

// GetRect method returns the rectangule for the screen.
func (s *Screen) GetRect() *api.Rect {
	rect := s.canvas.GetRect()
	rect.SetOrigin(s.origin)
	return rect
}

// GetOrigin method returns the origin point for the screen.
func (s *Screen) GetOrigin() *api.Point {
	return s.origin
}

// RenderCellAt method renders the cell in the screen canvas.
func (s *Screen) RenderCellAt(point *api.Point, cell *Cell) bool {
	if canvasCell := s.canvas.GetCellAt(point); canvasCell != nil {
		renderCell(canvasCell, cell)
		return true
	}
	// if the cell in the screen was nil, create a new one.
	s.canvas.SetCellAt(point, CloneCell(cell))
	return true
}

// SetDryRun method sets the dryRun variable to set dryRun flag which avoid any
// ncurses call.
func (s *Screen) SetDryRun(dryRun bool) {
	s.dryRun = dryRun
}

// Size method returns the screen size as width and height.
func (s *Screen) Size() *api.Size {
	return s.canvas.Size()
}

var _ IScreen = (*Screen)(nil)
