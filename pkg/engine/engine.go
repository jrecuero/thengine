// engine.go
//
// Package engine provides the core structures and methods to manage and run
// the application engine, handling screen management, events, and scene
// updates.
//
// Overview:
// The engine is responsible for initializing resources, managing the main
// event
// loop, handling input, drawing content to the screen, and coordinating
// updates and transitions within the application. The engine runs in a loop,
// processing input events, updating scenes, and rendering frames at a
// specified
// frames-per-second (fps) rate.
//
// Main Components:
// - Engine: Central struct that controls the application's core behavior,
//   including screen display, event handling, and scene management.
// - SceneManager: Manages scenes within the application, providing functions
//   for adding, updating, and switching between scenes.
// - ObserverManager: Manages observer instances for handling event
// notifications.
// - FocusManager: Manages focus for interactive elements within the scenes.
//
// Global Constants and Variables:
// - EngineMainSceneName: Default scene name for the main engine scene.
// - EngineSingleton: Singleton instance of the Engine, accessible via
// GetEngine().
//
// Core Methods:
// - Run(fps float64): Runs the engine in an infinite loop, processing events
//   and updating scenes at the specified fps.
// - Start(), Stop(): Methods to initialize and terminate engine resources.
// - CreateEngineScene(): Creates a main scene for the engine with the screen's
//   size as its dimensions.
// - Draw(), Update(), Consume(): Core rendering, updating, and message
//   consumption methods, respectively.
//
// Event Handling:
// The engine listens for keyboard and mouse events using tcell, an ncurses
// library for handling terminal-based graphical interfaces. Events like key
// presses (Escape, Ctrl+C) and window resize events are captured and processed
// during the event loop.
//
// Initialization and Resource Management:
// - Init(): Initializes resources needed to run the engine, like the screen.
// - InitResources(): Initializes low-level resources (like tcell screen).
// - SetDryRun(bool): Configures the engine to run in dry mode, bypassing
//   certain screen and input functionalities.
//
// Error Handling and Recovery:
// The Run method includes a panic recovery mechanism, logging stack traces and
// errors for debugging and safe application shutdown.
//
// Dependencies:
// This package requires tcell for terminal rendering and an external tools
// package for logging errors and debug information.
//
// Example Usage:
// To create and run the engine in a new application:
//     engine := engine.GetEngine()
//     engine.Init()
//     engine.Run(60.0) // Run at 60 FPS
//

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
	ctrlCh          chan bool
	dryRun          bool
	eventCh         chan tcell.Event
	focusManager    *FocusManager
	isRunning       bool
	observerManager *ObserverManager
	sceneManager    *SceneManager
	screen          tcell.Screen
}

// newEngine function creates a new Engine instance.
func newEngine() *Engine {
	engine := &Engine{
		ctrlCh:          make(chan bool, 2),
		eventCh:         make(chan tcell.Event),
		focusManager:    NewFocusManager(),
		isRunning:       true,
		observerManager: NewObserverManager(),
		sceneManager:    NewSceneManager(),
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
func (e *Engine) CreateEngineScene() (IScene, error) {

	if e.screen == nil {
		return nil, fmt.Errorf("Engine screen was not created. Engine.Init() has to be called")
	}

	// Create a new screen to be used by the engine scene with the same size as
	// the engine screen tcell.Screen.
	width, height := e.screen.Size()
	engineCamera := NewCamera(nil, api.NewSize(width, height))

	// Create a default scene for the engine that should be always present in
	// all applications.
	engineScene := NewScene(EngineMainSceneName, engineCamera)
	e.sceneManager.AddScene(engineScene)
	return engineScene, nil
}

// Draw method proceeds to draws all entities in visible scenes.
func (e *Engine) Draw() {
	e.screen.Clear()
	e.sceneManager.Draw(e.screen)
}

func (e *Engine) End() {
	e.isRunning = false
}

// EndTick methods calls any functionality required at the bottom of the tick.
func (e *Engine) EndTick() {
	e.sceneManager.EndTick()
}

// GetScreen method returns the tcell.Screen used by the engine.
func (e *Engine) GetScreen() tcell.Screen {
	return e.screen
}

// GetFocusManager method returns the focus manager instance.
func (e *Engine) GetFocusManager() *FocusManager {
	return e.focusManager
}

// GetObserverManager method returns the observer manager instance.
func (e *Engine) GetObserverManager() *ObserverManager {
	return e.observerManager
}

// GetSceneManager method returns the scene manager instance.
func (e *Engine) GetSceneManager() *SceneManager {
	return e.sceneManager
}

// Init method initializes are resources required to run the engine.
func (e *Engine) Init() {
	e.sceneManager.Init(e.screen)
}

// InitResources methos initializes all engine low level resources (tcell).
func (e *Engine) InitResources() {
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
}

// Run method runs the engine in an infinite loop.
func (e *Engine) Run(fps float64) {
	var event tcell.Event

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
			err := recover()
			if err != nil {
				tools.Logger.WithField("module", "engine").
					WithField("struct", "Engine").
					WithField("method", "Run").
					Errorf("panic %+v", err)
				tools.Logger.WithField("module", "engine").
					WithField("struct", "Engine").
					WithField("method", "Run").
					Errorf("%s", string(debug.Stack()))
				recoverStack = true
			}
			if !e.dryRun {
				e.screen.Fini()
			}
			if recoverStack {
				fmt.Printf("%+v", err)
				fmt.Printf("%s", string(debug.Stack()))
			}
			os.Exit(0)
		}()
	}

	for e.isRunning {
		nowTime := time.Now()
		// proceed with any action at the very start of the tick before event
		// is polled.
		e.StartTick()

		event = nil
		select {
		case event = <-e.eventCh:
			switch ev := event.(type) {
			case *tcell.EventResize:
				e.screen.Sync()
			case *tcell.EventMouse:
				tools.Logger.WithField("module", "engine").
					WithField("struct", "Engine").
					WithField("method", "Run").
					Debugf("mouse %+v", event)
			case *tcell.EventKey:
				switch ev.Key() {
				case tcell.KeyEscape:
					//h.isRunning = false
				case tcell.KeyCtrlC:
					e.isRunning = false
				case tcell.KeyTab:
					tools.Logger.WithField("module", "engine").
						WithField("struct", "Engine").
						WithField("method", "Run").
						Debugf("tab update focus")
					e.sceneManager.UpdateFocus()
				case tcell.KeyRune:
					tools.Logger.WithField("module", "engine").
						WithField("struct", "Engine").
						WithField("method", "Run").
						Debugf("rune %s", string(ev.Rune()))
				default:
					tools.Logger.WithField("module", "engine").
						WithField("struct", "Engine").
						WithField("method", "Run").
						Debugf("key %+v", ev.Key())
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

		// proceed with any action at the very end of the tick after everything
		// has been processed.
		e.EndTick()

		timeToSleep := (time.Until(nowTime).Seconds() * 1000.0) + 1000.0/fps
		time.Sleep(time.Duration(timeToSleep) * time.Millisecond)
	}

	// stop all engine resources.
	e.Stop()
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

// StartTick method calls any functionality required at the top of the tick.
func (e *Engine) StartTick() {
	e.sceneManager.StartTick()
}

// Stop method stops any engine resources.
func (e *Engine) Stop() {
	e.sceneManager.Stop()
}

// Update method proceeds to update all entities in active scenes.
func (e *Engine) Update(event tcell.Event) {
	e.sceneManager.Update(event)
}
