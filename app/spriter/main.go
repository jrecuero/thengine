package main

import (
	"strconv"

	"github.com/gdamore/tcell/v2"
	"github.com/jrecuero/thengine/pkg/api"
	"github.com/jrecuero/thengine/pkg/engine"
	"github.com/jrecuero/thengine/pkg/tools"
	"github.com/jrecuero/thengine/pkg/widgets"
)

// -----------------------------------------------------------------------------
// main private constants
// -----------------------------------------------------------------------------
const (
	theCameraWidth      = 90
	theCameraHeight     = 40
	theMenuBoxWidth     = theCameraWidth
	theMenuBoxHeight    = 3
	theDrawingBoxWidth  = theCameraWidth
	theDrawingBoxHeight = 27
	theEntityBoxWidth   = theCameraWidth
	theEntityBoxHeight  = 10
	theFPS              = 60.0
)

// -----------------------------------------------------------------------------
// main public constants
// -----------------------------------------------------------------------------
const (
	DrawingSceneName  = "scene/drawing/1"
	EntitySceneName   = "scene/entity/1"
	TopMenuName       = "menu/top/1"
	DrawingBoxName    = "entity/drawing-box/1"
	EntityBoxName     = "entity/entity-box/1"
	CursorPosTextName = "text/cursor-position/1"
)

// -----------------------------------------------------------------------------
// main private variables
// -----------------------------------------------------------------------------

var (
	theCamera = engine.NewCamera(api.NewPoint(0, 0), api.NewSize(theCameraWidth, theCameraHeight))
	theEngine = engine.GetEngine()
)

// -----------------------------------------------------------------------------
// main public variables
// -----------------------------------------------------------------------------

var (
	TheStyleWhiteOverBlack         = tcell.StyleDefault.Foreground(tcell.ColorWhite).Background(tcell.ColorBlack)
	TheStyleBlinkingWhiteOverBlack = tcell.StyleDefault.Foreground(tcell.ColorWhite).Background(tcell.ColorBlack).Attributes(tcell.AttrBlink)
	TheStyleBlackOverWhite         = tcell.StyleDefault.Foreground(tcell.ColorBlack).Background(tcell.ColorWhite)
	TheStyleBlinkingBlackOverWhite = tcell.StyleDefault.Foreground(tcell.ColorBlack).Background(tcell.ColorWhite).Attributes(tcell.AttrBlink)
	TheStyleBoldBlackOverWhite     = tcell.StyleDefault.Foreground(tcell.ColorBlack).Background(tcell.ColorWhite).Attributes(tcell.AttrBold)
	TheStyleBoldBlackOverGreen     = tcell.StyleDefault.Foreground(tcell.ColorBlack).Background(tcell.ColorGreen).Attributes(tcell.AttrBold)
	TheStyleBoldGreenOverBlack     = tcell.StyleDefault.Foreground(tcell.ColorGreen).Background(tcell.ColorBlack).Attributes(tcell.AttrBold)
	TheDrawingBoxOrigin            = api.NewPoint(0, theMenuBoxHeight)
	TheDrawingBoxSize              = api.NewSize(theDrawingBoxWidth, theDrawingBoxHeight)
	TheDrawingBoxRect              = api.NewRect(TheDrawingBoxOrigin, TheDrawingBoxSize)
	TheEntityBoxOrigin             = api.NewPoint(0, theMenuBoxHeight+theDrawingBoxHeight)
	TheEntityBoxSize               = api.NewSize(theEntityBoxWidth, theEntityBoxHeight)
)

// -----------------------------------------------------------------------------
// main private structs
// -----------------------------------------------------------------------------

type topmenu struct {
	*widgets.Menu
	LoseFocus bool
}

func (t *topmenu) EndTick(engine.IScene) {
	if t.LoseFocus {
		tools.Logger.WithField("module", "spriter").
			WithField("struct", "topmenu").
			WithField("function", "EndTick").
			Tracef("%s %t", t.GetName(), t.LoseFocus)
		engine.GetEngine().GetSceneManager().UpdateFocus()
		t.LoseFocus = false
	}
}

type builtin struct {
}

func (b *builtin) GetClassFromString(class string) engine.IEntity {
	if class == "Sprite" {
		return widgets.NewSprite("", nil, nil)
	}
	return engine.NewEmptyEntity()
}

type colorinput struct {
	Fg    *widgets.TextInput
	Bg    *widgets.TextInput
	Attrs *widgets.TextInput
}

// -----------------------------------------------------------------------------
// main private methods
// -----------------------------------------------------------------------------

func createDrawingBox(scene engine.IScene) {
	//drawingBox := engine.NewEntity(DrawingBoxName, TheDrawingBoxOrigin, TheDrawingBoxSize, &TheStyleWhiteOverBlack)
	//drawingBox.GetCanvas().WriteRectangleInCanvasAt(nil, nil, &TheStyleWhiteOverBlack, engine.CanvasRectSingleLine)
	drawingBox := widgets.NewBox(DrawingBoxName, TheDrawingBoxOrigin, TheDrawingBoxSize, &TheStyleWhiteOverBlack, widgets.BoxSingleLine)
	scene.AddEntity(drawingBox)

	cursor := NewCursor(api.NewPoint(1, theMenuBoxHeight+1))
	cursor.SetZLevel(1)
	scene.AddEntity(cursor)

	cursorPosText := widgets.NewText(CursorPosTextName, api.NewPoint(80, 29), api.NewSize(10, 1), &TheStyleWhiteOverBlack, "[1,1]")
	scene.AddEntity(cursorPosText)

	theHandler = NewHandler()
	scene.AddEntity(theHandler)

	engine.GetEngine().GetSceneManager().UpdateFocus()
}

func newSpriter(ent engine.IEntity, args ...any) bool {
	tools.Logger.WithField("module", "spriter").
		WithField("function", "newSpriter").
		Tracef("%s %+v", ent.GetName(), args)
	if menu, ok := ent.(*widgets.Menu); ok {
		menu.DisableMenuItemForIndex(0)
		menu.EnableMenuItemForIndex(3, 5)
		menu.EnableMenuItemsForLabel("Save")
		menu.SetSelectionToIndex(6)
		menu.Refresh()
	}
	if scene, ok := args[0].(engine.IScene); ok {
		createDrawingBox(scene)
	}

	return true
}

func load(entity engine.IEntity, args ...any) bool {
	scene := args[0].(engine.IScene)
	filename := args[1].(string)
	if theHandler == nil {
		newSpriter(entity, args...)
	}
	entities := engine.ImportEntitiesFromJSON(filename, TheDrawingBoxOrigin, &builtin{})
	for _, entity := range entities {
		tools.Logger.WithField("module", "import").
			WithField("function", "ImportEntitiesToJSON").
			Debugf("importing entity %+#v", entity.GetPosition())
		theHandler.entities = append(theHandler.entities, entity)
		scene.AddEntity(entity)
	}
	return true
}

func save(ent engine.IEntity, args ...any) bool {
	if theHandler != nil {
		theHandler.SaveEntities()
	}
	return true
}

func createSprite(ent engine.IEntity, args ...any) bool {
	var scene engine.IScene
	var menu *topmenu
	var menuItem *widgets.MenuItem
	var ok bool

	if scene, ok = args[0].(engine.IScene); !ok {
		return false
	}
	if menu, ok = args[1].(*topmenu); !ok {
		return false
	}
	if menuItem, ok = args[2].(*widgets.MenuItem); !ok {
		return false
	}
	tools.Logger.WithField("module", "spriter").
		WithField("function", "createSprite").
		Tracef("%s %+v %+#v %+#v", ent.GetName(), scene, menu, menuItem)

	menu.DisableMenuItemForIndex(3)
	menu.EnableMenuItemForIndex(4)
	menu.SetSelectionToIndex(4)
	menu.LoseFocus = true
	menu.Refresh()
	theHandler.CreateSprite(scene)

	return true
}

func saveSprite(ent engine.IEntity, args ...any) bool {
	var scene engine.IScene
	var menu *topmenu
	var menuItem *widgets.MenuItem
	var ok bool

	if scene, ok = args[0].(engine.IScene); !ok {
		return false
	}
	if menu, ok = args[1].(*topmenu); !ok {
		return false
	}
	if menuItem, ok = args[2].(*widgets.MenuItem); !ok {
		return false
	}
	tools.Logger.WithField("module", "spriter").
		WithField("function", "saveSprite").
		Tracef("%s %+v %+#v %+#v", ent.GetName(), scene, menu, menuItem)

	menu.DisableMenuItemForIndex(4)
	menu.EnableMenuItemForIndex(3)
	menu.SetSelectionToIndex(3)
	menu.Refresh()
	theHandler.SaveSprite()

	return true
}

func color(ent engine.IEntity, args ...any) bool {
	var scene engine.IScene
	var ok bool
	if scene, ok = args[0].(engine.IScene); !ok {
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
	tools.Logger.WithField("module", "spriter").
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
	tools.Logger.WithField("module", "spriter").
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

func exit(ent engine.IEntity, args ...any) bool {
	engine.GetEngine().End()
	return true
}

func main() {
	tools.Logger.WithField("module", "spriter").WithField("function", "main").Infof("Spriter App")
	drawingScene := engine.NewScene(DrawingSceneName, theCamera)

	createSpriteMenuItem := widgets.NewExtendedMenuItem("New SPR", false, nil, nil, nil)
	saveSpriteMenuItem := widgets.NewExtendedMenuItem("Save SPR", false, nil, nil, nil)
	topMenuItems := []*widgets.MenuItem{
		widgets.NewExtendedMenuItem("New", true, nil, newSpriter, []any{drawingScene}),
		widgets.NewExtendedMenuItem("Save", false, nil, save, nil),
		widgets.NewExtendedMenuItem("Load", true, nil, load, []any{drawingScene, "output_0_3.json"}),
		createSpriteMenuItem,
		saveSpriteMenuItem,
		widgets.NewExtendedMenuItem("Color", false, nil, color, []any{drawingScene}),
		widgets.NewExtendedMenuItem("Exit", true, nil, exit, nil),
	}
	topMenu := &topmenu{
		Menu: widgets.NewTopMenu(TopMenuName, api.NewPoint(0, 0), api.NewSize(theMenuBoxWidth, theMenuBoxHeight), &TheStyleWhiteOverBlack, topMenuItems, 0),
	}
	createSpriteMenuItem.SetCallback(createSprite, []any{drawingScene, topMenu, createSpriteMenuItem})
	saveSpriteMenuItem.SetCallback(saveSprite, []any{drawingScene, topMenu, saveSpriteMenuItem})
	drawingScene.AddEntity(topMenu)

	theEngine.InitResources()
	theEngine.GetSceneManager().AddScene(drawingScene)
	theEngine.GetSceneManager().SetSceneAsActive(drawingScene)
	theEngine.GetSceneManager().SetSceneAsVisible(drawingScene)
	theEngine.GetSceneManager().UpdateFocus()
	theEngine.Init()
	theEngine.Start()
	theEngine.Run(theFPS)
}
