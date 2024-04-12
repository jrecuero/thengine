// engine.go contains all structures and method required for handling the
// application engine.
package engine

import (
	"fmt"
	"os"
	"runtime/debug"
	"time"

	"github.com/gdamore/tcell/v2"
	"github.com/jrecuero/thengine/pkg/api"
	"github.com/jrecuero/thengine/pkg/tools"
)

// -----------------------------------------------------------------------------
// Package global constants
// -----------------------------------------------------------------------------

const (
	EngineMainSceneName = "engine/main"
)

// -----------------------------------------------------------------------------
// Package global variables
// -----------------------------------------------------------------------------

var (
	EngineSingleton *Engine
)

// -----------------------------------------------------------------------------
// Package global functions
// -----------------------------------------------------------------------------

// GetEngine funcion return the Engine instance singleton if it exists or
// creates the singleton instance.
func GetEngine() *Engine {
	if EngineSingleton == nil {
		EngineSingleton = newEngine()
	}
	return EngineSingleton
}

// -----------------------------------------------------------------------------
//
// Engine
//
// -----------------------------------------------------------------------------

// Engine struct contains all attributes required for handling the appplication
// engine.
// screen tcell.Screen instance used to display any application object.
type Engine struct {
	screen       tcell.Screen
	sceneManager *SceneManager
	focusManager *FocusManager
	ctrlCh       chan bool
	eventCh      chan tcell.Event
	dryRun       bool
}

// newEngine function creates a new Engine instance.
func newEngine() *Engine {
	engine := &Engine{
		sceneManager: NewSceneManager(),
		focusManager: NewFocusManager(),
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
			e.eventCh <- e.screen.PollEvent()
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

// Consume method calls all underneath instances to consume all messages from
// the mailbox.
func (e *Engine) Consume() {
	e.sceneManager.Consume()
}

// CreateEngineScene method proceeds to create a new an unique scene owned by
// the engine. The size of the scene is equal to the size of the tcell.Screen.
func (e *Engine) CreateEngineScene() error {

	if e.screen == nil {
		return fmt.Errorf("Engine screen was not created. Engine.Init() has to be called")
	}

	// Create a new screen to be used by the engine scene with the same size as
	// the engine screen tcell.Screen.
	width, height := e.screen.Size()
	engineCamera := NewCamera(nil, api.NewSize(width, height))

	// Create a default scene for the engine that should be always present in
	// all applications.
	engineScene := NewScene(EngineMainSceneName, engineCamera)
	e.sceneManager.AddScene(engineScene)
	return nil
}

// Draw method proceeds to draws all entities in visible scenes.
func (e *Engine) Draw() {
	e.screen.Clear()
	e.sceneManager.Draw(e.screen)
}

// GetScreen method returns the tcell.Screen used by the engine.
func (e *Engine) GetScreen() tcell.Screen {
	return e.screen
}

// GetFocusManager method returns the focus manager instance.
func (e *Engine) GetFocusManager() *FocusManager {
	return e.focusManager
}

// GetSceneManager method returns the scene manager instance.
func (e *Engine) GetSceneManager() *SceneManager {
	return e.sceneManager
}

// Init method initializes are resources required to run the engine.
func (e *Engine) Init() {
	// Initialize tcell Screen.
	if !e.dryRun {
		var err error
		e.screen, err = tcell.NewScreen()
		if err != nil {
			panic(err)
		}
		if err = e.screen.Init(); err != nil {
			panic(err)
		}
		defaultStyle := tcell.StyleDefault.Background(tcell.ColorBlack).Foreground(tcell.ColorWhite)
		e.screen.SetStyle(defaultStyle)
	}
	e.sceneManager.Init(e.screen)
}

// Run method runs the engine in an infinite loop.
func (e *Engine) Run(fps float64) {
	var event tcell.Event
	var isRunning bool = true

	if !e.dryRun {
		e.startEventPoll()
		defer e.stopEventPoll()

		// Enable Mouse & Focus.
		//e.screen.EnableMouse()
		//e.screen.EnableFocus()

		// Clear the screen.
		e.screen.Clear()

		// panic handler.
		defer func() {
			recoverStack := false
			if err := recover(); err != nil {
				tools.Logger.WithField("module", "engine").WithField("function", "Run").Infof("panic %+v", err)
				tools.Logger.WithField("module", "engine").WithField("function", "Run").Infof("%s", string(debug.Stack()))
				recoverStack = true
			}
			if !e.dryRun {
				e.screen.Fini()
			}
			if recoverStack {
				fmt.Printf("%s", string(debug.Stack()))
			}
			os.Exit(0)
		}()
	}

	for isRunning {
		nowTime := time.Now()
		event = nil
		select {
		case event = <-e.eventCh:
			switch ev := event.(type) {
			case *tcell.EventResize:
				e.screen.Sync()
			case *tcell.EventMouse:
				tools.Logger.WithField("module", "engine").
					WithField("function", "Run").
					Debugf("mouse %+v", event)
			case *tcell.EventKey:
				switch ev.Key() {
				case tcell.KeyEscape:
					isRunning = false
				case tcell.KeyTab:
					tools.Logger.WithField("module", "engine").
						WithField("function", "Run").
						Debugf("tab update focus")
					e.sceneManager.UpdateFocus()
				case tcell.KeyRune:
					tools.Logger.WithField("module", "engine").
						WithField("function", "Run").
						Debugf("rune %s", string(ev.Rune()))
				default:
				}
			}
		default:
		}

		// update all engine resources.
		e.Update(event)

		// consume all message in the mailbox
		e.Consume()

		// draw all engine resources.
		e.Draw()

		timeToSleep := (time.Until(nowTime).Seconds() * 1000.0) + 1000.0/fps
		time.Sleep(time.Duration(timeToSleep) * time.Millisecond)
	}
}

// SetDryRun method sets the dryRun variable to set dryRun flag which avoid any
// ncurses call.
func (e *Engine) SetDryRun(dryRun bool) {
	e.dryRun = dryRun
	e.sceneManager.SetDryRun(dryRun)
}

// Start method starts any required functionality for running the engine.
func (e *Engine) Start() {
	e.sceneManager.Start()
}

// Update method proceeds to update all entities in active scenes.
func (e *Engine) Update(event tcell.Event) {
	e.sceneManager.Update(event)
}
