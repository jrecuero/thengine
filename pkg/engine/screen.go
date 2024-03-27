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
	Init(tcell.Screen)
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
	origin  *api.Point
	size    *api.Size
	display tcell.Screen
	dryRun  bool
}

// NewScreen function creates a new screen with the given width and height.
func NewScreen(origin *api.Point, size *api.Size) *Screen {
	if origin == nil {
		origin = api.NewPoint(0, 0)
	}
	return &Screen{
		origin:  origin,
		size:    size,
		display: nil,
	}
}

// -----------------------------------------------------------------------------
// Package private methods
// -----------------------------------------------------------------------------

// -----------------------------------------------------------------------------
// Screen private methods
// -----------------------------------------------------------------------------

// -----------------------------------------------------------------------------
// Screen public methods
// -----------------------------------------------------------------------------

// Draw method draws the canvas content in the display.
func (s *Screen) Draw(flush bool, screen tcell.Screen) {
	if flush {
		if !s.dryRun {
			s.display.Show()
		}
	}
}

// GetOrigin method returns the origin point for the screen.
func (s *Screen) GetOrigin() *api.Point {
	return s.origin
}

// GetSize method returns the size for the screen.
func (s *Screen) GetSize() *api.Size {
	return s.size
}

// Init method initializes the screen instance.
func (s *Screen) Init(display tcell.Screen) {
	s.display = display
}

// RenderCellAt method renders the cell in the screen canvas.
func (s *Screen) RenderCellAt(point *api.Point, cell *Cell) bool {
	if !s.dryRun {
		col, row := point.Get()
		fg, bg, attrs := cell.Style.Decompose()
		style := tcell.StyleDefault.Background(bg).Foreground(fg).Attributes(attrs)
		s.display.SetContent(col+s.origin.X, row+s.origin.Y, cell.Rune, nil, style)
	}
	return true
}

// SetDryRun method sets the dryRun variable to set dryRun flag which avoid any
// ncurses call.
func (s *Screen) SetDryRun(dryRun bool) {
	s.dryRun = dryRun
}

var _ IScreen = (*Screen)(nil)
