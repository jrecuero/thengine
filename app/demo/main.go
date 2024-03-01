package main

import (
	"fmt"

	"github.com/gdamore/tcell/v2"
	"github.com/jrecuero/thengine/pkg/api"
	"github.com/jrecuero/thengine/pkg/engine"
)

func main() {
	fmt.Println("ThEngine test")
	screen := engine.NewScreen(api.NewSize(40, 80))
	defaultStyle := tcell.StyleDefault
	text := engine.NewCanvasFromString("Hello World", &defaultStyle)
	text.Render(screen)
	appEngine := engine.NewEngine()
	appEngine.Init()
	screen.Draw(true, appEngine.Screen)
	appEngine.Run()
}
