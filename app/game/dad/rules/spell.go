// spell.go Package provides a structure and associated methods to represent
// and manipulate spells in a Dungeons & Dragons game. Each spell includes
// various attributes such as name, level, school of magic, casting time,
// range, components, duration, description, and damage type. The package also
// includes methods to simulate casting a spell and to generate a detailed
// description of the spell.
package rules

import "github.com/jrecuero/thengine/app/game/dad/constants"

// -----------------------------------------------------------------------------
// Module public types
// -----------------------------------------------------------------------------

type MagicSchool string

type MagicComponent string

// -----------------------------------------------------------------------------
// Module public constants
// -----------------------------------------------------------------------------

//const (
//    AbjurationMagic    MagicSchool = "abjuration"
//    ConjurationMagic   MagicSchool = "conjuration"
//    DivinationMagic    MagicSchool = "divination"
//    EnchantmentMagic   MagicSchool = "enchantment"
//    EvocationMagic     MagicSchool = "evocation"
//    IllusionMagic      MagicSchool = "illusion"
//    NecromancyMagic    MagicSchool = "necromacy"
//    TransmutationMagic MagicSchool = "transmutation"

//    VerbalComponent  MagicComponent = "verbal"
//    SomaticComponent MagicComponent = "somatic"
//    MaterialCompoent MagicComponent = "material"
//)

// -----------------------------------------------------------------------------
//
// ISpell
//
// -----------------------------------------------------------------------------

// ISpell interface provides all methods any spell have to be implementing.
type ISpell interface {
	DieRoll(IUnit) int
	GetCastingTime() int
	GetComponents() []MagicComponent
	GetDamage() IDamage
	GetDescription() string
	GetDuration() int
	GetHigherLevel() IDamage
	GetName() string
	GetSchool() MagicSchool
	GetRange() int
	GetUName() string
	GetLevel() int
	RollCast(IUnit) int
	SetCastingTime(int)
	SetComponents([]MagicComponent)
	SetDamage(IDamage)
	SetDescription(string)
	SetDuration(int)
	SetHigherLevel(IDamage)
	SetName(string)
	SetSchool(MagicSchool)
	SetRange(int)
	SetUName(string)
	SetLevel(int)
}

// -----------------------------------------------------------------------------
//
// Spell
//
// -----------------------------------------------------------------------------

// Spell structure  represents a spell in Dungeons & Dragons. It includes
// various attributes that define the spell's characteristics and behavior.
//
// Attributes:
//   - name: The name of the spell (e.g., "Fireball").
//   - description: A detailed description of the spell's effects.
//   - level: The spell level, ranging from 0 (cantrips) to 9 (high-level
//     spells).
//   - school: The school of magic to which the spell belongs (e.g., Evocation,
//     Illusion).
//   - castingTime: The time required to cast the spell (e.g., "1 action").
//   - srange: The effective range of the spell (e.g., "150 feet").
//   - components: The components required to cast the spell, including
//     Verbal(V), Somatic (S), and Material (M).
//   - duration: The duration for which the spell's effects last (e.g.,
//     "Instantaneous").
//   - higherLevel: Additional effects or changes when the spell is
//     cast using a higher-level spell slot.
//   - damage: The type of damage the spell deals (e.g., Fire, Cold).
type Spell struct {
	name        string
	uname       string
	description string
	school      MagicSchool
	castingTime int
	srange      int
	components  []MagicComponent
	duration    int
	level       int
	higherLevel IDamage
	damage      IDamage
}

// NewSpell function creates a new Spell instance.
func NewSpell(name string, uname string, school MagicSchool, level int, damage IDamage) *Spell {
	spell := &Spell{
		name:        name,
		uname:       uname,
		description: name,
		school:      school,
		castingTime: 0,
		srange:      0,
		components:  []MagicComponent{constants.Verbal},
		duration:    0,
		higherLevel: nil,
		damage:      damage,
	}
	return spell
}

// -----------------------------------------------------------------------------
// Spell public methods
// -----------------------------------------------------------------------------

// DieRoll method returns any additional value to add to the die roll related
// with the specific spell.
func (s *Spell) DieRoll(IUnit) int {
	return 0
}

func (s *Spell) GetCastingTime() int {
	return s.castingTime
}

func (s *Spell) GetComponents() []MagicComponent {
	return s.components
}

func (s *Spell) GetDamage() IDamage {
	return s.damage
}

func (s *Spell) GetDescription() string {
	return s.description
}

func (s *Spell) GetDuration() int {
	return s.duration
}

func (s *Spell) GetHigherLevel() IDamage {
	return s.higherLevel
}

func (s *Spell) GetName() string {
	return s.name
}

func (s *Spell) GetSchool() MagicSchool {
	return s.school
}

func (s *Spell) GetRange() int {
	return s.srange
}

func (s *Spell) GetUName() string {
	return s.uname
}

func (s *Spell) GetLevel() int {
	return s.level
}

func (s *Spell) RollCast(IUnit) int {
	if s.damage != nil {
		s.damage.RollDamageValue()
	}
	return 0
}

func (s *Spell) SetCastingTime(castingTime int) {
	s.castingTime = castingTime
}

func (s *Spell) SetComponents(components []MagicComponent) {
	s.components = components
}

func (s *Spell) SetDamage(damage IDamage) {
	s.damage = damage
}

func (s *Spell) SetDescription(description string) {
	s.description = description
}

func (s *Spell) SetDuration(duration int) {
	s.duration = duration
}

func (s *Spell) SetHigherLevel(higherLevel IDamage) {
	s.higherLevel = higherLevel
}

func (s *Spell) SetName(name string) {
	s.name = name
}

func (s *Spell) SetSchool(school MagicSchool) {
	s.school = school
}

func (s *Spell) SetRange(srange int) {
	s.srange = srange
}

func (s *Spell) SetUName(uname string) {
	s.uname = uname
}

func (s *Spell) SetLevel(level int) {
	s.level = level
}

var _ ISpell = (*Spell)(nil)
