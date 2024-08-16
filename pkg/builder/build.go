package builder

import (
	"github.com/jrecuero/thengine/pkg/api"
	"github.com/jrecuero/thengine/pkg/engine"
	"github.com/jrecuero/thengine/pkg/widgets"
)

type EDoorPlace int

const (
	TopDoor EDoorPlace = iota
	BottomDoor
	LeftDoor
	RightDoor
	NoDoor
)

type EAxe int

const (
	AxeX EAxe = iota
	AxeY
)

type DoorHook struct {
	hookA *api.Point
	hookB *api.Point
}

func NewDoorHook(hookA *api.Point, hookB *api.Point) *DoorHook {
	return &DoorHook{
		hookA: hookA,
		hookB: hookB,
	}
}

type Door struct {
	place EDoorPlace
	wide  int
	hook  *DoorHook
}

func NewDoor(place EDoorPlace, wide int, hook *DoorHook) *Door {
	return &Door{
		place: place,
		wide:  wide,
		hook:  hook,
	}
}

func (d *Door) SetHooksInWall(wallOrigin *api.Point, wallLen int) {
	hookA := api.ClonePoint(wallOrigin)
	hookB := api.ClonePoint(wallOrigin)
	start, end := getDoorHooksInWall(wallLen, d.wide)
	switch d.place {
	case TopDoor:
		hookA.AddScale(start, -1)
		hookB.AddScale(end, -1)
	case BottomDoor:
		hookA.AddScale(start, 1)
		hookB.AddScale(end, 1)
	case LeftDoor:
		hookA.AddScale(-1, start)
		hookB.AddScale(-1, end)
	case RightDoor:
		hookA.AddScale(1, start)
		hookB.AddScale(1, end)
	case NoDoor:
		return
	}
	d.hook = NewDoorHook(hookA, hookB)
}

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

func isMiddle(x int, length int) bool {
	middle := length / 2
	//if (length % 2) != 0 {
	//    middle++
	//}
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

func BuildHWall(y int, w int, cell *engine.Cell, doorPlace EDoorPlace, doorWide int) ([]*widgets.SpriteCell, *Door) {
	//spriteCells := []*widgets.SpriteCell{}
	//var spriteCell *widgets.SpriteCell
	//var door *Door
	//if doorPlace != NoDoor {
	//    door = NewDoor(doorPlace, doorWide, nil)
	//}
	//for x := 0; x < w; x++ {
	//    if !(door != nil && isMiddle(x, w)) {
	//        spriteCell = widgets.NewSpriteCell(api.NewPoint(x, y), cell)
	//        spriteCells = append(spriteCells, spriteCell)
	//    }
	//}
	//return spriteCells, door
	return buildWall(0, y, w, cell, doorPlace, doorWide)

}

func BuildVWall(x int, h int, cell *engine.Cell, doorPlace EDoorPlace, doorWide int) ([]*widgets.SpriteCell, *Door) {
	//spriteCells := []*widgets.SpriteCell{}
	//var spriteCell *widgets.SpriteCell
	//var door *Door
	//if doorPlace != NoDoor {
	//    door = NewDoor(doorPlace, doorWide, nil)
	//}
	//for y := 0; y < h; y++ {
	//    if !(door != nil && isMiddle(y, h)) {
	//        spriteCell = widgets.NewSpriteCell(api.NewPoint(x, y), cell)
	//        spriteCells = append(spriteCells, spriteCell)
	//    }
	//}
	//return spriteCells, door
	return buildWall(1, x, h, cell, doorPlace, doorWide)
}

type RoomData struct {
	doorPlace  EDoorPlace
	axe        EAxe
	fix        int
	length     int
	wallOrigin *api.Point
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
	//posANDsize := [][]int{
	//    {int(TopDoor), 0, 0, size.W, x, y},
	//    {int(BottomDoor), 0, size.H - 1, size.W, x, y + size.H - 1},
	//    {int(LeftDoor), 1, 0, size.H, x, y},
	//    {int(RightDoor), 1, size.W - 1, size.H, x + size.W - 1, y},
	//}
	for _, entry := range roomData {
		doorPlace := NoDoor
		if isDoors[entry.doorPlace] {
			doorPlace = entry.doorPlace
		}
		wall, door := buildWall(entry.axe, entry.fix, entry.length, cell, doorPlace, doorsWide[entry.doorPlace])
		spriteCells = append(spriteCells, wall...)
		if door != nil {
			door.SetHooksInWall(entry.wallOrigin, entry.length)
			doors = append(doors, door)
		}
	}
	//for _, entry := range posANDsize {
	//    doorPlace := NoDoor
	//    edoorPlace := EDoorPlace(entry[0])
	//    if isDoors[edoorPlace] {
	//        doorPlace = edoorPlace
	//    }
	//    wall, door := buildWall(entry[1], entry[2], entry[3], cell, doorPlace, doorsWide[edoorPlace])
	//    spriteCells = append(spriteCells, wall...)
	//    if door != nil {
	//        wallOrigin := api.NewPoint(entry[4], entry[5])
	//        door.SetHooksInWall(wallOrigin, entry[3])
	//        doors = append(doors, door)
	//    }
	//}

	//topDoorPlace := NoDoor
	//bottomDoorPlace := NoDoor
	//leftDoorPlace := NoDoor
	//rightDoorPlace := NoDoor

	//if isDoors[TopDoor] {
	//    topDoorPlace = TopDoor
	//}
	//topWall, topDoor := BuildHWall(0, size.W, cell, topDoorPlace, doorsWide[TopDoor])
	//spriteCells = append(spriteCells, topWall...)
	//if topDoor != nil {
	//    doors = append(doors, topDoor)
	//}

	//if isDoors[BottomDoor] {
	//    bottomDoorPlace = BottomDoor
	//}
	//bottomWall, bottomDoor := BuildHWall(size.H-1, size.W, cell, bottomDoorPlace, doorsWide[BottomDoor])
	//spriteCells = append(spriteCells, bottomWall...)
	//if bottomDoor != nil {
	//    doors = append(doors, bottomDoor)
	//}

	//if isDoors[LeftDoor] {
	//    leftDoorPlace = LeftDoor
	//}
	//leftWall, leftDoor := BuildVWall(0, size.H, cell, leftDoorPlace, doorsWide[LeftDoor])
	//spriteCells = append(spriteCells, leftWall...)
	//if leftDoor != nil {
	//    doors = append(doors, leftDoor)
	//}

	//if isDoors[RightDoor] {
	//    rightDoorPlace = RightDoor
	//}
	//rightWall, rightDoor := BuildVWall(size.W-1, size.H, cell, rightDoorPlace, doorsWide[RightDoor])
	//spriteCells = append(spriteCells, rightWall...)
	//if rightDoor != nil {
	//    doors = append(doors, rightDoor)
	//}

	sprite := widgets.NewSprite(name, position, spriteCells)
	sprite.SetSolid(true)
	room := &Room{
		Sprite: sprite,
		doors:  doors,
	}
	return room

}

func BuildRoom(name string, position *api.Point, size *api.Size, cell *engine.Cell, opts ...any) *widgets.Sprite {
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

func BuildCorridor(name string, origin *api.Point, dest *api.Point, cell *engine.Cell, opts ...any) *widgets.Sprite {
	axeX, axeY := getAxe(origin, dest)
	originX, originY := origin.Get()
	destX, destY := dest.Get()
	if !axeX && !axeY {
		return nil
	}
	spriteCells := []*widgets.SpriteCell{}
	var spriteCell *widgets.SpriteCell
	if axeX {
		y := []int{originY - 1, destY + 1}
		for x := originX; x <= destX; x++ {
			spriteCell = widgets.NewSpriteCell(api.NewPoint(x, y[0]), cell)
			spriteCells = append(spriteCells, spriteCell)
			spriteCell = widgets.NewSpriteCell(api.NewPoint(x, y[1]), cell)
			spriteCells = append(spriteCells, spriteCell)
		}
	} else if axeY {
		x := []int{originX - 1, destX + 1}
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

func BuildLine(name string, origin *api.Point, dest *api.Point, cell *engine.Cell, opts ...any) *widgets.Sprite {
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
