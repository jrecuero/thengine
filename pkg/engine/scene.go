// scene.go contains all attributes and methods required to ahdnle a single
// scene in the application.
package engine

import "github.com/gdamore/tcell/v2"

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
	entities []IEntity
	screen   IScreen
}

// NewScreen function creates a new Scene instance.
func NewScene(name string, screen IScreen) *Scene {
	scene := &Scene{
		EObject:  NewEObject(name),
		entities: []IEntity{},
		screen:   screen,
	}
	return scene
}

// -----------------------------------------------------------------------------
// Scene public methods
// -----------------------------------------------------------------------------

// AddEntity methods adds a new entity to the scene.
func (s *Scene) AddEntity(entity IEntity) error {
	s.entities = append(s.entities, entity)
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
	for _, entity := range s.entities {
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

// Start method proceeds to starts all scene resources.
func (s *Scene) Start() {
}

// Update method proceeds to updates all scene resources.
func (s *Scene) Update(event tcell.Event) {
	for _, entity := range s.entities {
		entity.Update(event)
	}
}

var _ IScene = (*Scene)(nil)
