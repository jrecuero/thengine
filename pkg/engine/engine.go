// engine.go contains all structures and method required for handling the
// application engine.
package engine

import (
	"os"
	"time"

	"github.com/gdamore/tcell/v2"
)

// -----------------------------------------------------------------------------
//
// Engine
//
// -----------------------------------------------------------------------------

// Engine struct contains all attributes required for handling the appplication
// engine.
// display tcell.Screen instance used to display any application object.
type Engine struct {
	display      tcell.Screen
	sceneManager *SceneManager
	ctrlCh       chan bool
	eventCh      chan tcell.Event
}

// NewEngine function creates a new Engine instance.
func NewEngine(engineScreen IScreen) *Engine {
	engine := &Engine{
		sceneManager: NewSceneManager(engineScreen),
		ctrlCh:       make(chan bool, 2),
		eventCh:      make(chan tcell.Event),
	}
	return engine
}

// -----------------------------------------------------------------------------
// Engine private methods
// -----------------------------------------------------------------------------

// eventPoll method polls the keyboard for any entry.
func (e *Engine) eventPoll() {
loop:
	for {
		select {
		case <-e.ctrlCh:
			break loop
		default:
			e.eventCh <- e.display.PollEvent()
		}
	}
}

// startEventPoll method starts the keyboard polling mechanism
func (e *Engine) startEventPoll() {
	go e.eventPoll()
}

// stopEventPoll method stops the keyboard polling mechanism
func (e *Engine) stopEventPoll() {
	e.ctrlCh <- true
}

// -----------------------------------------------------------------------------
// Engine public methods
// -----------------------------------------------------------------------------

// Draw method proceeds to draws all entities in visible scenes.
func (e *Engine) Draw() {
	e.sceneManager.Draw(e.display)
}

// GetDisplay method returns the tcell.Screen used by the engine.
func (e *Engine) GetDisplay() tcell.Screen {
	return e.display
}

// GetSceneManager method returns the SceneManager instance.
func (e *Engine) GetSceneManager() *SceneManager {
	return e.sceneManager
}

// Init method initializes are resources required to run the engine.
func (e *Engine) Init() {
	var err error
	e.display, err = tcell.NewScreen()
	if err != nil {
		panic(err)
	}
	if err = e.display.Init(); err != nil {
		panic(err)
	}
	defaultStyle := tcell.StyleDefault.Background(tcell.ColorBlack).Foreground(tcell.ColorWhite)
	e.display.SetStyle(defaultStyle)
}

// Run method runs the engine in an infinite loop.
func (e *Engine) Run(fps float64) {
	var event tcell.Event
	var isRunning bool = true

	e.startEventPoll()
	defer e.stopEventPoll()

	for isRunning {
		//tools.Logger.WithField("module", "engine").WithField("function", "Run").Infof("loop")
		nowTime := time.Now()
		select {
		case event = <-e.eventCh:
			switch ev := event.(type) {
			case *tcell.EventResize:
				e.display.Sync()
			case *tcell.EventKey:
				if ev.Key() == tcell.KeyEscape {
					isRunning = false
				}
			}
		default:
		}

		// update all engine resources.
		e.Update(event)

		// draw all engine resources.
		e.Draw()

		timeToSleep := (time.Until(nowTime).Seconds() * 1000.0) + 1000.0/fps
		time.Sleep(time.Duration(timeToSleep) * time.Millisecond)
	}
	e.display.Fini()
	os.Exit(0)
}

// Update method proceeds to update all entities in active scenes.
func (e *Engine) Update(event tcell.Event) {
	e.sceneManager.Update(event)
}
