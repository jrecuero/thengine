package main

import (
	"github.com/jrecuero/thengine/pkg/engine"
	"github.com/jrecuero/thengine/pkg/tools"
)

func main() {
	tools.Logger.WithField("module", "main").WithField("function", "main").Infof("RhuneDice launched...")

	mainScene := engine.NewScene(TheMainSceneName, theCamera)
	gameHandler := NewGameHandler()

	buildBoxes(mainScene, gameHandler)

	buildDungeon(mainScene)

	mainScene.AddEntity(gameHandler)

	theEngine.InitResources()
	theEngine.GetSceneManager().AddScene(mainScene)
	theEngine.GetSceneManager().SetSceneAsActive(mainScene)
	theEngine.GetSceneManager().SetSceneAsVisible(mainScene)
	theEngine.GetSceneManager().UpdateFocus()
	theEngine.Init()
	theEngine.Start()
	theEngine.Run(theFPS)
}
