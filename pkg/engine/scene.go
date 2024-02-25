// scene.go contains all attributes and methods required to ahdnle a single
// scene in the application.
package engine

// -----------------------------------------------------------------------------
//
// Scene
//
// -----------------------------------------------------------------------------

type Scene struct {
	entities []IEntity
	screen   IScreen
}

func NewSceen(screen IScreen) *Scene {
	scene := &Scene{
		entities: []IEntity{},
		screen:   screen,
	}
	return scene
}
