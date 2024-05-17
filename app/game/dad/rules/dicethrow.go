package rules

// IDiceThrow interface defines all possible methods for any dice throw.
type IDiceThrow interface {
	GetName() string
	SetName(string)
	GetDescription() string
	SetDescription(string)
	GetScore() int
	SetScore(int) bool
	GetExtra() int
	SetExtra(int)
}

// DiceThrow structure contains all attributes required to define an dice throw.
type DiceThrow struct {
	name        string // dice throw name.
	shortName   string // dice throw short name.
	description string // dice throw description.
	score       int    // dice throw score.
	extra       int    // dice throw extra score.
}

// NewDiceThrow function creates a new DiceThrow instance.
func NewDiceThrow(name, shortname string, score int) *DiceThrow {
	return &DiceThrow{
		name:      name,
		shortName: shortname,
		score:     score,
	}
}

// GetName method returns dice throw name.
func (t *DiceThrow) GetName() string {
	return t.name
}

// SetName method sets dice throw name.
func (t *DiceThrow) SetName(name string) {
	t.name = name
}

// GetShortName method returns dice throw short name.
func (t *DiceThrow) GetShortName() string {
	return t.shortName
}

// SetShortName method sets dice throw short name.
func (t *DiceThrow) SetShortName(name string) {
	t.shortName = name
}

// GetDescription method returns dice throw description.
func (t *DiceThrow) GetDescription() string {
	return t.description
}

// SetDescription method sets dice throw description.
func (t *DiceThrow) SetDescription(desc string) {
	t.description = desc
}

// GetScore method returns dice throw score value.
func (t *DiceThrow) GetScore() int {
	score := t.score + t.GetExtra()
	if score > 30 {
		score = 30
	}
	return score
}

// SetScore method sets dice throw score value.
func (t *DiceThrow) SetScore(score int) bool {
	if score < 1 || score > 30 {
		return false
	}
	t.score = score
	return true
}

// GetExtra method returns dice throw extra score value.
func (t *DiceThrow) GetExtra() int {
	return t.extra
}

// SetExtra method sets dice throw extra score value.
func (t *DiceThrow) SetExtra(extra int) {
	t.extra = extra
}
