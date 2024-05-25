package main

import (
	"fmt"

	"github.com/jrecuero/thengine/pkg/api"
	"github.com/jrecuero/thengine/pkg/engine"
	"github.com/jrecuero/thengine/pkg/tools"
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

type EntityTextInput struct {
	ClassName *widgets.TextInput
	Position  *widgets.TextInput
	Size      *widgets.TextInput
	Style     *widgets.TextInput
	Rune      *widgets.TextInput
}

func NewEntityHandler(drawingScene engine.IScene, cursor *Cursor) *EntityHandler {
	handler := &EntityHandler{
		Entity:       engine.NewHandler(EntityHandlerName),
		cursor:       cursor,
		drawingScene: drawingScene,
		entityScene:  nil,
	}
	//handler.SetFocusType(engine.SingleFocus)
	handler.SetFocusEnable(false)
	handler.createEntityScene()
	return handler
}

func (h *EntityHandler) acceptCallback(entity engine.IEntity, args ...any) bool {
	textInput := args[0].(*EntityTextInput)
	tools.Logger.WithField("module", "entityhandler").WithField("method", "acceptCallback").Debugf("class-name %s", textInput.ClassName.GetInputText())
	tools.Logger.WithField("module", "entityhandler").WithField("method", "acceptCallback").Debugf("position %s", textInput.Position.GetInputText())
	tools.Logger.WithField("module", "entityhandler").WithField("method", "acceptCallback").Debugf("size %s", textInput.Size.GetInputText())
	tools.Logger.WithField("module", "entityhandler").WithField("method", "acceptCallback").Debugf("style %s", textInput.Style.GetInputText())
	tools.Logger.WithField("module", "entityhandler").WithField("method", "acceptCallback").Debugf("rune %s", textInput.Rune.GetInputText())

	theEngine := engine.GetEngine()
	sceneManager := theEngine.GetSceneManager()
	sceneManager.RemoveSceneAsActive(h.entityScene)
	sceneManager.RemoveSceneAsVisible(h.entityScene)
	sceneManager.RemoveScene(h.entityScene)
	sceneManager.SetSceneAsActive(h.drawingScene)
	sceneManager.SetSceneAsVisible(h.drawingScene)

	return true
}

func (h *EntityHandler) createEntityScene() {
	if h.drawingScene == nil {
		return
	}
	camera := h.drawingScene.GetCamera()
	entityScene := engine.NewScene(EntitySceneName, camera)

	//entityBox := engine.NewEntity(EntityBoxName, TheEntityBoxOrigin, TheEntityBoxSize, &TheStyleWhiteOverBlack)
	//entityBox.GetCanvas().WriteRectangleInCanvasAt(nil, nil, &TheStyleWhiteOverBlack, engine.CanvasRectSingleLine)
	entityBox := widgets.NewBox(EntityBoxName, TheEntityBoxOrigin, TheEntityBoxSize, &TheStyleWhiteOverBlack, widgets.BoxSingleLine)
	entityScene.AddEntity(entityBox)

	var ClassTextName = "text/class-name/1"
	var ClassTextInputName = "text-input/class-name/1"
	var PositionTextName = "text/position/1"
	var PositionTextInputName = "text-input/position/1"
	var SizeTextName = "text/size/1"
	var SizeTextInputName = "text-input/size/1"
	var StyleTextName = "text/style/1"
	var StyleTextInputName = "text-input/style/1"
	var RuneTextName = "text/rune/1"
	var RuneTextInputName = "text-input/rune/1"
	var AcceptButtonName = "button/accept/1"
	classText := widgets.NewText(ClassTextName,
		api.NewPoint(TheEntityBoxOrigin.X+1, TheEntityBoxOrigin.Y+1),
		api.NewSize(12, 1),
		&TheStyleWhiteOverBlack, "Class Name: ")
	entityScene.AddEntity(classText)

	classTextInput := widgets.NewTextInput(ClassTextInputName,
		api.NewPoint(TheEntityBoxOrigin.X+13, TheEntityBoxOrigin.Y+1),
		api.NewSize(20, 1),
		&TheStyleBlackOverWhite, "")
	entityScene.AddEntity(classTextInput)

	positionText := widgets.NewText(PositionTextName,
		api.NewPoint(TheEntityBoxOrigin.X+1, TheEntityBoxOrigin.Y+2),
		api.NewSize(12, 1),
		&TheStyleWhiteOverBlack, "Position  : ")
	entityScene.AddEntity(positionText)

	positionTextInput := widgets.NewTextInput(PositionTextInputName,
		api.NewPoint(TheEntityBoxOrigin.X+13, TheEntityBoxOrigin.Y+2),
		api.NewSize(20, 1),
		&TheStyleBlackOverWhite,
		fmt.Sprintf("%d,%d", h.cursor.GetPosition().X, h.cursor.GetPosition().Y))
	entityScene.AddEntity(positionTextInput)

	sizeText := widgets.NewText(SizeTextName,
		api.NewPoint(TheEntityBoxOrigin.X+1, TheEntityBoxOrigin.Y+3),
		api.NewSize(12, 1),
		&TheStyleWhiteOverBlack, "Size      : ")
	entityScene.AddEntity(sizeText)

	sizeTextInput := widgets.NewTextInput(SizeTextInputName,
		api.NewPoint(TheEntityBoxOrigin.X+13, TheEntityBoxOrigin.Y+3),
		api.NewSize(20, 1),
		&TheStyleBlackOverWhite,
		fmt.Sprintf("%d,%d", h.cursor.GetSize().W, h.cursor.GetSize().H))
	entityScene.AddEntity(sizeTextInput)

	styleText := widgets.NewText(StyleTextName,
		api.NewPoint(TheEntityBoxOrigin.X+1, TheEntityBoxOrigin.Y+4),
		api.NewSize(12, 1),
		&TheStyleWhiteOverBlack, "Style     : ")
	entityScene.AddEntity(styleText)

	fg, bg, attrs := h.cursor.GetStyle().Decompose()
	styleTextInput := widgets.NewTextInput(StyleTextInputName,
		api.NewPoint(TheEntityBoxOrigin.X+13, TheEntityBoxOrigin.Y+4),
		api.NewSize(20, 1),
		&TheStyleBlackOverWhite,
		fmt.Sprintf("%s,%s,%d", fg.String(), bg.String(), attrs))
	entityScene.AddEntity(styleTextInput)

	runeText := widgets.NewText(RuneTextName,
		api.NewPoint(TheEntityBoxOrigin.X+1, TheEntityBoxOrigin.Y+5),
		api.NewSize(12, 1),
		&TheStyleWhiteOverBlack, "Rune      : ")
	entityScene.AddEntity(runeText)

	runeTextInput := widgets.NewTextInput(RuneTextInputName,
		api.NewPoint(TheEntityBoxOrigin.X+13, TheEntityBoxOrigin.Y+5),
		api.NewSize(1, 1),
		&TheStyleBlackOverWhite,
		fmt.Sprintf("%s", string(h.cursor.GetCanvas().GetCellAt(nil).Rune)))
	entityScene.AddEntity(runeTextInput)

	newEntityTextInput := &EntityTextInput{
		ClassName: classTextInput,
		Position:  positionTextInput,
		Size:      sizeTextInput,
		Style:     styleTextInput,
		Rune:      runeTextInput,
	}
	acceptButton := widgets.NewButton(AcceptButtonName,
		api.NewPoint(TheEntityBoxOrigin.X+1, TheEntityBoxOrigin.Y+6),
		api.NewSize(20, 1),
		&TheStyleBoldGreenOverBlack, "Accept")
	acceptButton.SetWidgetCallback(h.acceptCallback, newEntityTextInput)
	entityScene.AddEntity(acceptButton)

	entityScene.AddEntity(h)

	theEngine := engine.GetEngine()
	sceneManager := theEngine.GetSceneManager()

	sceneManager.RemoveSceneAsActive(h.drawingScene)
	//sceneManager.RemoveSceneAsVisible(h.drawingScene)

	sceneManager.AddScene(entityScene)
	sceneManager.SetSceneAsActive(entityScene)
	sceneManager.SetSceneAsVisible(entityScene)
	sceneManager.UpdateFocus()

	h.entityScene = entityScene
}
