package main

import (
	"fmt"

	"github.com/jrecuero/thengine/pkg/api"
	"github.com/jrecuero/thengine/pkg/engine"
)

func main() {
	fmt.Println("ThEngine test")
	screen := engine.NewScreen(api.NewSize(40, 80))
	text := engine.NewCanvasFromString("Hello World", api.ColorBlackAndWhite)
	text.Render(screen)
	appEngine := engine.NewEngine()
	appEngine.Init()
	screen.Draw(true)
	appEngine.Run()
}
