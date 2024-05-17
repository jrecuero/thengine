package main

import (
	"github.com/gdamore/tcell/v2"
	"github.com/jrecuero/thengine/pkg/api"
	"github.com/jrecuero/thengine/pkg/engine"
	"github.com/jrecuero/thengine/pkg/tools"
)

var (
	theGameHandler *GameHandler
)

type GameHandler struct {
	*engine.Entity
}

// -----------------------------------------------------------------------------
// Module private methods
// -----------------------------------------------------------------------------

func getEnemiesInScene(scene engine.IScene) []engine.IEntity {
	var result []engine.IEntity
	for _, ent := range scene.GetEntities() {
		if _, ok := ent.(*Enemy); ok {
			result = append(result, ent)
		}
	}
	return result
}

func isAnyEnemyAdjacent(player engine.IEntity, enemies []engine.IEntity) engine.IEntity {
	for _, enemy := range enemies {
		if player.GetPosition().IsAdjacent(enemy.GetPosition()) {
			return enemy
		}
	}
	return nil
}

// -----------------------------------------------------------------------------
// GameHandler public methods
// -----------------------------------------------------------------------------

func NewGameHandler() *GameHandler {
	if theGameHandler == nil {
		tools.Logger.WithField("module", "gameHandler").WithField("method", "NewGameHandler").Infof("handler/game/1")
		theGameHandler = &GameHandler{
			Entity: engine.NewNamedEntity("handler/game/1"),
		}
		theGameHandler.SetFocusType(engine.MultiFocus)
		theGameHandler.SetFocusEnable(true)
	}
	return theGameHandler
}

func (h *GameHandler) Update(event tcell.Event, scene engine.IScene) {
	if !h.HasFocus() {
		return
	}
	player := scene.GetEntityByName(thePlayerName)
	if player == nil {
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
					tools.Logger.WithField("module", "gameHandler").WithField("method", "Update").Infof("player can attack to %s", enemy.GetName())
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
}
