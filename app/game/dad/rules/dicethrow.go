package rules

import "github.com/jrecuero/thengine/app/game/dad/dice"

var (
	DiceThrow1d2  = NewDiceThrow("dice-throw/1d2", "1d2", []dice.IDie{dice.DieTwo})
	DiceThrow1d3  = NewDiceThrow("dice-throw/1d3", "1d3", []dice.IDie{dice.DieThree})
	DiceThrow1d4  = NewDiceThrow("dice-throw/1d4", "1d4", []dice.IDie{dice.DieThree})
	DiceThrow1d5  = NewDiceThrow("dice-throw/1d5", "1d5", []dice.IDie{dice.DieThree})
	DiceThrow1d6  = NewDiceThrow("dice-throw/1d6", "1d6", []dice.IDie{dice.DieSix})
	DiceThrow1d8  = NewDiceThrow("dice-throw/1d8", "1d8", []dice.IDie{dice.DieEight})
	DiceThrow1d10 = NewDiceThrow("dice-throw/1d10", "1d10", []dice.IDie{dice.DieTen})
	DiceThrow1d12 = NewDiceThrow("dice-throw/1d12", "1d12", []dice.IDie{dice.DieTwelve})
	DiceThrow2d6  = NewDiceThrow("dice-throw/2d6", "2d6", []dice.IDie{dice.DieSix, dice.DieSix})
)

// -----------------------------------------------------------------------------
//
// IDiceThrow
//
// -----------------------------------------------------------------------------

// IDiceThrow interface defines all possible methods for any dice throw.
type IDiceThrow interface {
	GetDescription() string
	GetDices() []dice.IDie
	GetExtra() int
	GetName() string
	Roll() int
	SetDescription(string)
	SetDices([]dice.IDie)
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
	name        string      // dice throw name.
	shortName   string      // dice throw short name.
	description string      // dice throw description.
	dices       []dice.IDie // dice throw score.
	extra       int         // dice throw extra score.
}

// NewDiceThrow function creates a new DiceThrow instance.
func NewDiceThrow(name string, shortname string, dices []dice.IDie) *DiceThrow {
	return &DiceThrow{
		name:      name,
		shortName: shortname,
		dices:     dices,
	}
}

// -----------------------------------------------------------------------------
// DiceThrow public methods
// -----------------------------------------------------------------------------

// GetDescription method returns dice throw description.
func (d *DiceThrow) GetDescription() string {
	return d.description
}

func (d *DiceThrow) GetDices() []dice.IDie {
	return d.dices
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
	for _, dice := range d.dices {
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

func (d *DiceThrow) SetDices(dices []dice.IDie) {
	d.dices = dices
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
