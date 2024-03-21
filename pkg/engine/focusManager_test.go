package engine_test

import (
	"testing"

	"github.com/jrecuero/thengine/pkg/engine"
)

var (
	scene1   *engine.Scene  = engine.NewScene("scene/1", nil)
	scene2   *engine.Scene  = engine.NewScene("scene/2", nil)
	entity11 *engine.Entity = engine.NewNamedEntity("entity1/1")
	entity12 *engine.Entity = engine.NewNamedEntity("entity1/2")
	entity13 *engine.Entity = engine.NewNamedEntity("entity1/3")
	entity21 *engine.Entity = engine.NewNamedEntity("entity2/1")
	entity22 *engine.Entity = engine.NewNamedEntity("entity2/2")
	entity23 *engine.Entity = engine.NewNamedEntity("entity2/3")
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

func TestFocusManagerUpdateFocusForScene(t *testing.T) {
	// Update to the proper focus for all entities as single-focus.
	entity11.SetFocusEnable(true)
	entity11.SetFocusType(engine.SingleFocus)
	entity12.SetFocusEnable(true)
	entity12.SetFocusType(engine.SingleFocus)
	entity13.SetFocusEnable(true)
	entity13.SetFocusType(engine.SingleFocus)
	entity21.SetFocusEnable(true)
	entity21.SetFocusType(engine.SingleFocus)
	entity22.SetFocusEnable(true)
	entity22.SetFocusType(engine.SingleFocus)
	entity23.SetFocusEnable(true)
	entity23.SetFocusType(engine.SingleFocus)

	got := engine.NewFocusManager()
	_ = got.AddEntity(scene1, entity11)
	_ = got.AddEntity(scene1, entity12)
	_ = got.AddEntity(scene1, entity13)
	_ = got.AddEntity(scene2, entity21)
	_ = got.AddEntity(scene2, entity22)
	_ = got.AddEntity(scene2, entity23)

	// Update focus the first time for any scene.
	gotError := got.UpdateFocusForScene(scene1)
	if gotError != nil {
		t.Errorf("[1] UpdateFocusForScene Error exp:nil got:%+v", gotError)
		return
	}
	withFocus := got.GetEntitiesWithFocus()
	if gotWithFocus, ok := withFocus[scene1.GetName()]; ok {
		if len(gotWithFocus) != 1 {
			t.Errorf("[1] UpdateFocusForScene Error.LenWithFocus exp:1 got:%d", len(gotWithFocus))
		}
		if gotWithFocus[0] != entity11 {
			t.Errorf("[1] UpdateFocusForScene Error.EntityWitFocus exp:%s got:%s", entity11.GetName(), gotWithFocus[0].GetName())
		}
		entities := got.GetEntities()
		if len(entities[scene1.GetName()]) != 2 {
			t.Errorf("[1] UpdateFocusForScene Error.LenEntities exp:2 got:%d", len(entities))
		}
		for _, gotEntity := range entities[scene1.GetName()] {
			if gotEntity.GetName() == entity11.GetName() {
				t.Errorf("[1] UpdateFocusForScene Error.Entity exp:nil got:%s", gotEntity.GetName())
			}
		}
	} else {
		t.Errorf("[1] UpdateFocusForScene Error.Scene exp:[]IEntity got:nil")
		return
	}

	// Update focus again for the same scene.
	gotError = got.UpdateFocusForScene(scene1)
	if gotError != nil {
		t.Errorf("[2] UpdateFocusForScene Error exp:nil got:%+v", gotError)
		return
	}
	withFocus = got.GetEntitiesWithFocus()
	if gotWithFocus, ok := withFocus[scene1.GetName()]; ok {
		if len(gotWithFocus) != 1 {
			t.Errorf("[2] UpdateFocusForScene Error.LenWithFocus exp:2 got:%d", len(gotWithFocus))
		}
		if gotWithFocus[0] != entity12 {
			t.Errorf("[2] UpdateFocusForScene Error.EntityWitFocus exp:%s got:%s", entity12.GetName(), gotWithFocus[0].GetName())
		}
		entities := got.GetEntities()
		if len(entities[scene1.GetName()]) != 2 {
			t.Errorf("[2] UpdateFocusForScene Error.LenEntities exp:2 got:%d", len(entities))
		}
		if gotEntity := entities[scene1.GetName()][len(entities)-1].GetName(); gotEntity != entity11.GetName() {
			t.Errorf("[2] UpdateFocusForScene Error.LastEntity exp:%s got:%s", entity11.GetName(), gotEntity)
		}
		for _, gotEntity := range entities[scene1.GetName()] {
			if gotEntity.GetName() == entity12.GetName() {
				t.Errorf("[2] UpdateFocusForScene Error.Entity exp:nil got:%s", gotEntity.GetName())
			}
		}
	} else {
		t.Errorf("[2] UpdateFocusForScene Error.Scene exp:[]IEntity got:nil")
		return
	}

	// Update focus the first time for the other scene.
	gotError = got.UpdateFocusForScene(scene2)
	if gotError != nil {
		t.Errorf("[3] UpdateFocusForScene Error exp:nil got:%+v", gotError)
		return
	}
	withFocus = got.GetEntitiesWithFocus()
	if gotWithFocus, ok := withFocus[scene2.GetName()]; ok {
		if len(gotWithFocus) != 1 {
			t.Errorf("[3] UpdateFocusForScene Error.LenWithFocus exp:1 got:%d", len(gotWithFocus))
		}
		if gotWithFocus[0] != entity21 {
			t.Errorf("[3] UpdateFocusForScene Error.EntityWitFocus exp:%s got:%s", entity21.GetName(), gotWithFocus[0].GetName())
		}
		entities := got.GetEntities()
		if len(entities[scene2.GetName()]) != 2 {
			t.Errorf("[3] UpdateFocusForScene Error.LenEntities exp:2 got:%d", len(entities))
		}
		for _, gotEntity := range entities[scene2.GetName()] {
			if gotEntity.GetName() == entity21.GetName() {
				t.Errorf("[3] UpdateFocusForScene Error.Entity exp:nil got:%s", gotEntity.GetName())
			}
		}
	} else {
		t.Errorf("[3] UpdateFocusForScene Error.Scene exp:[]IEntity got:nil")
		return
	}

	// Update focus again for the second scene.
	gotError = got.UpdateFocusForScene(scene2)
	if gotError != nil {
		t.Errorf("[4] UpdateFocusForScene Error exp:nil got:%+v", gotError)
		return
	}
	withFocus = got.GetEntitiesWithFocus()
	if gotWithFocus, ok := withFocus[scene2.GetName()]; ok {
		if len(gotWithFocus) != 1 {
			t.Errorf("[4] UpdateFocusForScene Error.LenWithFocus exp:2 got:%d", len(gotWithFocus))
		}
		if gotWithFocus[0] != entity22 {
			t.Errorf("[4] UpdateFocusForScene Error.EntityWitFocus exp:%s got:%s", entity22.GetName(), gotWithFocus[0].GetName())
		}
		entities := got.GetEntities()
		if len(entities[scene2.GetName()]) != 2 {
			t.Errorf("[4] UpdateFocusForScene Error.LenEntities exp:2 got:%d", len(entities))
		}
		if gotEntity := entities[scene2.GetName()][len(entities)-1].GetName(); gotEntity != entity21.GetName() {
			t.Errorf("[4] UpdateFocusForScene Error.LastEntity exp:%s got:%s", entity21.GetName(), gotEntity)
		}
		for _, gotEntity := range entities[scene2.GetName()] {
			if gotEntity.GetName() == entity22.GetName() {
				t.Errorf("[4] UpdateFocusForScene Error.Entity exp:nil got:%s", gotEntity.GetName())
			}
		}
	} else {
		t.Errorf("[4] UpdateFocusForScene Error.Scene exp:[]IEntity got:nil")
		return
	}

	// Update to the proper focus for all entities as single-focus and
	// multi-focus.
	entity11.SetFocusEnable(true)
	entity11.SetFocusType(engine.SingleFocus)
	entity12.SetFocusEnable(true)
	entity12.SetFocusType(engine.MultiFocus)
	entity13.SetFocusEnable(true)
	entity13.SetFocusType(engine.SingleFocus)
	entity21.SetFocusEnable(true)
	entity21.SetFocusType(engine.SingleFocus)
	entity22.SetFocusEnable(true)
	entity22.SetFocusType(engine.SingleFocus)
	entity23.SetFocusEnable(true)
	entity23.SetFocusType(engine.MultiFocus)

	got = engine.NewFocusManager()
	_ = got.AddEntity(scene1, entity11)
	_ = got.AddEntity(scene1, entity12)
	_ = got.AddEntity(scene1, entity13)
	_ = got.AddEntity(scene2, entity21)
	_ = got.AddEntity(scene2, entity22)
	_ = got.AddEntity(scene2, entity23)

	// Update focus the first time for any scene with multi-focus entities.
	gotError = got.UpdateFocusForScene(scene1)
	if gotError != nil {
		t.Errorf("[5] UpdateFocusForScene Error exp:nil got:%+v", gotError)
		return
	}
	withFocus = got.GetEntitiesWithFocus()
	if gotWithFocus, ok := withFocus[scene1.GetName()]; ok {
		if len(gotWithFocus) != 2 {
			t.Errorf("[5] UpdateFocusForScene Error.LenWithFocus exp:2 got:%d", len(gotWithFocus))
		}
		if gotWithFocus[0] != entity12 {
			t.Errorf("[5] UpdateFocusForScene Error.EntityWitFocus exp:%s got:%s", entity12.GetName(), gotWithFocus[0].GetName())
		}
		if gotWithFocus[1] != entity11 {
			t.Errorf("[5] UpdateFocusForScene Error.EntityWitFocus exp:%s got:%s", entity11.GetName(), gotWithFocus[1].GetName())
		}
		entities := got.GetEntities()
		if len(entities[scene1.GetName()]) != 1 {
			t.Errorf("[5] UpdateFocusForScene Error.LenEntities exp:1 got:%d", len(entities))
		}
		for _, gotEntity := range entities[scene1.GetName()] {
			if gotEntity.GetName() == entity11.GetName() {
				t.Errorf("[5] UpdateFocusForScene Error.Entity exp:nil got:%s", gotEntity.GetName())
			}
			if gotEntity.GetName() == entity12.GetName() {
				t.Errorf("[5] UpdateFocusForScene Error.Entity exp:nil got:%s", gotEntity.GetName())
			}
		}
	} else {
		t.Errorf("[5] UpdateFocusForScene Error.Scene exp:[]IEntity got:nil")
		return
	}

	// Update focus again for first scene with multi-focus entities.
	gotError = got.UpdateFocusForScene(scene1)
	if gotError != nil {
		t.Errorf("[6] UpdateFocusForScene Error exp:nil got:%+v", gotError)
		return
	}
	withFocus = got.GetEntitiesWithFocus()
	if gotWithFocus, ok := withFocus[scene1.GetName()]; ok {
		if len(gotWithFocus) != 2 {
			t.Errorf("[6] UpdateFocusForScene Error.LenWithFocus exp:2 got:%d", len(gotWithFocus))
		}
		if gotWithFocus[0] != entity12 {
			t.Errorf("[6] UpdateFocusForScene Error.EntityWitFocus exp:%s got:%s", entity12.GetName(), gotWithFocus[0].GetName())
		}
		if gotWithFocus[1] != entity13 {
			t.Errorf("[6] UpdateFocusForScene Error.EntityWitFocus exp:%s got:%s", entity13.GetName(), gotWithFocus[1].GetName())
		}
		entities := got.GetEntities()
		if len(entities[scene1.GetName()]) != 1 {
			t.Errorf("[6] UpdateFocusForScene Error.LenEntities exp:1 got:%d", len(entities))
		}
		for _, gotEntity := range entities[scene1.GetName()] {
			if gotEntity.GetName() == entity13.GetName() {
				t.Errorf("[6] UpdateFocusForScene Error.Entity exp:nil got:%s", gotEntity.GetName())
			}
			if gotEntity.GetName() == entity12.GetName() {
				t.Errorf("[6] UpdateFocusForScene Error.Entity exp:nil got:%s", gotEntity.GetName())
			}
		}
	} else {
		t.Errorf("[6] UpdateFocusForScene Error.Scene exp:[]IEntity got:nil")
		return
	}
}
