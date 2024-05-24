package main

import (
	"github.com/jrecuero/thengine/pkg/api"
	"github.com/jrecuero/thengine/pkg/engine"
	"github.com/jrecuero/thengine/pkg/widgets"
)

const (
	EntityHandlerName = "entity/entity-handler/1"
)

type EntityHandler struct {
	*engine.Entity
	cursor       *Cursor
	drawingScene engine.IScene
	entityScene  engine.IScene
}

func NewEntityHandler(drawingScene engine.IScene, cursor *Cursor) *EntityHandler {
	handler := &EntityHandler{
		Entity:       engine.NewNamedEntity(EntityHandlerName),
		cursor:       cursor,
		drawingScene: drawingScene,
		entityScene:  nil,
	}
	handler.SetFocusType(engine.SingleFocus)
	handler.SetFocusEnable(true)
	handler.createEntityScene()
	return handler
}

func (h *EntityHandler) createEntityScene() {
	if h.drawingScene == nil {
		return
	}
	camera := h.drawingScene.GetCamera()
	entityScene := engine.NewScene(EntitySceneName, camera)

	entityBox := engine.NewEntity(EntityBoxName, TheEntityBoxOrigin, TheEntityBoxSize, &TheStyleWhiteOverBlack)
	entityBox.GetCanvas().WriteRectangleInCanvasAt(nil, nil, &TheStyleWhiteOverBlack, engine.CanvasRectSingleLine)
	entityScene.AddEntity(entityBox)

	var ClassTextName = "text/class-name/1"
	//var PositionTextName = "text/position/1"
	classText := widgets.NewText(ClassTextName, api.NewPoint(TheEntityBoxOrigin.X+1, TheEntityBoxOrigin.Y+1), api.NewSize(10, 1), &TheStyleWhiteOverBlack, "Class Name: ")
	entityScene.AddEntity(classText)

	entityScene.AddEntity(h)

	theEngine := engine.GetEngine()
	sceneManager := theEngine.GetSceneManager()

	sceneManager.RemoveSceneAsActive(h.drawingScene)

	sceneManager.AddScene(entityScene)
	sceneManager.SetSceneAsActive(entityScene)
	sceneManager.SetSceneAsVisible(entityScene)

	h.entityScene = entityScene
}
