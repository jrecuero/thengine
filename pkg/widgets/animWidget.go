// animWidget.go contains all resources required to create an animated widget
// based on a set of canvas instances.
package widgets

import (
	"github.com/gdamore/tcell/v2"
	"github.com/jrecuero/thengine/pkg/api"
	"github.com/jrecuero/thengine/pkg/engine"
)

// -----------------------------------------------------------------------------
//
// Frame
//
// -----------------------------------------------------------------------------

// Frame structure defines the frame to be used and how many ticks has to
// be mantained.
type Frame struct {
	canvas   *engine.Canvas
	maxTicks int
	ticks    int
}

// NewFrameInfo function creates a new Frame instance.
func NewFrame(canvas *engine.Canvas, maxTicks int) *Frame {
	return &Frame{
		canvas:   canvas,
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

// -----------------------------------------------------------------------------
//
// AnimWidget
//
// -----------------------------------------------------------------------------

// AnimWidget structure defines all attributes and methods for any basic
// animated widget.
type AnimWidget struct {
	*Widget
	frames        []*Frame
	frameTraverse int
}

// NewAnimWidget function creates a new AnimWidget instance.
func NewAnimWidget(name string, position *api.Point, size *api.Size, frames []*Frame, initFrame int) *AnimWidget {
	widget := &AnimWidget{
		Widget:        NewWidget(name, position, size, nil),
		frames:        frames,
		frameTraverse: initFrame,
	}
	widget.updateCanvas()
	return widget
}

// -----------------------------------------------------------------------------
// AnimWidget private methods
// -----------------------------------------------------------------------------

func (w *AnimWidget) updateCanvas() {
	canvas := w.frames[w.frameTraverse].GetCanvas()
	w.SetCanvas(canvas)
}

// -----------------------------------------------------------------------------
// AnimWidget public methods
// -----------------------------------------------------------------------------

// Update method updates the entity instance.
func (w *AnimWidget) Update(event tcell.Event, scene engine.IScene) {
	if w.IsActive() {
		frame := w.frames[w.frameTraverse]
		if frame.Inc() {
			frame.Reset()
			w.frameTraverse = (w.frameTraverse + 1) % len(w.frames)
			w.updateCanvas()
		}
	}
}

var _ engine.IObject = (*AnimWidget)(nil)
var _ engine.IFocus = (*AnimWidget)(nil)
var _ engine.IEntity = (*AnimWidget)(nil)
