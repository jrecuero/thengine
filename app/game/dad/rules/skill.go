// skill.go modules defines all requirements for skills.
// skills represent a character's ability to perform various tasks and actions
// that might come up during gameplay. Each skill is tied to one of the six
// core abilities: Strength, Dexterity, Constitution, Intelligence, Wisdom, and
// Charisma. Skills are used to determine the outcome of actions that a
// character attempts, such as climbing a wall, deciphering ancient texts, or
// persuading a guard.
package rules

import "github.com/jrecuero/thengine/app/game/dad/constants"

// -----------------------------------------------------------------------------
// Module public types
// -----------------------------------------------------------------------------

type SkillName string

// -----------------------------------------------------------------------------
// Module public constants
// -----------------------------------------------------------------------------

//const (
//    // Acrobatics (Dexterity): Balancing, tumbling, and other agile movements.
//    AcrobaticsSK SkillName = constants.Acrobatics

//    // Animal Handling (Wisdom): Calming, controlling, or training animals.
//    AnimalHandlingSK SkillName = constants.AnimalHandling

//    // Arcana (Intelligence): Knowledge of magical lore, spells, and history.
//    ArcanaSK SkillName = constants.Arcana

//    // Athletics (Strength): Physical activities such as climbing, jumping, and
//    // swimming.
//    AthleticsSK SkillName = constants.Athletics

//    // Deception (Charisma): Lying, misleading, or tricking others.
//    DeceptionSK SkillName = constants.Deception

//    // History (Intelligence): Knowledge of historical events, people, and places.
//    HistorySK SkillName = constants.History

//    // Insight (Wisdom): Understanding the motives and emotions of others.
//    InsightSK SkillName = constants.Insight

//    // Intimidation (Charisma): Coercing or frightening others.
//    IntimidationSK SkillName = constants.Intimidation

//    // Investigation (Intelligence): Finding hidden objects, solving puzzles, and
//    // analyzing clues.
//    InvestigationSK SkillName = constants.Investigation

//    // Medicine (Wisdom): Providing medical care and diagnosing illnesses.
//    MedicineSK SkillName = constants.Medicine

//    // Nature (Intelligence): Knowledge of natural environments, flora, and fauna.
//    NatureSK SkillName = constants.Nature

//    // Perception (Wisdom): Noticing details and being aware of surroundings.
//    PerceptionSK SkillName = constants.Perception

//    // Performance (Charisma): Entertaining others through music, dance, acting,
//    // etc.
//    PerformanceSK SkillName = constants.Performance

//    // Persuasion (Charisma): Influencing others through diplomacy and charm.
//    PersuasionSK SkillName = constants.Persuasion

//    // Religion (Intelligence): Knowledge of deities, religious practices, and
//    // theology.
//    ReligionSK SkillName = constants.Religion

//    // Sleight of Hand (Dexterity): Picking pockets, palming objects, and other
//    // dexterous acts.
//    SleightSK SkillName = constants.Sleight

//    // Stealth (Dexterity): Sneaking, hiding, and avoiding detection.
//    StealthSK SkillName = constants.Stealth

//    // Survival (Wisdom): Tracking, foraging, and enduring harsh environments.
//    SurvivalSK SkillName = constants.Survival
//)

// -----------------------------------------------------------------------------
// Module public functions
// -----------------------------------------------------------------------------

func CreateSkills() []ISkill {
	skills := []ISkill{
		NewSkill(constants.Acrobatics, constants.Dexterity, 0),
		NewSkill(constants.AnimalHandling, constants.Wisdom, 0),
		NewSkill(constants.Arcana, constants.Intelligence, 0),
		NewSkill(constants.Athletics, constants.Strength, 0),
		NewSkill(constants.Deception, constants.Charisma, 0),
		NewSkill(constants.History, constants.Intelligence, 0),
		NewSkill(constants.Insight, constants.Wisdom, 0),
		NewSkill(constants.Intimidation, constants.Charisma, 0),
		NewSkill(constants.Investigation, constants.Intelligence, 0),
		NewSkill(constants.Medicine, constants.Wisdom, 0),
		NewSkill(constants.Nature, constants.Intelligence, 0),
		NewSkill(constants.Perception, constants.Wisdom, 0),
		NewSkill(constants.Performance, constants.Charisma, 0),
		NewSkill(constants.Persuasion, constants.Charisma, 0),
		NewSkill(constants.Religion, constants.Intelligence, 0),
		NewSkill(constants.Sleight, constants.Dexterity, 0),
		NewSkill(constants.Stealth, constants.Dexterity, 0),
		NewSkill(constants.Survival, constants.Wisdom, 0),
	}
	return skills
}

// -----------------------------------------------------------------------------
//
// ISkill
//
// -----------------------------------------------------------------------------

// ISkill interface provides all methods any skill have to implement.
type ISkill interface {
	GetAbility() AbilityScore
	GetDescription() string
	GetName() SkillName
	GetProficienty() int
	SetAbility(AbilityScore)
	SetDescription(string)
	SetName(SkillName)
	SetProficienty(int)
}

// -----------------------------------------------------------------------------
//
// Skill
//
// -----------------------------------------------------------------------------

// Skill structure represents all attributes and methods for any generic skill.
type Skill struct {
	name        SkillName
	description string
	ability     AbilityScore
	proficiency int
}

// NewSkill function creates a new Skill instance.
func NewSkill(name SkillName, ability AbilityScore, proficiency int) *Skill {
	skill := &Skill{
		name:        name,
		description: string(name),
		ability:     ability,
		proficiency: proficiency,
	}
	return skill
}

// -----------------------------------------------------------------------------
// Skill public methods
// -----------------------------------------------------------------------------

func (s *Skill) GetAbility() AbilityScore {
	return s.ability
}

func (s *Skill) GetDescription() string {
	return s.description
}

func (s *Skill) GetName() SkillName {
	return s.name
}

func (s *Skill) GetProficienty() int {
	return s.proficiency
}

func (s *Skill) SetAbility(ability AbilityScore) {
	s.ability = ability
}

func (s *Skill) SetDescription(description string) {
	s.description = description
}

func (s *Skill) SetName(name SkillName) {
	s.name = name
}

func (s *Skill) SetProficienty(proficiency int) {
	s.proficiency = proficiency
}

var _ ISkill = (*Skill)(nil)
