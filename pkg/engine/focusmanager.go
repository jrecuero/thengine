// focusManager.go contains all data and methods required for handling focus
// between all application entities.
package engine

import (
	"fmt"

	"github.com/jrecuero/thengine/pkg/tools"
)

// -----------------------------------------------------------------------------
// Package private constants
// -----------------------------------------------------------------------------

const (
	entityNotInScene int = -1
)

// -----------------------------------------------------------------------------
//
// FocusManager
//
// -----------------------------------------------------------------------------

// FocusManager struct contains all attributes and methods required for
// handling focus between all application entities.
type FocusManager struct {
	entities  map[string][]IEntity
	withFocus map[string][]IEntity
	locked    bool
}

// -----------------------------------------------------------------------------
// New FocusManager function
// -----------------------------------------------------------------------------

// NewFocusManager function creates a new FocusManager instance.
func NewFocusManager() *FocusManager {
	return &FocusManager{
		locked:    false,
		entities:  make(map[string][]IEntity),
		withFocus: make(map[string][]IEntity),
	}
}

// -----------------------------------------------------------------------------
// FocusManager private methods
// -----------------------------------------------------------------------------

// acquireFocusToEntityInScene method acquires the focus for the given entity
// in the given scene with the given index.
func (m *FocusManager) acquireFocusToEntityInScene(sceneName string, entity IEntity, index int) error {
	if _, ok := m.withFocus[sceneName]; !ok {
		m.withFocus[sceneName] = []IEntity{}
	}

	m.withFocus[sceneName] = append(m.withFocus[sceneName], entity)
	tools.Logger.WithField("module", "focusmanager").
		WithField("method", "acquireFocusToEntityInScene").
		Debugf("withFocus %+v", m.withFocus)
	tools.Logger.WithField("module", "focusmanager").
		WithField("method", "acquireFocusToEntityInScene").
		Debugf("acquire focus entity %s", entity.GetName())
	entity.AcquireFocus()
	// Remove the entity from the list of entities so it can not take focus
	// again.
	m.entities[sceneName] = append(m.entities[sceneName][:index], m.entities[sceneName][index+1:]...)
	return nil
}

// indexForEntityInEntities method looks for the scene name and the index of
// the given entity in all lists of entities.
func (m *FocusManager) indexForEntityInEntities(entity IEntity) (string, int) {
	for sceneName, entities := range m.entities {
		for index, ent := range entities {
			if ent == entity {
				return sceneName, index
			}
		}
	}
	return "", entityNotInScene
}

// indexForEntityInScene method looks for the index of the given entity in the
// list of entities for the given scene.
func (m *FocusManager) indexForEntityInScene(scene IScene, entity IEntity) int {
	if entitiesInScene, ok := m.entities[scene.GetName()]; ok {
		for index, ent := range entitiesInScene {
			if ent == entity {
				return index
			}
		}
	}
	return entityNotInScene
}

// indexForEntityInSceneWithFocus method looks for the index of the given
// entity in the list of entities with focus for the given scene.
func (m *FocusManager) indexForEntityInSceneWithFocus(scene IScene, entity IEntity) int {
	if entitiesInScene, ok := m.withFocus[scene.GetName()]; ok {
		for index, ent := range entitiesInScene {
			if ent == entity {
				return index
			}
		}
	}
	return entityNotInScene
}

// indexForEntityInWithFocus method looks for the scene name and index of the
// given entity in all list of entities with focus.
func (m *FocusManager) indexForEntityInWithFocus(entity IEntity) (string, int) {
	for sceneName, entities := range m.withFocus {
		for index, ent := range entities {
			if ent == entity {
				return sceneName, index
			}
		}
	}
	return "", entityNotInScene
}

// releaseFocusFromEntityInScene method release the focus for the given entity
// in the given scene with the given index.
func (m *FocusManager) releaseFocusFromEntityInScene(sceneName string, entity IEntity, index int) error {
	m.withFocus[sceneName] = append(m.withFocus[sceneName][:index], m.withFocus[sceneName][index+1:]...)
	// Remove the scene entry if there are not any entities there.
	if len(m.withFocus[sceneName]) == 0 {
		delete(m.withFocus, sceneName)
	}
	entity.ReleaseFocus()
	// Add the entity to the end of the entities list.
	m.entities[sceneName] = append(m.entities[sceneName], entity)
	return nil
}

// -----------------------------------------------------------------------------
// FocusManager public methods
// -----------------------------------------------------------------------------

// AcquireFocusToEntity method acquires the focus to the given entity.
func (m *FocusManager) AcquireFocusToEntity(entity IEntity) error {
	// if focus manager is locked, return.
	if m.IsLocked() {
		return nil
	}
	// if the entity already has focus, return.
	if _, index := m.indexForEntityInWithFocus(entity); index != entityNotInScene {
		return nil
	}
	if sceneName, index := m.indexForEntityInEntities(entity); index != entityNotInScene {
		// look for any single-focus entity in the list of entities with focus
		// for the given scene and remove the focus for that entity.
		for index, entity := range m.withFocus[sceneName] {
			if entity.GetFocusType() == SingleFocus {
				tools.Logger.WithField("module", "focusmanager").
					WithField("method", "AcquireFocusToEntity").
					Debugf("release-focus entity %s", entity.GetName())
				m.releaseFocusFromEntityInScene(sceneName, entity, index)
				break
			}
		}
		m.acquireFocusToEntityInScene(sceneName, entity, index)
		return nil
	}
	return fmt.Errorf("entity %s not found", entity.GetName())
}

// AddEntity method adds a new entity to be handled in the focus manager.
func (m *FocusManager) AddEntity(scene IScene, entity IEntity) error {
	if !entity.IsFocusEnable() {
		return fmt.Errorf("entity %s in scene %s has focus disabled", entity.GetName(), scene.GetName())
	}
	tools.Logger.WithField("module", "focusmanager").
		WithField("struct", "FocusManager").
		WithField("method", "UpdateFocusForScene").
		Debugf("add entity %s", scene.GetName())
	m.entities[scene.GetName()] = append(m.entities[scene.GetName()], entity)
	return nil
}

// GetEntities method returns all entities in the focus manager.
func (m *FocusManager) GetEntities() map[string][]IEntity {
	return m.entities
}

// GetEntitiesWithFocus method returns all entities with focus.
func (m *FocusManager) GetEntitiesWithFocus() map[string][]IEntity {
	return m.withFocus
}

// IsLocked method checks if the focus manager is locked for switching focus to
// other entities.
func (m *FocusManager) IsLocked() bool {
	return m.locked
}

// NextEntityWithFocusInScene method looks for the next entity to give focus in
// the given scene.
func (m *FocusManager) NextEntityWithFocusInScene(scene IScene) (IEntity, int) {
	if entities, ok := m.entities[scene.GetName()]; ok {
		for index, entity := range entities {
			if entity.CanHaveFocus() {
				return entity, index
			}
		}
	}
	return nil, 0
}

// ReleaseFocusFromEntity method release the focus from the given entity.
func (m *FocusManager) ReleaseFocusFromEntity(entity IEntity) error {
	// if focus manager is locked, return.
	if m.IsLocked() {
		return nil
	}
	// look for the entity in the list of entities with focus.
	if sceneName, index := m.indexForEntityInWithFocus(entity); index != entityNotInScene {
		m.releaseFocusFromEntityInScene(sceneName, entity, index)
	}
	return nil
}

// RemoveEntity method removes the given entity to be handled in the focus
// manaer.
func (m *FocusManager) RemoveEntity(scene IScene, entity IEntity) error {
	if !entity.IsFocusEnable() {
		return fmt.Errorf("entity %s in scene %s has focus disabled", entity.GetName(), scene.GetName())
	}
	sceneName := scene.GetName()

	// If the entity is in the list of entities with focus, remove it from
	// there.
	if index := m.indexForEntityInScene(scene, entity); index != entityNotInScene {
		// Remove the entity from the list of entities in the given scene.
		m.entities[sceneName] = append(m.entities[sceneName][:index], m.entities[sceneName][index+1:]...)
		// Remove the scene entry if there are not any entities there.
		if len(m.entities[sceneName]) == 0 {
			delete(m.entities, sceneName)
		}
	}

	// If the entity is in the list of entities with focus, remove it from
	// there.
	if index := m.indexForEntityInSceneWithFocus(scene, entity); index != entityNotInScene {
		// Be sure the entity release its focus.
		entity.ReleaseFocus()
		m.withFocus[sceneName] = append(m.withFocus[sceneName][:index], m.withFocus[sceneName][index+1:]...)
		// Remove the scene entry if there are not any entities there.
		if len(m.withFocus[sceneName]) == 0 {
			delete(m.withFocus, sceneName)
		}
	}
	// after entity has been removed from all list, update the focus for the
	// scene for the next available entity.
	m.UpdateFocusForScene(scene)
	return nil
}

// RemoveEntitiesInScene method removes all entities from the given scene in
// the focus manager.
func (m *FocusManager) RemoveEntitiesInScene(scene IScene) error {
	m.entities[scene.GetName()] = []IEntity{}
	m.withFocus[scene.GetName()] = []IEntity{}
	return nil
}

// RemoveScene method removes all entities for the given scene.
func (m *FocusManager) RemoveScene(scene IScene) {
	delete(m.entities, scene.GetName())
	delete(m.withFocus, scene.GetName())
}

// SetLocked method locks or unlocks focus manager in order to select a new
// entity to focus.
func (m *FocusManager) SetLocked(locked bool) {
	m.locked = locked
}

// UpdateFocusForScene method proceeds to update entities with focus for the
// given entity.
// Searches for all entities with multi-focus which will be given focus.
// Remove focus for the active single-focus entity, if present in the scene..
// Acquire focus to the next single-focus entity in the scene.
func (m *FocusManager) UpdateFocusForScene(scene IScene) error {
	// if focus manager is locked, return.
	if m.IsLocked() {
		return nil
	}

	sceneName := scene.GetName()
	if entities, ok := m.entities[sceneName]; ok {
		// ensure that there is a list for entities with focus for the given
		// scene.
		if _, ok := m.withFocus[sceneName]; !ok {
			m.withFocus[sceneName] = []IEntity{}
		}
		// look for all multi-focus entities in the given scene.
		for index, entity := range entities {
			if entity.CanHaveFocus() && (entity.GetFocusType() == MultiFocus) {
				m.acquireFocusToEntityInScene(sceneName, entity, index)
			}
		}
		// look for any single-focus entity in the list of entities with focus
		// for the given scene and remove the focus for that entity.
		for index, entity := range m.withFocus[sceneName] {
			if entity.GetFocusType() == SingleFocus {
				tools.Logger.WithField("module", "focusmanager").
					WithField("method", "UpdateFocusForScene").
					Debugf("release-focus entity %s", entity.GetName())
				m.releaseFocusFromEntityInScene(sceneName, entity, index)
				break
			}
		}
		// look for the next single-focus entity in the list of entities for
		// the given scene.
		if entity, index := m.NextEntityWithFocusInScene(scene); entity != nil {
			m.acquireFocusToEntityInScene(sceneName, entity, index)
		}
	} else {
		message := fmt.Sprintf("scene %s not found", sceneName)
		tools.Logger.WithField("module", "focusmanager").
			WithField("method", "UpdateFocusForScene").
			Errorf(message)
		return fmt.Errorf(message)
	}
	return nil
}
