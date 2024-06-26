package rules

import (
	"fmt"

	"github.com/jrecuero/thengine/app/game/dad/battlelog"
	"github.com/jrecuero/thengine/app/game/dad/constants"
	"github.com/jrecuero/thengine/app/game/dad/dice"
	"github.com/jrecuero/thengine/pkg/tools"
)

const (
	Die20  = 20 // Die 20 (faces)
	BaseAC = 10 // Base Armor Class
)

// -----------------------------------------------------------------------------
// Package private functions
// -----------------------------------------------------------------------------

// -----------------------------------------------------------------------------
//
// IUnit
//
// -----------------------------------------------------------------------------

// IUnit interface  defines all methods required a unit has to implement.
type IUnit interface {
	AddActivable(IActivable)
	AddCondition(ICondition) error
	EndTurn()
	GetAbilities() IAbilities
	GetArmorClass() int // unit AC 10 + mod(dex) + mod(gear)
	GetAttacks() IAttacks
	GetClass() IClass
	GetConditions() []ICondition
	GetConditionResistances() map[string]int
	GetDescription() string
	GetDieRoll() int
	GetFeats() []IFeat
	GetGear() IGear
	GetHitDice() int // unit hit dice.
	GetHitting(IUnit) bool
	GetHitPoints() IHitPoints
	GetInitiativeRoll() int // unit initiative 1d20 + mod(dex)
	GetInventory() IInventory
	GetLanguages() []ILanguage
	GetLevel() ILevel
	GetProficiencies() []IProficiency
	GetProficiencyBonus() int // unit proficiency bonus.
	GetSkills() []ISkill
	GetSpeed() int // unit speed
	GetSpells() []ISpell
	GetRace() IRace
	GetTraits() []ITrait
	GetUName() string
	Populate(map[string]any, map[string]any)
	RemoveCondition(ICondition) error
	//RollAttack(int, IUnit) (bool, int)
	RollDamage(int, IUnit) (bool, int)
	RollDieRoll(int, IUnit) bool
	RollConditions() []int
	SetAbilities(IAbilities)
	SetAttacks(IAttacks)
	SetClass(IClass)
	SetConditions([]ICondition)
	SetConditionResistances(map[string]int)
	SetDescription(string)
	SetFeats([]IFeat)
	SetGear(IGear)
	SetHitPoints(IHitPoints)
	SetInventory(IInventory)
	SetLanguages([]ILanguage)
	SetLevel(ILevel)
	SetProficiencies([]IProficiency)
	SetRace(IRace)
	SetSkills([]ISkill)
	SetSpells([]ISpell)
	SetTraits([]ITrait)
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
	abilities            IAbilities   // unit abilities.
	activables           []IActivable // unit activables in a turn.
	attacks              IAttacks     // unit type of attacks
	class                IClass
	conditions           []ICondition   // unit conditions or status effects
	conditionResistances map[string]int // unit condition resistances
	description          string         // unit description.
	dieRoll              IDiceThrow     // unit die roll
	feats                []IFeat
	gear                 IGear       // unit gear
	hitPoints            IHitPoints  // unit hit points.
	inventory            IInventory  // unit inventory
	languages            []ILanguage // unit languages
	level                ILevel      // unit level.
	proficiencies        []IProficiency
	race                 IRace
	skills               []ISkill // unit skills
	spells               []ISpell // unit spells
	traits               []ITrait // unit personality traits
	uname                string   // unit name
}

func NewUnit(name string) *Unit {
	die20 := dice.DieTwenty
	return &Unit{
		abilities:            NewAbilities(),
		attacks:              NewAttacks(nil),
		class:                nil,
		conditions:           nil,
		conditionResistances: make(map[string]int),
		description:          name,
		dieRoll:              NewDiceThrow("dice-throw/die20", "dieroll", []dice.IDie{die20}),
		feats:                nil,
		gear:                 NewGear(),
		hitPoints:            NewHitPoints(0),
		inventory:            NewInventory("inventory/1"),
		languages:            nil,
		level:                NewLevel(0, 0, 0),
		proficiencies:        nil,
		race:                 nil,
		skills:               CreateSkills(),
		spells:               nil,
		traits:               nil,
		uname:                name,
	}
}

// -----------------------------------------------------------------------------
// Unit private methods
// -----------------------------------------------------------------------------

// getRollBonusForAction method checks in unit feats, proficiencies and traits
// if there is any of them impacting the giving action.
// TODO: condition/status effects are still missing in this check.
func (u *Unit) getRollBonusForAction(action string) []any {
	var result []any

	// Check weapons (handheld gear)

	// Check armor

	// Check accessories

	// Check any item

	// Check feats.
	feats := u.getFeatsRollBonusForAction(action)
	result = append(result, feats...)

	// Check proficiencies
	proficiencies := u.getProficiencyRollBonusForAction(action)
	result = append(result, proficiencies...)

	// Check traits
	traits := u.getTraitRollBonusForAction(action)
	result = append(result, traits...)

	return result
}

// checkProcessingToInt method process al entries in a slice of check where the
// final result should be an integer.
func (u *Unit) checkProcessingToInt(checks ...any) int {
	result := 0
	for _, check := range checks {
		switch check.(type) {
		case int:
			result += check.(int)
		case func(IUnit) int:
			result += check.(func(IUnit) int)(u)
		default:
		}
	}
	return result
}

// getFeatsRollBonusForAction method returns any bonus to be applied to any
// roll for every player character feat.
func (u *Unit) getFeatsRollBonusForAction(action string) []any {
	var result []any
	for _, feat := range u.feats {
		if r := feat.GetRollBonusForAction(action); r != nil {
			result = append(result, r)
		}
	}
	return result
}

// getConditionsRollBonusForAction method returns any bonus to be applied to
// any roll for every player character condition.
func (u *Unit) getConditionsRollBonusForAction(condition ICondition) int {
	if resistance, ok := u.conditionResistances[condition.GetName()]; ok {
		return resistance
	}
	return 0
}

// getProficiencyRollBonusForAction method returns any bonus to be applied to any
// roll for every player character proficiency.
func (u Unit) getProficiencyRollBonusForAction(action string) []any {
	var result []any
	for _, proficiency := range u.proficiencies {
		if r := proficiency.GetRollBonusForAction(action); r != nil {
			result = append(result, r)
		}
	}
	return result
}

// getTraitRollBonusForAction method returns any bonus to be applied to any roll for
// every player character trait.
func (u Unit) getTraitRollBonusForAction(action string) []any {
	var result []any
	for _, trait := range u.traits {
		if r := trait.GetRollBonusForAction(action); r != nil {
			result = append(result, r)
		}
	}
	return result
}

// -----------------------------------------------------------------------------
// Unit public methods
// -----------------------------------------------------------------------------

func (u *Unit) AddActivable(act IActivable) {
	u.activables = append(u.activables, act)
}

// AddCondition method adds given condition to the unit.
func (u *Unit) AddCondition(condition ICondition) error {
	if apply := condition.GetApply(); apply != nil {
		if err := apply(u); err != nil {
			return err
		}
	}
	u.conditions = append(u.conditions, condition)
	return nil
}

func (u *Unit) EndTurn() {
	for _, act := range u.activables {
		act.Clean()
	}
	u.activables = nil
}

// GetAbilities method returns unit abilities.
func (u *Unit) GetAbilities() IAbilities {
	return u.abilities
}

// GetAttacks method returns unit attacks.
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
	result := BaseAC + dexModifier + u.GetGear().AC()
	return result
}

// GetClass method returns the unit class.
func (u *Unit) GetClass() IClass {
	return u.class
}

func (u *Unit) GetConditions() []ICondition {
	return u.conditions
}

func (u *Unit) GetConditionResistances() map[string]int {
	return u.conditionResistances
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
	//strModifier := u.GetAbilities().GetStrength().GetModifier()
	strModifier := tools.NilToInt(u.GetAbilities().GetRollBonusForAction(constants.SavingThrowRollStrength))
	weaponStrModifier := 0
	if u.GetGear() != nil {
		weaponStrModifier = tools.NilToInt(u.GetGear().GetRollBonusForAction(constants.SavingThrowRollStrength))
	}
	result := die20 + strModifier + weaponStrModifier
	battlelog.BLog.PushDebug(fmt.Sprintf("[%s] die-roll %d+%d+%d", u.GetUName(), die20, strModifier, weaponStrModifier))
	tools.Logger.WithField("module", "unit").
		WithField("method", "GetDieRoll").
		Debugf("[%s] die-roll %d+%d+%d", u.GetUName(), die20, strModifier, weaponStrModifier)
	return result
}

// GetFeats method returns all unit feats.
func (u *Unit) GetFeats() []IFeat {
	feats := u.feats
	if u.race != nil && u.race.GetFeats() != nil {
		feats = append(feats, u.race.GetFeats()...)
	}
	if u.class != nil && u.class.GetFeats() != nil {
		feats = append(feats, u.class.GetFeats()...)
	}
	return feats
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

// GetInitiativeRoll method return the unit initiative.
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
func (u *Unit) GetInitiativeRoll() int {
	initiative := u.dieRoll.Roll()
	initiative += u.GetAbilities().GetDexterity().GetModifier()

	// TODO: any feats, traits or proficiencies related with initiative have to
	// be applied now.
	checks := u.getRollBonusForAction(constants.InitiativeRoll)
	tools.Logger.WithField("module", "unit").
		WithField("method", "GetInitiativeRoll").
		Tracef("initiative checks %+#v", checks)
	checksResult := u.checkProcessingToInt(checks...)
	initiative += checksResult
	return initiative
}

// GetInventory method returns the unit inventory.
func (u *Unit) GetInventory() IInventory {
	return u.inventory
}

// GetLanguages method returns the unit languages.
func (u *Unit) GetLanguages() []ILanguage {
	return u.languages
}

// GetLevel method returns the unit Level instance.
func (u *Unit) GetLevel() ILevel {
	return u.level
}

// GetProficiencies method returns all unit proficiencies.
func (u *Unit) GetProficiencies() []IProficiency {
	return u.proficiencies
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

// GetRace method returns unit race.
func (u *Unit) GetRace() IRace {
	return u.race
}

// GetSkills method returns unit skills.
func (u *Unit) GetSkills() []ISkill {
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

// GetSpells method return unit spells.
func (u *Unit) GetSpells() []ISpell {
	return u.spells
}

// GetTraits method returns unit treats.
func (u *Unit) GetTraits() []ITrait {
	traits := u.traits
	if u.race != nil && u.race.GetTraits() != nil {
		traits = append(traits, u.race.GetTraits()...)
	}
	if u.class != nil && u.class.GetTraits() != nil {
		traits = append(traits, u.class.GetTraits()...)
	}
	return traits
}

// GetUName method returns unit name.
func (u *Unit) GetUName() string {
	return u.uname
}

// Populate methods populate the unit with the given default values and content
// passed as a map[string]any.
func (u *Unit) Populate(defaults map[string]any, content map[string]any) {
	var hp int = defaults["hp"].(int)
	var strength int = defaults["strength"].(int)
	var dexterity int = defaults["dexterity"].(int)
	var constitution int = defaults["constitution"].(int)
	var intelligence int = defaults["intelligence"].(int)
	var wisdom int = defaults["wisdom"].(int)
	var charisma int = defaults["charisma"].(int)
	if _hp, ok := content["hp"].(float64); ok {
		hp = int(_hp)
	}
	if _strength, ok := content["strength"].(float64); ok {
		strength = int(_strength)
	}
	if _dexterity, ok := content["dexterity"].(float64); ok {
		dexterity = int(_dexterity)
	}
	if _constitution, ok := content["constitution"].(float64); ok {
		constitution = int(_constitution)
	}
	if _intelligence, ok := content["intelligence"].(float64); ok {
		intelligence = int(_intelligence)
	}
	if _wisdom, ok := content["wisdom"].(float64); ok {
		wisdom = int(_wisdom)
	}
	if _charisma, ok := content["charisma"].(float64); ok {
		charisma = int(_charisma)
	}
	u.GetHitPoints().SetMaxScore(hp)
	u.GetHitPoints().SetScore(hp)
	u.GetAbilities().GetStrength().SetScore(strength)
	u.GetAbilities().GetDexterity().SetScore(dexterity)
	u.GetAbilities().GetConstitution().SetScore(constitution)
	u.GetAbilities().GetIntelligence().SetScore(intelligence)
	u.GetAbilities().GetWisdom().SetScore(wisdom)
	u.GetAbilities().GetCharisma().SetScore(charisma)
}

// RemoveCondition method removes the given condition/status effect from the
// list of conditions in the unit.
func (u *Unit) RemoveCondition(condition ICondition) error {
	for i, unitCondition := range u.conditions {
		if unitCondition == condition {
			if remove := condition.GetRemove(); remove != nil {
				if err := remove(u); err != nil {
					return err
				}
			}
			u.conditions = append(u.conditions[:i], u.conditions[i+1:]...)
			break
		}
	}
	return nil
}

// RollConditions method roll damages for every condition applied to the unit.
func (u *Unit) RollConditions() []int {
	var result []int
	for _, condition := range u.conditions {
		resistance := u.getConditionsRollBonusForAction(condition)
		conditionRoll := condition.RollDamage()
		damage := (conditionRoll * resistance) / 100
		result = append(result, damage)
	}
	return result
}

func (u *Unit) RollDamage(index int, other IUnit) (bool, int) {
	attack := u.GetAttacks().GetAttacks()[index]
	attackIcon := ""
	switch index {
	case 0:
		attackIcon = "🗡"
	case 1:
		attackIcon = "⚒"
	case 2:
		attackIcon = "⚡"
	}
	damage := attack.Roll(u)
	stDamage := attack.RollSavingThrows(other)
	otherHp := other.GetHitPoints().GetScore()
	damage += stDamage
	// TODO: any feats, traits or proficiencies related with damage have to
	// be applied now.
	damageChecks := u.getRollBonusForAction(constants.DamageRoll)
	damageChecksResult := u.checkProcessingToInt(damageChecks...)
	damage += damageChecksResult
	otherHp -= damage
	other.GetHitPoints().SetScore(otherHp)
	//battlelog.BLog.PushInfo(fmt.Sprintf("[%s] %s roll:%dvs%d🛡️%d⚔%d⚁", u.GetUName(), attack.GetName(), dieRoll, ac, damage, stDamage))
	battlelog.BLog.PushInfo(fmt.Sprintf("[%s]\t%s\t%d⚔%d⚁", u.GetUName(), attackIcon, damage, stDamage))
	return true, damage
}

func (u *Unit) RollDieRoll(index int, other IUnit) bool {
	dieRoll := u.GetDieRoll()
	// Call attack die roll for any die roll effect provided by the attack that
	// can involve activation of any feat, trait, proficiency, ...
	attack := u.GetAttacks().GetAttacks()[index]
	tools.Logger.WithField("module", "unit").
		WithField("method", "RollAttack").
		Debugf("%+#v", attack)
	dieRoll += attack.DieRoll(u)
	// TODO: any feats, traits or proficiencies related with die roll have to
	// be applied now.
	dieRollChecks := u.getRollBonusForAction(constants.SavingThrowRoll)
	dieRollChecksResult := u.checkProcessingToInt(dieRollChecks...)
	dieRoll += dieRollChecksResult
	ac := other.GetArmorClass()
	tools.Logger.WithField("module", "unit").
		WithField("method", "RollAttack").
		Debug(dieRoll, ac)
	// if die-roll is greater than the other unit armor class, it is a hit.
	if dieRoll < ac {
		//battlelog.BLog.PushInfo(fmt.Sprintf("[%s] miss!\troll:%dvs%d🛡️", u.GetUName(), dieRoll, ac))
		battlelog.BLog.PushInfo(fmt.Sprintf("[%s]\t❌\troll:%dvs%d🛡️", u.GetUName(), dieRoll, ac))
		return false
	}
	battlelog.BLog.PushInfo(fmt.Sprintf("[%s]\troll:%dvs%d🛡️", u.GetUName(), dieRoll, ac))
	return true
}

// SetAbilities method sets unit abilitities.
func (u *Unit) SetAbilities(abilities IAbilities) {
	u.abilities = abilities
}

// SetAttacks method sets unit attacks.
func (u *Unit) SetAttacks(attacks IAttacks) {
	u.attacks = attacks
}

// SetClass method sets a new class to the unit.
func (u *Unit) SetClass(class IClass) {
	u.class = class
}

// SetConditions method sets unit conditions/status effects.
func (u *Unit) SetConditions(conditions []ICondition) {
	u.conditions = conditions
}

// SetConditionResistances method sets unit resistances to multiple
// conditions/status effects.
func (u *Unit) SetConditionResistances(resitances map[string]int) {
	u.conditionResistances = resitances
}

// SetDescription method sets the unit description.
func (u *Unit) SetDescription(desc string) {
	u.description = desc
}

// SetFeat method sets a new slice of feats for the unit.
func (u *Unit) SetFeats(feats []IFeat) {
	u.feats = feats
}

// SetGear method sets the unit gear.
func (u *Unit) SetGear(gear IGear) {
	u.gear = gear
}

// SetHitPoints method sets the unit hit points.
func (u *Unit) SetHitPoints(hp IHitPoints) {
	u.hitPoints = hp
}

// SetInventory method sets the unit inventory.
func (u *Unit) SetInventory(inventory IInventory) {
	u.inventory = inventory
}

// SetLanguages method sets the unit languages.
func (u *Unit) SetLanguages(langs []ILanguage) {
	u.languages = langs
}

// SetLevel method sets the unit Level instance.
func (u *Unit) SetLevel(level ILevel) {
	u.level = level
}

// SetProficiencies method sets the given slice of proficiencies.
func (u *Unit) SetProficiencies(proficiencies []IProficiency) {
	u.proficiencies = proficiencies
}

// SetRace method sets a new race for the unit.
func (u *Unit) SetRace(race IRace) {
	u.race = race
}

func (u *Unit) SetSkills(skills []ISkill) {
	u.skills = skills
}

func (u *Unit) SetSpells(spells []ISpell) {
	u.spells = spells
}

func (u *Unit) SetTraits(traits []ITrait) {
	u.traits = traits
}

func (u *Unit) SetUName(name string) {
	u.uname = name
}

var _ IUnit = (*Unit)(nil)
