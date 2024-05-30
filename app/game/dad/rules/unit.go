package rules

import (
	"fmt"

	"github.com/jrecuero/thengine/app/game/dad/battlelog"
	"github.com/jrecuero/thengine/app/game/dad/dice"
	"github.com/jrecuero/thengine/pkg/tools"
)

const (
	Die20  = 20 // Die 20 (faces)
	BaseAC = 10 // Base Armor Class
)

// -----------------------------------------------------------------------------
//
// IUnit
//
// -----------------------------------------------------------------------------

// IUnit interface  defines all methods required a unit has to implement.
type IUnit interface {
	GetAbilities() IAbilities
	GetArmorClass() int // unit AC 10 + mod(dex) + mod(gear)
	GetAttacks() IAttacks
	GetDescription() string
	GetDieRoll() int
	GetGear() IGear
	GetHitDice() int // unit hit dice.
	GetHitting(IUnit) bool
	GetHitPoints() IHitPoints
	GetInitiative() int // unit initiative 1d20 + mod(dex)
	GetLanguages() any
	GetProficiencyBonus() int // unit proficiency bonus.
	GetSavingThrows() ISavingThrows
	GetSkills() any
	GetSpeed() int // unit speed
	GetSpells() any
	GetTraits() any
	GetUName() string
	RollAttack(IUnit) (bool, int)
	SetAbilities(IAbilities)
	SetAttacks(IAttacks)
	SetDescription(string)
	SetGear(IGear)
	SetHitPoints(IHitPoints)
	SetLanguages(any)
	SetSavingThrows(ISavingThrows)
	SetSkills(any)
	SetSpells(any)
	SetTraits(any)
	SetUName(string)
}

// -----------------------------------------------------------------------------
//
// Unit
//
// -----------------------------------------------------------------------------

// Unit structure is the common and generic structure for any unit in the
// application.
//
// A unit or character refers to a single individual, creature, or object that
// participates in the game world.
//
// Players create and control their own characters, which can be any of the
// playable races and classes available in the game. These characters have their
// own set of abilities, skills, hit points, and gear, and are used to
// navigate the world and interact with other characters, creatures, and objects.
//
// Non-player characters (NPCs) are controlled by the game's Dungeon Master (DM)
// and can be friendly or hostile to the players' characters. NPCs can include
// shopkeepers, quest-givers, villains, and monsters, among others.
//
// In combat, each unit or character takes turns based on their initiative roll,
// with the highest roll going first. During their turn, a character can move a
// certain distance, take actions such as attacking or casting spells, and use
// bonus actions or reactions as appropriate.
//
// Overall, the unit or character is the primary unit of play in D&D,
// representing a unique entity with its own attributes, abilities, and role in
// the game world.
type Unit struct {
	uname        string        // unit name
	description  string        // unit description.
	hitPoints    IHitPoints    // unit hit points.
	level        ILevel        // unit level.
	abilities    IAbilities    // unit abilities.
	savingThrows ISavingThrows // unit saving throws.
	skills       any           // unit skills
	gear         IGear         // unit gear
	attacks      IAttacks      // unit type of attacks
	spells       any           // unit spells
	languages    any           // unit languages
	traits       any           // unit personality traits
	dieRoll      IDiceThrow    // unit die roll
}

func NewUnit(name string) *Unit {
	die20 := dice.DieTwenty
	return &Unit{
		uname:        name,
		hitPoints:    NewHitPoints(0),
		level:        NewLevel(0, 0),
		abilities:    NewAbilities(),
		savingThrows: NewSavingThrows(),
		attacks:      NewAttacks(nil),
		dieRoll:      NewDiceThrow("dice-throw/die20", "dieroll", []dice.IDie{die20}),
		gear:         NewGear(),
	}
}

// -----------------------------------------------------------------------------
// Unit public methods
// -----------------------------------------------------------------------------

func (u *Unit) GetAbilities() IAbilities {
	return u.abilities
}

func (u *Unit) GetAttacks() IAttacks {
	return u.attacks
}

// GetArmorClass method returns the unit armor class.
//
// The Armor Class (AC) in represents a character's ability to dodge attacks
// and avoid taking damage. The formula for calculating AC varies depending on
// the character's gear and abilities, but generally consists of the
// following components:
//
// Base AC = 10 + Dexterity Modifier + Shield Bonus (if applicable)
//
// Base AC: All characters have a base AC of 10, which can be increased or
// decreased based on other factors.
//
// Dexterity Modifier: The character's Dexterity modifier is determined by
// their Dexterity score and is added to their base AC. A higher Dexterity
// score generally results in a higher AC.
//
// Shield Bonus: If a character is using a shield, the shield bonus is added to
// their AC. This bonus can range from +1 for a light shield to +2 for a heavy
// shield.
//
// Armor Bonus: The bonus provided by a character's armor is added to their AC.
// Different types of armor provide different bonuses, with heavier armor
// providing higher bonuses but also reducing the character's mobility.
//
// Other Bonuses: Certain abilities, spells, and other effects may provide
// additional bonuses to a character's AC. These bonuses should be recorded on
// the character sheet and added to the AC calculation as necessary.
func (u *Unit) GetArmorClass() int {
	dexModifier := u.GetAbilities().GetDexterity().GetModifier()
	result := BaseAC + dexModifier
	return result
}

func (u *Unit) GetDescription() string {
	return u.description
}

// GetDieRoll method returns the die roll for a physical or magical attack with
// all extras.
//
// Physical Attack:
// **Die Roll:** You roll a 20-sided die (d20). This represents the randomness
// of the attack itself.
// **Ability Modifier:** Add your ability modifier to the roll. This depends
// on the weapon used:
//   - **Strength:** Modifier used for melee weapons like swords, axes, and
//     maces.
//   - **Dexterity:** Modifier used for finesse weapons (which can use
//     Dexterity instead of Strength), ranged weapons like bows and crossbows,
//     and thrown weapons.
//
// **Proficiency Bonus:** If you're proficient with the weapon you're using,
// add your proficiency bonus to the roll. This bonus increases as you gain
// levels.
//
// Magical Attack:
// **Die Roll:** Similar to physical attacks, you still roll a d20.
// **Ability Modifier:** This depends on the type of magic being used:
//   - **Spellcasting Ability:** For spells, you add the modifier of your
//     spellcasting ability (Intelligence for wizards, Wisdom for druids, etc.).
//   - **Saving Throws:** Some magical effects require the target to make a
//     saving throw instead. Here, the target rolls a d20 and adds their relevant
//     ability modifier (e.g., Dexterity for dodging an fireball).
func (u *Unit) GetDieRoll() int {
	die20 := u.dieRoll.Roll()
	// TODO: Ability modifier FIXED to strength.
	strengthModifier := u.GetAbilities().GetStrength().GetModifier()
	result := die20 + strengthModifier
	battlelog.BLog.Push(fmt.Sprintf("[%s] die-roll %d+%d", u.GetUName(), die20, strengthModifier))
	return result
}

func (u *Unit) GetGear() IGear {
	return u.gear
}

// GetHitDice method returns the unit hit dice.
//
// The hit dice is a term used to refer to the dice a character rolls to
// determine the amount of hit points they gain when they level up. The type
// and number of hit dice a character has is determined by their class.
// For example, a fighter has 1d10 hit dice, while a wizard has 1d6 hit dice.
// When a character levels up, they roll their hit dice and add their
// Constitution modifier to the result to determine their hit point increase
// for that level.
func (u *Unit) GetHitDice() int {
	// TODO: To be implemented.
	return 1
}

// GetHitting method returns if the attack hits the other unit.
//
// Physical Attack:
// You hit the target if the **total roll (d20 + ability modifier + proficiency
// bonus)** is **equal to or higher than** the target's **Armor Class (AC)**.
// - AC represents the target's defense against physical attacks.
//
// Magical Attack:
// **Spells:** For spells that require an attack roll, you hit if the total
// roll (d20 + ability modifier + proficiency bonus (if applicable)) is equal
// to or higher than the target's AC.
// **Saving Throws:** The target succeeds on the saving throw if their total
// (d20 + ability modifier) is equal to or higher than the **spell save DC**
// set by the caster. The Spell Save DC is calculated based on the caster's
// level and spellcasting ability modifier.
func (u *Unit) GetHitting(unit IUnit) bool {
	// TODO: To be implemented.
	return false
}

func (u *Unit) GetHitPoints() IHitPoints {
	return u.hitPoints
}

// GetInitiative method return the unit initiative.
//
// Initiative is calculated based on a character's Dexterity modifier. The
// formula for calculating initiative is as follows:
//
// Initiative = d20 roll + Dexterity modifier
//
// At the start of a combat encounter, each participant rolls a d20 (a 20-sided
// die) and adds their Dexterity modifier to the result. The Dexterity modifier
// is determined by the character's Dexterity score and can be found on the
// character sheet.
//
// For example, if a character has a Dexterity score of 16, their Dexterity
// modifier would be +3. If they roll a 12 on their d20 initiative roll, their
// total initiative would be 15 (12 + 3). This would mean that they would act
// in the combat encounter before creatures with lower initiative scores.
func (u *Unit) GetInitiative() int {
	// TODO: to be implemented.
	return 1
}

func (u *Unit) GetLanguages() any {
	return u.languages
}

// GetProficiencyBonus method returns the unit proficiency bonus.
//
// The proficiency bonus is a bonus to a character's attack rolls, skill checks,
// and saving throws that is determined by the character's level. It starts at
// +2 for 1st level characters and increases by 1 every 4 levels thereafter (to
// a maximum of +6 at level 17).
func (u *Unit) GetProficiencyBonus() int {
	// TODO: To be implemented.
	return 1
}

func (u *Unit) GetSavingThrows() ISavingThrows {
	return u.savingThrows
}

func (u *Unit) GetSkills() any {
	return u.skills
}

// GetSpeed method returns the unit speed.
//
// A character's speed is the distance they can move in a single round of
// combat or during normal movement. The base speed for most races is 30 feet
// per round, but this can be increased or decreased based on factors such as
// race, class, gear, and other abilities.
func (u *Unit) GetSpeed() int {
	// TODO: to be implemented.
	return 1
}

func (u *Unit) GetSpells() any {
	return u.spells
}

func (u *Unit) GetTraits() any {
	return u.traits
}

func (u *Unit) GetUName() string {
	return u.uname
}

func (u *Unit) RollAttack(other IUnit) (bool, int) {
	dieRoll := u.GetDieRoll()
	ac := other.GetArmorClass()
	tools.Logger.WithField("module", "unit").WithField("method", "Attack").Debug(dieRoll, ac)
	// if die-roll is greater than the other unit armor class, it is a hit.
	if dieRoll < ac {
		battlelog.BLog.Push(fmt.Sprintf("[%s] no hit ac:%d", u.GetUName(), ac))
		return false, 0
	}
	// TODO: fix to used the first attack, usually attack/weapon
	weaponAttack := u.GetAttacks().GetAttacks()[0]
	damage := weaponAttack.Roll()
	otherHp := other.GetHitPoints().GetScore()
	battlelog.BLog.Push(fmt.Sprintf("[%s] hit ac:%d damage:%d", u.GetUName(), ac, damage))
	otherHp -= damage
	other.GetHitPoints().SetScore(otherHp)
	return true, damage
}

func (u *Unit) SetAbilities(abilities IAbilities) {
	u.abilities = abilities
}

func (u *Unit) SetAttacks(attacks IAttacks) {
	u.attacks = attacks
}

func (u *Unit) SetDescription(desc string) {
	u.description = desc
}

func (u *Unit) SetGear(gear IGear) {
	u.gear = gear
}

func (u *Unit) SetHitPoints(hp IHitPoints) {
	u.hitPoints = hp
}

func (u *Unit) SetLanguages(langs any) {
	u.languages = langs
}

func (u *Unit) SetSavingThrows(savingThrows ISavingThrows) {
	u.savingThrows = savingThrows
}

func (u *Unit) SetSkills(skills any) {
	u.skills = skills
}

func (u *Unit) SetSpells(spells any) {
	u.spells = spells
}

func (u *Unit) SetTraits(traits any) {
	u.traits = traits
}

func (u *Unit) SetUName(name string) {
	u.uname = name
}

var _ IUnit = (*Unit)(nil)
