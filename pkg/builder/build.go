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

func buildWall(isXAxe EAxe, fixAxe int, length int, cell *engine.Cell,
	doorPlace EDoorPlace, doorWide int) ([]*widgets.SpriteCell, *Door) {

	spriteCells := []*widgets.SpriteCell{}
	var spriteCell *widgets.SpriteCell
	var door *Door
	if doorPlace != NoDoor {
		door = NewDoor(doorPlace, doorWide, nil)
	}
	for x := 0; x < length; x++ {
		if !(door != nil && isInDoorSpace(x, length, doorWide)) {
			if isXAxe == AxeX {
				spriteCell = widgets.NewSpriteCell(api.NewPoint(x, fixAxe), cell)
			} else {
				spriteCell = widgets.NewSpriteCell(api.NewPoint(fixAxe, x), cell)
			}

			spriteCells = append(spriteCells, spriteCell)
		}
	}
	return spriteCells, door
}

func BuildHWall(y int, w int, cell *engine.Cell, doorPlace EDoorPlace,
	doorWide int) ([]*widgets.SpriteCell, *Door) {

	return buildWall(AxeX, y, w, cell, doorPlace, doorWide)

}

func BuildVWall(x int, h int, cell *engine.Cell, doorPlace EDoorPlace,
	doorWide int) ([]*widgets.SpriteCell, *Door) {
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
	cell *engine.Cell, isDoors []bool, doorsWide []int) *Room {

	spriteCells := []*widgets.SpriteCell{}
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
		spriteCells = append(spriteCells, wall...)
		if door != nil {
			door.SetHooksInWall(entry.wallOrigin, entry.length)
			doors = append(doors, door)
		}
	}

	sprite := widgets.NewSprite(name, position, spriteCells)
	sprite.SetSolid(true)
	room := &Room{
		Sprite: sprite,
		doors:  doors,
	}
	return room

}

func ConnectRooms(name string, doorA *Door, doorB *Door, cell *engine.Cell) *widgets.Sprite {
	spriteA := BuildLine("", doorA.hook.hookA, doorB.hook.hookA, cell)
	spriteB := BuildLine("", doorA.hook.hookB, doorB.hook.hookB, cell)
	spriteCells := spriteA.GetSpriteCells()
	spriteCells = append(spriteCells, spriteB.GetSpriteCells()...)
	sprite := widgets.NewSprite(name, api.NewPoint(0, 0), spriteCells)
	sprite.SetSolid(true)
	return sprite
}

func BuildRoom(name string, position *api.Point, size *api.Size,
	cell *engine.Cell, opts ...any) *widgets.Sprite {

	var doors []bool
	if len(opts) != 0 {
		doors = opts[0].([]bool)
	}
	spriteCells := []*widgets.SpriteCell{}
	var spriteCell *widgets.SpriteCell
	w, h := size.Get()
	for x := 0; x < w; x++ {
		if !(doors[0] && isMiddle(x, w)) {
			spriteCell = widgets.NewSpriteCell(api.NewPoint(x, 0), cell)
			spriteCells = append(spriteCells, spriteCell)
		}
		if !(doors[1] && isMiddle(x, w)) {
			spriteCell = widgets.NewSpriteCell(api.NewPoint(x, h-1), cell)
			spriteCells = append(spriteCells, spriteCell)
		}
	}
	for y := 1; y < h-1; y++ {
		if !(doors[2] && isMiddle(y, h)) {
			spriteCell = widgets.NewSpriteCell(api.NewPoint(0, y), cell)
			spriteCells = append(spriteCells, spriteCell)
		}
		if !(doors[3] && isMiddle(y, h)) {
			spriteCell = widgets.NewSpriteCell(api.NewPoint(w-1, y), cell)
			spriteCells = append(spriteCells, spriteCell)
		}
	}
	sprite := widgets.NewSprite(name, position, spriteCells)
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
	cell *engine.Cell, opts ...any) *widgets.Sprite {

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

	spriteCells := []*widgets.SpriteCell{}
	var spriteCell *widgets.SpriteCell
	if axeX {
		y := []int{originY - wideA, destY + wideB}
		for x := originX; x <= destX; x++ {
			spriteCell = widgets.NewSpriteCell(api.NewPoint(x, y[0]), cell)
			spriteCells = append(spriteCells, spriteCell)
			spriteCell = widgets.NewSpriteCell(api.NewPoint(x, y[1]), cell)
			spriteCells = append(spriteCells, spriteCell)
		}
	} else if axeY {
		x := []int{originX - wideA, destX + wideB}
		for y := originY; y <= destY; y++ {
			spriteCell = widgets.NewSpriteCell(api.NewPoint(x[0], y), cell)
			spriteCells = append(spriteCells, spriteCell)
			spriteCell = widgets.NewSpriteCell(api.NewPoint(x[1], y), cell)
			spriteCells = append(spriteCells, spriteCell)
		}
	}
	sprite := widgets.NewSprite(name, api.NewPoint(0, 0), spriteCells)
	sprite.SetSolid(true)
	return sprite
}

func BuildLine(name string, origin *api.Point, dest *api.Point,
	cell *engine.Cell, opts ...any) *widgets.Sprite {

	axeX, axeY := getAxe(origin, dest)
	originX, originY := origin.Get()
	destX, destY := dest.Get()
	if !axeX && !axeY {
		return nil
	}
	spriteCells := []*widgets.SpriteCell{}
	var spriteCell *widgets.SpriteCell
	if axeX {
		y := originY
		for x := originX; x <= destX; x++ {
			spriteCell = widgets.NewSpriteCell(api.NewPoint(x, y), cell)
			spriteCells = append(spriteCells, spriteCell)
		}
	} else if axeY {
		x := originX
		for y := originY; y <= destY; y++ {
			spriteCell = widgets.NewSpriteCell(api.NewPoint(x, y), cell)
			spriteCells = append(spriteCells, spriteCell)
		}
	}
	sprite := widgets.NewSprite(name, api.NewPoint(0, 0), spriteCells)
	sprite.SetSolid(true)
	return sprite
}
