package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/gdamore/tcell/v2"
	"github.com/jrecuero/thengine/pkg/api"
	"github.com/jrecuero/thengine/pkg/engine"
	"github.com/jrecuero/thengine/pkg/tools"
	"github.com/jrecuero/thengine/pkg/widgets"
)

const (
	EntityHandlerName = "handler/entity-handler/1"
)

type EntityHandler struct {
	*engine.Entity
	cursor       *Cursor
	drawingScene engine.IScene
	entityScene  engine.IScene
	response     func(engine.IEntity)
}

type EntityTextInput struct {
	ClassName *widgets.TextInput
	Name      *widgets.TextInput
	Position  *widgets.TextInput
	Size      *widgets.TextInput
	Style     *widgets.TextInput
	Rune      *widgets.TextInput
}

func NewEntityHandler(drawingScene engine.IScene, cursor *Cursor, response func(engine.IEntity)) *EntityHandler {
	handler := &EntityHandler{
		Entity:       engine.NewHandler(EntityHandlerName),
		cursor:       cursor,
		drawingScene: drawingScene,
		entityScene:  nil,
		response:     response,
	}
	//handler.SetFocusType(engine.SingleFocus)
	handler.SetFocusEnable(false)
	handler.createEntityScene()
	return handler
}

func (h *EntityHandler) processEntityTextInput(entityTextInput *EntityTextInput) engine.IEntity {
	result := engine.NewEmptyEntity()
	result.SetClassName(entityTextInput.ClassName.GetInputText())
	result.SetName(entityTextInput.Name.GetInputText())

	// Process position.
	posTmp := entityTextInput.Position.GetInputText()
	posSlice := strings.Split(posTmp, ",")
	x, _ := strconv.Atoi(posSlice[0])
	y, _ := strconv.Atoi(posSlice[1])
	position := api.NewPoint(x, y)
	result.SetPosition(position)

	// Process size.
	sizeTmp := entityTextInput.Size.GetInputText()
	sizeSlice := strings.Split(sizeTmp, ",")
	width, _ := strconv.Atoi(sizeSlice[0])
	height, _ := strconv.Atoi(sizeSlice[1])
	size := api.NewSize(width, height)
	result.SetSize(size)

	// Process style
	styleTmp := entityTextInput.Style.GetInputText()
	styleSlice := strings.Split(styleTmp, ",")
	fg := styleSlice[0]
	bg := styleSlice[1]
	attrs, _ := strconv.Atoi(styleSlice[2])
	style := tcell.StyleDefault.
		Foreground(tcell.GetColor(fg)).
		Background(tcell.GetColor(bg)).
		Attributes(tcell.AttrMask(attrs))

	// Process cell
	canvas := engine.NewCanvas(size)
	ch := entityTextInput.Rune.GetInputText()
	cell := engine.NewCell(&style, rune(ch[0]))
	if size.IsOneSize() {
		canvas.SetCellAt(nil, cell)
	} else {
		canvas.FillWithCell(cell)
	}
	result.SetCanvas(canvas)
	result.SetStyle(&style)

	return result
}

func (h *EntityHandler) acceptCallback(entity engine.IEntity, args ...any) bool {
	textInput := args[0].(*EntityTextInput)
	tools.Logger.WithField("module", "entityhandler").
		WithField("method", "acceptCallback").
		Debugf("create new entity")
	result := h.processEntityTextInput(textInput)

	theEngine := engine.GetEngine()
	sceneManager := theEngine.GetSceneManager()
	//sceneManager.RemoveSceneAsActive(h.entityScene)
	//sceneManager.RemoveSceneAsVisible(h.entityScene)
	sceneManager.RemoveScene(h.entityScene)
	sceneManager.SetSceneAsActive(h.drawingScene)
	sceneManager.SetSceneAsVisible(h.drawingScene)

	if h.response != nil {
		h.response(result)
		h.drawingScene = nil
		h.entityScene = nil
	}

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
	var NameTextName = "text/name/1"
	var NameTextInputName = "text-input/name/1"
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

	nameText := widgets.NewText(NameTextName,
		api.NewPoint(TheEntityBoxOrigin.X+1, TheEntityBoxOrigin.Y+2),
		api.NewSize(12, 1),
		&TheStyleWhiteOverBlack, "Name      : ")
	entityScene.AddEntity(nameText)

	nameTextInput := widgets.NewTextInput(NameTextInputName,
		api.NewPoint(TheEntityBoxOrigin.X+13, TheEntityBoxOrigin.Y+2),
		api.NewSize(20, 1),
		&TheStyleBlackOverWhite, "")
	entityScene.AddEntity(nameTextInput)

	positionText := widgets.NewText(PositionTextName,
		api.NewPoint(TheEntityBoxOrigin.X+1, TheEntityBoxOrigin.Y+3),
		api.NewSize(12, 1),
		&TheStyleWhiteOverBlack, "Position  : ")
	entityScene.AddEntity(positionText)

	positionTextInput := widgets.NewTextInput(PositionTextInputName,
		api.NewPoint(TheEntityBoxOrigin.X+13, TheEntityBoxOrigin.Y+3),
		api.NewSize(20, 1),
		&TheStyleBlackOverWhite,
		fmt.Sprintf("%d,%d", h.cursor.GetPosition().X, h.cursor.GetPosition().Y))
	entityScene.AddEntity(positionTextInput)

	sizeText := widgets.NewText(SizeTextName,
		api.NewPoint(TheEntityBoxOrigin.X+1, TheEntityBoxOrigin.Y+4),
		api.NewSize(12, 1),
		&TheStyleWhiteOverBlack, "Size      : ")
	entityScene.AddEntity(sizeText)

	sizeTextInput := widgets.NewTextInput(SizeTextInputName,
		api.NewPoint(TheEntityBoxOrigin.X+13, TheEntityBoxOrigin.Y+4),
		api.NewSize(20, 1),
		&TheStyleBlackOverWhite,
		fmt.Sprintf("%d,%d", h.cursor.GetSize().W, h.cursor.GetSize().H))
	entityScene.AddEntity(sizeTextInput)

	styleText := widgets.NewText(StyleTextName,
		api.NewPoint(TheEntityBoxOrigin.X+1, TheEntityBoxOrigin.Y+5),
		api.NewSize(12, 1),
		&TheStyleWhiteOverBlack, "Style     : ")
	entityScene.AddEntity(styleText)

	fg, bg, attrs := h.cursor.GetStyle().Decompose()
	styleTextInput := widgets.NewTextInput(StyleTextInputName,
		api.NewPoint(TheEntityBoxOrigin.X+13, TheEntityBoxOrigin.Y+5),
		api.NewSize(20, 1),
		&TheStyleBlackOverWhite,
		fmt.Sprintf("%s,%s,%d", fg.String(), bg.String(), attrs))
	entityScene.AddEntity(styleTextInput)

	runeText := widgets.NewText(RuneTextName,
		api.NewPoint(TheEntityBoxOrigin.X+1, TheEntityBoxOrigin.Y+6),
		api.NewSize(12, 1),
		&TheStyleWhiteOverBlack, "Rune      : ")
	entityScene.AddEntity(runeText)

	runeTextInput := widgets.NewTextInput(RuneTextInputName,
		api.NewPoint(TheEntityBoxOrigin.X+13, TheEntityBoxOrigin.Y+6),
		api.NewSize(1, 1),
		&TheStyleBlackOverWhite,
		fmt.Sprintf("%s", string(h.cursor.GetCanvas().GetCellAt(nil).GetRune())))
	entityScene.AddEntity(runeTextInput)

	newEntityTextInput := &EntityTextInput{
		ClassName: classTextInput,
		Name:      nameTextInput,
		Position:  positionTextInput,
		Size:      sizeTextInput,
		Style:     styleTextInput,
		Rune:      runeTextInput,
	}
	acceptButton := widgets.NewButton(AcceptButtonName,
		api.NewPoint(TheEntityBoxOrigin.X+1, TheEntityBoxOrigin.Y+7),
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
