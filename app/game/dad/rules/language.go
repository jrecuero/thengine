// language.go
package rules

// -----------------------------------------------------------------------------
//
// ILanguage
//
// -----------------------------------------------------------------------------

type ILanguage interface {
	GetDescription() string
	GetName() string
	SetDescription(string)
	SetName(string)
}

// -----------------------------------------------------------------------------
//
// Language
//
// -----------------------------------------------------------------------------

type Language struct {
	name        string
	description string
}

// NewLanguage function creates a new Language instance.
func NewLanguage(name string) *Language {
	l := &Language{
		description: name,
		name:        name,
	}
	return l
}

// -----------------------------------------------------------------------------
// Language public methods
// -----------------------------------------------------------------------------

func (l *Language) GetDescription() string {
	return l.description
}

func (l *Language) GetName() string {
	return l.name
}

func (l *Language) SetDescription(description string) {
	l.description = description
}

func (l *Language) SetName(name string) {
	l.name = name
}

var _ ILanguage = (*Language)(nil)
