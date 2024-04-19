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
	Consume()
	Draw(IScene)
	GetCanvas() *Canvas
	GetCollider() *Collider
	GetPLevel() int
	GetPosition() *api.Point
	GetRect() *api.Rect
	GetSize() *api.Size
	GetStyle() *tcell.Style
	GetZLevel() int
	Init(tcell.Screen)
	IsSolid() bool
	Refresh()
	SetCanvas(*Canvas)
	SetCustomInit(func())
	SetCustomStart(func())
	SetCustomStop(func())
	SetPLevel(int)
	SetPosition(*api.Point)
	SetSize(*api.Size)
	SetSolid(bool)
	SetStyle(*tcell.Style)
	SetZLevel(int)
	Start()
	Stop()
	Update(tcell.Event, IScene)
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
	canvas      *Canvas
	position    *api.Point
	size        *api.Size
	style       *tcell.Style
	screen      tcell.Screen
	zLevel      int
	pLevel      int
	solid       bool
	customInit  func()
	customStart func()
	customStop  func()
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
	//return e.IsFocusEnable() && e.IsVisible() && e.IsActive()
	return e.IsFocusEnable() && e.IsActive()
}

func (e *Entity) Draw(scene IScene) {
	if e.IsVisible() && e.GetCanvas() != nil {
		e.canvas.RenderAt(scene.GetCamera(), e.position)
	}
}

func (e *Entity) GetCollider() *Collider {
	if rect := e.GetRect(); rect != nil {
		return &Collider{
			rect:   rect,
			points: nil,
		}
	}
	return nil
}

// Consume method consume all messages from the mailbox.
func (e *Entity) Consume() {
}

func (e *Entity) GetCanvas() *Canvas {
	return e.canvas
}

func (e *Entity) GetScreen() tcell.Screen {
	return e.screen
}

func (e *Entity) GetPLevel() int {
	return e.pLevel
}

func (e *Entity) GetPosition() *api.Point {
	return e.position
}

func (e *Entity) GetRect() *api.Rect {
	var rect *api.Rect
	if e.GetCanvas() != nil {
		rect = e.GetCanvas().GetRect()
		rect.Origin.X += e.GetPosition().X
		rect.Origin.Y += e.GetPosition().Y
	}
	return rect
}

func (e *Entity) GetSize() *api.Size {
	return e.size
}

func (e *Entity) GetStyle() *tcell.Style {
	return e.style
}

func (e *Entity) GetZLevel() int {
	return e.zLevel
}

func (e *Entity) Init(screen tcell.Screen) {
	e.screen = screen
	if e.customInit != nil {
		e.customInit()
	}
}

func (e *Entity) IsSolid() bool {
	return e.solid
}

func (e *Entity) Refresh() {
}

func (e *Entity) SetCanvas(canvas *Canvas) {
	e.canvas = canvas
}

func (e *Entity) SetCustomInit(f func()) {
	e.customInit = f
}

func (e *Entity) SetCustomStart(f func()) {
	e.customStart = f
}

func (e *Entity) SetCustomStop(f func()) {
	e.customStop = f
}

func (e *Entity) SetPLevel(level int) {
	e.pLevel = level
}

func (e *Entity) SetPosition(position *api.Point) {
	e.position = position
}

func (e *Entity) SetSize(size *api.Size) {
	e.size = size
}

func (e *Entity) SetSolid(solid bool) {
	e.solid = solid
}

func (e *Entity) SetStyle(style *tcell.Style) {
	if e.GetCanvas() == nil {
		return
	}
	e.style = style
	for _, rows := range e.GetCanvas().Rows {
		for _, cell := range rows.Cols {
			if cell != nil {
				cell.Style = e.style
			}
		}
	}
}

func (e *Entity) SetZLevel(level int) {
	e.zLevel = level
}

func (e *Entity) Start() {
	if e.customStart != nil {
		e.customStart()
	}
}

func (e *Entity) Stop() {
	if e.customStop != nil {
		e.customStop()
	}
}

func (e *Entity) Update(tcell.Event, IScene) {
	if e.IsActive() {
	}
}

var _ IObject = (*Entity)(nil)
var _ IFocus = (*Entity)(nil)
var _ IEntity = (*Entity)(nil)
