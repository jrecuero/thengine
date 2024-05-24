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
	theWidth  = 90
	theHeight = 30
	theFPS    = 60.0
)

// -----------------------------------------------------------------------------
// main public constants
// -----------------------------------------------------------------------------
const (
	TopMenuName = "menu/top/1"
)

// -----------------------------------------------------------------------------
// main private variables
// -----------------------------------------------------------------------------

var (
	theCamera              = engine.NewCamera(api.NewPoint(0, 0), api.NewSize(theWidth, theHeight))
	theEngine              = engine.GetEngine()
	theStyleWhiteOverBlack = tcell.StyleDefault.Foreground(tcell.ColorWhite).Background(tcell.ColorBlack)
	theStyleBlackOverWhite = tcell.StyleDefault.Foreground(tcell.ColorBlack).Background(tcell.ColorWhite)
)

// -----------------------------------------------------------------------------
// main private methods
// -----------------------------------------------------------------------------

func main() {
	tools.Logger.WithField("module", "spriter").WithField("function", "main").Infof("Spriter App")
	drawingScene := engine.NewScene("scene/drawing/1", theCamera)

	topMenuItems := []*widgets.MenuItem{
		widgets.NewMenuItem("New"),
		widgets.NewExtendedMenuItem("Save", false, nil, nil, nil),
		widgets.NewMenuItem("Load"),
		widgets.NewMenuItem("New Rune"),
	}
	topMenu := widgets.NewTopMenu(TopMenuName, api.NewPoint(0, 0), api.NewSize(theWidth, 3), &theStyleBlackOverWhite, topMenuItems, 0)
	topMenu.GetCanvas().WriteRectangleInCanvasAt(nil, nil, &theStyleWhiteOverBlack, engine.CanvasRectSingleLine)
	drawingScene.AddEntity(topMenu)

	theEngine.InitResources()
	theEngine.GetSceneManager().AddScene(drawingScene)
	theEngine.GetSceneManager().SetSceneAsActive(drawingScene)
	theEngine.GetSceneManager().SetSceneAsVisible(drawingScene)
	theEngine.Init()
	theEngine.Start()
	theEngine.Run(theFPS)
}
