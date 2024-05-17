package rules

// ISavingThrows interfaces defines all abilities methods to be implemented.
type ISavingThrows interface {
	GetThrowByName(string) IDiceThrow
	GetConstitution() IDiceThrow
	GetStrength() IDiceThrow
	GetDexteriry() IDiceThrow
	GetIntelligence() IDiceThrow
	GetWisdom() IDiceThrow
	GetCharisma() IDiceThrow
}

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
		constitution: NewDiceThrow("constitution", "con", 0),
		strength:     NewDiceThrow("strength", "str", 0),
		dexterity:    NewDiceThrow("dexterity", "dex", 0),
		intelligence: NewDiceThrow("intelligence", "int", 0),
		charisma:     NewDiceThrow("charisma", "char", 0),
		wisdom:       NewDiceThrow("wisdom", "wis", 0),
	}
}

// GetSavingThrowName method return the saving throw for the given name.
func (a *SavingThrows) GetSavingThrowName(name string) IDiceThrow {
	result := (IDiceThrow)(nil)
	switch name {
	case ConstitutionStr:
		result = a.constitution
	case StrengthStr:
		result = a.strength
	case DexterityStr:
		result = a.dexterity
	case IntelligenceStr:
		result = a.intelligence
	case WisdomStr:
		result = a.wisdom
	case CharismaStr:
		result = a.charisma
	}
	return result
}

// GetConstitution method returns constitution saving throw.
func (t *SavingThrows) GetConstitution() IDiceThrow {
	return t.constitution
}

// GetStrength method returns strength saving throw.
func (t *SavingThrows) GetStrength() IDiceThrow {
	return t.strength
}

// GetDexterity method returns dexterity saving throw.
func (t *SavingThrows) GetDexterity() IDiceThrow {
	return t.dexterity
}

// GetIntelligence method returns intelligence saving throw.
func (t *SavingThrows) GetIntelligence() IDiceThrow {
	return t.intelligence
}

// GetWisdom method returns wisdom saving throw.
func (t *SavingThrows) GetWisdom() IDiceThrow {
	return t.wisdom
}

// GetCharisma method returns charisma saving throw.
func (t *SavingThrows) GetCharisma() IDiceThrow {
	return t.charisma
}
