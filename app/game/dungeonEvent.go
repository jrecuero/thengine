package main

import (
	"github.com/gdamore/tcell/v2"
	"github.com/jrecuero/thengine/app/game/dad/rules"
	"github.com/jrecuero/thengine/pkg/api"
	"github.com/jrecuero/thengine/pkg/constants"
	"github.com/jrecuero/thengine/pkg/engine"
	"github.com/jrecuero/thengine/pkg/widgets"
)

// -----------------------------------------------------------------------------
//
// IDungeonEvent
//
// -----------------------------------------------------------------------------

type IDungeonEvent interface {
	IsDungeonEvent() bool
}

// -----------------------------------------------------------------------------
//
// OpenDoorAction
//
// -----------------------------------------------------------------------------

type OpenDoorAction struct {
	*rules.Action
}

func NewOpenDoorAction() *OpenDoorAction {
	return &OpenDoorAction{}
}

func (a *OpenDoorAction) Execute(e rules.IEvent, args ...any) (rules.UniqueActionId, error) {
	parentScene := args[0].(engine.IScene)
	door := args[1].(*DoorEvent)
	modalDialog := widgets.NewModalDialog(parentScene)
	question := widgets.NewText("text/open-door/question/1", nil, nil, &constants.WhiteOverBlack, "Open Door?")
	buttonYes := widgets.NewButton("button/ope-door/yes/1", nil, nil, &constants.WhiteOverBlack, "YES")
	buttonNo := widgets.NewButton("button/ope-door/no/1", nil, nil, &constants.WhiteOverBlack, "NO")

	yesCallback := func(entity engine.IEntity, args ...any) bool {
		modalDialog.Close()
		door.SetSolid(false)
		return true
	}
	buttonYes.SetWidgetCallback(yesCallback)

	noCallback := func(entity engine.IEntity, args ...any) bool {
		modalDialog.Close()
		return true
	}
	buttonNo.SetWidgetCallback(noCallback)

	dialog := widgets.NewDialog("dialog/open-door/1", api.NewPoint(1, 1), api.NewSize(20, 5), &constants.WhiteOverBlack,
		modalDialog.GetDialogScene(),
		[]*widgets.Text{question},
		nil,
		[]*widgets.Button{buttonYes, buttonNo})

	modalDialog.Open(dialog)
	return rules.UAIdEnd, nil
}

var _ rules.IAction = (*OpenDoorAction)(nil)

// -----------------------------------------------------------------------------
//
// DoorEvent
//
// -----------------------------------------------------------------------------

type DoorEvent struct {
	*engine.Entity
	Event rules.IEvent
}

func NewDoorEvent(name string, pos *api.Point, size *api.Size, style *tcell.Style, parentScene engine.IScene) *DoorEvent {
	e := &DoorEvent{
		Entity: engine.NewEntity(name, pos, size, style),
		Event:  rules.NewEvent(name, nil),
	}
	e.SetSolid(true)
	cell := engine.NewCell(style, 'D')
	e.GetCanvas().SetCellAt(nil, cell)
	e.Event.SetActions([]rules.IAction{NewOpenDoorAction()})
	return e
}

func (e *DoorEvent) IsDungeonEvent() bool {
	return true
}

var _ IDungeonEvent = (*DoorEvent)(nil)
