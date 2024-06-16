package main

import (
	"strconv"

	"github.com/gdamore/tcell/v2"
	"github.com/jrecuero/thengine/pkg/api"
	"github.com/jrecuero/thengine/pkg/engine"
	"github.com/jrecuero/thengine/pkg/tools"
	"github.com/jrecuero/thengine/pkg/widgets"
)

func updateColor(ent engine.IEntity, args ...any) bool {
	var scene engine.IScene
	var ok bool
	if scene, ok = args[1].(engine.IScene); !ok {
		return false
	}

	camera := scene.GetCamera()
	colorScene := engine.NewScene("scene/color/1", camera)

	background := engine.NewEntity("entity/color/background/1", api.NewPoint(20, 5), api.NewSize(25, 7), &TheStyleBoldBlackOverGreen)
	background.GetCanvas().FillWithCell(engine.NewCell(&TheStyleWhiteOverBlack, ' '))
	background.GetCanvas().WriteRectangleInCanvasAt(nil, nil, &TheStyleBoldGreenOverBlack, engine.CanvasRectSingleLine)
	colorScene.AddEntity(background)

	c := scene.GetEntityByName(CursorName)
	if c == nil {
		return false
	}
	cursor, ok := c.(*Cursor)
	if !ok {
		return false
	}
	cursorFg, cursorBg, cursorAttrs := cursor.GetStyle().Decompose()

	fg := widgets.NewText("text/color/fg/1",
		api.NewPoint(22, 7), api.NewSize(10, 1), &TheStyleBoldGreenOverBlack, "Fg   : ")
	colorScene.AddEntity(fg)

	fgInput := widgets.NewTextInput("text-input/color/fg/1",
		api.NewPoint(32, 7), api.NewSize(10, 1), &TheStyleBoldBlackOverGreen, cursorFg.String())
	colorScene.AddEntity(fgInput)

	bg := widgets.NewText("text/color/bg/1", api.NewPoint(22, 8),
		api.NewSize(10, 1), &TheStyleBoldGreenOverBlack, "Bg   : ")
	colorScene.AddEntity(bg)

	bgInput := widgets.NewTextInput("text-input/color/bg/1",
		api.NewPoint(32, 8), api.NewSize(10, 1), &TheStyleBoldBlackOverGreen, cursorBg.String())
	colorScene.AddEntity(bgInput)

	attrs := widgets.NewText("text/color/attrs/1",
		api.NewPoint(22, 9), api.NewSize(10, 1), &TheStyleBoldGreenOverBlack, "attrs: ")
	colorScene.AddEntity(attrs)

	attrsInput := widgets.NewTextInput("text-input/color/attrs/1",
		api.NewPoint(32, 9), api.NewSize(10, 1), &TheStyleBoldBlackOverGreen, strconv.Itoa(int(cursorAttrs)))
	colorScene.AddEntity(attrsInput)

	input := &colorinput{
		Fg:    fgInput,
		Bg:    bgInput,
		Attrs: attrsInput,
	}

	accept := widgets.NewButton("button/color/accept/1",
		api.NewPoint(23, 11), api.NewSize(10, 1), &TheStyleBoldGreenOverBlack, "Accept")
	accept.SetWidgetCallback(acceptColor, colorScene, scene, input)
	colorScene.AddEntity(accept)

	cancel := widgets.NewButton("button/color/cancel/1",
		api.NewPoint(37, 11), api.NewSize(10, 1), &TheStyleBoldGreenOverBlack, "Cancel")
	cancel.SetWidgetCallback(cancelColor, colorScene, scene)
	colorScene.AddEntity(cancel)

	theEngine := engine.GetEngine()
	sceneManager := theEngine.GetSceneManager()

	sceneManager.RemoveSceneAsActive(scene)
	sceneManager.AddScene(colorScene)
	sceneManager.SetSceneAsActive(colorScene)
	sceneManager.SetSceneAsVisible(colorScene)
	sceneManager.UpdateFocus()

	return true
}

func acceptColor(entity engine.IEntity, args ...any) bool {
	tools.Logger.WithField("module", "colorhandler").
		WithField("function", "acceptColor").
		Tracef("%s args: %+#v", entity.GetName(), args)

	colorScene := args[0].(engine.IScene)
	drawingScene := args[1].(engine.IScene)
	input := args[2].(*colorinput)
	fg := input.Fg.GetInputText()
	bg := input.Bg.GetInputText()
	tmp := input.Attrs.GetInputText()
	attrs, _ := strconv.Atoi(tmp)
	style := tcell.StyleDefault.
		Foreground(tcell.GetColor(fg)).
		Background(tcell.GetColor(bg)).
		Attributes(tcell.AttrMask(attrs))

	c := drawingScene.GetEntityByName(CursorName)
	if c == nil {
		return false
	}
	cursor, ok := c.(*Cursor)
	if !ok {
		return false
	}
	cursor.SetStyle(&style)

	theEngine := engine.GetEngine()
	sceneManager := theEngine.GetSceneManager()
	sceneManager.RemoveScene(colorScene)
	sceneManager.SetSceneAsActive(drawingScene)
	sceneManager.SetSceneAsVisible(drawingScene)
	sceneManager.UpdateFocus()
	return true
}

func cancelColor(entity engine.IEntity, args ...any) bool {
	tools.Logger.WithField("module", "colorhandler").
		WithField("function", "cancelColor").
		Tracef("%s args: %+#v", entity.GetName(), args)

	colorScene := args[0].(engine.IScene)
	drawingScene := args[1].(engine.IScene)

	theEngine := engine.GetEngine()
	sceneManager := theEngine.GetSceneManager()
	sceneManager.RemoveScene(colorScene)
	sceneManager.SetSceneAsActive(drawingScene)
	sceneManager.SetSceneAsVisible(drawingScene)
	sceneManager.UpdateFocus()
	return true
}
