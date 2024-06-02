package engine_test

import (
	"reflect"
	"testing"

	"github.com/jrecuero/thengine/pkg/engine"
)

type Wall struct {
	*engine.Entity
}

func NewWall() *Wall {
	return &Wall{
		Entity: engine.NewEmptyEntity(),
	}
}

type BuiltInTest struct {
	engine.IBuiltIn
}

func (b *BuiltInTest) GetClassFromString(className string) engine.IEntity {
	switch className {
	case "Wall":
		return NewWall()
	default:
		return engine.NewEmptyEntity()
	}
}

func TestImportGetEntitiesFromJSON(t *testing.T) {
	filename := "assets/test/entities.json"
	got := engine.ImportEntitiesFromJSON(filename, nil, &BuiltInTest{})
	if len(got) != 1 {
		t.Errorf("[0] GetEntitiesFromJSON len error exp:%d got:%d", 1, len(got))
	}
	if _, ok := got[0].(*Wall); !ok {
		t.Errorf("[0] GetEntitiesFromJSON class error exp:%s got:%s", "*Wall", reflect.TypeOf(got[0]).String())
	}
	if got[0].GetName() != "entity/name/1" {
		t.Errorf("[0] GetEntitiesFromJSON name error exp:%s got:%s", "entity/name/1", got[0].GetName())
	}
}
