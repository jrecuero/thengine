package main

import (
	"fmt"

	"github.com/gdamore/tcell/v2"
	"github.com/jrecuero/thengine/pkg/api"
	"github.com/jrecuero/thengine/pkg/engine"
)

func demoOne() {
	fmt.Println("ThEngine demo-one")
	screen := engine.NewScreen(api.NewSize(40, 80))
	defaultStyle := tcell.StyleDefault
	text := engine.NewCanvasFromString("Hello World", &defaultStyle)
	text.Render(screen)
	appEngine := engine.NewEngine(nil)
	appEngine.Init()
	screen.Draw(true, appEngine.GetDisplay())
	appEngine.Run(60.0)
}

func demoTwo() {
	fmt.Println("ThEngine demo-two")
	screen := engine.NewScreen(api.NewSize(40, 80))
	//defaultStyle := tcell.StyleDefault.Foreground(tcell.Color104).Background(tcell.ColorBlack).Attributes(tcell.AttrBlink)
	styleOne := tcell.StyleDefault.Foreground(tcell.Color100).Background(tcell.ColorBlack)
	styleTwo := tcell.StyleDefault.Foreground(tcell.Color101).Background(tcell.ColorBlack)
	scene := engine.NewScene("scene", screen)
	textOne := engine.NewEntity("text-one", api.NewPoint(0, 0), api.NewSize(1, 1), &styleOne)
	textOneCanvas := engine.NewCanvasFromString("Hello World!!!", &styleOne)
	textOne.SetCanvas(textOneCanvas)
	scene.AddEntity(textOne)
	textTwo := engine.NewEntity("text-two", api.NewPoint(0, 1), api.NewSize(1, 1), &styleTwo)
	textTwoCanvas := engine.NewCanvasFromString("Hello World******", &styleTwo)
	textTwo.SetCanvas(textTwoCanvas)
	scene.AddEntity(textTwo)
	appEngine := engine.NewEngine(nil)
	if !appEngine.GetSceneManager().AddScene(scene) {
		panic(fmt.Sprintf("can not add scene %s", scene.GetName()))
	}
	if !appEngine.GetSceneManager().SetSceneAsActive(scene) {
		panic(fmt.Sprintf("can not set scene %s as active", scene.GetName()))
	}
	if !appEngine.GetSceneManager().SetSceneAsVisible(scene) {
		panic(fmt.Sprintf("can not set scene %s as visible", scene.GetName()))
	}
	appEngine.Init()
	appEngine.Run(60.0)
}

func main() {
	demoTwo()
}
