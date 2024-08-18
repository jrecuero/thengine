package api

import "fmt"

type IDiceSet interface {
	GetDice() []IDice
	GetName() string
	Roll() []IFace
	SetDice([]IDice)
	SetName(string)
	String() string
}

type DiceSet struct {
	dice []IDice
	name string
}

func NewDiceSet(name string, dice []IDice) *DiceSet {
	return &DiceSet{
		dice: dice,
		name: name,
	}
}

func (d *DiceSet) GetDice() []IDice {
	return d.dice
}

func (d *DiceSet) GetName() string {
	return d.name
}

func (d *DiceSet) Roll() []IFace {
	result := make([]IFace, len(d.dice))
	for i, dice := range d.dice {
		result[i] = dice.Roll()
	}
	return result
}

func (d *DiceSet) SetDice(dice []IDice) {
	d.dice = dice
}

func (d *DiceSet) SetName(name string) {
	d.name = name
}

func (d *DiceSet) String() string {
	return fmt.Sprintf("%s %s", d.dice, d.name)
}

var _ IDiceSet = (*DiceSet)(nil)
