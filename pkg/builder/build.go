package builder

import (
	"github.com/jrecuero/thengine/pkg/api"
	"github.com/jrecuero/thengine/pkg/engine"
	"github.com/jrecuero/thengine/pkg/widgets"
)

func isMiddle(x int, length int) bool {
	middle := length / 2
	//if (length % 2) != 0 {
	//    middle++
	//}
	return x == middle
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
