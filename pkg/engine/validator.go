// validator.go module contains all interfaces, structures and logic in order
// to implement input validators to be used in any Entity.
package engine

import "github.com/gdamore/tcell/v2"

// -----------------------------------------------------------------------------
// Module public types
// -----------------------------------------------------------------------------

type Validation func(any, ...any) error

// -----------------------------------------------------------------------------
//
// IValidator
//
// -----------------------------------------------------------------------------

type IValidator interface {
	GetErrorStyle() *tcell.Style
	GetName() string
	Validate(any, ...any) error
}

// -----------------------------------------------------------------------------
//
// Validator
//
// -----------------------------------------------------------------------------

type Validator struct {
	name       string
	validation Validation
	errorStyle *tcell.Style
}

func NewValidator(name string, validation Validation, errorStyle *tcell.Style) *Validator {
	v := &Validator{
		name:       name,
		validation: validation,
		errorStyle: errorStyle,
	}
	return v
}

// -----------------------------------------------------------------------------
// Validator public methods
// -----------------------------------------------------------------------------

func (v *Validator) GetErrorStyle() *tcell.Style {
	return v.errorStyle
}

func (v *Validator) GetName() string {
	return v.name
}

func (v *Validator) Validate(data any, args ...any) error {
	if v.validation != nil {
		return v.validation(data, args...)
	}
	return nil
}

var _ IValidator = (*Validator)(nil)
