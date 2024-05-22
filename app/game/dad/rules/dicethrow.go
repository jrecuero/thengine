package rules

import "github.com/jrecuero/thengine/app/game/dad/dices"

// -----------------------------------------------------------------------------
//
// IDiceThrow
//
// -----------------------------------------------------------------------------

// IDiceThrow interface defines all possible methods for any dice throw.
type IDiceThrow interface {
	GetDescription() string
	GetDices() []dices.IDice
	GetExtra() int
	GetName() string
	Roll() int
	SetDescription(string)
	SetDices([]dices.IDice)
	SetName(string)
	SetExtra(int)
	SureRoll() int
}

// -----------------------------------------------------------------------------
//
// DiceThrow
//
// -----------------------------------------------------------------------------

// DiceThrow structure contains all attributes required to define an dice throw.
type DiceThrow struct {
	name        string        // dice throw name.
	shortName   string        // dice throw short name.
	description string        // dice throw description.
	dizes       []dices.IDice // dice throw score.
	extra       int           // dice throw extra score.
}

// NewDiceThrow function creates a new DiceThrow instance.
func NewDiceThrow(name, shortname string, dizes []dices.IDice) *DiceThrow {
	return &DiceThrow{
		name:      name,
		shortName: shortname,
		dizes:     dizes,
	}
}

// -----------------------------------------------------------------------------
// DiceThrow public methods
// -----------------------------------------------------------------------------

// GetDescription method returns dice throw description.
func (d *DiceThrow) GetDescription() string {
	return d.description
}

func (d *DiceThrow) GetDices() []dices.IDice {
	return d.dizes
}

// GetExtra method returns dice throw extra score value.
func (d *DiceThrow) GetExtra() int {
	return d.extra
}

// GetName method returns dice throw name.
func (d *DiceThrow) GetName() string {
	return d.name
}

// GetShortName method returns dice throw short name.
func (d *DiceThrow) GetShortName() string {
	return d.shortName
}

// Roll method returns dice throw score value.
func (d *DiceThrow) Roll() int {
	score := 0
	for _, dice := range d.dizes {
		score += dice.Roll()
	}
	score += d.GetExtra()
	//if score > 30 {
	//    score = 30
	//}
	return score
}

// SetDescription method sets dice throw description.
func (d *DiceThrow) SetDescription(desc string) {
	d.description = desc
}

func (d *DiceThrow) SetDices(dizes []dices.IDice) {
	d.dizes = dizes
}

// SetExtra method sets dice throw extra score value.
func (d *DiceThrow) SetExtra(extra int) {
	d.extra = extra
}

// SetName method sets dice throw name.
func (d *DiceThrow) SetName(name string) {
	d.name = name
}

// SetShortName method sets dice throw short name.
func (d *DiceThrow) SetShortName(name string) {
	d.shortName = name
}

func (d *DiceThrow) SureRoll() int {
	result := d.Roll()
	if result == 0 {
		result = 1
	}
	return result
}

var _ IDiceThrow = (*DiceThrow)(nil)
