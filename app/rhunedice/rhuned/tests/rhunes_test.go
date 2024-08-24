package tests

import (
	"fmt"
	"testing"

	"github.com/jrecuero/thengine/app/rhunedice/rhuned/api"
	"github.com/jrecuero/thengine/app/rhunedice/rhuned/avatars"
	"github.com/jrecuero/thengine/app/rhunedice/rhuned/rhunes"
)

func TestClothesRhune(t *testing.T) {
	// create map from stat name to stat value
	statsmap := map[string]int{
		api.AttackName:  2,
		api.DefenseName: 1,
		api.SkillName:   3,
	}

	// create default avatar
	avatar := avatars.DefaultAvatar("test/1", statsmap)
	fmt.Println("default-avatar ", avatar)

	// create clothes-rhune
	clothesRhune := rhunes.ClothesRhune

	// execute clothes-rhune with the given avatar
	clothesRhune.Execute(avatar)

	avatar.StartTurn()
	fmt.Println("default-avatar ", avatar)
}
