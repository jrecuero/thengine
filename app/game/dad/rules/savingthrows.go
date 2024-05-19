package rules

// -----------------------------------------------------------------------------
//
// ISavingThrows
//
// -----------------------------------------------------------------------------

// ISavingThrows interfaces defines all abilities methods to be implemented.
type ISavingThrows interface {
	GetSavingThrowByName(string) IDiceThrow
	GetConstitution() IDiceThrow
	GetStrength() IDiceThrow
	GetDexterity() IDiceThrow
	GetIntelligence() IDiceThrow
	GetWisdom() IDiceThrow
	GetCharisma() IDiceThrow
}

// -----------------------------------------------------------------------------
//
// SavingThrows
//
// -----------------------------------------------------------------------------

// SavingThrows struct contains all attributes and methods required all saving
// throws required for any unit.
//
// Saving throws, also known as saves, are rolls made by characters to resist
// the effects of spells, traps, poisons, and other hazards that can harm or
// incapacitate them.
//
// When a character is subjected to an effect that allows a saving throw, they
// roll a d20 and add their relevant saving throw bonus to the result. If the
// total is equal to or greater than the DC (Difficulty Class) of the effect,
// the character succeeds the saving throw and avoids or reduces the effect.
//
// There are three types of saving throws in D&D:
//
// Strength Saving Throws (STR): used to resist physical effects such as
// grappling, shoving, or being pushed back.
//
// Dexterity Saving Throws (DEX): used to resist effects that require quick
// reflexes, such as dodging a trap or avoiding a spell.
//
// Constitution Saving Throws (CON): used to resist effects that target a
// character's health, such as poisons or diseases.
//
// In addition to these three primary types of saving throws, some effects may
// require a character to make a saving throw based on their Intelligence (INT),
// Wisdom (WIS), or Charisma (CHA) scores. The rules for each effect specify
// which type of saving throw is required.
type SavingThrows struct {
	constitution IDiceThrow // unit constitution.
	strength     IDiceThrow // unit strength.
	dexterity    IDiceThrow // unit dexterity.
	intelligence IDiceThrow // unit intelligence.
	wisdom       IDiceThrow // unit wisdom.
	charisma     IDiceThrow // unit charisma.
}

// NewSavingThrows function creates a new SavingThrows instance.
func NewSavingThrows() *SavingThrows {
	return &SavingThrows{
		constitution: NewDiceThrow("constitution", "con", nil),
		strength:     NewDiceThrow("strength", "str", nil),
		dexterity:    NewDiceThrow("dexterity", "dex", nil),
		intelligence: NewDiceThrow("intelligence", "int", nil),
		charisma:     NewDiceThrow("charisma", "char", nil),
		wisdom:       NewDiceThrow("wisdom", "wis", nil),
	}
}

// -----------------------------------------------------------------------------
// SavingThrows public methods
// -----------------------------------------------------------------------------

// GetSavingThrowName method return the saving throw for the given name.
func (s *SavingThrows) GetSavingThrowByName(name string) IDiceThrow {
	result := (IDiceThrow)(nil)
	switch name {
	case ConstitutionStr:
		result = s.constitution
	case StrengthStr:
		result = s.strength
	case DexterityStr:
		result = s.dexterity
	case IntelligenceStr:
		result = s.intelligence
	case WisdomStr:
		result = s.wisdom
	case CharismaStr:
		result = s.charisma
	}
	return result
}

// GetConstitution method returns constitution saving throw.
func (s *SavingThrows) GetConstitution() IDiceThrow {
	return s.constitution
}

// GetStrength method returns strength saving throw.
func (s *SavingThrows) GetStrength() IDiceThrow {
	return s.strength
}

// GetDexterity method returns dexterity saving throw.
func (s *SavingThrows) GetDexterity() IDiceThrow {
	return s.dexterity
}

// GetIntelligence method returns intelligence saving throw.
func (s *SavingThrows) GetIntelligence() IDiceThrow {
	return s.intelligence
}

// GetWisdom method returns wisdom saving throw.
func (s *SavingThrows) GetWisdom() IDiceThrow {
	return s.wisdom
}

// GetCharisma method returns charisma saving throw.
func (s *SavingThrows) GetCharisma() IDiceThrow {
	return s.charisma
}

var _ ISavingThrows = (*SavingThrows)(nil)
