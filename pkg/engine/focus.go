// focus.go contains all data and methods required for any entity that can
// acquire keyboard focus
package engine

import "fmt"

// -----------------------------------------------------------------------------
// Package types
// -----------------------------------------------------------------------------

// FocusType type defines the type for the focus type enums.
type FocusType int

const (
	NoFocus FocusType = iota
	SingleFocus
	MultiFocus
)

// -----------------------------------------------------------------------------
//
// IFocus
//
// -----------------------------------------------------------------------------

// IFocus interface defines all methods that should be implemented by any
// entity with focus.
type IFocus interface {
	AcquireFocus() (bool, error)
	CanHaveFocus() bool
	GetFocusType() FocusType
	HasFocus() bool
	IsFocusEnable() bool
	ReleaseFocus() (bool, error)
	SetFocusEnable(bool)
}

// -----------------------------------------------------------------------------
//
// Focus
//
// -----------------------------------------------------------------------------

// Focus structure defines a baseline for any entity with focus.
type Focus struct {
	focus     bool
	enable    bool
	focusType FocusType
}

// -----------------------------------------------------------------------------
// New Focus functions
// -----------------------------------------------------------------------------

// NewFocus function creates a new Focus instance.
func NewFocus(focusType FocusType) *Focus {
	return &Focus{
		focus:     false,
		enable:    true,
		focusType: focusType,
	}
}

// NewDisableFocus function creates a new disabled Focus instance.
func NewDisableFocus() *Focus {
	return &Focus{
		focus:     false,
		enable:    false,
		focusType: NoFocus,
	}
}

// -----------------------------------------------------------------------------
// Focus interface methods
// -----------------------------------------------------------------------------

// AcquireFocus method acquires focus for the entity.
func (f *Focus) AcquireFocus() (bool, error) {
	if f.enable {
		f.focus = true
		return true, nil
	}
	return false, fmt.Errorf("focus disabled")
}

// CanHaveFocus method returns if the entity can have focus. This depends on
// the focus enable flag and any other attribute in the entity, so this most
// likely to be rewritten in the entity.
func (f *Focus) CanHaveFocus() bool {
	return f.enable
}

// GetFocusType method returns the focus type.
func (f *Focus) GetFocusType() FocusType {
	return f.focusType
}

// HasFocus method checks if the entity has the focus.
func (f *Focus) HasFocus() bool {
	return f.focus
}

// IsFocusEnable method checks if the entity can have focus.
func (f *Focus) IsFocusEnable() bool {
	return f.enable
}

// ReleaseFocus method release the focus for the entity.
func (f *Focus) ReleaseFocus() (bool, error) {
	if f.enable {
		f.focus = false
		return true, nil
	}
	return false, fmt.Errorf("focus disabled")
}

// SetFocusEnable method sets the enable flag for the focus.
func (f *Focus) SetFocusEnable(enable bool) {
	f.enable = enable
}
