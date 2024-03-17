package engine_test

import (
	"testing"

	"github.com/jrecuero/thengine/pkg/engine"
)

var (
	scene1   *engine.Scene  = engine.NewScene("scene/1", nil)
	scene2   *engine.Scene  = engine.NewScene("scene/2", nil)
	entity11 *engine.Entity = engine.NewNamedEntity("entity1/1")
	entity21 *engine.Entity = engine.NewNamedEntity("entity2/1")
)

func TestFocusManagerNewFocusManager(t *testing.T) {
	got := engine.NewFocusManager()
	if got == nil {
		t.Errorf("NewFocusManager Error exp:*FocusManager got:nil")
	}
	if got.IsLocked() {
		t.Errorf("NewFocusManager Error.IsLocked exp:false got:%t", got.IsLocked())
	}
	if len(got.GetEntities()) != 0 {
		t.Errorf("NewFocusManager Error.GetEntities exp:0 got:%d", len(got.GetEntities()))
	}
	if len(got.GetEntitiesWithFocus()) != 0 {
		t.Errorf("NewFocusManager Error.GetEntitiesWithFocus exp:0 got:%d", len(got.GetEntitiesWithFocus()))
	}
}

func TestFocusManagerAddEntity(t *testing.T) {
	got := engine.NewFocusManager()

	// Add one entity in one scene.
	gotError := got.AddEntity(scene1, entity11)
	if gotError != nil {
		t.Errorf("[1] AddEntity Error exp:nil got:%t", gotError)
	}
	gotEntities := got.GetEntities()
	if len(gotEntities) != 1 {
		t.Errorf("[1] AddEntity Error.GetEntities exp:1 got:%d", len(gotEntities))
	}
	if entities, ok := gotEntities[scene1.GetName()]; ok {
		if len(entities) == 1 {
			if entities[0] != entity11 {
				t.Errorf("[1] AddEntity Error.Entity exp:%s got:%s", entity11.GetName(), entities[0].GetName())
			}
		} else {
			t.Errorf("[1] AddEntity Error.Entities exp:1 got:%d", len(entities))
		}
	} else {
		t.Errorf("[1] AddEntity Error.Scene exp:%s got:nil", scene1.GetName())
	}

	// Add other entity in same scene.
	gotError = got.AddEntity(scene1, entity21)
	if gotError != nil {
		t.Errorf("[2] AddEntity Error exp:nil got:%t", gotError)
	}
	gotEntities = got.GetEntities()
	if len(gotEntities) != 1 {
		t.Errorf("[2] AddEntity Error.GetEntities exp:2 got:%d", len(gotEntities))
	}
	if entities, ok := gotEntities[scene1.GetName()]; ok {
		if len(entities) == 2 {
			if entities[1] != entity21 {
				t.Errorf("[2] AddEntity Error.Entity exp:%s got:%s", entity21.GetName(), entities[0].GetName())
			}
		} else {
			t.Errorf("[2] AddEntity Error.Entities exp:2 got:%d", len(entities))
		}
	} else {
		t.Errorf("[2] AddEntity Error.Scene exp:%s got:nil", scene1.GetName())
	}

	// Add other entity in other scene.
	gotError = got.AddEntity(scene2, entity21)
	if gotError != nil {
		t.Errorf("[3] AddEntity Error exp:nil got:%t", gotError)
	}
	gotEntities = got.GetEntities()
	if len(gotEntities) != 2 {
		t.Errorf("[3] AddEntity Error.GetEntities exp:2 got:%d", len(gotEntities))
	}
	if entities, ok := gotEntities[scene2.GetName()]; ok {
		if len(entities) == 1 {
			if entities[0] != entity21 {
				t.Errorf("[3] AddEntity Error.Entity exp:%s got:%s", entity21.GetName(), entities[0].GetName())
			}
		} else {
			t.Errorf("[3] AddEntity Error.Entities exp:1 got:%d", len(entities))
		}
	} else {
		t.Errorf("[3] AddEntity Error.Scene exp:%s got:nil", scene2.GetName())
	}

	// Add entity with focus disable.
	entity11.SetFocusEnable(false)
	gotError = got.AddEntity(scene1, entity11)
	if gotError == nil {
		t.Errorf("[4] AddEntity Error exp:error got:%v", gotError)
	}
}

func TestFocusManagerRemoveEntity(t *testing.T) {

	// Remove entity with focus disable
	got := engine.NewFocusManager()
	_ = got.AddEntity(scene1, entity11)
	entity11.SetFocusEnable(false)
	gotError := got.RemoveEntity(scene1, entity11)
	if gotError == nil {
		t.Errorf("[1] RemoveEntity Error exp:error got:%v", gotError)
	}
	entity11.SetFocusEnable(true)

	// Add entity in one scene and remove it.
	got = engine.NewFocusManager()
	_ = got.AddEntity(scene1, entity11)
	gotError = got.RemoveEntity(scene1, entity11)
	if gotError != nil {
		t.Errorf("[2] RemoveEntity Error exp:nil got:%v", gotError)
	}
	if len(got.GetEntities()) != 0 {
		t.Errorf("[2] RemoveEntity Error.GetEntities exp:0 got:%d", len(got.GetEntities()))
	}

	// Add two entities in same scene and remove the first one.
	_ = got.AddEntity(scene1, entity11)
	_ = got.AddEntity(scene1, entity21)
	gotError = got.RemoveEntity(scene1, entity11)
	gotEntities := got.GetEntities()
	if len(gotEntities) != 1 {
		t.Errorf("[3] RemoveEntity Error.GetEntities exp:1 got:%d", len(gotEntities))
	}
	if entities, ok := gotEntities[scene1.GetName()]; ok {
		if len(entities) == 1 {
			if entities[0] != entity21 {
				t.Errorf("[3] RemoveEntity Error.Entity exp:%s got:%s", entity21.GetName(), entities[0].GetName())
			}
		} else {
			t.Errorf("[3] RemoveEntity Error.Entities exp:1 got:%d", len(entities))
		}
	} else {
		t.Errorf("[3] RemoveEntity Error.Scene exp:%s got:nil", scene1.GetName())
	}
}

func TestFocusManagerRemoveScene(t *testing.T) {
	got := engine.NewFocusManager()
	_ = got.AddEntity(scene1, entity11)
	_ = got.AddEntity(scene2, entity21)

	got.RemoveScene(scene1)
	gotEntities := got.GetEntities()
	if len(gotEntities) != 1 {
		t.Errorf("[1] RemoveScene Error.GetEntities exp:2 got:%d", len(gotEntities))
	}
}
