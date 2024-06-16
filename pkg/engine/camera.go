// camera.go contains all structures and methods required to handle the camera
// that will be displayed. Camera is mostly composed from a canvas with the
// size of the camera.
package engine

import (
	"github.com/gdamore/tcell/v2"
	"github.com/jrecuero/thengine/pkg/api"
	"github.com/jrecuero/thengine/pkg/tools"
)

// -----------------------------------------------------------------------------
//
// ICamera
//
// -----------------------------------------------------------------------------

// ICamera interface defines all functions a Camera has to implement.
type ICamera interface {
	//Draw(bool, tcell.Screen)
	GetOrigin() *api.Point
	Init(tcell.Screen)
	RenderCellAt(*api.Point, *Cell) bool
	SetDryRun(bool)
}

// -----------------------------------------------------------------------------
//
// Camera
//
// -----------------------------------------------------------------------------

// Camera struct contains all required information to display all application
// data in the display.
// oldCanvas Canvas instance contains the last canvas being flushed.
// Canvas Canvas instance contains the latest canvas to be flushed.
// DryRun bool flag is set true for testing where termbox is not called.
// TODO: Camera requires an origin point to be used as offset in the engine
// display tcell.Screen.
type Camera struct {
	origin *api.Point
	size   *api.Size
	screen tcell.Screen
	dryRun bool
}

// NewCamera function creates a new camera with the given width and height.
func NewCamera(origin *api.Point, size *api.Size) *Camera {
	if origin == nil {
		origin = api.NewPoint(0, 0)
	}
	return &Camera{
		origin: origin,
		size:   size,
		screen: nil,
	}
}

// -----------------------------------------------------------------------------
// Package private methods
// -----------------------------------------------------------------------------

// -----------------------------------------------------------------------------
// Camera private methods
// -----------------------------------------------------------------------------

// -----------------------------------------------------------------------------
// Camera public methods
// -----------------------------------------------------------------------------

// Draw method draws the canvas content in the screen.
//func (s *Camera) Draw(flush bool, screen tcell.Screen) {
//    if flush {
//        if !s.dryRun {
//            s.screen.Show()
//        }
//    }
//}

// GetOrigin method returns the origin point for the camera.
func (s *Camera) GetOrigin() *api.Point {
	return s.origin
}

// GetSize method returns the size for the camera.
func (s *Camera) GetSize() *api.Size {
	return s.size
}

// Init method initializes the camera instance.
func (s *Camera) Init(screen tcell.Screen) {
	tools.Logger.WithField("module", "camera").
		WithField("method", "Init").
		Debugf("set screen to %v", screen)
	s.screen = screen
}

// RenderCellAt method renders the cell in the camera canvas.
func (s *Camera) RenderCellAt(point *api.Point, cell *Cell) bool {
	if !s.dryRun {
		col, row := point.Get()
		fg, bg, attrs := cell.Style.Decompose()
		style := tcell.StyleDefault.Background(bg).Foreground(fg).Attributes(attrs)
		s.screen.SetContent(col+s.origin.X, row+s.origin.Y, cell.Rune, nil, style)
	}
	return true
}

// SetDryRun method sets the dryRun variable to set dryRun flag which avoid any
// ncurses call.
func (s *Camera) SetDryRun(dryRun bool) {
	s.dryRun = dryRun
}

var _ ICamera = (*Camera)(nil)
