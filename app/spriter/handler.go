package main

import (
	"github.com/gdamore/tcell/v2"
	"github.com/jrecuero/thengine/pkg/api"
	"github.com/jrecuero/thengine/pkg/engine"
	"github.com/jrecuero/thengine/pkg/tools"
)

const (
	HandlerName = "entity/handler/1"
)

var (
	theHandler *Handler
)

type Handler struct {
	*engine.Entity
	entities []engine.IEntity
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

func (h *Handler) entityHandlerResponse(scene engine.IScene) func(engine.IEntity) {
	return func(respEntity engine.IEntity) {
		tools.Logger.WithField("module", "handler").
			WithField("method", "entityHandlerResponse").
			Debugf("%s %s %s", scene.GetName(), respEntity.GetClassName(), respEntity.GetName())
		scene.AddEntity(respEntity)
		h.entities = append(h.entities, respEntity)
	}
}

func (h *Handler) SaveEntities() {
	for _, ent := range h.entities {
		tools.Logger.WithField("module", "handler").
			WithField("method", "SaveEntites").
			Debugf("saving %+#v", ent)
	}
	if err := engine.ExportEntitiesToJSON("output.json", h.entities, nil); err != nil {
		tools.Logger.WithField("module", "handler").
			WithField("method", "SaveEntites").
			Errorf("error %s", err.Error())
	}
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
			NewEntityHandler(scene, cursor, h.entityHandlerResponse(scene))
		case tcell.KeyRune:
			h.updateCursorRune(cursor, ev.Rune())
		}
	}
	if TheDrawingBoxRect.IsInside(cursorNewPosition) {
		cursor.SetPosition(cursorNewPosition)
	}
}
