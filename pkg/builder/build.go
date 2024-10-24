package builder

import (
	"github.com/jrecuero/thengine/pkg/api"
	"github.com/jrecuero/thengine/pkg/engine"
	"github.com/jrecuero/thengine/pkg/widgets"
)

type RoomData struct {
	doorPlace  EDoorPlace
	axe        EAxe
	fix        int
	length     int
	wallOrigin *api.Point
}

func isMiddle(x int, length int) bool {
	middle := length / 2
	return x == middle
}

// getDoorSpaceInWall function returns start and end indexes, as zero base,
// where a door starts and end in a wall.
func getDoorSpaceInWall(wallLen int, doorWide int) (int, int) {
	start := (wallLen - doorWide) / 2 // door start index
	end := start + doorWide - 1       // door end index
	return start, end
}

// getDoorHooksInWall function returns start and end indexex, as zero base,
// where door hooks should be placed.
func getDoorHooksInWall(wallLen int, doorWide int) (int, int) {
	start, end := getDoorSpaceInWall(wallLen, doorWide)
	return start - 1, end + 1
}

// isInDoorSpace function returns if the given index is in the door space for
// the given wall length and with the given door wide. Index is zero base.
func isInDoorSpace(x int, length int, wide int) bool {
	start, end := getDoorSpaceInWall(length, wide)
	return (x >= start) && (x <= end)
}

func buildWall(isXAxe EAxe, fixAxe int, length int, cell engine.ICell,
	doorPlace EDoorPlace, doorWide int) (engine.CellGroup, *Door) {

	cells := engine.CellGroup{}
	var newcell engine.ICell
	var door *Door
	if doorPlace != NoDoor {
		door = NewDoor(doorPlace, doorWide, nil)
	}
	for x := 0; x < length; x++ {
		if !(door != nil && isInDoorSpace(x, length, doorWide)) {
			if isXAxe == AxeX {
				newcell = engine.NewCellAt(cell.GetStyle(), cell.GetRune(), api.NewPoint(x, fixAxe))
			} else {
				newcell = engine.NewCellAt(cell.GetStyle(), cell.GetRune(), api.NewPoint(fixAxe, x))
			}

			cells = append(cells, newcell)
		}
	}
	return cells, door
}

func BuildHWall(y int, w int, cell engine.ICell, doorPlace EDoorPlace,
	doorWide int) (engine.CellGroup, *Door) {

	return buildWall(AxeX, y, w, cell, doorPlace, doorWide)

}

func BuildVWall(x int, h int, cell engine.ICell, doorPlace EDoorPlace,
	doorWide int) (engine.CellGroup, *Door) {
	return buildWall(AxeY, x, h, cell, doorPlace, doorWide)
}

func NewRoomData(doorPlace EDoorPlace, axe EAxe, fix int, length int,
	wallOrigin *api.Point) *RoomData {
	return &RoomData{
		doorPlace:  doorPlace,
		axe:        axe,
		fix:        fix,
		length:     length,
		wallOrigin: wallOrigin,
	}
}

func BuildRoomWithDoors(name string, position *api.Point, size *api.Size,
	cell engine.ICell, isDoors []bool, doorsWide []int) *Room {

	cells := engine.CellGroup{}
	doors := []*Door{}
	x, y := position.Get()
	w, h := size.Get()
	roomData := []*RoomData{
		NewRoomData(TopDoor, AxeX, 0, w, api.NewPoint(x, y)),
		NewRoomData(BottomDoor, AxeX, h-1, w, api.NewPoint(x, y+h-1)),
		NewRoomData(LeftDoor, AxeY, 0, h, api.NewPoint(x, y)),
		NewRoomData(RightDoor, AxeY, w-1, h, api.NewPoint(x+w-1, y)),
	}
	for _, entry := range roomData {
		doorPlace := NoDoor
		if isDoors[entry.doorPlace] {
			doorPlace = entry.doorPlace
		}
		wall, door := buildWall(entry.axe, entry.fix, entry.length, cell, doorPlace,
			doorsWide[entry.doorPlace])
		cells = append(cells, wall...)
		if door != nil {
			door.SetHooksInWall(entry.wallOrigin, entry.length)
			doors = append(doors, door)
		}
	}

	sprite := widgets.NewSprite(name, position, cells)
	sprite.SetSolid(true)
	room := &Room{
		Sprite: sprite,
		doors:  doors,
	}
	return room

}

func ConnectRooms(name string, doorA *Door, doorB *Door, cell engine.ICell) *widgets.Sprite {
	spriteA := BuildLine("", doorA.hook.hookA, doorB.hook.hookA, cell)
	spriteB := BuildLine("", doorA.hook.hookB, doorB.hook.hookB, cell)
	spriteCells := spriteA.GetCells()
	spriteCells = append(spriteCells, spriteB.GetCells()...)
	sprite := widgets.NewSprite(name, api.NewPoint(0, 0), spriteCells)
	sprite.SetSolid(true)
	return sprite
}

func BuildRoom(name string, position *api.Point, size *api.Size,
	cell engine.ICell, opts ...any) *widgets.Sprite {

	var doors []bool
	if len(opts) != 0 {
		doors = opts[0].([]bool)
	}
	cells := engine.CellGroup{}
	var newcell engine.ICell
	w, h := size.Get()
	for x := 0; x < w; x++ {
		if !(doors[0] && isMiddle(x, w)) {
			newcell = engine.NewCellAt(cell.GetStyle(), cell.GetRune(), api.NewPoint(x, 0))
			cells = append(cells, newcell)
		}
		if !(doors[1] && isMiddle(x, w)) {
			newcell = engine.NewCellAt(cell.GetStyle(), cell.GetRune(), api.NewPoint(x, h-1))
			cells = append(cells, newcell)
		}
	}
	for y := 1; y < h-1; y++ {
		if !(doors[2] && isMiddle(y, h)) {
			newcell = engine.NewCellAt(cell.GetStyle(), cell.GetRune(), api.NewPoint(0, y))
			cells = append(cells, newcell)
		}
		if !(doors[3] && isMiddle(y, h)) {
			newcell = engine.NewCellAt(cell.GetStyle(), cell.GetRune(), api.NewPoint(w-1, y))
			cells = append(cells, newcell)
		}
	}
	sprite := widgets.NewSprite(name, position, cells)
	sprite.SetSolid(true)
	return sprite
}

func getAxe(origin *api.Point, dest *api.Point) (bool, bool) {
	axeX, axeY := false, false
	originX, originY := origin.Get()
	destX, destY := dest.Get()
	if originX == destX {
		axeY = true
	} else if originY == destY {
		axeX = true
	} else {
		return false, false
	}
	return axeX, axeY
}

func BuildCorridor(name string, origin *api.Point, dest *api.Point,
	cell engine.ICell, opts ...any) *widgets.Sprite {

	axeX, axeY := getAxe(origin, dest)
	originX, originY := origin.Get()
	destX, destY := dest.Get()
	if !axeX && !axeY {
		return nil
	}
	wideA, wideB := 1, 1
	if len(opts) == 1 {
		wideA = opts[0].(int)
		wideB = opts[0].(int)
	} else if len(opts) == 2 {
		wideA = opts[0].(int)
		wideB = opts[1].(int)
	}

	cells := engine.CellGroup{}
	var newcell engine.ICell
	if axeX {
		y := []int{originY - wideA, destY + wideB}
		for x := originX; x <= destX; x++ {
			newcell = engine.NewCellAt(cell.GetStyle(), cell.GetRune(), api.NewPoint(x, y[0]))
			cells = append(cells, newcell)
			newcell = engine.NewCellAt(cell.GetStyle(), cell.GetRune(), api.NewPoint(x, y[1]))
			cells = append(cells, newcell)
		}
	} else if axeY {
		x := []int{originX - wideA, destX + wideB}
		for y := originY; y <= destY; y++ {
			newcell = engine.NewCellAt(cell.GetStyle(), cell.GetRune(), api.NewPoint(x[0], y))
			cells = append(cells, newcell)
			newcell = engine.NewCellAt(cell.GetStyle(), cell.GetRune(), api.NewPoint(x[1], y))
			cells = append(cells, newcell)
		}
	}
	sprite := widgets.NewSprite(name, api.NewPoint(0, 0), cells)
	sprite.SetSolid(true)
	return sprite
}

func BuildLine(name string, origin *api.Point, dest *api.Point,
	cell engine.ICell, opts ...any) *widgets.Sprite {

	axeX, axeY := getAxe(origin, dest)
	originX, originY := origin.Get()
	destX, destY := dest.Get()
	if !axeX && !axeY {
		return nil
	}
	cells := engine.CellGroup{}
	var newcell engine.ICell
	if axeX {
		y := originY
		for x := originX; x <= destX; x++ {
			newcell = engine.NewCellAt(cell.GetStyle(), cell.GetRune(), api.NewPoint(x, y))
			cells = append(cells, newcell)
		}
	} else if axeY {
		x := originX
		for y := originY; y <= destY; y++ {
			newcell = engine.NewCellAt(cell.GetStyle(), cell.GetRune(), api.NewPoint(x, y))
			cells = append(cells, newcell)
		}
	}
	sprite := widgets.NewSprite(name, api.NewPoint(0, 0), cells)
	sprite.SetSolid(true)
	return sprite
}
