// iobject.go contains the interface and the struct for the most basic
// application object.
package engine

import (
	"fmt"

	"github.com/google/uuid"
)

// -----------------------------------------------------------------------------
// Local package functions
// -----------------------------------------------------------------------------

// newID function returns a unique identification based in the UUID module.
func newID() string {
	if uuid, err := uuid.NewUUID(); err == nil {
		return uuid.String()
	} else {
		panic(fmt.Sprintf("uuid generation error: %+v", err))
	}
}

// -----------------------------------------------------------------------------
//
// IObject
//
// -----------------------------------------------------------------------------

// IObject interface defines all function that should be implemented by the
// most basic and common application object.
type IObject interface {
	GetClassName() string
	GetID() string
	GetName() string
	IsActive() bool
	IsVisible() bool
	SetActive(bool)
	SetClassName(string)
	SetName(string)
	SetVisible(bool)
}

// -----------------------------------------------------------------------------
//
// EObject
//
// -----------------------------------------------------------------------------

// EObject structure defines the most basic and common application object.
type EObject struct {
	id        string
	name      string
	active    bool
	visible   bool
	className string
}

// NewEObject function creates a new EObject instance with the given name.
func NewEObject(name string) *EObject {
	return &EObject{
		name:      name,
		id:        newID(),
		active:    true,
		visible:   true,
		className: "",
	}
}

// -----------------------------------------------------------------------------
//
// EObject interface methods
//
// -----------------------------------------------------------------------------

// GetClassName method returns the instance class name attribute.
func (o *EObject) GetClassName() string {
	return o.className
}

// GetID method returns the instance unique identification.
func (o *EObject) GetID() string {
	return o.id
}

// GetName method returns the instance name.
func (o *EObject) GetName() string {
	return o.name
}

// IsActive method returns if the instance is active or not.
func (o *EObject) IsActive() bool {
	return o.active
}

// IsVisible method returns if the instance is visible or not.
func (o *EObject) IsVisible() bool {
	return o.visible
}

// SetActive method sets the instance active with the given value.
func (o *EObject) SetActive(active bool) {
	o.active = active
}

// SetClassName method sets the instance class name attribute.
func (e *EObject) SetClassName(className string) {
	e.className = className
}

// SetName method sets the instance name with the given value.
func (o *EObject) SetName(name string) {
	o.name = name
}

// SetVisible method sets the instance visible with the given value.
func (o *EObject) SetVisible(visible bool) {
	o.visible = visible
}

var _ IObject = (*EObject)(nil)
