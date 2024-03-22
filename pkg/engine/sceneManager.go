// sceneManager.go contains all logic required for handling all scenes in the
// application.
package engine

import "github.com/gdamore/tcell/v2"

// -----------------------------------------------------------------------------
// Package global constants
// -----------------------------------------------------------------------------

const (
	InvalidSceneIndex = -1
)

// -----------------------------------------------------------------------------
//
// SceneManager
//
// -----------------------------------------------------------------------------

// SceneManager structure defines attributes and functions to handle multiple
// scenes in the application.
type SceneManager struct {
	scenes        []IScene
	activeScenes  []IScene
	visibleScenes []IScene
}

// NewSceneManager function creates a new SceneManager instance.
func NewSceneManager() *SceneManager {
	mgr := &SceneManager{
		scenes:        make([]IScene, 0),
		activeScenes:  make([]IScene, 0),
		visibleScenes: make([]IScene, 0),
	}
	return mgr
}

// -----------------------------------------------------------------------------
// Package private methods
// -----------------------------------------------------------------------------

// getSceneInSlice function looks for the given scene in the given slice of
// scenes.
func getSceneInSlice(scene IScene, slice []IScene) int {
	for index, sc := range slice {
		if sc == scene {
			return index
		}
	}
	return InvalidSceneIndex
}

// -----------------------------------------------------------------------------
// SceneManager public methods
// -----------------------------------------------------------------------------

// AddScene method adds a new Scene to be handled by the manager.
func (m *SceneManager) AddScene(scene IScene) bool {
	m.scenes = append(m.scenes, scene)
	return true
}

// Draw method is called by the engine to draw all visible scenes in the scene
// manager.
func (m *SceneManager) Draw(display tcell.Screen) {
	for _, scene := range m.visibleScenes {
		scene.Draw()
		scene.GetScreen().Draw(true, display)
	}
}

// GetSceneByIndex method finds a scene with the given index. If the index is
// -1 it retreive the last scene.
func (m *SceneManager) GetSceneByIndex(index int) IScene {
	if index == -1 {
		return m.scenes[index]
	}
	if (index >= 0) && (index < len(m.scenes)) {
		return m.scenes[index]
	}
	return nil
}

// GetAllActiveScene method returns all active scenes in the scene manager.
func (m *SceneManager) GetAllActiveScenes() []IScene {
	return m.activeScenes
}

// GetAllScenes method returns all available scenes in the scene manager.
func (m *SceneManager) GetAllScenes() []IScene {
	return m.scenes
}

// GetAllVisibleScenes method returns all visible scenes in the scene manager.
func (m *SceneManager) GetAllVisibleScenes() []IScene {
	return m.visibleScenes
}

// GetSceneByName method finds a scene with the given name.
func (m *SceneManager) GetSceneByName(name string) IScene {
	for _, scene := range m.scenes {
		if scene.GetName() == name {
			return scene
		}
	}
	return nil
}

// GetActiveSceneIndex method returns the index of the scene in the list of
// active scenes.
func (m *SceneManager) GetActiveSceneIndex(scene IScene) int {
	return getSceneInSlice(scene, m.activeScenes)
}

// GetSceneIndex method returns the index of the scene in list of available
// scenes.
func (m *SceneManager) GetSceneIndex(scene IScene) int {
	return getSceneInSlice(scene, m.scenes)
}

// GetVisibleSceneIndex method returns the index of the scene in the list of
// visible scenes.
func (m *SceneManager) GetVisibleSceneIndex(scene IScene) int {
	return getSceneInSlice(scene, m.visibleScenes)
}

// Init method is called to initializes all scene manager resources.
func (m *SceneManager) Init() {
}

// IsSceneAvailable method finds the given scene in the list of all scenes available.
func (m *SceneManager) IsSceneAvailable(scene IScene) bool {
	if index := m.GetSceneIndex(scene); index != InvalidSceneIndex {
		return true
	}
	return false
}

// IsSceneActive method finds if the given scene is in the list of all active
// scenes.
func (m *SceneManager) IsSceneActive(scene IScene) bool {
	if index := m.GetActiveSceneIndex(scene); index != InvalidSceneIndex {
		return true
	}
	return false
}

// IsSceneVisible method finds if the given scene in in the list of all visible
// scenes.
func (m *SceneManager) IsSceneVisible(scene IScene) bool {
	if index := m.GetVisibleSceneIndex(scene); index != InvalidSceneIndex {
		return true
	}
	return false
}

// PushActiveSceneAsFirst method pushes the given scene as the first in the
// slice of active scenes.
func (m *SceneManager) PushActiveSceneAsFirst(scene IScene) bool {
	if m.IsSceneAvailable(scene) {
		m.activeScenes = append([]IScene{scene}, m.activeScenes...)
		return true
	}
	return false
}

// PushActiveSceneAsLast method pushes the given scene as the last in the slice
// of active scenes.
func (m *SceneManager) PushActiveSceneAsLast(scene IScene) bool {
	if m.IsSceneAvailable(scene) {
		m.activeScenes = append(m.activeScenes, scene)
		return true
	}
	return false
}

// PushVisibleSceneAsFirst method pushes the given scene as the first in  the
// slice of visible scenes.
func (m *SceneManager) PushVisibleSceneAsFirst(scene IScene) bool {
	if m.IsSceneAvailable(scene) {
		m.visibleScenes = append([]IScene{scene}, m.visibleScenes...)
		return true
	}
	return false
}

// PushVisibleSceneAsLast method pushes the given scene as the last in the
// slice of visble scenes.
func (m *SceneManager) PushVisibleSceneAsLast(scene IScene) bool {
	if m.IsSceneAvailable(scene) {
		m.visibleScenes = append(m.visibleScenes, scene)
		return true
	}
	return false
}

// RemoveScene method removes the given scene from all scene slices.
func (m *SceneManager) RemoveScene(scene IScene) bool {
	if activeIndex := m.GetActiveSceneIndex(scene); activeIndex != InvalidSceneIndex {
		m.activeScenes = append(m.activeScenes[:activeIndex], m.scenes[activeIndex+1:]...)
	}
	if visibleIndex := m.GetVisibleSceneIndex(scene); visibleIndex != InvalidSceneIndex {
		m.visibleScenes = append(m.visibleScenes[:visibleIndex], m.visibleScenes[visibleIndex+1:]...)
	}
	if index := m.GetSceneIndex(scene); index != InvalidSceneIndex {
		m.scenes = append(m.scenes[:index], m.scenes[index+1:]...)
		return true
	}
	return false
}

// SetDryRun method sets the dryRun variable to set dryRun flag which avoid any
// ncurses call.
func (m *SceneManager) SetDryRun(dryRun bool) {
	for _, scene := range m.scenes {
		screen := scene.GetScreen()
		screen.SetDryRun(dryRun)
	}
}

// SetSceneAsActive methos sets the given scene as an active scene.
func (m *SceneManager) SetSceneAsActive(scene IScene) bool {
	if !m.IsSceneAvailable(scene) {
		return false
	}
	if m.IsSceneActive(scene) {
		return true
	}
	m.activeScenes = append(m.activeScenes, scene)
	return true
}

// SetSceneAsVisible method sets the given scene as a visible scene.
func (m *SceneManager) SetSceneAsVisible(scene IScene) bool {
	if !m.IsSceneAvailable(scene) {
		return false
	}
	if m.IsSceneVisible(scene) {
		return true
	}
	m.visibleScenes = append(m.visibleScenes, scene)
	return true
}

// Start method is called by the engine to start all scene manager resources.
func (m *SceneManager) Start() {
}

// Update method is called by the engine to update all scene manager scenes.
func (m *SceneManager) Update(event tcell.Event) {
	for _, scene := range m.activeScenes {
		scene.Update(event)
	}
}
