package main

import (
	"github.com/jrecuero/thengine/pkg/api"
	"github.com/jrecuero/thengine/pkg/engine"
	"github.com/jrecuero/thengine/pkg/widgets"
)

func save(entity engine.IEntity, args ...any) bool {
	menuItem := args[0].(*widgets.MenuItem)
	scene := args[1].(engine.IScene)
	pos := menuItem.GetPosition()
	options := widgets.NewListBox("menu/list-box/save/1",
		pos, api.NewSize(10, 4), &TheStyleWhiteOverBlack,
		[]string{"json", "code"}, 0)
	options.SetWidgetCallback(saveEntities, scene)
	scene.AddEntity(options)

	engine.GetEngine().GetFocusManager().AcquireFocusToEntity(options)
	engine.GetEngine().GetFocusManager().SetLocked(true)

	return true
}

func saveEntities(entity engine.IEntity, args ...any) bool {
	scene := args[0].(engine.IScene)
	listbox := entity.(*widgets.ListBox)
	selection := listbox.GetSelectionIndex()
	switch selection {
	case 0:
		saveEntitiesToJSON()
	case 1:
		saveEntitiesToCode()
	}

	engine.GetEngine().GetFocusManager().SetLocked(false)
	scene.RemoveEntity(listbox)
	return true
}

func saveEntitiesToJSON() {
	if theHandler != nil {
		theHandler.SaveEntitiesToJSON()
	}
}

func saveEntitiesToCode() {
	if theHandler != nil {
		theHandler.SaveEntitiesToCode()
	}
}
