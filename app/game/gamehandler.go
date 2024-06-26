package main

import (
	"fmt"

	"github.com/gdamore/tcell/v2"
	"github.com/jrecuero/thengine/app/game/assets"
	"github.com/jrecuero/thengine/app/game/dad/battlelog"
	"github.com/jrecuero/thengine/pkg/api"
	"github.com/jrecuero/thengine/pkg/engine"
	"github.com/jrecuero/thengine/pkg/tools"
	"github.com/jrecuero/thengine/pkg/widgets"
)

var (
	theGameHandler *GameHandler
)

// -----------------------------------------------------------------------------
// Private private types
// -----------------------------------------------------------------------------

type attackInfo struct {
	index int
	name  string
}

type inputAction struct {
	selAttack   bool
	attack      *attackInfo
	pos         *api.Point
	consumeItem bool
}

func newInputActionWithSelectAttack() *inputAction {
	return &inputAction{
		selAttack:   true,
		attack:      nil,
		pos:         nil,
		consumeItem: false,
	}
}

func newInputActionWithAttack(index int, name string) *inputAction {
	return &inputAction{
		selAttack: false,
		attack: &attackInfo{
			index: index,
			name:  name,
		},
		pos:         nil,
		consumeItem: false,
	}
}

func newInputActionWithPosition(pos *api.Point) *inputAction {
	return &inputAction{
		selAttack:   false,
		attack:      nil,
		pos:         pos,
		consumeItem: false,
	}
}

func newInputActionWithConsumeItem() *inputAction {
	return &inputAction{
		selAttack:   false,
		attack:      nil,
		pos:         nil,
		consumeItem: true,
	}
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
	enemyText.SetText(fmt.Sprintf("%s\t[AC:%d]", enemy.GetUName(), enemy.GetArmorClass()))
	enemyHealthBar.SetVisible(true)
	enemyHealthBar.SetTotal(enemy.GetHitPoints().GetMaxScore())
	enemyHealthBar.UpdateStyle(enemy.GetHitPoints().GetScore())
	enemyHealthBar.SetCompleted(enemy.GetHitPoints().GetScore())
}

func getEnemiesInScene(scene engine.IScene) []engine.IEntity {
	var result []engine.IEntity
	for _, entity := range scene.GetEntities() {
		if _, ok := entity.(*Enemy); ok {
			result = append(result, entity)
		}
	}
	return result
}

func getTrapsInScene(scene engine.IScene) []assets.ITrap {
	var result []assets.ITrap
	for _, entity := range scene.GetEntities() {
		if trap, ok := entity.(*assets.Trap); ok {
			result = append(result, trap)
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

func isAnyTrapAdjacent(player engine.IEntity, traps []assets.ITrap) assets.ITrap {
	for _, trap := range traps {
		if player.GetPosition().IsAdjacent(trap.GetPosition()) {
			return trap
		}
	}
	return nil
}

func readFromBattleLog(scene engine.IScene) {
	for battlelog.BLog.IsAny() {
		if str := battlelog.BLog.PopInfo(); str != "" {
			writeToCommandLine(scene, fmt.Sprintf("\n> %s", str))
		}
	}
}

func updateDataBox(scene engine.IScene, player *Player) {
	if tmp := scene.GetEntityByName(PlayerNameTextName); tmp != nil {
		if playerNameText, ok := tmp.(*widgets.Text); ok {
			playerLevel := player.GetLevel()
			playerStr := fmt.Sprintf("%s L%d [%d]", player.GetUName(), playerLevel.GetScore(), playerLevel.GetExperience())
			playerNameText.SetText(playerStr)
		}
	}
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
	if tmp := scene.GetEntityByName(InventoryTextName); tmp != nil {
		if inventoryText, ok := tmp.(*widgets.Text); ok {
			inventoryStr := "Inventory\n"
			inventory := player.GetInventory()
			for _, consumable := range inventory.GetConsumables() {
				inventoryStr += consumable.GetUName() + "\n"
			}
			inventoryText.SetText(inventoryStr)
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
//
// GameHandler
//
// -----------------------------------------------------------------------------

type GameHandler struct {
	*engine.Entity
	player             *Player
	enemy              *Enemy
	playerActionOption int
}

func NewGameHandler() *GameHandler {
	if theGameHandler == nil {
		tools.Logger.WithField("module", "gamehandler").
			WithField("function", "NewGameHandler").
			Debugf("handler/game/1")
		theGameHandler = &GameHandler{
			Entity:             engine.NewHandler("handler/game/1"),
			player:             nil,
			enemy:              nil,
			playerActionOption: -1,
		}
		theGameHandler.SetFocusType(engine.SingleFocus)
		theGameHandler.SetFocusEnable(true)
	}
	return theGameHandler
}

// -----------------------------------------------------------------------------
// GameHandler private methods
// -----------------------------------------------------------------------------

func (h *GameHandler) playerSelection(entity engine.IEntity, args ...any) bool {
	scene := args[0].(engine.IScene)
	tools.Logger.WithField("module", "gamehandler").
		WithField("method", "playerSelection").
		Debugf("selection %d", entity.(*widgets.ListBox).GetSelectionIndex())
	lb := entity.(*widgets.ListBox)
	scene.RemoveEntity(entity)
	newInput := newInputActionWithAttack(lb.GetSelectionIndex(), lb.GetSelection())
	h.RunStateMachineTurn(scene, newInput)
	return true
}

// -----------------------------------------------------------------------------
// GameHandler public methods
// -----------------------------------------------------------------------------

func (h *GameHandler) EnemyAttack(scene engine.IScene) {
	if h.enemy != nil {
		// Roll Damage only if enemy gets a hit from the die roll.
		if hit := h.enemy.RollDieRoll(0, h.player); hit {
			h.enemy.RollDamage(0, h.player)
		}
		updateDataBox(scene, h.player)
		tools.Logger.WithField("module", "gamehandler").
			WithField("method", "Update").
			Debugf("Enemy %s to %s", h.enemy.GetName(), h.player.GetName())
		readFromBattleLog(scene)
	}
}

func (h *GameHandler) PlayerAttack(scene engine.IScene, attack *attackInfo) {
	enemies := getEnemiesInScene(scene)
	if enemy := isAnyEnemyAdjacent(h.player, enemies); enemy != nil {
		if e, ok := enemy.(*Enemy); ok {
			h.enemy = e
			// Roll Damage only if player gets a hit from the die roll.
			if hit := h.player.RollDieRoll(attack.index, e); hit {
				h.player.RollDamage(attack.index, e)
			}
			readFromBattleLog(scene)
			// if enemy has equal or lower than 0 hit points, remove it from
			// the scene.
			if e.GetHitPoints().GetScore() <= 0 {
				exp := e.GetLevel().GetToGive()
				h.player.GetLevel().IncExperience(exp)
				scene.RemoveEntity(e)
				h.enemy = nil
				writeToCommandLine(scene, fmt.Sprintf("\n> Enemy %s defeated", e.GetUName()))
			}
			tools.Logger.WithField("module", "gamehandler").
				WithField("method", "Update").
				Debugf("player %s to %s", attack.name, enemy.GetName())
		}
	} else {
		writeToCommandLine(scene, fmt.Sprintf("\n> Player attack not available"))
	}
}

func (h *GameHandler) PlayerMove(scene engine.IScene, playerNewPosition *api.Point) {
	playerX, playerY := h.player.GetPosition().Get()
	h.player.SetPosition(playerNewPosition)
	collisions := scene.CheckCollisionWith(h.player)
	for _, entity := range collisions {
		switch entity.(type) {
		case IDungeonEvent:
			h.player.SetPosition(api.NewPoint(playerX, playerY))
			doorEvent := entity.(*DoorEvent)
			doorEvent.Event.Run(scene, doorEvent)
		case *Wall:
			h.player.SetPosition(api.NewPoint(playerX, playerY))
		case *Enemy:
			h.player.SetPosition(api.NewPoint(playerX, playerY))
		case assets.ITrap:
			// if the trap is not visible, it was not detected, so damage has
			// to be applied to the player.
			// if the trap is visible, it was detected, so try to disarm it, if
			// it fails, trap is removed and damage is applied to the player
			// but player does not update the position in that case and trap is
			// being removed.
			// if trap is properly disarmed, trap is removed and player update
			// the position.
			if trap, ok := entity.(*assets.Trap); ok {
				damage := 0
				if entity.IsVisible() {
					if disarm := trap.CanDisarm(h.player); disarm {
						writeToCommandLine(scene, fmt.Sprintf("\n> player %s disarm trap %s",
							h.player.GetName(), trap.GetName()))
					} else {
						h.player.SetPosition(api.NewPoint(playerX, playerY))
						damage = trap.RollDamageValue()
						writeToCommandLine(scene, fmt.Sprintf("\n> player %s damaged disarming trap %s for %d",
							h.player.GetName(), trap.GetName(), damage))
					}
				} else {
					h.player.SetPosition(api.NewPoint(playerX, playerY))
					damage = trap.RollDamageValue()
					writeToCommandLine(scene, fmt.Sprintf("\n> player %s damaged by trap %s for %d",
						h.player.GetName(), trap.GetName(), damage))
				}
				if damage != 0 {
					h.player.GetHitPoints().Dec(damage)
				}
				scene.RemoveEntity(entity)
			}
		case *widgets.Sprite:
			h.player.SetPosition(api.NewPoint(playerX, playerY))
		}
	}

	// check for any enemy that is adjacent to the final player position.
	if tmp := scene.GetEntityByName(PlayerPosTextName); tmp != nil {
		if cursorPosText, ok := tmp.(*widgets.Text); ok {
			pos := api.ClonePoint(h.player.GetPosition())
			pos.Subtract(TheGameBoxOrigin)
			cursorPosText.SetText(fmt.Sprintf("[%d,%d]", pos.X, pos.Y))
		}
	}

	// check for any trap that is adjacent to the final player position.
	traps := getTrapsInScene(scene)
	if t := isAnyTrapAdjacent(h.player, traps); t != nil {
		if trap, ok := t.(*assets.Trap); ok {
			if pass := trap.CanDetect(h.player); pass {
				t.SetVisible(true)
				writeToCommandLine(scene, fmt.Sprintf("\n> player %s detect trap %s",
					h.player.GetName(), trap.GetName()))
			}
		}
	}
}

func (h *GameHandler) RunEndTurn(scene engine.IScene, input *inputAction) {
	tools.Logger.WithField("module", "gamehandler").
		WithField("method", "RunEndTurn").
		Debugf("END TURN %+v", input)
	h.player = nil
	h.enemy = nil
	h.RunStateMachineTurn(scene, input)
}

func (h *GameHandler) RunEnemyTurn(scene engine.IScene, input *inputAction) {
	tools.Logger.WithField("module", "gamehandler").
		WithField("method", "EnemyTurn").
		Debugf("ENEMY TURN %+v", input)
	h.EnemyAttack(scene)
	h.RunStateMachineTurn(scene, input)
}

func (h *GameHandler) RunPlayerTurn(scene engine.IScene, input *inputAction) {
	tools.Logger.WithField("module", "gamehandler").
		WithField("method", "PlayerTurn").
		Debugf("PLAYER TURN %+v", input)
	if input != nil && input.attack != nil {
		h.PlayerAttack(scene, input.attack)
	}
	if input != nil && input.pos != nil {
		h.PlayerMove(scene, input.pos)
	}
	updateDataBox(scene, h.player)
	h.RunStateMachineTurn(scene, input)
}

func (h *GameHandler) RunStartTurn(scene engine.IScene, input *inputAction) {
	tools.Logger.WithField("module", "gamehandler").
		WithField("method", "StartTurn").
		Debugf("START TURN %+v", input)
	enemies := getEnemiesInScene(scene)
	// TODO: initiative should be checked for player and enemies in order to
	// define who should be take action first and in which order.
	if input != nil {
		if input.consumeItem {
			consumables := h.player.GetInventory().GetConsumables()
			if len(consumables) != 0 {
				potion := consumables[0]
				potion.Consume(h.player)
				h.player.GetInventory().RemoveConsumable(potion)
				updateDataBox(scene, h.player)
			}
		} else if input.selAttack {
			if enemy := isAnyEnemyAdjacent(h.player, enemies); enemy != nil {
				x, y := h.player.GetPosition().Get()
				options := widgets.NewListBox("list-box/player-options/1",
					api.NewPoint(x, y), api.NewSize(10, 5), &theStyleBlueOverBlack,
					[]string{"weapon", "power", "magical"}, 0)
				options.SetZLevel(1)
				options.SetWidgetCallback(h.playerSelection, scene)
				scene.AddEntity(options)
				theEngine.GetSceneManager().UpdateFocus()
				return
			}
		}
	}
	h.RunStateMachineTurn(scene, input)
}

func (h *GameHandler) RunStateMachineTurn(scene engine.IScene, input *inputAction) {
	TheStateMachine.Next()
	state := TheStateMachine.GetState()
	switch state {
	case WaitingSM:
		// do nothing
		h.RunWaitingTurn(scene, input)
	case StartSM:
		// run everything required at the start of the turn
		h.RunStartTurn(scene, input)
	case PlayerSM:
		// run player action (move or attack)
		h.RunPlayerTurn(scene, input)
	case EnemySM:
		// run selected enemy action (attack)
		h.RunEnemyTurn(scene, input)
	case EndSM:
		// run everything required at the end of the turn
		h.RunEndTurn(scene, input)
	}
	//tools.Logger.WithField("module", "gamehandler").
	//    WithField("method", "Next State").
	//    Debugf("state is %d", state)
}

func (h *GameHandler) RunWaitingTurn(scene engine.IScene, input *inputAction) {
	tools.Logger.WithField("module", "gamehandler").
		WithField("method", "WaitingTurn").
		Debugf("WAITING TURN %+v", input)
}

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
	if !TheStateMachine.IsWaiting() {
		return
	}
	h.player = player

	playerX, playerY := player.GetPosition().Get()
	var input *inputAction
	switch ev := event.(type) {
	case *tcell.EventKey:
		switch ev.Key() {
		case tcell.KeyUp:
			input = newInputActionWithPosition(api.NewPoint(playerX, playerY-1))
		case tcell.KeyDown:
			input = newInputActionWithPosition(api.NewPoint(playerX, playerY+1))
		case tcell.KeyLeft:
			input = newInputActionWithPosition(api.NewPoint(playerX-1, playerY))
		case tcell.KeyRight:
			input = newInputActionWithPosition(api.NewPoint(playerX+1, playerY))
		case tcell.KeyRune:
			switch ev.Rune() {
			case 'A', 'a':
				input = newInputActionWithSelectAttack()
			case '1':
				input = newInputActionWithConsumeItem()
			}
		}
	}
	if input != nil {
		tools.Logger.WithField("module", "gamehandler").
			WithField("method", "Update").
			Debugf("CALL RunStateMachineTurn %+v", input)
		h.RunStateMachineTurn(scene, input)
	}

	enemies := getEnemiesInScene(scene)
	if enemy := isAnyEnemyAdjacent(player, enemies); enemy != nil {
		displayEnemyHealthBar(scene, enemy)
	} else {
		hideEnemyHealthBar(scene)
	}
}
