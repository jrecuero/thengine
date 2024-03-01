// engine.go contains all structures and method required for handling the
// application engine.
package engine

import (
	"os"

	"github.com/gdamore/tcell"
)

// -----------------------------------------------------------------------------
//
// Engine
//
// -----------------------------------------------------------------------------

type Engine struct {
	Screen tcell.Screen
}

func NewEngine() *Engine {
	engine := &Engine{}
	return engine
}

// -----------------------------------------------------------------------------
// Engine public methods
// -----------------------------------------------------------------------------

func (e *Engine) Init() {
	// if err := termbox.Init(); err != nil {
	// 	tools.Logger.WithField("module", "engine").Infof("termbox init error=%+v", err)
	// 	panic(err)
	// }
	// termbox.SetOutputMode(termbox.Output256)
	// termbox.SetInputMode(termbox.InputEsc | termbox.InputMouse)
	// e.Input.Start()
	var err error
	e.Screen, err = tcell.NewScreen()
	if err != nil {
		panic(err)
	}
	if err = e.Screen.Init(); err != nil {
		panic(err)
	}
	defaultStyle := tcell.StyleDefault.Background(tcell.ColorBlack).Foreground(tcell.ColorWhite)
	e.Screen.SetStyle(defaultStyle)
}

func (e *Engine) Run() {
	// defer termbox.Close()
	// defer e.Input.Stop()
	// isRunning := true

	// for isRunning {
	// 	nowTime := time.Now()
	// 	select {
	// 	case event := <-e.Input.EventQ:
	// 		if event.Key == termbox.KeyCtrlC {
	// 			isRunning = false
	// 		}
	// 	default:
	// 	}
	// 	fps := 60.0
	// 	timeToSleep := (time.Until(nowTime).Seconds() * 1000.0) + 1000.0/fps
	// 	time.Sleep(time.Duration(timeToSleep) * time.Millisecond)
	// }

	for {
		switch event := e.Screen.PollEvent().(type) {
		case *tcell.EventResize:
			e.Screen.Sync()
		case *tcell.EventKey:
			if event.Key() == tcell.KeyEscape {
				e.Screen.Fini()
				os.Exit(0)
			}
		}
	}
}
