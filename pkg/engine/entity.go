// entity.go contains all data and methods required for handling an entity
// in the application. An entity is the basic object that engine handles.
package engine

import (
	"encoding/json"
	"fmt"
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
	IObjectUI
	IFocus
	IObserver
	Consume()
	Draw(IScene)
	EndTick(IScene)
	GetCache() api.ICache
	GetCanvas() *Canvas
	GetCollider() *Collider
	GetPLevel() int
	GetValidator() IValidator
	GetZLevel() int
	Init(tcell.Screen)
	IsSolid() bool
	MarshalJSON() ([]byte, error)
	MarshalMap(*api.Point) (map[string]any, error)
	MarshalCode(*api.Point) (string, error)
	Refresh()
	SetBehaviorFor(string, any)
	SetCache(api.ICache)
	SetCanvas(*Canvas)
	SetPLevel(int)
	SetSolid(bool)
	SetValidator(IValidator)
	SetZLevel(int)
	Start()
	StartTick(IScene)
	Stop()
	Update(tcell.Event, IScene)
	UnmarshalMap(map[string]any, *api.Point) error
	Validate(any, ...any) error
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
	*ObjectUI
	*Focus
	behavior  IBehavior
	cache     api.ICache
	canvas    *Canvas
	pLevel    int
	screen    tcell.Screen
	solid     bool
	validator IValidator
	zLevel    int
}

// -----------------------------------------------------------------------------
// New Entity functions
// -----------------------------------------------------------------------------

// NewEntity function creates a new Entity instance with all given attributes.
func NewEntity(name string, position *api.Point, size *api.Size, style *tcell.Style) *Entity {
	entity := &Entity{
		ObjectUI:  NewObjectUI(name, position, size, style),
		Focus:     NewDisableFocus(),
		behavior:  NewBehavior(),
		cache:     api.NewCache(),
		canvas:    NewCanvas(size),
		pLevel:    0,
		screen:    nil,
		solid:     false,
		validator: nil,
		zLevel:    0,
	}
	return entity
}

// NewEmptyEntity function creates a new Entity instance with all attributes
// as default values.
func NewEmptyEntity() *Entity {
	return &Entity{
		ObjectUI:  NewObjectUI("", nil, nil, nil),
		Focus:     NewDisableFocus(),
		behavior:  NewBehavior(),
		cache:     api.NewCache(),
		canvas:    nil,
		pLevel:    0,
		screen:    nil,
		solid:     false,
		validator: nil,
		zLevel:    0,
	}
}

// NewNamedEntity function creates a new Entity instance with all default
// attributes but the given name.
func NewNamedEntity(name string) *Entity {
	return &Entity{
		ObjectUI:  NewObjectUI(name, nil, nil, nil),
		Focus:     NewDisableFocus(),
		behavior:  NewBehavior(),
		cache:     api.NewCache(),
		canvas:    nil,
		pLevel:    0,
		screen:    nil,
		solid:     false,
		validator: nil,
		zLevel:    0,
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
	//    WithField("struct", "Entity").
	//    WithField("method", "CanHaveFocus").
	//    Debugf("entity %s %d %t %t", e.GetName(), int(e.GetFocusType()), e.IsFocusEnable(), e.IsActive())
	return e.IsFocusEnable() && e.IsActive()
}

// Draw method renders the entity in the screen.
func (e *Entity) Draw(scene IScene) {
	if b := e.behavior.GetBehaviorFor(BehaviorDraw); b != nil {
		if behavior, ok := b.(func(IScene)); ok {
			defer behavior(scene)
		}
	}
	if e.IsVisible() && e.GetCanvas() != nil {
		e.canvas.RenderAt(scene.GetCamera(), e.position)
	}
}

func (e *Entity) EndTick(IScene) {
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
	if b := e.behavior.GetBehaviorFor(BehaviorConsume); b != nil {
		if behavior, ok := b.(func()); ok {
			defer behavior()
		}
	}
}

// GetCanvas method returns the entity canvas instance.
func (e *Entity) GetCanvas() *Canvas {
	return e.canvas
}
func (e *Entity) GetCache() api.ICache {
	return e.cache
}

// GetScreen method returns the entity screen instance.
func (e *Entity) GetScreen() tcell.Screen {
	return e.screen
}

// GetPLevel method returns the entity p-level value.
func (e *Entity) GetPLevel() int {
	return e.pLevel
}

// GetRect method returns the entity rectangle instance.
//func (e *Entity) GetRect() *api.Rect {
//    var rect *api.Rect
//    if e.GetCanvas() != nil {
//        rect = e.GetCanvas().GetRect()
//        rect.Origin.X += e.GetPosition().X
//        rect.Origin.Y += e.GetPosition().Y
//    }
//    return rect
//}

func (e *Entity) GetValidator() IValidator {
	return e.validator
}

// GetZLevel method returns the entity z-level value.
func (e *Entity) GetZLevel() int {
	return e.zLevel
}

// Init methos initialize the entity instance.
func (e *Entity) Init(screen tcell.Screen) {
	if b := e.behavior.GetBehaviorFor(BehaviorInit); b != nil {
		if behavior, ok := b.(func()); ok {
			defer behavior()
		}
	}
	e.screen = screen
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
		"ch":       string(cell.GetRune()),
	}
	return content, nil
}

// MarshalCode method is the custom marshal method to generate pseudocode for
// the the instance.
func (e *Entity) MarshalCode(origin *api.Point) (string, error) {
	result := ""
	position := api.ClonePoint(e.position)
	if origin != nil {
		position.Subtract(origin)
	}
	fg, bg, attrs := e.style.Decompose()
	cell := e.GetCanvas().GetCellAt(nil)
	result += fmt.Sprintf("// entity: %s:%s\n", e.GetClassName(), e.GetName())
	result += fmt.Sprintf("style := tcell.StyleDefault.Foreground(tcell.GetColor(%s)).Background(tcell.GetColor(%s)).Attributes(tcell.AttrMask(%d))\n", fg, bg, attrs)
	result += fmt.Sprintf("entity := New%s(%s, api.Point(%d, %d), api.NewSize(%d, %d), &style)\n",
		e.GetClassName(), e.GetName(), position.X, position.Y, e.GetSize().W, e.GetSize().H)
	result += fmt.Sprintf("cell := engine.NewCell(&style, %d)\n", cell.GetRune())
	result += fmt.Sprintf("entity.GetCanvas().FillWithCell(cell)\n")
	result += fmt.Sprintf("\n")
	return result, nil
}

func (e *Entity) Notify(subjectID any, message any) {
	if b := e.behavior.GetBehaviorFor(BehaviorNotify); b != nil {
		if behavior, ok := b.(func(any, any)); ok {
			defer behavior(subjectID, message)
		}
	}
}

// Refresh method refreshes the entity instance.
func (e *Entity) Refresh() {
}

func (e *Entity) SetCache(cache api.ICache) {
	e.cache = cache
}

func (e *Entity) SetBehaviorFor(name string, f any) {
	e.behavior.SetBehaviorFor(name, f)
}

// SetCanvas method sets a new value for the entity canvas.
func (e *Entity) SetCanvas(canvas *Canvas) {
	e.canvas = canvas
}

// SetPLevel method sets a new value for the entity p-level.
func (e *Entity) SetPLevel(level int) {
	e.pLevel = level
}

// SetSolid method sets a new value for the entity solid attribute.
func (e *Entity) SetSolid(solid bool) {
	e.solid = solid
}

// SetStyle method sets a new value for the entity style.
func (e *Entity) SetStyle(style *tcell.Style) {
	e.ObjectUI.SetStyle(style)
	if e.GetCanvas() == nil {
		return
	}
	for _, rows := range e.GetCanvas().Rows {
		for _, cell := range rows.Cols {
			if cell != nil {
				cell.SetStyle(e.style)
			}
		}
	}
}

func (e *Entity) SetValidator(validator IValidator) {
	e.validator = validator
}

// SetZLevel method sets a new value for the entity z-level.
func (e *Entity) SetZLevel(level int) {
	e.zLevel = level
}

// Start method starts the entity instance.
func (e *Entity) Start() {
	if b := e.behavior.GetBehaviorFor(BehaviorStart); b != nil {
		if behavior, ok := b.(func()); ok {
			defer behavior()
		}
	}
}

func (e *Entity) StartTick(IScene) {
}

// Stop method stops the entity instance.
func (e *Entity) Stop() {
	if b := e.behavior.GetBehaviorFor(BehaviorStop); b != nil {
		if behavior, ok := b.(func()); ok {
			defer behavior()
		}
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
func (e *Entity) Update(event tcell.Event, scene IScene) {
	if b := e.behavior.GetBehaviorFor(BehaviorUpdate); b != nil {
		if behavior, ok := b.(func(tcell.Event, IScene)); ok {
			defer behavior(event, scene)
		}
	}
	if e.IsActive() {
	}
}

func (e *Entity) Validate(data any, args ...any) error {
	if e.validator != nil {
		return e.validator.Validate(data, args...)
	}
	return nil
}

var _ IObject = (*Entity)(nil)
var _ IObjectUI = (*Entity)(nil)
var _ IFocus = (*Entity)(nil)
var _ IObserver = (*Entity)(nil)
var _ IEntity = (*Entity)(nil)
