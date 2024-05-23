// import.go module contains all code related with importing data from JSON
// files into entities.
package engine

import (
	"encoding/json"
	"fmt"
	"os"
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
func ImportEntitiesFromJSON(filename string, builtin IBuiltIn) []IEntity {
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
		ent := builtin.GetClassFromString(mapEntity["class"].(string))
		if err := ent.UnmarshalMap(mapEntity); err != nil {
			panic(fmt.Sprintf("Error unmarshaling entitys %s:%s", filename, err.Error()))
		}
		canvas := NewCanvas(ent.GetSize())
		ch := mapEntity["ch"].(string)
		if len(ch) != 1 {
			canvas.WriteStringInCanvas(mapEntity["ch"].(string), ent.GetStyle())
		} else {
			cell := NewCell(ent.GetStyle(), rune(ch[0]))
			canvas.FillWithCell(cell)
		}
		ent.SetCanvas(canvas)
		result = append(result, ent)
	}

	return result
}
