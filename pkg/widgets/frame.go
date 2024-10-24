// frame.go contains all information to create frames used for animation
// sprites or widgets.
package widgets

import "github.com/jrecuero/thengine/pkg/engine"

// -----------------------------------------------------------------------------
//
// IFrame
//
// -----------------------------------------------------------------------------

// IFrame interface is the common interface for any frame used in animated
// widgets or sprites.
type IFrame interface {
	GetCanvas() *engine.Canvas
	GetCells() engine.CellGroup
	Inc() bool
	Reset()
}

// -----------------------------------------------------------------------------
//
// Frame
//
// -----------------------------------------------------------------------------

// Frame structure defines the frame to be used and how many ticks has to
// be mantained.
type Frame struct {
	canvas   *engine.Canvas
	cells    engine.CellGroup
	maxTicks int
	ticks    int
}

// NewFrame function creates a new empty Frame instance without frames.
func NewFrame(maxTicks int) *Frame {
	return &Frame{
		canvas:   nil,
		cells:    nil,
		maxTicks: maxTicks,
		ticks:    0,
	}
}

// NewFrameWithCanvas function creates a new Frame instance using Canvas for
// frames.
func NewFrameWithCanvas(canvas *engine.Canvas, maxTicks int) *Frame {
	return &Frame{
		canvas:   canvas,
		cells:    nil,
		maxTicks: maxTicks,
		ticks:    0,
	}
}

// NewFrameWithCells function creates a new Frame instance using Cells for
// frames.
func NewFrameWithCells(cells engine.CellGroup, maxTicks int) *Frame {
	return &Frame{
		canvas:   nil,
		cells:    cells,
		maxTicks: maxTicks,
		ticks:    0,
	}
}

// -----------------------------------------------------------------------------
// Frame public methods
// -----------------------------------------------------------------------------

// GetCanvas method returns the canvas instance number.
func (f *Frame) GetCanvas() *engine.Canvas {
	return f.canvas
}

func (f *Frame) GetCells() engine.CellGroup {
	return f.cells
}

// Inc method increase the actual counter for the frame instance and returns if
// that counter has reached the maxTicks.
func (f *Frame) Inc() bool {
	f.ticks++
	return f.ticks >= f.maxTicks
}

// Reset method resets the frame counter value.
func (f *Frame) Reset() {
	f.ticks = 0
}

func (f *Frame) SetCenvas(canvas *engine.Canvas) {
	f.canvas = canvas
}

func (f *Frame) SetCells(cells engine.CellGroup) {
	f.cells = cells
}

var _ IFrame = (*Frame)(nil)
