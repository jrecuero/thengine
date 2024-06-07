// condition.go defines interface and structure related with a condition or a
// status effect, which alter a chracter abilities and actions in various ways.
package rules

const (
	ForEverDuration = -1
)

// -----------------------------------------------------------------------------
//
// ICondition
//
// -----------------------------------------------------------------------------

// ICondition interface defines all methods any condition structure should be
// implementing.
type ICondition interface {
	GetDescription() string
	GetDiceThrows() []IDiceThrow
	GetDuration() int
	GetName() string
	GetUName() string
	SetDescription(string)
	SetDiceThrows([]IDiceThrow)
	SetDuration(int)
	SetName(string)
	SetUName(string)
}

// -----------------------------------------------------------------------------
//
// Condition
//
// -----------------------------------------------------------------------------

// Condition structure defines all attributes and methos for any standard
// condition or status effect.
type Condition struct {
	name        string
	uname       string
	description string
	diceThrows  []IDiceThrow
	duration    int
}

func NewCondition(name string, uname string, diceThrows []IDiceThrow, duration int) *Condition {
	return &Condition{
		name:        name,
		uname:       uname,
		description: "",
		diceThrows:  diceThrows,
		duration:    duration,
	}
}

// -----------------------------------------------------------------------------
// Condition public methods
// -----------------------------------------------------------------------------

func (c *Condition) GetDescription() string {
	return c.description
}

func (c *Condition) GetDiceThrows() []IDiceThrow {
	return c.diceThrows
}

func (c *Condition) GetDuration() int {
	return c.duration
}

func (c *Condition) GetName() string {
	return c.name
}

func (c *Condition) GetUName() string {
	return c.uname
}

func (c *Condition) SetDescription(description string) {
	c.description = description
}

func (c *Condition) SetDiceThrows(diceThrows []IDiceThrow) {
	c.diceThrows = diceThrows
}

func (c *Condition) SetDuration(duration int) {
	c.duration = duration
}

func (c *Condition) SetName(name string) {
	c.name = name
}

func (c *Condition) SetUName(uname string) {
	c.uname = uname
}

var _ ICondition = (*Condition)(nil)
