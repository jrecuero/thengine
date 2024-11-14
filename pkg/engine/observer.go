// observer.go
//
// Package observer implements the Observer design pattern to manage
// communication between multiple subjects and their observers. This package
// provides a central mechanism for managing updates and notifying registered
// observers when a subject's state changes.
//
// The key components of this package are:
//
// - **Observer**: An interface that defines a method for receiving updates
// from subjects.
// - **Subject**: An interface that defines the methods for updating and
// notifying observers.
// - **Manager**: A concrete struct that maintains the mapping between subjects
// and their observers. It handles registering, unregistering, and notifying
// observers when a subject's state is updated.
//
// Usage:
//
// 1. Define a subject that implements the `Subject` interface and triggers
// updates.
// 2. Define observers that implement the `Observer` interface and react to
// notifications.
// 3. Use the `Manager` to register observers to subjects and handle
// notifications.
//
// Example:
//
//	manager := NewManager()  // Create a new manager to handle observer
//	registrations
//
//	subject1 := &MyStruct{id: "Subject1"}  // Create subjects
//	subject2 := &MyStruct{id: "Subject2"}
//
//	observer1 := &AnotherStruct{id: "Observer1"}  // Create observers
//	observer2 := &AnotherStruct{id: "Observer2"}
//
//	manager.RegisterObserver(subject1.ID(), observer1)  // Register observers
//	with the manager
//	manager.RegisterObserver(subject1.ID(), observer2)
//	manager.RegisterObserver(subject2.ID(), observer1)
//
//	subject1.Update(manager, "NewData1")  // Trigger updates and notify
//	observers
//	subject2.Update(manager, "NewData2")
//
// This pattern allows for decoupling the subject and its observers,
// facilitating a flexible and scalable way of handling updates across various
// components of the system.

package engine

// -----------------------------------------------------------------------------
//
// Observer
//
// -----------------------------------------------------------------------------

// IObserver interface to be implemented by all observers.
type IObserver interface {
	Notify(any, any) // notify with a subject identifier
}

// -----------------------------------------------------------------------------
//
// Subject
//
// -----------------------------------------------------------------------------

// ISubject interface that triggers updates.
type ISubject interface {
	GetSubjectID() any            // returns a unique identifier for the subject.
	Update(*ObserverManager, any) // trigger an update adn notify via ObserverManager.
}

// -----------------------------------------------------------------------------
//
// ObserverManager
//
// -----------------------------------------------------------------------------

// ObserverManager struct is responsible for managing the registration and
// notification of observers in an Observer design pattern implementation. It
// serves as a central hub for subjects and observers, ensuring that observers
// are notified when the state of a subject changes.
type ObserverManager struct {
	name      string
	observers map[any][]IObserver // map subject IDs to observers
}

// -----------------------------------------------------------------------------
// ObserverManager create functions.
// -----------------------------------------------------------------------------

// NewObserverManager function creates and initializes a new ObserverManager
// instance.
func NewObserverManager(name string) *ObserverManager {
	return &ObserverManager{
		name:      name,
		observers: make(map[any][]IObserver),
	}
}

// -----------------------------------------------------------------------------
// ObserverManager private methods.
// -----------------------------------------------------------------------------

// findObserver method looks up for a given observer for an specific subject.
func (m *ObserverManager) findObserver(subjectID any, observer IObserver) (int, bool) {
	if _, exists := m.observers[subjectID]; exists {
		for index, obs := range m.observers[subjectID] {
			if obs == observer {
				return index, true
			}
		}
	}
	return -1, false
}

// -----------------------------------------------------------------------------
// ObserverManager public methods.
// -----------------------------------------------------------------------------

// GetName method returns the ObserveManager instance name.
func (m *ObserverManager) GetName() string {
	return m.name
}

// NotifyObservers method notifies all observers registered to a given subject.
func (m *ObserverManager) NotifyObservers(subjectID any, message any) {
	if observers, exists := m.observers[subjectID]; exists {
		for _, observer := range observers {
			observer.Notify(subjectID, message)
		}
	}
}

// RegisterObserver method adds an observer for an specific subject.
func (m *ObserverManager) RegisterObserver(subjectID any, observer IObserver) {
	if _, found := m.findObserver(subjectID, observer); !found {
		m.observers[subjectID] = append(m.observers[subjectID], observer)
	}
}

// UnregisterObserver method removes an observer for a specific subject.
func (m *ObserverManager) UnregisterObserver(subjectID any, observer IObserver) {
	if index, found := m.findObserver(subjectID, observer); found {
		m.observers[subjectID] = append(m.observers[subjectID][:index], m.observers[subjectID][index+1:]...)
	}
}
