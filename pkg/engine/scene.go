// scene.go contains all attributes and methods required to ahdnle a single
// scene in the application.
package engine

import (
	"sort"

	"github.com/gdamore/tcell/v2"
	"github.com/jrecuero/thengine/pkg/tools"
)

// -----------------------------------------------------------------------------
// Package public constants
// -----------------------------------------------------------------------------
const (
	InvalidEntityIndex = -1
)

// -----------------------------------------------------------------------------
//
// IScene
//
// -----------------------------------------------------------------------------

// IScene interface defines all methods any Scene structure should implement.
type IScene interface {
	IObject
	AddEntity(IEntity) error
	Draw()
	GetEntities() []IEntity
	GetScreen() IScreen
	Init(tcell.Screen)
	Update(tcell.Event)
	Start()
}

// -----------------------------------------------------------------------------
//
// Scene
//
// -----------------------------------------------------------------------------

// Scene struct contains all attribute required for handling an application
// scene.
type Scene struct {
	*EObject
	entities       []IEntity
	zLevelEntities []IEntity
	pLevelEntities []IEntity
	screen         IScreen
}

// NewScreen function creates a new Scene instance.
func NewScene(name string, screen IScreen) *Scene {
	scene := &Scene{
		EObject:        NewEObject(name),
		entities:       []IEntity{},
		zLevelEntities: []IEntity{},
		pLevelEntities: []IEntity{},
		screen:         screen,
	}
	return scene
}

// -----------------------------------------------------------------------------
// Scene private methods
// -----------------------------------------------------------------------------

// findEntity methods finds the given entity in the list of entities.
func (s *Scene) findEntity(entity IEntity) int {
	for index, ent := range s.entities {
		if ent == entity {
			return index
		}
	}
	return InvalidEntityIndex
}

// sortEntities method sorts zLevelEntites and pLevelEntities.
func (s *Scene) sortEntities() {
	// copy and sort zLevelEntities. Entities with lower zLevel are drawed
	// first.
	s.zLevelEntities = make([]IEntity, len(s.entities))
	copy(s.zLevelEntities, s.entities)
	sort.Slice(s.zLevelEntities, func(i, j int) bool {
		return s.zLevelEntities[i].GetZLevel() < s.zLevelEntities[j].GetZLevel()
	})

	// copy and soert pLevelEntities. Entities with higher pLevel are being
	// called first.
	s.pLevelEntities = make([]IEntity, len(s.entities))
	copy(s.pLevelEntities, s.entities)
	sort.Slice(s.pLevelEntities, func(i, j int) bool {
		return s.pLevelEntities[i].GetPLevel() > s.pLevelEntities[j].GetPLevel()
	})

	tools.Logger.WithField("module", "scene").
		WithField("function", "sortEntities").
		Debugf("zLevelEntities %+v, pLevelEntities %+v", s.zLevelEntities, s.pLevelEntities)
}

// -----------------------------------------------------------------------------
// Scene public methods
// -----------------------------------------------------------------------------

// AddEntity methods adds a new entity to the scene.
func (s *Scene) AddEntity(entity IEntity) error {
	s.entities = append(s.entities, entity)
	s.sortEntities()
	focusManager := GetEngine().GetFocusManager()
	focusManager.AddEntity(s, entity)
	return nil
}

// Draw method proceeds to draw all entities registered and visible in the
// scene at the scene screen.
func (s *Scene) Draw() {
	if s.screen == nil {
		return
	}
	// Draw entites by its zLevel.
	for _, entity := range s.zLevelEntities {
		entity.Draw(s.screen)
	}
}

// GetEntities method returns all entities in the scene.
func (s *Scene) GetEntities() []IEntity {
	return s.entities
}

// GetScreeen method returns the screen instance related to the scene.
func (s *Scene) GetScreen() IScreen {
	return s.screen
}

// Init method proceeds to initialize all scene resources.
func (s *Scene) Init(display tcell.Screen) {
	s.screen.Init(display)
	for _, entity := range s.entities {
		entity.Init(display)
	}
}

// Start method proceeds to start all scene resources.
func (s *Scene) Start() {
}

// RemoveEntity method proceeds to remove the given entity from the scene.
func (s *Scene) RemoveEntity(entity IEntity) error {
	if index := s.findEntity(entity); index != InvalidEntityIndex {
		s.entities = append(s.entities[:index], s.entities[index+1:]...)
		s.sortEntities()
	}
	return nil
}

// Update method proceeds to updates all scene resources.
func (s *Scene) Update(event tcell.Event) {
	// update entities by its pLevel.
	for _, entity := range s.pLevelEntities {
		entity.Update(event)
	}
}

var _ IScene = (*Scene)(nil)
