package main

import (
	"fmt"

	"github.com/gdamore/tcell/v2"
	"github.com/jrecuero/thengine/pkg/api"
	"github.com/jrecuero/thengine/pkg/engine"
	"github.com/jrecuero/thengine/pkg/tools"
	"github.com/jrecuero/thengine/pkg/widgets"
)

var (
	theCamera             = engine.NewCamera(api.NewPoint(0, 0), api.NewSize(90, 30))
	theEngine             = engine.GetEngine()
	theStyleBlueOverBlack = tcell.StyleDefault.Foreground(tcell.ColorBlue).Background(tcell.ColorBlack)
	theStyleWhiteOverRed  = tcell.StyleDefault.Foreground(tcell.ColorWhite).Background(tcell.ColorRed)
	theFPS                = 60.0
	thePlayerName         = "player/hero/1"
)

const (
	GameBoxEntityName      = "entity/game-box/1"
	DataBoxEntityName      = "entity/data-box/1"
	InfoBoxEntityName      = "entity/info-box/1"
	CommandLineTextName    = "text/command-line/1"
	PlayerLiveTextName     = "text/player/live/1"
	PlayerStrengthTextName = "text/player/strength/1"
)

func main() {
	tools.Logger.WithField("module", "game/main").Infof("The Game")
	mainScene := engine.NewScene("scene/main/1", theCamera)

	gameBox := engine.NewEntity(GameBoxEntityName, api.NewPoint(0, 0), api.NewSize(80, 20), &theStyleBlueOverBlack)
	gameBox.GetCanvas().WriteRectangleInCanvasAt(nil, nil, &theStyleBlueOverBlack, engine.CanvasRectSingleLine)
	mainScene.AddEntity(gameBox)

	dataBox := engine.NewEntity(DataBoxEntityName, api.NewPoint(80, 0), api.NewSize(20, 20), &theStyleBlueOverBlack)
	dataBox.GetCanvas().WriteRectangleInCanvasAt(nil, nil, &theStyleBlueOverBlack, engine.CanvasRectSingleLine)
	mainScene.AddEntity(dataBox)

	infoBox := engine.NewEntity(InfoBoxEntityName, api.NewPoint(0, 20), api.NewSize(100, 10), &theStyleBlueOverBlack)
	infoBox.GetCanvas().WriteRectangleInCanvasAt(nil, nil, &theStyleBlueOverBlack, engine.CanvasRectSingleLine)
	mainScene.AddEntity(infoBox)

	player := NewPlayer(thePlayerName, api.NewPoint(1, 1), &theStyleBlueOverBlack)
	mainScene.AddEntity(player)

	topWall := NewWall("widget/wall/top/1", api.NewPoint(0, 0), api.NewSize(80, 1), nil)
	topWall.SetVisible(false)
	mainScene.AddEntity(topWall)

	bottomWall := NewWall("widget/wall/bottom/1", api.NewPoint(0, 19), api.NewSize(80, 1), nil)
	bottomWall.SetVisible(false)
	mainScene.AddEntity(bottomWall)

	leftWall := NewWall("widget/wall/left/1", api.NewPoint(0, 1), api.NewSize(1, 18), nil)
	leftWall.SetVisible(false)
	mainScene.AddEntity(leftWall)

	rightWall := NewWall("widget/wall/right/1", api.NewPoint(79, 1), api.NewSize(1, 18), nil)
	rightWall.SetVisible(false)
	mainScene.AddEntity(rightWall)

	middleWall := NewWall("widget/wall/middle/1", api.NewPoint(2, 2), api.NewSize(76, 1), &theStyleBlueOverBlack)
	mainScene.AddEntity(middleWall)

	enemy := NewEnemy("widget/enemy/1", api.NewPoint(5, 5), &theStyleWhiteOverRed)
	mainScene.AddEntity(enemy)

	//commandLine := widgets.NewText(CommandLineTextName, api.NewPoint(1, 21), api.NewSize(98, 8), &theStyleBlueOverBlack, ">")
	commandLine := NewCommandLine(CommandLineTextName, api.NewPoint(1, 21), api.NewSize(98, 8), &theStyleBlueOverBlack)
	mainScene.AddEntity(commandLine)

	hpText := fmt.Sprintf("HP:  %d", player.GetHitPoints().GetScore())
	playerLiveText := widgets.NewText(PlayerLiveTextName, api.NewPoint(81, 1), api.NewSize(10, 1), &theStyleBlueOverBlack, hpText)
	mainScene.AddEntity(playerLiveText)

	strText := fmt.Sprintf("STR: %d", player.GetAbilities().GetStrength().GetScore())
	playerStrengthText := widgets.NewText(PlayerLiveTextName, api.NewPoint(81, 2), api.NewSize(10, 1), &theStyleBlueOverBlack, strText)
	mainScene.AddEntity(playerStrengthText)

	gameHandler := NewGameHandler()
	mainScene.AddEntity(gameHandler)

	theEngine.InitResources()
	theEngine.GetSceneManager().AddScene(mainScene)
	theEngine.GetSceneManager().SetSceneAsActive(mainScene)
	theEngine.GetSceneManager().SetSceneAsVisible(mainScene)
	theEngine.Init()
	theEngine.Start()
	theEngine.Run(theFPS)
}
