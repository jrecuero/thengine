// input.go contains logic required to capture any events from termbox which
// are used by the engine.
package engine

import "github.com/nsf/termbox-go"

// -----------------------------------------------------------------------------
//
// Input
//
// -----------------------------------------------------------------------------

// Input structure defines all attributes required for handling keyboard input
// events triggered by termbox.
// EventQ channel keeps all events triggered by termbox and it used to pass them
// to the engine.
// Ctrl channel is used to terminate all input functionality.
type Input struct {
	EventQ chan termbox.Event
	Ctrl   chan bool
}

// NewInput function creates a new Input instance.
func NewInput() *Input {
	return &Input{
		EventQ: make(chan termbox.Event),
		Ctrl:   make(chan bool, 2),
	}
}

// -----------------------------------------------------------------------------
// Module Private methods
// -----------------------------------------------------------------------------

// poll function is an infinite loop looking at a channel for any termbox event.
func poll(i *Input) {
loop:
	for {
		select {
		case <-i.Ctrl:
			break loop
		default:
			i.EventQ <- termbox.PollEvent()
		}
	}
}

// -----------------------------------------------------------------------------
// Input Public methods
// -----------------------------------------------------------------------------

// Start method starts the input mechanism polling termbox events.
func (i *Input) Start() {
	go poll(i)
}

// Stop method ends the input mechanism.
func (i *Input) Stop() {
	i.Ctrl <- true
}
