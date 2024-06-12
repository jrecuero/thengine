package main

import (
	"fmt"

	"github.com/gdamore/tcell/v2"
	"github.com/jrecuero/thengine/app/game/assets"
	"github.com/jrecuero/thengine/pkg/api"
	"github.com/jrecuero/thengine/pkg/engine"
	"github.com/jrecuero/thengine/pkg/tools"
	"github.com/jrecuero/thengine/pkg/widgets"
)

// -----------------------------------------------------------------------------
// Module private constants
// -----------------------------------------------------------------------------

var (
	theCamera              = engine.NewCamera(api.NewPoint(0, 0), api.NewSize(90, 30))
	theEngine              = engine.GetEngine()
	theStyleBlueOverBlack  = tcell.StyleDefault.Foreground(tcell.ColorBlue).Background(tcell.ColorBlack)
	theStyleWhiteOverRed   = tcell.StyleDefault.Foreground(tcell.ColorWhite).Background(tcell.ColorRed)
	theStyleGreenOverBlack = tcell.StyleDefault.Foreground(tcell.ColorGreen).Background(tcell.ColorBlack)
	theStyleRedOverBlack   = tcell.StyleDefault.Foreground(tcell.ColorRed).Background(tcell.ColorBlack)
	theFPS                 = 60.0
	thePlayerName          = "player/hero/1"
)

// -----------------------------------------------------------------------------
// Module public constants
// -----------------------------------------------------------------------------

var (
	TheGameBoxOrigin        = api.NewPoint(0, 0)
	TheGameBoxSize          = api.NewSize(80, 20)
	TheDataBoxOrigin        = api.NewPoint(80, 0)
	TheDataBoxSize          = api.NewSize(20, 20)
	TheCommandLineBoxOrigin = api.NewPoint(0, 20)
	TheCommandLineBoxSize   = api.NewSize(100, 10)
	TheEnemies              = []*Enemy{}
)

// -----------------------------------------------------------------------------
// Module public constants
// -----------------------------------------------------------------------------

const (
	GameBoxEntityName        = "entity/game-box/1"
	DataBoxEntityName        = "entity/data-box/1"
	CommandLineBoxEntityName = "entity/command-line-box/1"
	CommandLineTextName      = "text/command-line/1"
	PlayerLiveTextName       = "text/player/live/1"
	PlayerStrengthTextName   = "text/player/strength/1"
	PlayerDexterityTextName  = "text/player/dexteriry/1"
	PlayerACTextName         = "text/player/ac/1"
	PlayerNameTextName       = "text/player/name/1"
	PlayerHealthBar          = "health-bar/player/live/1"
	EnemyNameTextName        = "text/enemy/name/1"
	EnemyHealthBarName       = "health-bar/enemy/live/1"
)

// -----------------------------------------------------------------------------
// Module public structures
// -----------------------------------------------------------------------------

type BuiltIn struct {
	engine.IBuiltIn
}

// -----------------------------------------------------------------------------
// Module private methods
// -----------------------------------------------------------------------------

// -----------------------------------------------------------------------------
// Module public methods
// -----------------------------------------------------------------------------

func GenerateEnemyName() string {
	enemiesLen := len(TheEnemies) + 1
	result := fmt.Sprintf("widget/enemy/%d", enemiesLen)
	return result
}

func (b *BuiltIn) GetClassFromString(className string) engine.IEntity {
	switch className {
	case "Wall":
		return NewEmptyWall()
	case "Enemy":
		tools.Logger.WithField("module", "game/main").
			WithField("struct", "BuiltIn").
			WithField("method", "GetClassFromString").
			Infof("Created a new empty enemy")
		return NewEmptyEnemy()
	default:
		return engine.NewEmptyEntity()
	}
}

func main() {
	tools.Logger.WithField("module", "game/main").Infof("The Game")
	mainScene := engine.NewScene("scene/main/1", theCamera)

	gameBox := engine.NewEntity(GameBoxEntityName, TheGameBoxOrigin, TheGameBoxSize, &theStyleBlueOverBlack)
	gameBox.GetCanvas().WriteRectangleInCanvasAt(nil, nil, &theStyleBlueOverBlack, engine.CanvasRectSingleLine)
	mainScene.AddEntity(gameBox)

	dataBox := engine.NewEntity(DataBoxEntityName, TheDataBoxOrigin, TheDataBoxSize, &theStyleBlueOverBlack)
	dataBox.GetCanvas().WriteRectangleInCanvasAt(nil, nil, &theStyleBlueOverBlack, engine.CanvasRectSingleLine)
	mainScene.AddEntity(dataBox)

	comandLineBox := engine.NewEntity(CommandLineBoxEntityName, TheCommandLineBoxOrigin, TheCommandLineBoxSize, &theStyleBlueOverBlack)
	comandLineBox.GetCanvas().WriteRectangleInCanvasAt(nil, nil, &theStyleBlueOverBlack, engine.CanvasRectSingleLine)
	mainScene.AddEntity(comandLineBox)

	player := NewPlayer(thePlayerName, api.NewPoint(2, 2), &theStyleGreenOverBlack)
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

	entities := engine.ImportEntitiesFromJSON("app/game/assets/first_map.json", api.NewPoint(1, 1), &BuiltIn{})
	for _, ent := range entities {
		mainScene.AddEntity(ent)
	}

	enemy := NewEnemy(GenerateEnemyName(), api.NewPoint(5, 5), &theStyleWhiteOverRed)
	mainScene.AddEntity(enemy)
	TheEnemies = append(TheEnemies, enemy)

	//commandLine := widgets.NewText(CommandLineTextName, api.NewPoint(1, 21), api.NewSize(98, 8), &theStyleBlueOverBlack, ">")
	commandLine := NewCommandLine(CommandLineTextName, api.NewPoint(1, 21), api.NewSize(98, 8), &theStyleBlueOverBlack)
	mainScene.AddEntity(commandLine)

	hpText := fmt.Sprintf("HP:  %d", player.GetHitPoints().GetScore())
	playerLiveText := widgets.NewText(PlayerLiveTextName, api.NewPoint(81, 1), api.NewSize(10, 1), &theStyleBlueOverBlack, hpText)
	mainScene.AddEntity(playerLiveText)

	strText := fmt.Sprintf("STR: %d", player.GetAbilities().GetStrength().GetScore())
	playerStrengthText := widgets.NewText(PlayerStrengthTextName, api.NewPoint(81, 2), api.NewSize(10, 1), &theStyleBlueOverBlack, strText)
	mainScene.AddEntity(playerStrengthText)

	dexText := fmt.Sprintf("DEX: %d", player.GetAbilities().GetDexterity().GetScore())
	playerDexterityText := widgets.NewText(PlayerDexterityTextName, api.NewPoint(81, 3), api.NewSize(10, 1), &theStyleBlueOverBlack, dexText)
	mainScene.AddEntity(playerDexterityText)

	acText := fmt.Sprintf("AC:  %d", player.GetArmorClass())
	playerACText := widgets.NewText(PlayerACTextName, api.NewPoint(81, 4), api.NewSize(10, 1), &theStyleBlueOverBlack, acText)
	mainScene.AddEntity(playerACText)

	playerNameText := widgets.NewText(PlayerNameTextName, api.NewPoint(81, 9), api.NewSize(18, 1), &theStyleBlueOverBlack, player.GetUName())
	mainScene.AddEntity(playerNameText)

	playerHealthBar := NewHealthBar(PlayerHealthBar, api.NewPoint(81, 10), api.NewSize(18, 1), player.GetHitPoints().GetScore())
	playerHealthBar.SetCompleted(player.GetHitPoints().GetScore())
	mainScene.AddEntity(playerHealthBar)

	enemyText := fmt.Sprintf("%s\t[AC:%d]", enemy.GetUName(), enemy.GetArmorClass())
	enemyNameText := widgets.NewText(EnemyNameTextName, api.NewPoint(81, 11), api.NewSize(18, 1), &theStyleBlueOverBlack, enemyText)
	enemyNameText.SetVisible(false)
	mainScene.AddEntity(enemyNameText)

	enemyHealthBar := NewHealthBar(EnemyHealthBarName, api.NewPoint(81, 12), api.NewSize(18, 1), enemy.GetHitPoints().GetScore())
	enemyHealthBar.SetCompleted(enemy.GetHitPoints().GetScore())
	enemyHealthBar.SetVisible(false)
	mainScene.AddEntity(enemyHealthBar)

	//diceOneAnimWidget := widgets.NewAnimWidget("anim-widget/dice/1", api.NewPoint(65, 13), api.NewSize(6, 6), assets.NewAsciiFramesForAllNumbers(10), 0)
	//diceOneAnimWidget.Shuffle()
	//mainScene.AddEntity(diceOneAnimWidget)
	//diceTwoAnimWidget := widgets.NewAnimWidget("anim-widget/dice/1", api.NewPoint(72, 13), api.NewSize(6, 6), assets.NewAsciiFramesForAllNumbers(10), 0)
	//diceTwoAnimWidget.Shuffle()
	//mainScene.AddEntity(diceTwoAnimWidget)

	diceOneAnimWidget := assets.NewShuffleDieWidget("widget/anim-dice/1", api.NewPoint(72, 16), &theStyleRedOverBlack, 6, 10)
	mainScene.AddEntity(diceOneAnimWidget)
	diceTwoAnimWidget := assets.NewShuffleDieWidget("widget/anim-dice/1", api.NewPoint(75, 16), &theStyleBlueOverBlack, 20, 10)
	mainScene.AddEntity(diceTwoAnimWidget)

	gameHandler := NewGameHandler()
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
