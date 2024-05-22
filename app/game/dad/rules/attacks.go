package rules

import "github.com/jrecuero/thengine/app/game/dad/dices"

// -----------------------------------------------------------------------------
//
// IAttack
//
// -----------------------------------------------------------------------------

// IAttack interface defines all methods any attack structure should be
// implementing.
type IAttack interface {
	GetAttack() int
	GetName() string
	Roll() int
}

// -----------------------------------------------------------------------------
//
// Attack
//
// -----------------------------------------------------------------------------

// Attack struct defines the common and generic structure for any attack.
type Attack struct {
	name      string
	diceThrow IDiceThrow
}

func NewAttack(name string, diceThrow IDiceThrow) *Attack {
	return &Attack{
		name:      name,
		diceThrow: diceThrow,
	}
}

func NewDefaultAttack(score int) *Attack {
	dice := dices.NewDice("dice/attack", score)
	diceThrow := NewDiceThrow("dice-throw/attack", "attack", []dices.IDice{dice})
	attack := NewAttack("attack/default", diceThrow)
	return attack
}

// -----------------------------------------------------------------------------
// Attack public methods
// -----------------------------------------------------------------------------

func (a *Attack) GetAttack() int {
	attack := a.diceThrow.Roll()
	return attack
}

func (a *Attack) GetName() string {
	return a.name
}

func (a *Attack) Roll() int {
	return a.diceThrow.SureRoll()
}

// -----------------------------------------------------------------------------
//
// IAttacks
//
// -----------------------------------------------------------------------------

// IAttacks interface defines all method required to handle a set of attacks.
type IAttacks interface {
	AddAttack(IAttack)
	GetAttackByName(string) IAttack
	GetAttacks() []IAttack
	RemoveAttack(IAttack)
}

// -----------------------------------------------------------------------------
//
// Attacks
//
// -----------------------------------------------------------------------------

// Attacks structure defines the basic attributes and methods to handle a set
// of attacks.
type Attacks struct {
	attacks []IAttack
}

func NewAttacks(attacks []IAttack) *Attacks {
	return &Attacks{
		attacks: attacks,
	}
}

// -----------------------------------------------------------------------------
// Attacks public methods
// -----------------------------------------------------------------------------

func (a *Attacks) AddAttack(attack IAttack) {
	a.attacks = append(a.attacks, attack)
}

func (a *Attacks) GetAttackByName(name string) IAttack {
	for _, attack := range a.attacks {
		if attack.GetName() == name {
			return attack
		}
	}
	return nil
}

func (a *Attacks) GetAttacks() []IAttack {
	return a.attacks
}

func (a *Attacks) RemoveAttack(attack IAttack) {
	for index, att := range a.attacks {
		if att == attack {
			a.attacks = append(a.attacks[:index], a.attacks[index+1:]...)
			return
		}
	}
}

var _ IAttack = (*Attack)(nil)
var _ IAttacks = (*Attacks)(nil)
