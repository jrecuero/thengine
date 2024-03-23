// entify.go contains all data and methods required for handling an entity
// in the application. An entity is the basic object that engine handles.
package engine

import (
	"github.com/gdamore/tcell/v2"
	"github.com/jrecuero/thengine/pkg/api"
)

// -----------------------------------------------------------------------------
//
// IEntity
//
// -----------------------------------------------------------------------------

// IEntity interface defines all methods any Entity structure should implement.
type IEntity interface {
	IObject
	IFocus
	Draw(IScreen)
	GetCanvas() *Canvas
	GetPosition() *api.Point
	GetSize() *api.Size
	GetStyle() *tcell.Style
	Init()
	SetCanvas(*Canvas)
	SetPosition(*api.Point)
	SetSize(*api.Size)
	SetStyle(*tcell.Style)
	Start()
	Update(tcell.Event)
}

// -----------------------------------------------------------------------------
//
// Entity
//
// -----------------------------------------------------------------------------

// Entity structure defines all attributes and methods for the basic
// application object.
// zLevel represents the z coordinate with allows to prioritize entities to be
// displayed before.
// pLevel represents the update priority of the entity which allows to update
// entities before.
type Entity struct {
	*EObject
	*Focus
	canvas   *Canvas
	position *api.Point
	size     *api.Size
	style    *tcell.Style
	zLevel   int
	pLevel   int
}

// NewEntity function creates a new Entity instance with all given attributes.
func NewEntity(name string, position *api.Point, size *api.Size, style *tcell.Style) *Entity {
	entity := &Entity{
		EObject:  NewEObject(name),
		Focus:    NewDisableFocus(),
		canvas:   NewCanvas(size),
		position: position,
		size:     size,
		style:    style,
	}
	return entity
}

// NewEmptyEntity function creates a new Entity instance with all attributes
// as default values.
func NewEmptyEntity() *Entity {
	return &Entity{
		EObject: NewEObject(""),
		Focus:   NewFocus(NoFocus),
	}
}

// NewNamedEntity function creates a new Entity instance with all default
// attributes but the given name.
func NewNamedEntity(name string) *Entity {
	return &Entity{
		EObject: NewEObject(name),
		Focus:   NewFocus(NoFocus),
	}
}

// -----------------------------------------------------------------------------
// Entity public methods
// -----------------------------------------------------------------------------

func (e *Entity) CanHaveFocus() bool {
	return e.IsFocusEnable() && e.IsVisible() && e.IsActive()
}

func (e *Entity) Draw(screen IScreen) {
	if e.IsVisible() {
		e.canvas.RenderAt(screen, e.position)
	}
}

func (e *Entity) GetCanvas() *Canvas {
	return e.canvas
}

func (e *Entity) GetPosition() *api.Point {
	return e.position
}

func (e *Entity) GetSize() *api.Size {
	return e.size
}

func (e *Entity) GetStyle() *tcell.Style {
	return e.style
}

func (e *Entity) Init() {

}

func (e *Entity) SetCanvas(canvas *Canvas) {
	e.canvas = canvas
}

func (e *Entity) SetPosition(position *api.Point) {
	e.position = position
}

func (e *Entity) SetSize(size *api.Size) {
	e.size = size
}

func (e *Entity) SetStyle(style *tcell.Style) {
	e.style = style
	for _, rows := range e.canvas.Rows {
		for _, cell := range rows.Cols {
			if cell != nil {
				cell.Style = e.style
			}
		}
	}
}

func (e *Entity) Start() {

}

func (e *Entity) Update(tcell.Event) {
	if e.IsActive() {
		//tools.Logger.WithField("module", "entity").WithField("function", "update").Infof("!!!")
	}
}

var _ IObject = (*Entity)(nil)
var _ IFocus = (*Entity)(nil)
var _ IEntity = (*Entity)(nil)
