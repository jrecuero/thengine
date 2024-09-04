package tests

import (
	"fmt"
	"testing"

	"github.com/jrecuero/thengine/app/rhunedice/rhuned/api"
	"github.com/jrecuero/thengine/app/rhunedice/rhuned/avatars"
	"github.com/jrecuero/thengine/pkg/tools"
)

func TestTurn(t *testing.T) {
	// create player stats map.
	playerStatsMap := map[string]int{
		api.AttackName:  2,
		api.DefenseName: 1,
		api.SkillName:   3,
	}

	// create player avatar.
	player := avatars.DefaultAvatar("player/1", playerStatsMap)
	player.SetActive(true)

	// create enemy stats map.
	enemyStatsMap := map[string]int{
		api.AttackName:  1,
		api.DefenseName: 1,
		api.SkillName:   1,
	}

	// create enemy avatar.
	enemy := avatars.DefaultEnemy("enemy/1", enemyStatsMap)
	enemy.SetActive(true)

	// create turn handler.
	turnHandler := api.NewTurnHandler(player, []api.IAvatar{enemy})

	fmt.Println("player ", turnHandler.GetPlayer())
	fmt.Println("enmeies ", turnHandler.GetEnemies())

	turnHandler.SetState(api.InitTurn)

	// turn-handler init -> start
	turnHandler.Run()
	fmt.Println("turn-handler state ", turnHandler.GetState())

	// turn-handler start -> avatar-update
	turnHandler.Run()
	fmt.Println("turn-handler state ", turnHandler.GetState())

	// turn-handler avatar-update -> avatar-update-bucket
	turnHandler.Run()
	fmt.Println("turn-handler state ", turnHandler.GetState())

	// turn-handler avatar-update-bucket -> roll-dice
	turnHandler.Run()
	fmt.Println("turn-handler state ", turnHandler.GetState())

	// turn-handler roll-dice -> bucket-selection
	turnHandler.Run()
	fmt.Println("turn-handler state ", turnHandler.GetState())

	// player: from the roll-dice-buckets, select first and second in the slice
	// as the user-selected buckets.
	if buckets := player.GetRollDiceBuckets(); buckets != nil {
		limit := tools.Min(2, len(buckets))
		selected := buckets[0:limit]
		player.SetSelected(selected)
		fmt.Println("player select buckets ", selected)
	}

	// enemy: from the roll-dice buckets, select first in the slice.
	if buckets := enemy.GetRollDiceBuckets(); buckets != nil {
		selected := buckets[:1]
		enemy.SetSelected(selected)
		fmt.Println("enemy select buckets ", selected)
	}

	fmt.Println("player buckets", player.GetBuckets())
	fmt.Println("enemy buckets", enemy.GetBuckets())

	// turn-handler bucket-selection -> execute-bucket
	turnHandler.Run()
	fmt.Println("turn-handler state ", turnHandler.GetState())

	// turn-handler execute-bucket -> end
	turnHandler.Run()
	fmt.Println("turn-handler state ", turnHandler.GetState())

	// turn-handler end -> start
	turnHandler.Run()
	fmt.Println("turn-handler state ", turnHandler.GetState())
}
