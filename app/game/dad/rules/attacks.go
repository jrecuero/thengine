package rules

import "github.com/jrecuero/thengine/app/game/dad/dice"

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
	RollSavingThrows(IUnit) int
}

// -----------------------------------------------------------------------------
//
// Attack
//
// -----------------------------------------------------------------------------

// Attack struct defines the common and generic structure for any attack.
type Attack struct {
	name       string
	diceThrows []IDiceThrow
}

func NewAttack(name string, diceThrows []IDiceThrow) *Attack {
	return &Attack{
		name:       name,
		diceThrows: diceThrows,
	}
}

func NewDefaultAttack(score int) *Attack {
	die := dice.NewDie("dice/attack", score)
	diceThrows := NewDiceThrow("dice-throw/attack", "attack", []dice.IDie{die})
	attack := NewAttack("attack/default", []IDiceThrow{diceThrows})
	return attack
}

// -----------------------------------------------------------------------------
// Attack public methods
// -----------------------------------------------------------------------------

func (a *Attack) GetAttack() int {
	result := 0
	for _, diceThrow := range a.diceThrows {
		result += diceThrow.Roll()
	}
	return result
}

func (a *Attack) GetName() string {
	return a.name
}

func (a *Attack) Roll() int {
	result := 0
	for _, diceThrow := range a.diceThrows {
		result += diceThrow.SureRoll()
	}
	return result
}

func (a *Attack) RollSavingThrows(unit IUnit) int {
	return 0
}

var _ IAttack = (*Attack)(nil)

// -----------------------------------------------------------------------------
//
// WeaponAttack
//
// -----------------------------------------------------------------------------

type WeaponAttack struct {
	*Attack
	gear IGear
}

func NewWeaponAttack(gear IGear) *WeaponAttack {
	return &WeaponAttack{
		Attack: NewAttack("attack/weapon", nil),
		gear:   gear,
	}
}

// -----------------------------------------------------------------------------
// WeaponAttack public methods
// -----------------------------------------------------------------------------

func (a *WeaponAttack) Roll() int {
	//return a.gear.RollDamage()
	if a.gear.GetMainHand() != nil {
		return a.gear.GetMainHand().GetDamage().RollDamageValue()
	}
	return 0
}

func (a *WeaponAttack) RollSavingThrows(unit IUnit) int {
	if a.gear.GetMainHand() != nil {
		if savingThrows := a.gear.GetMainHand().GetSavingThrows(); savingThrows != nil {
			return a.gear.GetMainHand().GetDamage().RollSavingThrowsDamage(unit)
		}
	}
	return 0
}

var _ IAttack = (*WeaponAttack)(nil)

// -----------------------------------------------------------------------------
//
// PowerAttack
//
// -----------------------------------------------------------------------------

type PowerAttack struct {
	*Attack
	gear IGear
}

func NewPowerAttack(gear IGear) *PowerAttack {
	return &PowerAttack{
		Attack: NewAttack("attack/skill/power", nil),
		gear:   gear,
	}
}

// -----------------------------------------------------------------------------
// WeaponAttack public methods
// -----------------------------------------------------------------------------

func (a *PowerAttack) Roll() int {
	// Power attack rolls x2 weapon damage.
	if a.gear.GetMainHand() != nil {
		result := a.gear.GetMainHand().GetDamage().RollDamageValue()
		result += a.gear.GetMainHand().GetDamage().RollDamageValue()
		return result
	}
	return 0
}

func (a *PowerAttack) RollSavingThrows(unit IUnit) int {
	if a.gear.GetMainHand() != nil {
		if savingThrows := a.gear.GetMainHand().GetSavingThrows(); savingThrows != nil {
			return a.gear.GetMainHand().GetDamage().RollSavingThrowsDamage(unit)
		}
	}
	return 0
}

var _ IAttack = (*PowerAttack)(nil)

// -----------------------------------------------------------------------------
//
// MagicalAttack
//
// -----------------------------------------------------------------------------

type MagicalAttack struct {
	*Attack
	damage IDamage
	gear   IGear
}

func NewMagicalAttack(damage IDamage, gear IGear) *MagicalAttack {
	return &MagicalAttack{
		Attack: NewAttack("attack/magical", nil),
		damage: damage,
		gear:   gear,
	}
}

// -----------------------------------------------------------------------------
// MagicalAttack public methods
// -----------------------------------------------------------------------------

func (a *MagicalAttack) Roll() int {
	if a.damage != nil {
		return a.damage.RollDamageValue()
	}
	return 0
}

func (a *MagicalAttack) RollSavingThrows(unit IUnit) int {
	if a.damage != nil {
		if savingThrows := a.damage.GetSavingThrows(); savingThrows != nil {
			return a.damage.RollSavingThrowsDamage(unit)
		}
	}
	return 0
}

var _ IAttack = (*WeaponAttack)(nil)

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

var _ IAttacks = (*Attacks)(nil)
