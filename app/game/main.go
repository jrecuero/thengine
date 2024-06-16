package main

import (
	"fmt"

	"github.com/gdamore/tcell/v2"
	"github.com/jrecuero/thengine/app/game/assets"
	dad_constants "github.com/jrecuero/thengine/app/game/dad/constants"
	"github.com/jrecuero/thengine/app/game/dad/rules"
	"github.com/jrecuero/thengine/pkg/api"
	"github.com/jrecuero/thengine/pkg/builder"
	"github.com/jrecuero/thengine/pkg/constants"
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
	PlayerPosTextName        = "text/player-position/1"
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
		tools.Logger.WithField("module", "main").
			WithField("struct", "BuiltIn").
			WithField("method", "GetClassFromString").
			Infof("Created a new empty enemy")
		return NewEmptyEnemy()
	default:
		return engine.NewEmptyEntity()
	}
}

func buildDungeon(scene engine.IScene) {

	cell := engine.NewCell(&constants.YellowOverBlack, '#')
	room1 := builder.BuildRoom("room/1", api.NewPoint(1, 1), api.NewSize(10, 7), cell,
		[]bool{false, false, false, true})
	scene.AddEntity(room1)

	corridor1 := builder.BuildCorridor("corridor/1", api.NewPoint(11, 4), api.NewPoint(15, 4), cell)
	scene.AddEntity(corridor1)

	room2 := builder.BuildRoom("room/2", api.NewPoint(16, 2), api.NewSize(5, 5), cell,
		[]bool{false, true, true, false})
	scene.AddEntity(room2)

	corridor2 := builder.BuildCorridor("corridor/2", api.NewPoint(18, 7), api.NewPoint(18, 10), cell)
	scene.AddEntity(corridor2)

	room3 := builder.BuildRoom("room/3", api.NewPoint(12, 11), api.NewSize(13, 7), cell,
		[]bool{true, false, false, true})
	scene.AddEntity(room3)

	corridor3 := builder.BuildCorridor("corridor/3", api.NewPoint(25, 14), api.NewPoint(30, 14), cell)
	scene.AddEntity(corridor3)

	room4 := builder.BuildRoom("room/4", api.NewPoint(31, 13), api.NewSize(3, 3), cell,
		[]bool{true, false, true, false})
	scene.AddEntity(room4)

	corridor4 := builder.BuildCorridor("corridor/4", api.NewPoint(32, 10), api.NewPoint(32, 12), cell)
	scene.AddEntity(corridor4)

	room5 := builder.BuildRoom("room/5", api.NewPoint(28, 4), api.NewSize(9, 7), cell,
		[]bool{false, true, false, true})
	scene.AddEntity(room5)

	corridor5 := builder.BuildCorridor("corridor/5", api.NewPoint(37, 7), api.NewPoint(50, 7), cell)
	scene.AddEntity(corridor5)

	room6 := builder.BuildRoom("room/6", api.NewPoint(51, 2), api.NewSize(20, 11), cell,
		[]bool{false, false, true, false})
	scene.AddEntity(room6)
}

func buildBoxesAndWalls(scene engine.IScene) {

	gameBox := engine.NewEntity(GameBoxEntityName, TheGameBoxOrigin,
		TheGameBoxSize, &theStyleBlueOverBlack)
	gameBox.GetCanvas().WriteRectangleInCanvasAt(nil, nil, &theStyleBlueOverBlack, engine.CanvasRectSingleLine)
	scene.AddEntity(gameBox)

	dataBox := engine.NewEntity(DataBoxEntityName, TheDataBoxOrigin,
		TheDataBoxSize, &theStyleBlueOverBlack)
	dataBox.GetCanvas().WriteRectangleInCanvasAt(nil, nil, &theStyleBlueOverBlack, engine.CanvasRectSingleLine)
	scene.AddEntity(dataBox)

	comandLineBox := engine.NewEntity(CommandLineBoxEntityName, TheCommandLineBoxOrigin,
		TheCommandLineBoxSize, &theStyleBlueOverBlack)
	comandLineBox.GetCanvas().WriteRectangleInCanvasAt(nil, nil, &theStyleBlueOverBlack, engine.CanvasRectSingleLine)
	scene.AddEntity(comandLineBox)

	topWall := NewWall("widget/wall/top/1", api.NewPoint(0, 0), api.NewSize(80, 1), nil)
	topWall.SetVisible(false)
	scene.AddEntity(topWall)

	bottomWall := NewWall("widget/wall/bottom/1", api.NewPoint(0, 19), api.NewSize(80, 1), nil)
	bottomWall.SetVisible(false)
	scene.AddEntity(bottomWall)

	leftWall := NewWall("widget/wall/left/1", api.NewPoint(0, 1), api.NewSize(1, 18), nil)
	leftWall.SetVisible(false)
	scene.AddEntity(leftWall)

	rightWall := NewWall("widget/wall/right/1", api.NewPoint(79, 1), api.NewSize(1, 18), nil)
	rightWall.SetVisible(false)
	scene.AddEntity(rightWall)

	//commandLine := widgets.NewText(CommandLineTextName, api.NewPoint(1, 21),
	//    api.NewSize(98, 8), &theStyleBlueOverBlack, ">")
	commandLine := NewCommandLine(CommandLineTextName, api.NewPoint(1, 21),
		api.NewSize(98, 8), &theStyleBlueOverBlack)
	scene.AddEntity(commandLine)
}

func buildUI(scene engine.IScene, player *Player, enemy *Enemy) {

	hpText := fmt.Sprintf("HP:  %d", player.GetHitPoints().GetScore())
	playerLiveText := widgets.NewText(PlayerLiveTextName, api.NewPoint(81, 1),
		api.NewSize(10, 1), &theStyleBlueOverBlack, hpText)
	scene.AddEntity(playerLiveText)

	strText := fmt.Sprintf("STR: %d", player.GetAbilities().GetStrength().GetScore())
	playerStrengthText := widgets.NewText(PlayerStrengthTextName, api.NewPoint(81, 2),
		api.NewSize(10, 1), &theStyleBlueOverBlack, strText)
	scene.AddEntity(playerStrengthText)

	dexText := fmt.Sprintf("DEX: %d", player.GetAbilities().GetDexterity().GetScore())
	playerDexterityText := widgets.NewText(PlayerDexterityTextName, api.NewPoint(81, 3),
		api.NewSize(10, 1), &theStyleBlueOverBlack, dexText)
	scene.AddEntity(playerDexterityText)

	acText := fmt.Sprintf("AC:  %d", player.GetArmorClass())
	playerACText := widgets.NewText(PlayerACTextName, api.NewPoint(81, 4),
		api.NewSize(10, 1), &theStyleBlueOverBlack, acText)
	scene.AddEntity(playerACText)

	playerNameText := widgets.NewText(PlayerNameTextName, api.NewPoint(81, 9),
		api.NewSize(18, 1), &theStyleBlueOverBlack, player.GetUName())
	scene.AddEntity(playerNameText)

	playerHealthBar := NewHealthBar(PlayerHealthBar, api.NewPoint(81, 10),
		api.NewSize(18, 1), player.GetHitPoints().GetScore())
	playerHealthBar.SetCompleted(player.GetHitPoints().GetScore())
	scene.AddEntity(playerHealthBar)

	enemyText := fmt.Sprintf("%s\t[AC:%d]", enemy.GetUName(), enemy.GetArmorClass())
	enemyNameText := widgets.NewText(EnemyNameTextName, api.NewPoint(81, 11),
		api.NewSize(18, 1), &theStyleBlueOverBlack, enemyText)
	enemyNameText.SetVisible(false)
	scene.AddEntity(enemyNameText)

	enemyHealthBar := NewHealthBar(EnemyHealthBarName, api.NewPoint(81, 12),
		api.NewSize(18, 1), enemy.GetHitPoints().GetScore())
	enemyHealthBar.SetCompleted(enemy.GetHitPoints().GetScore())
	enemyHealthBar.SetVisible(false)
	scene.AddEntity(enemyHealthBar)

	playerPosText := widgets.NewText(PlayerPosTextName, api.NewPoint(70, 19),
		api.NewSize(10, 1), &constants.WhiteOverBlack, "[2,2]")
	playerPosText.SetZLevel(1)
	scene.AddEntity(playerPosText)
}

func newTrap(name string, pos *api.Point, style *tcell.Style) *assets.Trap {
	trap := assets.NewTrap(name, pos, api.NewSize(1, 1), style)
	trapDC := 8
	trapDamage := &rules.SavingThrowDamage{
		SavingThrow: rules.NewSavingThrow(dad_constants.Perception, trapDC),
		Damage:      rules.NewDamage(rules.DiceThrow1d3, dad_constants.Poison),
	}
	trap.Damage = rules.NewNoDamage()
	trap.Damage.SetSavingThrows([]*rules.SavingThrowDamage{trapDamage})
	trap.GetCanvas().SetCellAt(nil, engine.NewCell(&constants.RedOverBlack, '*'))
	return trap
}

func main() {
	tools.Logger.WithField("module", "main").Infof("The Game")
	mainScene := engine.NewScene("scene/main/1", theCamera)

	buildBoxesAndWalls(mainScene)

	player := NewPlayer(thePlayerName, api.NewPoint(2, 2), &theStyleGreenOverBlack)
	mainScene.AddEntity(player)

	//entities := engine.ImportEntitiesFromJSON("app/game/assets/first_map.json",
	//    api.NewPoint(1, 1), &BuiltIn{})
	//for _, ent := range entities {
	//    mainScene.AddEntity(ent)
	//}

	enemy1 := NewEnemy(GenerateEnemyName(), api.NewPoint(5, 5), &theStyleWhiteOverRed)
	mainScene.AddEntity(enemy1)
	TheEnemies = append(TheEnemies, enemy1)

	enemy2 := NewEnemy(GenerateEnemyName(), api.NewPoint(67, 7), &constants.AquaOverWhite)
	mainScene.AddEntity(enemy2)
	TheEnemies = append(TheEnemies, enemy2)

	buildUI(mainScene, player, enemy1)

	//diceOneAnimWidget := widgets.NewAnimWidget("anim-widget/dice/1", api.NewPoint(65, 13),
	//    api.NewSize(6, 6), assets.NewAsciiFramesForAllNumbers(10), 0)
	//diceOneAnimWidget.Shuffle()
	//mainScene.AddEntity(diceOneAnimWidget)
	//diceTwoAnimWidget := widgets.NewAnimWidget("anim-widget/dice/1", api.NewPoint(72, 13),
	//    api.NewSize(6, 6), assets.NewAsciiFramesForAllNumbers(10), 0)
	//diceTwoAnimWidget.Shuffle()
	//mainScene.AddEntity(diceTwoAnimWidget)

	throne := widgets.NewSprite("sprite/throne/1", api.NewPoint(67, 6), nil)
	throne.StringToSprite("--\\\n #|\n--/", &constants.MaroonOverBlack)
	throne.SetSolid(true)
	mainScene.AddEntity(throne)

	throneRoom := widgets.NewSprite("sprite/throne-room/1", api.NewPoint(52, 3), nil)
	//throneRoom.StringToSprite("Throne Room", &constants.RedOverBlack)
	throneRoom.StringToSprite("Throne .", &constants.RedOverBlack, false)
	throneRoom.StringToSpriteAtEnd("Room", &constants.AquaOverBlack)
	//throne.SetSolid(true)
	mainScene.AddEntity(throneRoom)

	diceOneAnimWidget := assets.NewShuffleDieWidget("widget/anim-dice/1",
		api.NewPoint(72, 16), &theStyleRedOverBlack, 6, 10)
	mainScene.AddEntity(diceOneAnimWidget)
	diceTwoAnimWidget := assets.NewShuffleDieWidget("widget/anim-dice/1",
		api.NewPoint(75, 16), &theStyleBlueOverBlack, 20, 10)
	mainScene.AddEntity(diceTwoAnimWidget)

	buildDungeon(mainScene)

	//trap := assets.NewTrap("widget/trap/1", api.NewPoint(14, 4), api.NewSize(1, 1), &constants.RedOverBlack)
	//trapDC := 8
	//trapDamage := &rules.SavingThrowDamage{
	//    SavingThrow: rules.NewSavingThrow(dad_constants.Wisdom, trapDC),
	//    Damage:      rules.NewDamage(rules.DiceThrow1d3, dad_constants.Poison),
	//}
	//trap.Damage = rules.NewNoDamage()
	//trap.Damage.SetSavingThrows([]*rules.SavingThrowDamage{trapDamage})
	//trap.GetCanvas().SetCellAt(nil, engine.NewCell(&constants.RedOverBlack, '*'))
	trap := newTrap("widget/trap/1", api.NewPoint(14, 4), &constants.RedOverBlack)
	mainScene.AddEntity(trap)

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
