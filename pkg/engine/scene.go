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
	CheckCollisionWith(IEntity) []IEntity
	Clean()
	Consume()
	Draw()
	GetEntities() []IEntity
	GetCamera() ICamera
	Init(tcell.Screen)
	RemoveEntity(IEntity) error
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
	camera         ICamera
}

// NewCamera function creates a new Scene instance.
func NewScene(name string, camera ICamera) *Scene {
	scene := &Scene{
		EObject:        NewEObject(name),
		entities:       []IEntity{},
		zLevelEntities: []IEntity{},
		pLevelEntities: []IEntity{},
		camera:         camera,
	}
	tools.Logger.WithField("module", "scene").WithField("function", "NewScene").Debugf("new scene %s", scene.GetName())
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

	//tools.Logger.WithField("module", "scene").
	//    WithField("function", "sortEntities").
	//    Debugf("zLevelEntities %+v, pLevelEntities %+v", s.zLevelEntities, s.pLevelEntities)
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

// CheckCollisionWith method checks if the given entity has a collision with
// any other solid entity in the scene.
func (s *Scene) CheckCollisionWith(entity IEntity) []IEntity {
	solidEntities := []IEntity{}
	for _, ent := range s.entities {
		if ent.IsActive() && ent.IsSolid() {
			solidEntities = append(solidEntities, ent)
		}
	}
	return CheckCollisionWith(entity, solidEntities)
}

// Clean method cleans all resources for the scene in order to set it up as a
// brand new screen.
func (s *Scene) Clean() {
	s.entities = []IEntity{}
	s.zLevelEntities = []IEntity{}
	s.pLevelEntities = []IEntity{}
}

// Consume method calls all entity instances to consume all messages from
// the mailbox.
func (s *Scene) Consume() {
	for _, entity := range s.pLevelEntities {
		entity.Consume()
	}
}

// Draw method proceeds to draw all entities registered and visible in the
// scene at the scene camera.
func (s *Scene) Draw() {
	if s.camera == nil {
		return
	}
	// Draw entites by its zLevel.
	for _, entity := range s.zLevelEntities {
		entity.Draw(s)
	}
}

// GetEntities method returns all entities in the scene.
func (s *Scene) GetEntities() []IEntity {
	return s.entities
}

// GetScreeen method returns the camera instance related to the scene.
func (s *Scene) GetCamera() ICamera {
	return s.camera
}

// Init method proceeds to initialize all scene resources.
func (s *Scene) Init(display tcell.Screen) {
	s.camera.Init(display)
	for _, entity := range s.entities {
		entity.Init(display)
	}
}

// Start method proceeds to start all scene resources.
func (s *Scene) Start() {
	for _, entity := range s.entities {
		entity.Start()
	}
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
		entity.Update(event, s)
	}
}

var _ IScene = (*Scene)(nil)
