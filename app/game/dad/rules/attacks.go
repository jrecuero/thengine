package rules

// -----------------------------------------------------------------------------
// Module public types
// -----------------------------------------------------------------------------

type AttackType string

// -----------------------------------------------------------------------------
// Module public constants
// -----------------------------------------------------------------------------

const (
	WeaponMeleeAttackType       AttackType = "attack/weapon/melee"
	WeaponRangedAttackType      AttackType = "attack/weapon/ranged"
	SpellMeleeAttackType        AttackType = "attack/spell/melee"
	SpellRangedAttackType       AttackType = "attack/spell/ranged"
	SpellAreaOfEffectAttackType AttackType = "attack/spell/aoe"
	SpecialAttackType           AttackType = "attack/special"
)

// -----------------------------------------------------------------------------
//
// IAttackRolls
//
// -----------------------------------------------------------------------------

// IAttackRolls interface defines all required methods for any custom attack
// roll to be used as an Attack.
type IAttackRoll interface {
	DieRoll(IUnit) int
	Roll(IUnit) int
	RollSavingThrows(IUnit) int
}

// -----------------------------------------------------------------------------
//
// IAttack
//
// -----------------------------------------------------------------------------

// IAttack interface defines all methods any attack structure should be
// implementing.
type IAttack interface {
	IAttackRoll
	GetName() string
	SetGear(IGear)
	SetSpecial(IAttackRoll)
	SetSpell(ISpell)
}

// -----------------------------------------------------------------------------
//
// Attack
//
// -----------------------------------------------------------------------------

// Attack struct defines the common and generic structure for any attack.
type Attack struct {
	name    string
	atype   AttackType
	gear    IGear
	spell   ISpell
	special IAttackRoll
}

func NewAttack(name string) *Attack {
	return &Attack{
		name:    name,
		atype:   SpecialAttackType,
		gear:    nil,
		spell:   nil,
		special: nil,
	}
}

func NewWeaponMeleeAttack(name string, gear IGear) *Attack {
	attack := NewAttack(name)
	attack.atype = WeaponMeleeAttackType
	attack.gear = gear
	return attack
}

func NewWeaponRangedAttack(name string, gear IGear) *Attack {
	attack := NewAttack(name)
	attack.atype = WeaponRangedAttackType
	attack.gear = gear
	return attack
}

func NewSpellMeleeAttack(name string, spell ISpell) *Attack {
	attack := NewAttack(name)
	attack.atype = SpellMeleeAttackType
	attack.spell = spell
	return attack
}

func NewSpellRangedAttack(name string, spell ISpell) *Attack {
	attack := NewAttack(name)
	attack.atype = SpellRangedAttackType
	attack.spell = spell
	return attack
}

func NewSpellAreaOfEffectAttack(name string, spell ISpell) *Attack {
	attack := NewAttack(name)
	attack.atype = SpellAreaOfEffectAttackType
	attack.spell = spell
	return attack
}

func NewSpecialAttack(name string, special IAttackRoll) *Attack {
	attack := NewAttack(name)
	attack.atype = SpecialAttackType
	attack.special = special
	return attack
}

// -----------------------------------------------------------------------------
// Attack private methods
// -----------------------------------------------------------------------------

// dieRoll method returns any additional value to add to the die roll related
// with an specific weapon attack.
func (a *Attack) dieRoll(IUnit) int {
	result := 0
	if a.gear.GetMainHand() != nil {
	}
	if a.gear.GetOffHand() != nil {
	}
	return result
}

func (a *Attack) mainHandWeaponRoll() int {
	if a.gear.GetMainHand() != nil {
		return a.gear.GetMainHand().GetDamage().RollDamageValue()
	}
	return 0
}

func (a *Attack) offHandWeaponRoll() int {
	if a.gear.GetOffHand() != nil {
		return a.gear.GetOffHand().GetDamage().RollDamageValue()
	}
	return 0
}

// weaponRoll method returns the weapon damage for the specific weapon attack.
func (a *Attack) weaponRoll(IUnit) int {
	result := 0
	if a.gear.GetMainHand() != nil {
		result += a.gear.GetMainHand().GetDamage().RollDamageValue()
	}
	if a.gear.GetOffHand() != nil {
		result += a.gear.GetOffHand().GetDamage().RollDamageValue()
	}
	return result
}

func (a *Attack) mainHandWeaponRollSavingThrows(unit IUnit) int {
	if a.gear.GetMainHand() != nil {
		return a.gear.GetMainHand().GetDamage().RollSavingThrowsDamage(unit)
	}
	return 0
}

func (a *Attack) offHandWeaponRollSavingThrows(unit IUnit) int {
	if a.gear.GetOffHand() != nil {
		return a.gear.GetOffHand().GetDamage().RollSavingThrowsDamage(unit)
	}
	return 0
}

// weaponRollSavingThrows method returns the weapon saving throws damage for
// the specific weapon attack.
func (a *Attack) weaponRollSavingThrows(unit IUnit) int {
	result := 0
	if a.gear.GetMainHand() != nil {
		result += a.gear.GetMainHand().GetDamage().RollSavingThrowsDamage(unit)
	}
	if a.gear.GetOffHand() != nil {
		result += a.gear.GetOffHand().GetDamage().RollSavingThrowsDamage(unit)
	}
	return result
}

// -----------------------------------------------------------------------------
// Attack public methods
// -----------------------------------------------------------------------------

func (a *Attack) DieRoll(unit IUnit) int {
	result := 0
	switch a.atype {
	case WeaponMeleeAttackType:
		fallthrough
	case WeaponRangedAttackType:
		if a.gear != nil {
			result += a.dieRoll(unit)
		}
	case SpellMeleeAttackType:
		fallthrough
	case SpellRangedAttackType:
		fallthrough
	case SpellAreaOfEffectAttackType:
		if a.spell != nil {
			result += a.spell.DieRoll(unit)
		}
	case SpecialAttackType:
		if a.special != nil {
			result += a.special.DieRoll(unit)
		}
	}
	return result
}

func (a *Attack) GetName() string {
	return a.name
}

func (a *Attack) Roll(unit IUnit) int {
	result := 0
	switch a.atype {
	case WeaponMeleeAttackType:
		fallthrough
	case WeaponRangedAttackType:
		if a.gear != nil {
			result += a.weaponRoll(unit)
		}
	case SpellMeleeAttackType:
		fallthrough
	case SpellRangedAttackType:
		fallthrough
	case SpellAreaOfEffectAttackType:
		if a.spell != nil {
			result += a.spell.RollCast(unit)
		}
	case SpecialAttackType:
		if a.special != nil {
			result += a.special.Roll(unit)
		}
	}
	return result
}

func (a *Attack) RollSavingThrows(unit IUnit) int {
	result := 0
	switch a.atype {
	case WeaponMeleeAttackType:
		fallthrough
	case WeaponRangedAttackType:
		if a.gear != nil {
			result += a.weaponRollSavingThrows(unit)
		}
	case SpellMeleeAttackType:
		fallthrough
	case SpellRangedAttackType:
		fallthrough
	case SpellAreaOfEffectAttackType:
		if a.spell != nil {
		}
	case SpecialAttackType:
		if a.special != nil {
			result += a.special.RollSavingThrows(unit)
		}
	}
	return result
}

func (a *Attack) SetGear(gear IGear) {
	a.gear = gear
}

func (a *Attack) SetSpecial(special IAttackRoll) {
	a.special = special
}

func (a *Attack) SetSpell(spell ISpell) {
	a.spell = spell
}

var _ IAttack = (*Attack)(nil)

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
