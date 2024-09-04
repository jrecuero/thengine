package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/jrecuero/thengine/app/rhunedice/rhuned/api"
	"github.com/jrecuero/thengine/app/rhunedice/rhuned/avatars"
)

var (
	turnHandler *api.TurnHandler
	reader      = bufio.NewReader(os.Stdin)
)

func waitForKeyPressed(message string, key byte) {
	fmt.Println(message)
	_, _ = reader.Discard(reader.Buffered())
	reader.ReadBytes(key)
}

func selectBuckets() (int, int) {
	fmt.Println("select 2 buckets...")
	var result1, result2 int
	fmt.Scan(&result1, &result2)
	//reader.Discard(reader.Buffered())
	return result1, result2
}

func setupTurn() {
	// create player stats map.
	playerStatsMap := map[string]int{
		api.AttackName:  2,
		api.DefenseName: 1,
		api.SkillName:   3,
		api.HealthName:  10,
	}

	// create player avatar.
	player := avatars.DefaultAvatar("player/1", playerStatsMap)
	player.SetActive(true)

	// create enemy stats map.
	enemyStatsMap := map[string]int{
		api.AttackName:  1,
		api.DefenseName: 1,
		api.SkillName:   1,
		api.HealthName:  5,
	}

	// create enemy avatar.
	enemy := avatars.DefaultEnemy("enemy/1", enemyStatsMap)
	enemy.SetActive(true)

	// create turn handler.
	turnHandler = api.NewTurnHandler(player, []api.IAvatar{enemy})

	fmt.Println("player ", turnHandler.GetPlayer())
	fmt.Println("enmeies ", turnHandler.GetEnemies())
	fmt.Println()

	turnHandler.SetState(api.InitTurn)

	// turn-handler init -> start
	turnHandler.Run()
	fmt.Println("turn-handler state ", turnHandler.GetState())
}

func runTurn() bool {

	player := turnHandler.GetPlayer()
	enemy := turnHandler.GetEnemies()[0]

	// turn-handler start -> avatar-update
	turnHandler.Run()
	fmt.Println("turn-handler state ", turnHandler.GetState())

	// turn-handler avatar-update -> avatar-update-bucket
	turnHandler.Run()
	fmt.Println("turn-handler state ", turnHandler.GetState())

	// turn-handler avatar-update-bucket -> roll-dice
	turnHandler.Run()
	fmt.Println("turn-handler state ", turnHandler.GetState())

	fmt.Println("player buckets", player.GetBuckets())
	fmt.Println("enemy buckets", enemy.GetBuckets())
	fmt.Println()

	waitForKeyPressed("Press 'r' to roll dice..", 'r')

	// turn-handler roll-dice -> bucket-selection
	turnHandler.Run()
	fmt.Println("turn-handler state ", turnHandler.GetState())

	// player: from the roll-dice-buckets, select first and second in the slice
	// as the user-selected buckets.
	if buckets := player.GetRollDiceBuckets(); buckets != nil {
		//limit := tools.Min(2, len(buckets))
		//selected := buckets[0:limit]
		//player.SetSelected(selected)
		//fmt.Println("player select buckets ", selected)
		fmt.Println("player select buckets ", buckets)
		for i, b := range buckets {
			fmt.Printf("[%d] %s\n", i, b)
		}
		index1, index2 := selectBuckets()
		selected := []api.IBucket{buckets[index1], buckets[index2]}
		player.SetSelected(selected)
		fmt.Println("player select buckets ", selected)
	}

	// enemy: from the roll-dice buckets, select first in the slice.
	if buckets := enemy.GetRollDiceBuckets(); buckets != nil {
		selected := buckets[:1]
		enemy.SetSelected(selected)
		fmt.Println("enemy select buckets ", selected)
	}

	// turn-handler bucket-selection -> execute-bucket
	turnHandler.Run()
	fmt.Println("turn-handler state ", turnHandler.GetState())

	// turn-handler execute-bucket -> end
	turnHandler.Run()
	fmt.Println("turn-handler state ", turnHandler.GetState())

	// turn-handler end -> start
	turnHandler.Run()
	fmt.Println("turn-handler state ", turnHandler.GetState())

	fmt.Println("player buckets", player.GetBuckets())
	fmt.Println("enemy buckets", enemy.GetBuckets())

	waitForKeyPressed("Press 'n' to end turn...", 'n')

	return (enemy.GetBuckets().GetBucketByName(api.HealthName).GetValue() > 0)
}

func main() {
	fmt.Println("turner")
	setupTurn()
	for run := true; run; run = runTurn() {
	}
	fmt.Println("game over")
}
