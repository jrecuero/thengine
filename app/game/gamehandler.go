package main

import (
	"fmt"

	"github.com/gdamore/tcell/v2"
	"github.com/jrecuero/thengine/app/game/dad/battlelog"
	"github.com/jrecuero/thengine/pkg/api"
	"github.com/jrecuero/thengine/pkg/engine"
	"github.com/jrecuero/thengine/pkg/tools"
	"github.com/jrecuero/thengine/pkg/widgets"
)

var (
	theGameHandler *GameHandler
)

type GameHandler struct {
	*engine.Entity
}

func NewGameHandler() *GameHandler {
	if theGameHandler == nil {
		tools.Logger.WithField("module", "gameHandler").WithField("function", "NewGameHandler").Debugf("handler/game/1")
		theGameHandler = &GameHandler{
			Entity: engine.NewHandler("handler/game/1"),
		}
		theGameHandler.SetFocusType(engine.MultiFocus)
		theGameHandler.SetFocusEnable(true)
	}
	return theGameHandler
}

// -----------------------------------------------------------------------------
// Module private methods
// -----------------------------------------------------------------------------

func displayEnemyHealthBar(scene engine.IScene, ent engine.IEntity) {
	enemy, _ := ent.(*Enemy)
	tmpText := scene.GetEntityByName(EnemyNameTextName)
	enemyText, _ := tmpText.(*widgets.Text)
	tmpHealthBar := scene.GetEntityByName(EnemyHealthBarName)
	enemyHealthBar, _ := tmpHealthBar.(*HealthBar)
	enemyText.SetVisible(true)
	enemyText.SetText(enemy.GetUName())
	enemyHealthBar.SetVisible(true)
	enemyHealthBar.SetTotal(enemy.GetHitPoints().GetMaxScore())
	enemyHealthBar.UpdateStyle(enemy.GetHitPoints().GetScore())
	enemyHealthBar.SetCompleted(enemy.GetHitPoints().GetScore())
}

func getEnemiesInScene(scene engine.IScene) []engine.IEntity {
	var result []engine.IEntity
	for _, ent := range scene.GetEntities() {
		if _, ok := ent.(*Enemy); ok {
			result = append(result, ent)
		}
	}
	return result
}

func hideEnemyHealthBar(scene engine.IScene) {
	tmpText := scene.GetEntityByName(EnemyNameTextName)
	enemyText, _ := tmpText.(*widgets.Text)
	tmpHealthBar := scene.GetEntityByName(EnemyHealthBarName)
	enemyHealthBar, _ := tmpHealthBar.(*HealthBar)
	enemyText.SetVisible(false)
	enemyHealthBar.SetVisible(false)
}

func isAnyEnemyAdjacent(player engine.IEntity, enemies []engine.IEntity) engine.IEntity {
	for _, enemy := range enemies {
		if player.GetPosition().IsAdjacent(enemy.GetPosition()) {
			return enemy
		}
	}
	return nil
}

func updateDataBox(scene engine.IScene, player *Player) {
	if tmp := scene.GetEntityByName(PlayerLiveTextName); tmp != nil {
		if playerLiveText, ok := tmp.(*widgets.Text); ok {
			hpText := fmt.Sprintf("HP:  %d", player.GetHitPoints().GetScore())
			playerLiveText.SetText(hpText)
		}
	}
	if tmp := scene.GetEntityByName(PlayerHealthBar); tmp != nil {
		if playerHealthBar, ok := tmp.(*HealthBar); ok {
			playerHealthBar.UpdateStyle(player.GetHitPoints().GetScore())
			playerHealthBar.SetCompleted(player.GetHitPoints().GetScore())
		}
	}
}

func writeToCommandLine(scene engine.IScene, str string) {
	commandLine := scene.GetEntityByName(CommandLineTextName)
	if commandLine != nil {
		if cl, ok := commandLine.(*CommandLine); ok {
			cl.AddText(str)
		}
	}
}

// -----------------------------------------------------------------------------
// GameHandler public methods
// -----------------------------------------------------------------------------

func (h *GameHandler) Update(event tcell.Event, scene engine.IScene) {
	if !h.HasFocus() {
		return
	}
	p := scene.GetEntityByName(thePlayerName)
	if p == nil {
		return
	}
	player, ok := p.(*Player)
	if !ok {
		return
	}
	playerX, playerY := player.GetPosition().Get()
	var playerNewPosition *api.Point
	switch ev := event.(type) {
	case *tcell.EventKey:
		switch ev.Key() {
		case tcell.KeyUp:
			playerNewPosition = api.NewPoint(playerX, playerY-1)
		case tcell.KeyDown:
			playerNewPosition = api.NewPoint(playerX, playerY+1)
		case tcell.KeyLeft:
			playerNewPosition = api.NewPoint(playerX-1, playerY)
		case tcell.KeyRight:
			playerNewPosition = api.NewPoint(playerX+1, playerY)
		case tcell.KeyRune:
			switch ev.Rune() {
			case 'A', 'a':
				enemies := getEnemiesInScene(scene)
				if enemy := isAnyEnemyAdjacent(player, enemies); enemy != nil {
					if e, ok := enemy.(*Enemy); ok {
						player.RollAttack(e)
						e.RollAttack(player)
						writeToCommandLine(scene, fmt.Sprintf("\n> %s [%d] attack to %s [%d]",
							player.GetName(), player.GetHitPoints().GetScore(),
							enemy.GetName(), e.GetHitPoints().GetScore()))
						//writeToCommandLine(scene, fmt.Sprintf("\n> player attack with damage %d", damage))
						updateDataBox(scene, player)
						tools.Logger.WithField("module", "gameHandler").WithField("method", "Update").Debugf("player can attack to %s", enemy.GetName())
						for battlelog.BLog.IsAny() {
							writeToCommandLine(scene, fmt.Sprintf("\n> %s", battlelog.BLog.Pop()))
						}
					}
				} else {
					writeToCommandLine(scene, fmt.Sprintf("\n> Player attack not available"))
				}
			}
		}
	}
	if playerNewPosition != nil {
		player.SetPosition(playerNewPosition)
		collisions := scene.CheckCollisionWith(player)
		for _, ent := range collisions {
			switch ent.(type) {
			case *Wall:
				player.SetPosition(api.NewPoint(playerX, playerY))
			case *Enemy:
				player.SetPosition(api.NewPoint(playerX, playerY))
			}
		}
	}
	enemies := getEnemiesInScene(scene)
	if enemy := isAnyEnemyAdjacent(player, enemies); enemy != nil {
		displayEnemyHealthBar(scene, enemy)
	} else {
		hideEnemyHealthBar(scene)
	}
}
