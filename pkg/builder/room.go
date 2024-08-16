package builder

import (
	"github.com/jrecuero/thengine/pkg/api"
	"github.com/jrecuero/thengine/pkg/engine"
	"github.com/jrecuero/thengine/pkg/widgets"
)

type Room struct {
	*widgets.Sprite
	doors []*Door
}

func NewRoom(name string, position *api.Point, size *api.Size, cell *engine.Cell, opts ...any) *Room {
	room := &Room{
		Sprite: nil,
		doors:  nil,
	}
	return room
}

func (r *Room) GetDoorAt(place EDoorPlace) *Door {
	for _, door := range r.doors {
		if door.GetPlace() == place {
			return door
		}
	}
	return nil
}

func (r *Room) GetDoors() []*Door {
	return r.doors
}
