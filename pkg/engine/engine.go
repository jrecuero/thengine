// engine.go contains all structures and method required for handling the
// application engine.
package engine

import (
	"time"

	"github.com/jrecuero/thengine/pkg/tools"
	"github.com/nsf/termbox-go"
)

// -----------------------------------------------------------------------------
//
// Engine
//
// -----------------------------------------------------------------------------

type Engine struct {
	Input *Input
}

func NewEngine() *Engine {
	engine := &Engine{
		Input: NewInput(),
	}
	return engine
}

// -----------------------------------------------------------------------------
// Engine public methods
// -----------------------------------------------------------------------------

func (e *Engine) Init() {
	if err := termbox.Init(); err != nil {
		tools.Logger.WithField("module", "engine").Infof("termbox init error=%+v", err)
		panic(err)
	}
	termbox.SetOutputMode(termbox.Output256)
	termbox.SetInputMode(termbox.InputEsc | termbox.InputMouse)
	e.Input.Start()
}

func (e *Engine) Run() {
	defer termbox.Close()
	defer e.Input.Stop()
	isRunning := true

	for isRunning {
		nowTime := time.Now()
		select {
		case event := <-e.Input.EventQ:
			if event.Key == termbox.KeyCtrlC {
				isRunning = false
			}
		default:
		}
		fps := 60.0
		timeToSleep := (time.Until(nowTime).Seconds() * 1000.0) + 1000.0/fps
		time.Sleep(time.Duration(timeToSleep) * time.Millisecond)
	}
}
