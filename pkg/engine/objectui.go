package engine

import (
	"github.com/gdamore/tcell/v2"
	"github.com/jrecuero/thengine/pkg/api"
)

// -----------------------------------------------------------------------------
//
// IObjectUI
//
// -----------------------------------------------------------------------------

type IObjectUI interface {
	IObject
	GetPosition() *api.Point
	GetRect() *api.Rect
	GetSize() *api.Size
	GetStyle() *tcell.Style
	SetPosition(*api.Point)
	SetSize(*api.Size)
	SetStyle(*tcell.Style)
}

// -----------------------------------------------------------------------------
//
// ObjectUI
//
// -----------------------------------------------------------------------------

type ObjectUI struct {
	*EObject
	position *api.Point
	size     *api.Size
	style    *tcell.Style
}

func NewObjectUI(name string, position *api.Point, size *api.Size, style *tcell.Style) *ObjectUI {
	return &ObjectUI{
		EObject:  NewEObject(name),
		position: position,
		size:     size,
		style:    style,
	}
}

// -----------------------------------------------------------------------------
// ObjectUI public methods
// -----------------------------------------------------------------------------

// GetPosition method returns the object origin position.
func (o *ObjectUI) GetPosition() *api.Point {
	return o.position
}

// GetRect method returns the object rectangle instance.
func (o *ObjectUI) GetRect() *api.Rect {
	return api.NewRect(o.position, o.size)
}

// GetSize method returns the object size instance.
func (o *ObjectUI) GetSize() *api.Size {
	return o.size
}

// GetStyle method returns the object style instance.
func (o *ObjectUI) GetStyle() *tcell.Style {
	return o.style
}

// SetPosition method sets a new value for the object position.
func (o *ObjectUI) SetPosition(position *api.Point) {
	o.position = position
}

// SetSize method sets a new value for the object size.
func (o *ObjectUI) SetSize(size *api.Size) {
	o.size = size
}

// SetStyle method sets a new value for the object style.
func (o *ObjectUI) SetStyle(style *tcell.Style) {
	o.style = style
}

var _ IObject = (*ObjectUI)(nil)
