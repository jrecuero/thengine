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
type Entity struct {
	*EObject
	canvas   *Canvas
	position *api.Point
	size     *api.Size
	style    *tcell.Style
}

// NewEntity function creates a new Entity instance with all given attributes.
func NewEntity(name string, position *api.Point, size *api.Size, style *tcell.Style) *Entity {
	entity := &Entity{
		EObject:  NewEObject(name),
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
	return &Entity{}
}

// -----------------------------------------------------------------------------
// Entity public methods
// -----------------------------------------------------------------------------

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
}

func (e *Entity) Start() {

}

func (e *Entity) Update(tcell.Event) {
	if e.IsActive() {
		//tools.Logger.WithField("module", "entity").WithField("function", "update").Infof("!!!")
	}
}

var _ IEntity = (*Entity)(nil)
