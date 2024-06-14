// timer.go module contains all logic required for timers based on the engine
// tick update mechanism.
package widgets

import (
	"time"

	"github.com/gdamore/tcell/v2"
	"github.com/jrecuero/thengine/pkg/engine"
	"github.com/jrecuero/thengine/pkg/tools"
)

// -----------------------------------------------------------------------------
// Package public constant
// -----------------------------------------------------------------------------
const (
	ForeverTimer   int = -1
	OneTimeTimer   int = 1
	CancelledTimer int = 0
)

// -----------------------------------------------------------------------------
//
// Timer
//
// -----------------------------------------------------------------------------

// Timer structure defines a baseline for any timer entity.
type Timer struct {
	*Widget
	interval      time.Duration
	time          time.Time
	count         int
	originalCount int
	elapsed       time.Duration
	running       bool
}

// NewTimer function creates a new Timer instance.
func NewTimer(name string, interval time.Duration, count int) *Timer {
	tools.Logger.WithField("module", "timer").
		WithField("function", "NewTimer").
		Debugf("%s", name)
	return &Timer{
		Widget:        NewWidget(name, nil, nil, nil),
		interval:      interval,
		count:         count,
		originalCount: count,
		elapsed:       0,
		running:       false,
	}
}

// -----------------------------------------------------------------------------
// Timer public methods
// -----------------------------------------------------------------------------

// CancelTimer method cancels the timer.
func (t *Timer) CancelTimer() {
	t.count = CancelledTimer
	t.running = false
}

// Draw method draws nothing.
func (t *Timer) Draw(engine.IScene) {
}

// Start methos starts the timer.
func (t *Timer) Start() {
	t.StartTimer()
}

// StartTimer method starts the timer.
func (t *Timer) StartTimer() {
	t.time = time.Now()
	t.running = true
}

// RestartTimer method re-starts the timer.
func (t *Timer) RestartTimer() {
	t.StartTimer()
	t.count = t.originalCount
}

// StopTimer method stops the timer
func (t *Timer) StopTimer() {
	now := time.Now()
	elapsed := now.Sub(t.time)
	t.elapsed += elapsed
	t.running = false
}

// Udpate method executes timer functionality. It check if timer has expired
// and command has to be called.
func (t *Timer) Update(event tcell.Event, scene engine.IScene) {
	defer t.Entity.Update(event, scene)
	if !t.running {
		return
	}
	if t.count == 0 {
		return
	}
	now := time.Now()
	if elapsed := now.Sub(t.time) + t.elapsed; elapsed < t.interval {
		//tools.Logger.WithField("module", "timer").WithField("function", "Update").Debugf("%d < %d", elapsed, t.interval)
		return
	}
	//tools.Logger.WithField("module", "timer").WithField("function", "Update").Debugf("%s callback", t.GetName())
	t.RunCallback(t)
	if t.count == ForeverTimer {
		t.time = time.Now()
		return
	}
	if t.count = t.count - 1; t.count != 0 {
		t.time = time.Now()
	}
}

var _ engine.IObject = (*Timer)(nil)
var _ engine.IFocus = (*Timer)(nil)
var _ engine.IEntity = (*Timer)(nil)
