// gauge.go module contains all attributes and functionality to implement a
// gauge widget.
package widgets

import (
	"strings"
	"time"

	"github.com/gdamore/tcell/v2"
	"github.com/jrecuero/thengine/pkg/api"
	"github.com/jrecuero/thengine/pkg/engine"
	"github.com/jrecuero/thengine/pkg/tools"
)

// -----------------------------------------------------------------------------
// Package private methods
// -----------------------------------------------------------------------------

// updateCanvas function updates the canvas for a widget instance as a gauge.
func updateCanvas(widget engine.IEntity, total int, completed int) {
	var step float64 = float64(total) / float64(widget.GetSize().W)
	var completedSteps float64 = float64(completed) / step
	completedPattern := strings.Repeat(" ", int(completedSteps))
	basePattern := strings.Repeat(" ", widget.GetSize().W)
	widget.GetCanvas().WriteStringInCanvas(basePattern, widget.GetStyle())
	widget.GetCanvas().WriteStringInCanvas(completedPattern, tools.ReverseStyle(widget.GetStyle()))
	//tools.Logger.WithField("module", "gauge").
	//    WithField("function", "updateCanvas").
	//    Debugf("%d/%d %f %f %s", total, completed, step, completedSteps, tools.StyleToString(widget.GetStyle()))
}

// -----------------------------------------------------------------------------
//
// Gauge
//
// -----------------------------------------------------------------------------

// Gauge structure defines a baseline for any gauge entity.
type Gauge struct {
	*Widget
	total     int
	completed int
}

// NewGauge function creates a new Gauge instance.
func NewGauge(name string, position *api.Point, size *api.Size, style *tcell.Style, total int) *Gauge {
	gauge := &Gauge{
		Widget:    NewWidget(name, position, size, style),
		total:     total,
		completed: 0,
	}
	gauge.updateCanvas()
	return gauge
}

// -----------------------------------------------------------------------------
// Gauge private methods
// -----------------------------------------------------------------------------

// updateCanvas method updates the gauge widget canvas with latest completed
// value.
func (g *Gauge) updateCanvas() {
	updateCanvas(g, g.total, g.completed)
}

// -----------------------------------------------------------------------------
// Gauge public methods
// -----------------------------------------------------------------------------

// GetCompleted method returns the gauge completed attribute.
func (g *Gauge) GetCompleted() int {
	return g.completed
}

// GetTotal method returns the gauge total attribute.
func (g *Gauge) GetTotal() int {
	return g.total
}

// IncCompleted method increase gauge completed attribute the given steps.
func (g *Gauge) IncCompleted(steps int) int {
	g.completed += steps
	g.updateCanvas()
	return g.completed
}

// SetCompleted method sets a new value for the completed attribute.
func (g *Gauge) SetCompleted(completed int) {
	// completed gauge can not be lower than zero.
	if completed < 0 {
		completed = 0
	}
	g.completed = completed
	g.updateCanvas()
}

// SetTotal method sets a new value for the total attribute.
func (g *Gauge) SetTotal(total int) {
	g.total = total
}

var _ engine.IObject = (*Gauge)(nil)
var _ engine.IFocus = (*Gauge)(nil)
var _ engine.IEntity = (*Gauge)(nil)

// -----------------------------------------------------------------------------
//
// TimerGauge
//
// -----------------------------------------------------------------------------

// TimerGauge structure defines a baseline for a TimerGauge, where a gauge is
// increased every time the timer expires.
type TimerGauge struct {
	*Timer
	total     int
	completed int
}

// NewTimerGauge function creates a new TimerGauge instance.
func NewTimerGauge(name string, position *api.Point, size *api.Size, style *tcell.Style, interval time.Duration, total int) *TimerGauge {
	gauge := &TimerGauge{
		Timer:     NewTimer(name, interval, ForeverTimer),
		completed: 0,
		total:     total,
	}
	gauge.Timer.Widget = NewWidget(name, position, size, style)
	gauge.updateCanvas()
	return gauge
}

// -----------------------------------------------------------------------------
// TimerGauge private methods
// -----------------------------------------------------------------------------

// updateCanvas method updates the gauge widget canvas with latest completed
// value.
func (g *TimerGauge) updateCanvas() {
	updateCanvas(g, g.total, g.completed)
	//tools.Logger.WithField("module", "TimerGauge").
	//    WithField("function", "updateCanvas").
	//    Debugf("completed %d/%d %d '%s'", g.completed, g.total, completedSteps, pattern)
}

// -----------------------------------------------------------------------------
// TimerGauge public methods
// -----------------------------------------------------------------------------

// Draw method draws nothing.
func (g *TimerGauge) Draw(scene engine.IScene) {
	g.Timer.Widget.Draw(scene)
}

// RestartTimer method re-starts the timer.
func (g *TimerGauge) RestartTimer() {
	g.Timer.RestartTimer()
	g.completed = 0
	g.updateCanvas()
}

// Udpate method executes timer gauge functionality.
func (g *TimerGauge) Update(event tcell.Event, scene engine.IScene) {
	if !g.running {
		return
	}
	if g.count == 0 {
		return
	}
	now := time.Now()
	if elapsed := now.Sub(g.time) + g.elapsed; elapsed < g.interval {
		return
	}
	g.completed++
	g.updateCanvas()
	if g.completed < g.total {
		g.time = time.Now()
		return
	}
	g.RunCallback(g)
	g.CancelTimer()
}

var _ engine.IObject = (*TimerGauge)(nil)
var _ engine.IFocus = (*TimerGauge)(nil)
var _ engine.IEntity = (*TimerGauge)(nil)
