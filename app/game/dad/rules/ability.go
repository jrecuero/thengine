package rules

// -----------------------------------------------------------------------------
//
// IAbility
//
// -----------------------------------------------------------------------------

// IAbility interface defines all possible methods for any ability.
type IAbility interface {
	GetName() string
	SetName(string)
	GetShortName() string
	SetShortName(string)
	GetDescription() string
	SetDescription(string)
	GetScore() int
	SetScore(int) bool
	GetExtra() int
	SetExtra(int)
	GetModifier() int
	GetScorePoint() int
}

// -----------------------------------------------------------------------------
//
// Ability
//
// -----------------------------------------------------------------------------

// Ability structure contains all attributes required to define an ability.
type Ability struct {
	name        string // ability name.
	shortName   string // ability short name.
	description string // ability description.
	score       int    // ability score.
	extra       int    // ability extra score.
}

// NewAbility function creates a new Ability instance.
func NewAbility(name, shortname string, score int) *Ability {
	return &Ability{
		name:      name,
		shortName: shortname,
		score:     score,
	}
}

// -----------------------------------------------------------------------------
// Ability public methods
// -----------------------------------------------------------------------------

// GetName method returns ability name.
func (a *Ability) GetName() string {
	return a.name
}

// SetName method sets ability name.
func (a *Ability) SetName(name string) {
	a.name = name
}

// GetShortName method returns ability short name.
func (a *Ability) GetShortName() string {
	return a.shortName
}

// SetShortName method sets ability short name.
func (a *Ability) SetShortName(name string) {
	a.shortName = name
}

// GetDescription method returns ability description.
func (a *Ability) GetDescription() string {
	return a.description
}

// SetDescription method sets ability description.
func (a *Ability) SetDescription(desc string) {
	a.description = desc
}

// GetScore method returns ability score value.
func (a *Ability) GetScore() int {
	score := a.score + a.GetExtra()
	if score > 30 {
		score = 30
	}
	return score
}

// SetScore method sets ability score value.
func (a *Ability) SetScore(score int) bool {
	if score < 1 || score > 30 {
		return false
	}
	a.score = score
	return true
}

// GetExtra method returns ability extra score value.
func (a *Ability) GetExtra() int {
	return a.extra
}

// SetExtra method sets ability extra score value.
func (a *Ability) SetExtra(extra int) {
	a.extra = extra
}

// GetModifier method returns the ability modifier value based on score.
//
// Ability modifiers are numerical values derived from a character's ability
// scores. Ability scores represent a character's innate abilities, such as
// strength, dexterity, constitution, intelligence, wisdom, and charisma, and
// are determined by rolling dice or using a point-buy system during character
// creation.
//
// The ability modifier is calculated by subtracting 10 from the ability score,
// dividing the result by 2 (rounding down), and then rounding down again. This
// means that an ability score of 10 or 11 has a modifier of +0, while an ability
// score of 12 or 13 has a modifier of +1, and so on.
func (a *Ability) GetModifier() int {
	modifier := 0
	switch a.GetScore() {
	case 1:
		modifier = -5
	case 2, 3:
		modifier = -4
	case 4, 5:
		modifier = -3
	case 6, 7:
		modifier = -2
	case 8, 9:
		modifier = -1
	case 10, 11:
		modifier = 0
	case 12, 13:
		modifier = 1
	case 14, 15:
		modifier = 2
	case 16, 17:
		modifier = 3
	case 18, 19:
		modifier = 4
	case 20, 21:
		modifier = 5
	case 22, 23:
		modifier = 6
	case 24, 25:
		modifier = 7
	case 26, 27:
		modifier = 8
	case 28, 29:
		modifier = 9
	case 30:
		modifier = 10
	}
	return modifier
}

// GetScorePoint method returns the number of points required to increse the
// ability score.
func (a *Ability) GetScorePoint() int {
	result := 0
	switch a.GetScore() {
	case 1, 2, 3, 4, 5, 6, 7, 8:
		result = 0
	case 9:
		result = 1
	case 10:
		result = 2
	case 11:
		result = 3
	case 12:
		result = 4
	case 13:
		result = 5
	case 14:
		result = 7
	case 15:
		result = 9
	default:
		result = 11
	}
	return result
}

// -----------------------------------------------------------------------------
//
// IAbilities
//
// -----------------------------------------------------------------------------

// IAbilities interfaces defines all abilities methods to be implemented.
type IAbilities interface {
	GetAbilityByName(string) IAbility
	GetConstitution() IAbility
	GetStrength() IAbility
	GetDexterity() IAbility
	GetIntelligence() IAbility
	GetWisdom() IAbility
	GetCharisma() IAbility
}

// -----------------------------------------------------------------------------
//
// Abilities
//
// -----------------------------------------------------------------------------

// Abilities struct contains all attributes and methods required all abilities
// required for any unit.
//
// There are six abilities, also known as ability scores, which represent a
// character's innate abilities and potential. They are:
//
// Constitution (CON): Measures a character's health, toughness, and endurance.
//
// Strength (STR): Measures a character's physical power, athletic ability, and
// raw physicality.
//
// Dexterity (DEX): Measures a character's agility, reflexes, and coordination.
//
// Intelligence (INT): Measures a character's knowledge, memory, and reasoning
// ability.
//
// Wisdom (WIS): Measures a character's perception, intuition, and insight.
//
// Charisma (CHA): Measures a character's personality, persuasiveness, and
// force of personality.
//
// Each ability score ranges from 1 to 20, with 10 representing an average score
// for a human. During character creation, players assign a score to each
// ability based on the character they are creating. These scores are then used
// to determine ability modifiers, which are used to calculate bonuses and
// penalties for various actions such as attack rolls, saving throws, and
// ability checks.
type Abilities struct {
	Constitution IAbility `json:"-"` // unit constitution.
	Strength     IAbility `json:"-"` // unit strength.
	Dexterity    IAbility `json:"-"` // unit dexterity.
	Intelligence IAbility `json:"-"` // unit intelligence.
	Wisdom       IAbility `json:"-"` // unit wisdom.
	Charisma     IAbility `json:"-"` // unit charisma.
}

// NewAbilities function creates a new Abilities instance.
func NewAbilities() *Abilities {
	return &Abilities{
		Constitution: NewAbility("constitution", "con", 0),
		Strength:     NewAbility("strength", "str", 0),
		Dexterity:    NewAbility("dexterity", "dex", 0),
		Intelligence: NewAbility("intelligence", "int", 0),
		Wisdom:       NewAbility("wisdom", "wis", 0),
		Charisma:     NewAbility("charisma", "char", 0),
	}
}

// -----------------------------------------------------------------------------
// Abilities public methods
// -----------------------------------------------------------------------------

// GetAbilityByName method return the ability for the given name.
func (a *Abilities) GetAbilityByName(name string) IAbility {
	result := (IAbility)(nil)
	switch name {
	case ConstitutionStr:
		result = a.Constitution
	case StrengthStr:
		result = a.Strength
	case DexterityStr:
		result = a.Dexterity
	case IntelligenceStr:
		result = a.Intelligence
	case WisdomStr:
		result = a.Wisdom
	case CharismaStr:
		result = a.Charisma
	}
	return result
}

// GetConstitution method returns constitution ability.
func (a *Abilities) GetConstitution() IAbility {
	return a.Constitution
}

// GetStrength method returns strength ability.
func (a *Abilities) GetStrength() IAbility {
	return a.Strength
}

// GetDexterity method returns dexterity ability.
func (a *Abilities) GetDexterity() IAbility {
	return a.Dexterity
}

// GetIntelligence method returns intelligence ability.
func (a *Abilities) GetIntelligence() IAbility {
	return a.Intelligence
}

// GetWisdom method returns wisdom ability.
func (a *Abilities) GetWisdom() IAbility {
	return a.Wisdom
}

// GetCharisma method returns charisma ability.
func (a *Abilities) GetCharisma() IAbility {
	return a.Charisma
}

var _ IAbility = (*Ability)(nil)
var _ IAbilities = (*Abilities)(nil)
