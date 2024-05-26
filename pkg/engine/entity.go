// entify.go contains all data and methods required for handling an entity
// in the application. An entity is the basic object that engine handles.
package engine

import (
	"encoding/json"
	"strconv"

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
	MarshalJSON() ([]byte, error)
	MarshalMap(*api.Point) (map[string]any, error)
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
	UnmarshalMap(map[string]any, *api.Point) error
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
		EObject:     NewEObject(name),
		Focus:       NewDisableFocus(),
		canvas:      NewCanvas(size),
		position:    position,
		size:        size,
		style:       style,
		screen:      nil,
		zLevel:      0,
		pLevel:      0,
		solid:       false,
		customInit:  nil,
		customStart: nil,
		customStop:  nil,
	}
	return entity
}

// NewEmptyEntity function creates a new Entity instance with all attributes
// as default values.
func NewEmptyEntity() *Entity {
	return &Entity{
		EObject:     NewEObject(""),
		Focus:       NewDisableFocus(),
		canvas:      nil,
		position:    nil,
		size:        nil,
		style:       nil,
		screen:      nil,
		zLevel:      0,
		pLevel:      0,
		solid:       false,
		customInit:  nil,
		customStart: nil,
		customStop:  nil,
	}
}

// NewNamedEntity function creates a new Entity instance with all default
// attributes but the given name.
func NewNamedEntity(name string) *Entity {
	return &Entity{
		EObject:     NewEObject(name),
		Focus:       NewDisableFocus(),
		canvas:      nil,
		position:    nil,
		size:        nil,
		style:       nil,
		screen:      nil,
		zLevel:      0,
		pLevel:      0,
		solid:       false,
		customInit:  nil,
		customStart: nil,
		customStop:  nil,
	}
}

// NewHandler function creates a new Entity as a handler. A Handler does not
// have any position or size by default, it is not a solid object and it is
// not visible.
func NewHandler(name string) *Entity {
	handler := NewNamedEntity(name)
	return handler
}

// -----------------------------------------------------------------------------
// Entity public methods
// -----------------------------------------------------------------------------

// CanHaveFocus method checks if the entity can receive and have focus.
func (e *Entity) CanHaveFocus() bool {
	//tools.Logger.WithField("module", "entity").
	//    WithField("function", "CanHaveFocus").
	//    Debugf("entity %s %d %t %t", e.GetName(), int(e.GetFocusType()), e.IsFocusEnable(), e.IsActive())
	return e.IsFocusEnable() && e.IsActive()
}

// Draw method renders the entity in the screen.
func (e *Entity) Draw(scene IScene) {
	if e.IsVisible() && e.GetCanvas() != nil {
		e.canvas.RenderAt(scene.GetCamera(), e.position)
	}
}

// GetCollider method returns the Collider instance for the entity.
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

// GetCanvas method returns the entity canvas instance.
func (e *Entity) GetCanvas() *Canvas {
	return e.canvas
}

// GetScreen method returns the entity screen instance.
func (e *Entity) GetScreen() tcell.Screen {
	return e.screen
}

// GetPLevel method returns the entity p-level value.
func (e *Entity) GetPLevel() int {
	return e.pLevel
}

// GetPosition method returns the entity origin position.
func (e *Entity) GetPosition() *api.Point {
	return e.position
}

// GetRect method returns the entity rectangle instance.
func (e *Entity) GetRect() *api.Rect {
	var rect *api.Rect
	if e.GetCanvas() != nil {
		rect = e.GetCanvas().GetRect()
		rect.Origin.X += e.GetPosition().X
		rect.Origin.Y += e.GetPosition().Y
	}
	return rect
}

// GetSize method returns the entity size instance.
func (e *Entity) GetSize() *api.Size {
	return e.size
}

// GetStyle method returns the entity style instance.
func (e *Entity) GetStyle() *tcell.Style {
	return e.style
}

// GetZLevel method returns the entity z-level value.
func (e *Entity) GetZLevel() int {
	return e.zLevel
}

// Init methos initialize the entity instance.
func (e *Entity) Init(screen tcell.Screen) {
	e.screen = screen
	if e.customInit != nil {
		e.customInit()
	}
}

// IsSolid method returns if the entity is solid or not.
func (e *Entity) IsSolid() bool {
	return e.solid
}

// MarshalJSON method is the custom marshal method to generate JSON from an
// instance.
func (e *Entity) MarshalJSON() ([]byte, error) {
	content, err := e.MarshalMap(nil)
	if err != nil {
		return nil, err
	}
	return json.Marshal(content)
}

// MarshalMap method is the custom marshal method to generate a map[string]any
// from an instance.
func (e *Entity) MarshalMap(origin *api.Point) (map[string]any, error) {
	position := api.ClonePoint(e.position)
	if origin != nil {
		position.Subtract(origin)
	}
	fg, bg, attrs := e.style.Decompose()
	cell := e.GetCanvas().GetCellAt(nil)
	content := map[string]any{
		"class":    e.GetClassName(),
		"name":     e.name,
		"position": []int{position.X, position.Y},
		"size":     []int{e.size.W, e.size.H},
		"style":    []string{fg.String(), bg.String(), strconv.Itoa(int(attrs))},
		"ch":       string(cell.Rune),
	}
	return content, nil
}

// Refresh method refreshes the entity instance.
func (e *Entity) Refresh() {
}

// SetCanvas method sets a new value for the entity canvas.
func (e *Entity) SetCanvas(canvas *Canvas) {
	e.canvas = canvas
}

// SetCustomInit method sets a new value for the custom init function.
func (e *Entity) SetCustomInit(f func()) {
	e.customInit = f
}

// SetCustomStart method sets a new value for the custon start function.
func (e *Entity) SetCustomStart(f func()) {
	e.customStart = f
}

// SetCustomStop method sets a new value for the custom stop function.
func (e *Entity) SetCustomStop(f func()) {
	e.customStop = f
}

// SetPLevel method sets a new value for the entity p-level.
func (e *Entity) SetPLevel(level int) {
	e.pLevel = level
}

// SetPosition method sets a new value for the entity position.
func (e *Entity) SetPosition(position *api.Point) {
	e.position = position
}

// SetSize method sets a new value for the entity size.
func (e *Entity) SetSize(size *api.Size) {
	e.size = size
}

// SetSolid method sets a new value for the entity solid attribute.
func (e *Entity) SetSolid(solid bool) {
	e.solid = solid
}

// SetStyle method sets a new value for the entity style.
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

// SetZLevel method sets a new value for the entity z-level.
func (e *Entity) SetZLevel(level int) {
	e.zLevel = level
}

// Start method starts the entity instance.
func (e *Entity) Start() {
	if e.customStart != nil {
		e.customStart()
	}
}

// Stop method stops the entity instance.
func (e *Entity) Stop() {
	if e.customStop != nil {
		e.customStop()
	}
}

// UnmarshalJSON method is the custom method to unmarshal JSON data into an
// instance.
func (e *Entity) UnmarshalJSON(data []byte) error {
	var content map[string]any
	if err := json.Unmarshal(data, &content); err != nil {
		return err
	}
	return e.UnmarshalMap(content, nil)
}

// UnmarshalMap method is the custom method to unmarshal a map[string]any data
// into an instance.
func (e *Entity) UnmarshalMap(content map[string]any, origin *api.Point) error {
	if name, ok := content["name"].(string); ok {
		e.name = name
	}
	if position, ok := content["position"].([]any); ok {
		e.position = api.NewPoint(int(position[0].(float64)), int(position[1].(float64)))
		if origin != nil {
			e.position.Add(origin)
		}
	}
	if size, ok := content["size"].([]any); ok {
		e.size = api.NewSize(int(size[0].(float64)), int(size[1].(float64)))
	}
	if style, ok := content["style"].([]any); ok {
		tcellStyle := tcell.StyleDefault.
			Foreground(tcell.GetColor(style[0].(string))).
			Background(tcell.GetColor(style[1].(string)))
		e.style = &tcellStyle
	}
	return nil
}

// Update method updates the entity instance.
func (e *Entity) Update(tcell.Event, IScene) {
	if e.IsActive() {
	}
}

var _ IObject = (*Entity)(nil)
var _ IFocus = (*Entity)(nil)
var _ IEntity = (*Entity)(nil)
