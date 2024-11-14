package engine_test

import (
	"fmt"
	"testing"

	"github.com/jrecuero/thengine/pkg/engine"
)

type TestSubject struct {
	name string
}

func (s *TestSubject) GetSubjectID() any {
	return s.name
}

func (s *TestSubject) Update(m *engine.ObserverManager, message any) {
	m.NotifyObservers(s.GetSubjectID(), message)
}

type TestObserver struct {
	name   string
	result string
}

func (o *TestObserver) Notify(subjectID any, message any) {
	o.result = message.(string)
	fmt.Printf("observer %s subject %v message %v\n", o.name, subjectID, message)
}

func TestObserveManager(t *testing.T) {
	// Create ObserverManager
	m := engine.NewObserverManager("manager/observer/1")
	if m == nil {
		t.Errorf("NewObserverManager Error exp:*ObserveManager got:nil")
		return
	}

	// Create subject
	subject := &TestSubject{name: "test/subject/1"}

	// Create observers
	observer1 := &TestObserver{name: "test/observer/1"}
	observer2 := &TestObserver{name: "test/observer/2"}

	// Register observers to ObserverManager
	m.RegisterObserver(subject.GetSubjectID(), observer1)
	m.RegisterObserver(subject.GetSubjectID(), observer2)

	// trigger subject notification.
	message1 := "test/data/1"
	subject.Update(m, message1)
	if observer1.result != message1 {
		t.Errorf("Notify for %s failed exp: %s got: %s", observer1.name, message1, observer1.result)
	}
	if observer2.result != message1 {
		t.Errorf("Notify for %s failed exp: %s got: %s", observer2.name, message1, observer2.result)
	}

	message2 := "test/data/2"
	subject.Update(m, message2)
	if observer1.result != message2 {
		t.Errorf("Notify for %s failed exp: %s got: %s", observer1.name, message2, observer1.result)
	}
	if observer2.result != message2 {
		t.Errorf("Notify for %s failed exp: %s got: %s", observer2.name, message2, observer2.result)
	}
}
