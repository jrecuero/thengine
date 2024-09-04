package dicesets

import (
	"fmt"

	"github.com/jrecuero/thengine/app/rhunedice/rhuned/api"
	"github.com/jrecuero/thengine/app/rhunedice/rhuned/dice"
)

func NewFiveAndOneDiceSet(name string) *api.DiceSet {
	numbOfBaseDice := 5
	numbOfExtraDice := 1
	numbOfDice := numbOfBaseDice + numbOfExtraDice
	dices := make([]api.IDice, numbOfDice)
	for i := 0; i < numbOfBaseDice; i++ {
		diceName := fmt.Sprintf("base-dice/%d", i)
		dices[i] = dice.NewDefaultBaseDice(diceName)
	}
	for i := numbOfBaseDice; i < numbOfDice; i++ {
		diceName := fmt.Sprintf("extra-dice/%d", i)
		dices[i] = dice.NewDefaultExtraDice(diceName)
	}
	diceSetName := fmt.Sprintf("diceset/five-and-one/%s", name)
	diceset := api.NewDiceSet(diceSetName, dices)
	return diceset
}

func NewThreeDiceSet(name string) *api.DiceSet {
	numbOfBaseDice := 3
	dices := make([]api.IDice, numbOfBaseDice)
	for i := 0; i < numbOfBaseDice; i++ {
		diceName := fmt.Sprintf("base-dice/%d", i)
		dices[i] = dice.NewDefaultBaseDice(diceName)
	}
	diceSetName := fmt.Sprintf("diceset/three/%s", name)
	diceset := api.NewDiceSet(diceSetName, dices)
	return diceset
}
