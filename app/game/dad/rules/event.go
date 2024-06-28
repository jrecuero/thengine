// event.go package contains all data and logic required for adding events to
// the application. An event can trigger multiple actions and they can follow a
// logical path.
package rules

import (
	"fmt"
	"sync"

	"github.com/google/uuid"
)

// -----------------------------------------------------------------------------
// Module global types
// -----------------------------------------------------------------------------

type UniqueActionId string

// -----------------------------------------------------------------------------
// Module global constants
// -----------------------------------------------------------------------------

const (
	UAIdNone UniqueActionId = "uaid/**none**"
	UAIdEnd  UniqueActionId = "uaid/**end**"
)

// -----------------------------------------------------------------------------
// Module private functions
// -----------------------------------------------------------------------------

// generateUniqueActionId function create a new Unique Action ID.
func generateUniqueActionId() UniqueActionId {
	if uuid, err := uuid.NewUUID(); err == nil {
		return UniqueActionId(fmt.Sprintf("uaid/%s", uuid.String()))
	} else {
		panic(fmt.Sprintf("uuid generation error: %+v", err))
	}
}

// -----------------------------------------------------------------------------
//
// cache
//
// -----------------------------------------------------------------------------

type cache struct {
	mu    sync.RWMutex
	store map[string]any
}

func newCache() *cache {
	return &cache{
		store: make(map[string]any),
	}
}

func (c *cache) Delete(key string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	delete(c.store, key)
}

func (c *cache) Get(key string) (any, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()
	value, found := c.store[key]
	return value, found
}

func (c *cache) Set(key string, value any) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.store[key] = value
}

// -----------------------------------------------------------------------------
//
// IAction
//
// -----------------------------------------------------------------------------

// IAction interface defines all methods any action that is executed insde an
// event has to implement.
type IAction interface {
	Execute(IEvent, ...any) (UniqueActionId, error)
	GetUniqueActionId() UniqueActionId
}

// -----------------------------------------------------------------------------
//
// Action
//
// -----------------------------------------------------------------------------

type Action struct {
	uaid UniqueActionId
}

func NewAction() *Action {
	return &Action{
		uaid: generateUniqueActionId(),
	}
}

func NewEndAction() *Action {
	return &Action{
		uaid: UAIdEnd,
	}
}

func (a *Action) GetUniqueActionId() UniqueActionId {
	return a.uaid
}

func (a *Action) Execute(IEvent, ...any) (UniqueActionId, error) {
	return UAIdNone, nil
}

// -----------------------------------------------------------------------------
//
// IEvent
//
// -----------------------------------------------------------------------------

type IEvent interface {
	GetActions() []IAction
	GetCache() *cache
	GetDescription() string
	GetName() string
	Run(...any) error
	RunFrom(UniqueActionId, ...any) error
	SetActions([]IAction)
	SetCache(*cache)
	SetDescription(string)
	SetName(string)
}

// -----------------------------------------------------------------------------
//
// Event
//
// -----------------------------------------------------------------------------

// Event structure defines all data and logic required for implementing an
// event as a sequence of actions to run.
// Event runner starts running actions sequentially, if action does not return
// any error or any addition information, it moves to the next action.
// If the action returns a unique action id, it looks for that particular
// action in the list of actions an continue running from that point.
type Event struct {
	actions     []IAction
	cache       *cache
	description string
	name        string
}

func NewEvent(name string, actions []IAction) *Event {
	return &Event{
		actions:     actions,
		cache:       newCache(),
		description: name,
		name:        name,
	}
}

// -----------------------------------------------------------------------------
// Event public methods
// -----------------------------------------------------------------------------

func (e *Event) GetActions() []IAction {
	return e.actions
}

func (e *Event) GetActionByUAId(id UniqueActionId) (int, IAction) {
	for i, action := range e.actions {
		if action.GetUniqueActionId() == id {
			return i, action
		}
	}
	return 0, nil
}

func (e *Event) GetCache() *cache {
	return e.cache
}

func (e *Event) GetDescription() string {
	return e.description
}

func (e *Event) GetName() string {
	return e.name
}

func (e *Event) Run(args ...any) error {
	return e.runActions(e.actions, args...)
}

// RunFrom method runs all actions starting with the give action by its unique
// action ID.
func (e *Event) RunFrom(uaid UniqueActionId, args ...any) error {
	if index, a := e.GetActionByUAId(uaid); a != nil {
		return e.runActions(e.actions[index:], args...)
	}
	return nil
}

func (e *Event) runActions(actions []IAction, args ...any) error {
	for _, action := range actions {
		result, err := action.Execute(e, args...)
		if err != nil {
			return err
		} else if result == UAIdEnd {
			return nil
		} else if result != UAIdNone {
			if index, a := e.GetActionByUAId(result); a != nil {
				return e.runActions(actions[index:], args...)
			}
		}
	}
	return nil
}

func (e *Event) SetActions(actions []IAction) {
	e.actions = actions
}

func (e *Event) SetCache(c *cache) {
	e.cache = c
}

func (e *Event) SetDescription(description string) {
	e.description = description
}

func (e *Event) SetName(name string) {
	e.name = name
}

var _ IEvent = (*Event)(nil)
