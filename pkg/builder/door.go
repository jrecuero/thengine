package builder

import "github.com/jrecuero/thengine/pkg/api"

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
	hook  *DoorHook
	place EDoorPlace
	wide  int
}

func NewDoor(place EDoorPlace, wide int, hook *DoorHook) *Door {
	return &Door{
		place: place,
		wide:  wide,
		hook:  hook,
	}
}

func (d *Door) GetHook() *DoorHook {
	return d.hook
}

func (d *Door) GetPlace() EDoorPlace {
	return d.place
}

func (d *Door) GetWide() int {
	return d.wide
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
