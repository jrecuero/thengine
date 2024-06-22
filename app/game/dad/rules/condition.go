// condition.go package status provides structures and functions for
// managing status effects in a Dungeons & Dragons (D&D) game.
//
// Status effects represent various conditions that can affect characters,
// such as being poisoned, asleep, or paralyzed. This package includes methods
// to apply, remove, and check these effects on characters.
//
// # Structures
//
// ## Condition
// The Condition structure defines the properties of a status effect.
// - Name: The name of the status effect (e.g., "Poisoned", "Asleep").
// - Duration: The duration of the status effect in turns.
// - ApplyEffect: A function that applies the effect to a character.
// - RemoveEffect: A function that removes the effect from a character.
//
// ## Character
// The Character structure represents a player character or NPC in the game.
//   - Name: The name of the character.
//   - Health: The health points of the character.
//   - StatusEffects: A slice of Condition objects currently affecting the
//     character.
//
// # Functions
//
// ## NewStatusEffect
// `func NewStatusEffect(name string, duration int, apply, remove
// func(*Character)) Condition`
// Creates a new status effect with the given name, duration, and effect
// functions.
//
// ## ApplyStatusEffect
// `func (c *Character) ApplyStatusEffect(effect Condition)`
// Applies the given status effect to the character.
//
// ## RemoveStatusEffect
// `func (c *Character) RemoveStatusEffect(name string)`
// Removes the specified status effect from the character.
//
// ## CheckStatusEffect
// `func (c *Character) CheckStatusEffect(name string) bool`
// Checks if the character is currently affected by the specified status
// effect.
//
// Common status effects in D&D include:
//
//   - Blinded: A blinded creature can’t see and automatically fails any
//     ability check that requires sight.
//   - Charmed: A charmed creature can’t attack the charmer or target the
//     charmer with harmful abilities or magical effects.
//   - Deafened: A deafened creature can’t hear and automatically fails any
//     ability check that requires hearing.
//   - Frightened: A frightened creature has disadvantage on ability checks and
//     attack rolls while the source of its fear is within line of sight.
//   - Grappled: A grappled creature’s speed becomes 0, and it can’t benefit
//     from any bonus to its speed.
//   - Incapacitated: An incapacitated creature can’t take actions or reactions.
//   - Invisible: An invisible creature is impossible to see without the aid of
//     magic or a special sense.
//   - Paralyzed: A paralyzed creature is incapacitated and can’t move or speak.
//   - Petrified: A petrified creature is transformed, along with any nonmagical
//     object it is wearing or carrying, into a solid inanimate substance (usually
//     stone).
//   - Poisoned: A poisoned creature has disadvantage on attack rolls and ability
//     checks.
//   - Prone: A prone creature’s only movement option is to crawl, unless it
//     stands up and thereby ends the condition.
//   - Restrained: A restrained creature’s speed becomes 0, and it can’t benefit
//     from any bonus to its speed.
//   - Stunned: A stunned creature is incapacitated, can’t move, and can speak
//     only falteringly.
//   - Unconscious: An unconscious creature is incapacitated, can’t move or speak,
//     and is unaware of its surroundings.
package rules

import "github.com/jrecuero/thengine/app/game/dad/constants"

// -----------------------------------------------------------------------------
//
// ICondition
//
// -----------------------------------------------------------------------------

type ICondition interface {
	GetApply() func(IUnit) error
	GetDescription() string
	GetDiceThrow() IDiceThrow
	GetDuration() int
	GetName() string
	GetRemove() func(IUnit) error
	IsForever() bool
	RollDamage() int
	SetApply(func(IUnit) error)
	SetDescription(string)
	SetDiceThrow(IDiceThrow)
	SetDuration(int)
	SetName(string)
	SetRemove(func(IUnit) error)
}

// -----------------------------------------------------------------------------
//
// Condition
//
// -----------------------------------------------------------------------------

type Condition struct {
	apply       func(IUnit) error
	description string
	dicethrow   IDiceThrow
	duration    int
	name        string
	remove      func(IUnit) error
}

// NewCondition function creates a new Condition instance.
func NewCondition(name string, dicethrow IDiceThrow, duration int) *Condition {
	c := &Condition{
		apply:       nil,
		description: name,
		dicethrow:   dicethrow,
		duration:    duration,
		name:        name,
		remove:      nil,
	}
	return c
}

// -----------------------------------------------------------------------------
// Condition private methods
// -----------------------------------------------------------------------------

// -----------------------------------------------------------------------------
// Condition public methods
// -----------------------------------------------------------------------------

func (s *Condition) GetApply() func(IUnit) error {
	return s.apply
}

func (s *Condition) GetDescription() string {
	return s.description
}

func (s *Condition) GetDiceThrow() IDiceThrow {
	return s.dicethrow
}

func (s *Condition) GetDuration() int {
	return s.duration
}

func (s *Condition) GetName() string {
	return s.name
}

func (s *Condition) GetRemove() func(IUnit) error {
	return s.remove
}

func (s *Condition) IsForever() bool {
	return s.duration == constants.IsForever
}

func (s *Condition) RollDamage() int {
	result := 0
	if s.dicethrow != nil {
		result = s.dicethrow.Roll()
	}
	return result
}

func (s *Condition) SetApply(apply func(IUnit) error) {
	s.apply = apply
}

func (s *Condition) SetDescription(description string) {
	s.description = description
}

func (s *Condition) SetDiceThrow(dicethrow IDiceThrow) {
	s.dicethrow = dicethrow
}

func (s *Condition) SetDuration(duration int) {
	s.duration = duration
}

func (s *Condition) SetName(name string) {
	s.name = name
}

func (s *Condition) SetRemove(remove func(IUnit) error) {
	s.remove = remove
}

var _ ICondition = (*Condition)(nil)
