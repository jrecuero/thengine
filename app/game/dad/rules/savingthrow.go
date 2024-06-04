package rules

import (
	"fmt"

	"github.com/jrecuero/thengine/pkg/tools"
)

// -----------------------------------------------------------------------------
//
// ISavingThrow
//
// -----------------------------------------------------------------------------

// ISavingThrow interfaces defines all abilities methods to be implemented.
type ISavingThrow interface {
	GetAS() AbilityScore
	GetDC() int
	Pass(IUnit) bool
	Roll() int
	SetAS(AbilityScore)
	SetDC(int)
	ToString() string
}

// -----------------------------------------------------------------------------
//
// SavingThrow
//
// -----------------------------------------------------------------------------

// SavingThrow struct contains all attributes and methods required all saving
// throw required for any unit.
//
// Saving throw, also known as saves, are rolls made by characters to resist
// the effects of spells, traps, poisons, and other hazards that can harm or
// incapacitate them.
//
// When a character is subjected to an effect that allows a saving throw, they
// roll a d20 and add their relevant saving throw bonus to the result. If the
// total is equal to or greater than the DC (Difficulty Class) of the effect,
// the character succeeds the saving throw and avoids or reduces the effect.
//
// There are three types of saving throw in D&D:
//
// Strength Saving Throw (STR): used to resist physical effects such as
// grappling, shoving, or being pushed back.
//
// Dexterity Saving Throw (DEX): used to resist effects that require quick
// reflexes, such as dodging a trap or avoiding a spell.
//
// Constitution Saving Throw (CON): used to resist effects that target a
// character's health, such as poisons or diseases.
//
// In addition to these three primary types of saving throw, some effects may
// require a character to make a saving throw based on their Intelligence (INT),
// Wisdom (WIS), or Charisma (CHA) scores. The rules for each effect specify
// which type of saving throw is required.
type SavingThrow struct {
	as        AbilityScore // ability score
	dc        int          // dificulty class
	diceThrow IDiceThrow   // dice throw: 1d20
}

// NewSavingThrow function creates a new SavingThrow instance.
func NewSavingThrow(as AbilityScore, dc int) *SavingThrow {
	return &SavingThrow{
		as:        as,
		dc:        dc,
		diceThrow: DiceThrow1d20,
	}
}

// -----------------------------------------------------------------------------
// SavingThrow public methods
// -----------------------------------------------------------------------------

func (s *SavingThrow) GetAS() AbilityScore {
	return s.as
}

func (s *SavingThrow) GetDC() int {
	return s.dc
}

// Pass method checks if the given unit is able to pass the saving throw for
// the required difficulty class.
func (s *SavingThrow) Pass(unit IUnit) bool {
	score := unit.GetAbilities().GetAbilityByName(s.GetAS()).GetScorePoint()
	roll := s.Roll()
	tools.Logger.WithField("module", "savingThrow").
		WithField("method", "Pass").
		Debugf("saving throw roll:%d score:%d dc:%d", roll, score, s.dc)
	return (roll + score) > s.GetDC()
}

func (s *SavingThrow) Roll() int {
	return s.diceThrow.Roll()
}

func (s *SavingThrow) SetAS(as AbilityScore) {
	s.as = as
}

func (s *SavingThrow) SetDC(dc int) {
	s.dc = dc
}

func (s *SavingThrow) ToString() string {
	return fmt.Sprintf("saving-throw %s:%d %s", s.as, s.dc, s.diceThrow.ToString())
}

var _ ISavingThrow = (*SavingThrow)(nil)
