package main

import (
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
	DrawingSceneName = "scene/drawing/1"
	EntitySceneName  = "scene/entity/1"
	TopMenuName      = "menu/top/1"
	DrawingBoxName   = "entity/drawing-box/1"
	EntityBoxName    = "entity/entity-box/1"
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
// main private methods
// -----------------------------------------------------------------------------

func createDrawingBox(scene engine.IScene) {
	//drawingBox := engine.NewEntity(DrawingBoxName, TheDrawingBoxOrigin, TheDrawingBoxSize, &TheStyleWhiteOverBlack)
	//drawingBox.GetCanvas().WriteRectangleInCanvasAt(nil, nil, &TheStyleWhiteOverBlack, engine.CanvasRectSingleLine)
	drawingBox := widgets.NewBox(DrawingBoxName, TheDrawingBoxOrigin, TheDrawingBoxSize, &TheStyleWhiteOverBlack, widgets.BoxSingleLine)
	scene.AddEntity(drawingBox)

	cursor := NewCursor(api.NewPoint(1, theMenuBoxHeight+1))
	scene.AddEntity(cursor)

	theHandler = NewHandler()
	scene.AddEntity(theHandler)

	engine.GetEngine().GetSceneManager().UpdateFocus()
}

func newSpriter(ent engine.IEntity, args ...any) bool {
	tools.Logger.WithField("module", "spriter").WithField("function", "newSpriter").Tracef("%s %+v", ent.GetName(), args)
	if menu, ok := ent.(*widgets.Menu); ok {
		menu.DisableMenuItemForIndex(0)
		menu.EnableMenuItemForIndex(3)
		menu.EnableMenuItemsForLabel("Save")
		menu.SetSelectionToIndex(3)
		menu.Refresh()
	}
	if scene, ok := args[0].(engine.IScene); ok {
		createDrawingBox(scene)
	}

	return true
}

func load(ent engine.IEntity, args ...any) bool {
	scene := args[0].(engine.IScene)
	filename := args[1].(string)
	if theHandler == nil {
		newSpriter(ent, args...)
	}
	entities := engine.ImportEntitiesFromJSON(filename, TheDrawingBoxOrigin, nil)
	for _, ent := range entities {
		scene.AddEntity(ent)
	}
	return true
}

func save(ent engine.IEntity, args ...any) bool {
	if theHandler != nil {
		theHandler.SaveEntities()
	}
	return true
}

func exit(ent engine.IEntity, args ...any) bool {
	engine.GetEngine().End()
	return true
}

func main() {
	tools.Logger.WithField("module", "spriter").WithField("function", "main").Infof("Spriter App")
	drawingScene := engine.NewScene(DrawingSceneName, theCamera)

	topMenuItems := []*widgets.MenuItem{
		widgets.NewExtendedMenuItem("New", true, nil, newSpriter, []any{drawingScene}),
		widgets.NewExtendedMenuItem("Save", false, nil, save, nil),
		widgets.NewExtendedMenuItem("Load", true, nil, load, []any{drawingScene, "output_0_3.json"}),
		widgets.NewExtendedMenuItem("Exit", true, nil, exit, nil),
	}
	topMenu := widgets.NewTopMenu(TopMenuName, api.NewPoint(0, 0), api.NewSize(theMenuBoxWidth, theMenuBoxHeight), &TheStyleBlackOverWhite, topMenuItems, 0)
	topMenu.GetCanvas().WriteRectangleInCanvasAt(nil, nil, &TheStyleWhiteOverBlack, engine.CanvasRectSingleLine)
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
