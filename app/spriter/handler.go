package main

import (
	"github.com/gdamore/tcell/v2"
	"github.com/jrecuero/thengine/pkg/api"
	"github.com/jrecuero/thengine/pkg/engine"
)

const (
	HandlerName = "entity/handler/1"
)

var (
	theHandler *Handler
)

type Handler struct {
	*engine.Entity
}

func NewHandler() *Handler {
	if theHandler == nil {
		theHandler = &Handler{
			Entity: engine.NewHandler(HandlerName),
		}
		theHandler.SetFocusType(engine.SingleFocus)
		theHandler.SetFocusEnable(true)
	}
	return theHandler
}

// -----------------------------------------------------------------------------
// Handler public methods
// -----------------------------------------------------------------------------

func (h *Handler) updateCursorRune(cursor *Cursor, ch rune) {
	cell := engine.NewCell(&TheStyleWhiteOverBlack, ch)
	cursor.GetCanvas().SetCellAt(nil, cell)
}

func (h *Handler) Update(event tcell.Event, scene engine.IScene) {
	if !h.HasFocus() {
		return
	}
	c := scene.GetEntityByName(CursorName)
	if c == nil {
		return
	}
	cursor, ok := c.(*Cursor)
	if !ok {
		return
	}
	cursorX, cursorY := cursor.GetPosition().Get()
	cursorNewPosition := api.NewPoint(cursorX, cursorY)
	switch ev := event.(type) {
	case *tcell.EventKey:
		switch ev.Key() {
		case tcell.KeyUp:
			cursorNewPosition = api.NewPoint(cursorX, cursorY-1)
		case tcell.KeyDown:
			cursorNewPosition = api.NewPoint(cursorX, cursorY+1)
		case tcell.KeyLeft:
			cursorNewPosition = api.NewPoint(cursorX-1, cursorY)
		case tcell.KeyRight:
			cursorNewPosition = api.NewPoint(cursorX+1, cursorY)
		case tcell.KeyEnter:
			NewEntityHandler(scene, cursor)
		case tcell.KeyRune:
			h.updateCursorRune(cursor, ev.Rune())
		}
	}
	if TheDrawingBoxRect.IsInside(cursorNewPosition) {
		cursor.SetPosition(cursorNewPosition)
	}
}
