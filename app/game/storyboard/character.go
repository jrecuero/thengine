// character.go package contains all data and logic related with any character
// involved in the storyboard.
package storyboard

// -----------------------------------------------------------------------------
//
// ICharacter
//
// -----------------------------------------------------------------------------

type ICharacter interface {
	GetAge() int
	GetDescription() string
	GetName() string
	GetRole() string
	SetAge(int)
	SetDescription(string)
	SetName(string)
	SetRole(string)
}

// -----------------------------------------------------------------------------
//
// Character
//
// -----------------------------------------------------------------------------

type Character struct {
	age         int
	description string
	name        string
	role        string
}

// -----------------------------------------------------------------------------
// Character public methods.
// -----------------------------------------------------------------------------
func (c *Character) GetAge() int {
	return c.age
}

func (c *Character) GetDescription() string {
	return c.description
}

func (c *Character) GetName() string {
	return c.name
}

func (c *Character) GetRole() string {
	return c.role
}

func (c *Character) SetAge(age int) {
	c.age = age
}

func (c *Character) SetDescription(description string) {
	c.description = description
}

func (c *Character) SetName(name string) {
	c.name = name
}

func (c *Character) SetRole(role string) {
	c.role = role
}

var _ ICharacter = (*Character)(nil)
