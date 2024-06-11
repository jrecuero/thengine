// import.go module contains all code related with importing data from JSON
// files into entities.
package engine

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/jrecuero/thengine/pkg/api"
	"github.com/jrecuero/thengine/pkg/tools"
)

// -----------------------------------------------------------------------------
//
// IBuiltIn
//
// -----------------------------------------------------------------------------

// IBuiltIn interface defines methods required to generate an Entity class from
// a string.
type IBuiltIn interface {
	GetClassFromString(string) IEntity
}

// -----------------------------------------------------------------------------
// Module public methods
// -----------------------------------------------------------------------------

// ImportEntitiesFromJSON function reads all entities in the given JSON file
// and it returns an array of IEntity instances.
func ImportEntitiesFromJSON(filename string, origin *api.Point, builtin IBuiltIn) []IEntity {
	var result []IEntity

	jsonContent, err := os.ReadFile(filename)
	if err != nil {
		panic(fmt.Sprintf("Error reading %s:%s", filename, err.Error()))
	}
	var content []map[string]any
	if err := json.Unmarshal(jsonContent, &content); err != nil {
		panic(fmt.Sprintf("Error unmarshaling %s:%s", filename, err.Error()))
	}
	for _, mapEntity := range content {
		var entity IEntity
		if builtin == nil {
			entity = NewEmptyEntity()
		} else {
			entity = builtin.GetClassFromString(mapEntity["class"].(string))
		}
		if err := entity.UnmarshalMap(mapEntity, origin); err != nil {
			panic(fmt.Sprintf("Error unmarshaling entitys %s:%s", filename, err.Error()))
		}
		canvas := NewCanvas(entity.GetSize())
		ch := mapEntity["ch"].(string)
		if len(ch) != 1 {
			canvas.WriteStringInCanvas(mapEntity["ch"].(string), entity.GetStyle())
		} else {
			cell := NewCell(entity.GetStyle(), rune(ch[0]))
			canvas.FillWithCell(cell)
		}
		entity.SetCanvas(canvas)
		result = append(result, entity)
	}

	return result
}

// ExportEntitiesToJSON function exports given entites to the given JSON file
// in JSON format.
func ExportEntitiesToJSON(filename string, entities []IEntity, origin *api.Point, builtin IBuiltIn) error {
	var result []map[string]any
	for _, entity := range entities {
		if entity == nil {
			continue
		}
		if resultMap, err := entity.MarshalMap(origin); err == nil {
			tools.Logger.WithField("module", "import").
				WithField("function", "ExportEntitiesToJSON").
				Debug(resultMap)
			result = append(result, resultMap)
		} else {
			return err
		}
	}
	if jsonData, err := json.Marshal(result); err == nil {
		if origin != nil {
			x, y := origin.Get()
			filename = fmt.Sprintf("%s_%d_%d.json", filename, x, y)
		} else {
			filename = fmt.Sprintf("%s.json", filename)
		}
		if err = os.WriteFile(filename, jsonData, 0644); err != nil {
			return err
		}
	} else {
		return err
	}
	return nil
}
