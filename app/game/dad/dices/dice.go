package dices

import "math/rand"

// -----------------------------------------------------------------------------
//
// IDice
//
// -----------------------------------------------------------------------------

// IDice interface defines all possible methods any dice struct should
// implement.
type IDice interface {
	GetName() string // returns the name of the dice.
	GetFaces() int   // returns the number of faces for the dice.
	Roll() int       // returns a roll dice.
}

// -----------------------------------------------------------------------------
//
// Dice
//
// -----------------------------------------------------------------------------

// Dice structure is the common and generic structure for any dice.
type Dice struct {
	name  string // dice name.
	faces int    // number of faces in the dice.
}

// NewDice function create a new Dice instance.
func NewDice(name string, faces int) *Dice {
	return &Dice{
		name:  name,
		faces: faces,
	}
}

// -----------------------------------------------------------------------------
// Dice public methods
// -----------------------------------------------------------------------------

// GetName method returns the name of the dice.
func (d *Dice) GetName() string {
	return d.name
}

// GetFaces method returns the numbed of faces for the dice.
func (d *Dice) GetFaces() int {
	return d.faces
}

// Roll method returns a roll dice.
func (d *Dice) Roll() int {
	return rand.Intn(d.GetFaces())
}

var _ IDice = (*Dice)(nil)
