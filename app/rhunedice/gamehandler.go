package main

import (
	"github.com/gdamore/tcell/v2"
	"github.com/jrecuero/thengine/pkg/engine"
	"github.com/jrecuero/thengine/pkg/tools"
)

var (
	theGameHandler *GameHandler
)

// -----------------------------------------------------------------------------
//
// GameHandler
//
// -----------------------------------------------------------------------------

type GameHandler struct {
	*engine.Entity
}

func NewGameHandler() *GameHandler {
	if theGameHandler == nil {
		name := "handler/game/1"
		tools.Logger.WithField("module", "gamehandler").
			WithField("function", "NewGameHandler").
			Debugf(name)
		theGameHandler = &GameHandler{
			Entity: engine.NewHandler(name),
		}
		theGameHandler.SetFocusType(engine.SingleFocus)
		theGameHandler.SetFocusEnable(true)
	}
	return theGameHandler
}

// -----------------------------------------------------------------------------
// GameHandler private methods
// -----------------------------------------------------------------------------

// -----------------------------------------------------------------------------
// GameHandler public methods
// -----------------------------------------------------------------------------

func (h *GameHandler) Update(event tcell.Event, scene engine.IScene) {
	if !h.HasFocus() {
		return
	}
	switch ev := event.(type) {
	case *tcell.EventKey:
		switch ev.Key() {
		case tcell.KeyUp:
		case tcell.KeyDown:
		case tcell.KeyLeft:
		case tcell.KeyRight:
		case tcell.KeyRune:
			switch ev.Rune() {
			case 'A', 'a':
			case '1':
			}
		}
	}
}
