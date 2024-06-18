package dice

import (
	"fmt"

	"github.com/jrecuero/thengine/pkg/tools"
)

var (
	DieOne       = NewDie("die/one", 1)
	DieTwo       = NewDie("die/two", 2)
	DieThree     = NewDie("die/three", 3)
	DieFour      = NewDie("die/four", 4)
	DieFive      = NewDie("die/five", 5)
	DieSix       = NewDie("die/six", 6)
	DieSeven     = NewDie("die/seven", 7)
	DieEight     = NewDie("die/eight", 8)
	DieNine      = NewDie("die/nine", 9)
	DieTen       = NewDie("die/ten", 10)
	DieEleven    = NewDie("die/eleven", 11)
	DieTwelve    = NewDie("die/twelve", 12)
	DieThirdteen = NewDie("die/thirdteen", 13)
	DieFourteen  = NewDie("die/fourteen", 14)
	DieFifteen   = NewDie("die/fifteen", 15)
	DieSixteen   = NewDie("die/sixteen", 16)
	DieSeventeen = NewDie("die/seventeen", 17)
	DieEighteen  = NewDie("die/eighteen", 18)
	DieNineteen  = NewDie("die/nineteen", 19)
	DieTwenty    = NewDie("die/twenty", 20)
)

// -----------------------------------------------------------------------------
//
// IDie
//
// -----------------------------------------------------------------------------

// IDie interface defines all possible methods any die struct should
// implement.
type IDie interface {
	AdvantageRoll() int        // returns an advantage roll
	AdvantageSureRoll() int    // returns an advantage roll without zero.
	DisadvantageRoll() int     // return a disadvantage roll
	DisadvantageSureRoll() int // return a disadvantage roll without zero.
	DoubleRoll() (int, int)
	DoubleSureRoll() (int, int)
	GetName() string // returns the name of the die.
	GetFaces() int   // returns the number of faces for the die.
	Roll() int       // returns a roll die.
	SureRoll() int   // returns a roll die without a zero
	ToString() string
}

// -----------------------------------------------------------------------------
//
// Die
//
// -----------------------------------------------------------------------------

// Die structure is the common and generic structure for any die.
type Die struct {
	name     string // die name.
	faces    int    // number of faces in the die.
	loaded   int    // loaded die value
	isLoaded bool   // is a loaded die that always return same value
}

// NewDie function create a new Die instance.
func NewDie(name string, faces int) *Die {
	return &Die{
		name:     name,
		faces:    faces,
		isLoaded: false,
		loaded:   0,
	}
}

// NewLoadedDie create a new Die instance that always return the same value.
func NewLoadedDie(name string, loaded int) *Die {
	return &Die{
		name:     name,
		faces:    0,
		isLoaded: true,
		loaded:   loaded,
	}
}

// -----------------------------------------------------------------------------
// Die public methods
// -----------------------------------------------------------------------------

// AdvantageRoll method rolls two dice and returns the highest value between
// [0-faces].
func (d *Die) AdvantageRoll() int {
	rollOne, rollTwo := d.DoubleRoll()
	return tools.Max(rollOne, rollTwo)
}

// AdvantageSureRoll method rolls two dice and returns the highest value
// between [1-faces].
func (d *Die) AdvantageSureRoll() int {
	rollOne, rollTwo := d.DoubleSureRoll()
	return tools.Max(rollOne, rollTwo)
}

// DisadvantageRoll method rolls two dice and returns the lowest value between
// [0-faces].
func (d *Die) DisadvantageRoll() int {
	rollOne, rollTwo := d.DoubleRoll()
	return tools.Min(rollOne, rollTwo)
}

// DisadvantageSureRoll method rolls two dice and returns the lowest value
// between [1-faces].
func (d *Die) DisadvantageSureRoll() int {
	rollOne, rollTwo := d.DoubleSureRoll()
	return tools.Min(rollOne, rollTwo)
}

// DoubleRoll method rolls two dice between [0-faces].
func (d *Die) DoubleRoll() (int, int) {
	rollOne := d.Roll()
	rollTwo := d.Roll()
	return rollOne, rollTwo
}

// DoubleSureRoll method rolls two dice without zero [1-faces].
func (d *Die) DoubleSureRoll() (int, int) {
	rollOne := d.SureRoll()
	rollTwo := d.SureRoll()
	return rollOne, rollTwo
}

// GetName method returns the name of the die.
func (d *Die) GetName() string {
	return d.name
}

// GetFaces method returns the numbed of faces for the die.
func (d *Die) GetFaces() int {
	return d.faces
}

// Roll method returns a roll die [0-faces]
func (d *Die) Roll() int {
	if d.isLoaded {
		return d.loaded
	}
	return tools.RandomRing.Intn(d.GetFaces() + 1)
}

// SureRoll method returns a roll die without a zero [1-faces]
func (d *Die) SureRoll() int {
	if d.isLoaded {
		return d.loaded
	}
	return tools.RandomRing.Intn(d.GetFaces()) + 1
}

// ToString method returns the dice struct as a string.
func (d *Die) ToString() string {
	return fmt.Sprintf("die %s %d/%d", d.name, d.faces, d.loaded)
}

var _ IDie = (*Die)(nil)
