// tileMap.go contains all attributes and methods required to implement a
// generic tilemap widget. A timemap has a canvas at a given position and size
// and camera, defined by an offset and a size that will display only a part of
// the tileMap.
package widgets

import (
	"github.com/gdamore/tcell/v2"
	"github.com/jrecuero/thengine/pkg/api"
	"github.com/jrecuero/thengine/pkg/engine"
)

// -----------------------------------------------------------------------------
//
// TileMap
//
// -----------------------------------------------------------------------------

// TileMap structure defines all attributes and methods required for any
// generic tilemap widget.
type TileMap struct {
	*Widget
	cameraOffset *api.Point
	cameraSize   *api.Size
}

// NewTimeMap function creates a new TileMap instance widget.
func NewTileMap(name string, position *api.Point, size *api.Size, style *tcell.Style, cameraOffset *api.Point, cameraSize *api.Size) *TileMap {
	tileMap := &TileMap{
		Widget:       NewWidget(name, position, size, style),
		cameraOffset: cameraOffset,
		cameraSize:   cameraSize,
	}
	return tileMap
}

// -----------------------------------------------------------------------------
// TileMap private methods
// -----------------------------------------------------------------------------

// -----------------------------------------------------------------------------
// TileMap public methods
// -----------------------------------------------------------------------------

func (t *TileMap) DistanceToTileMapEdgesX(tileMapPos *api.Point) (bool, int, int) {
	var x, y int
	ok := t.IsTileMapPosInside(tileMapPos)
	if ok {
		x, y = tileMapPos.X, t.GetSize().W-tileMapPos.X
	}
	return ok, x, y
}

func (t *TileMap) DistanceToTileMapEdgesY(tileMapPos *api.Point) (bool, int, int) {
	var x, y int
	ok := t.IsTileMapPosInside(tileMapPos)
	if ok {
		x, y = tileMapPos.Y, t.GetSize().H-tileMapPos.Y
	}
	return ok, x, y
}

func (t *TileMap) DistanceToCameraEdgesX(tileMapPos *api.Point) (bool, int, int) {
	var x, y int
	ok := t.IsTileMapPosInside(tileMapPos)
	if ok {
		x, y = tileMapPos.X-t.cameraOffset.X, t.cameraOffset.X+t.cameraSize.W-tileMapPos.X
	}
	return ok, x, y
}

func (t *TileMap) DistanceToCameraEdgesY(tileMapPos *api.Point) (bool, int, int) {
	var x, y int
	ok := t.IsTileMapPosInside(tileMapPos)
	if ok {
		x, y = tileMapPos.Y-t.cameraOffset.Y, t.cameraOffset.Y+t.cameraSize.H-tileMapPos.Y
	}
	return ok, x, y
}

func (t *TileMap) GetCameraOffset() *api.Point {
	return t.cameraOffset
}

func (t *TileMap) GetCameraSize() *api.Size {
	return t.cameraSize
}

func (t *TileMap) GetTileMapPosFromScreenPos(position *api.Point) *api.Point {
	screenX, screenY := position.Get()
	tileMapOriginX, tileMapOriginY := t.GetPosition().Get()
	if (screenX < tileMapOriginX) || (screenY < tileMapOriginY) {
		return nil
	}
	offsetX, offsetY := t.cameraOffset.Get()
	x := screenX - tileMapOriginX + offsetX
	y := screenY - tileMapOriginY + offsetY
	if (x >= t.GetSize().W) || (y >= t.GetSize().H) {
		return nil
	}
	return api.NewPoint(x, y)
}

func (t *TileMap) GetScreenPosFromTileMapPos(position *api.Point) *api.Point {
	if !t.IsTileMapPosInside(position) {
		return nil
	}
	tileMapPosX, tileMapPosY := position.Get()
	offsetX, offsetY := t.cameraOffset.Get()
	tileMapOriginX, tileMapOriginY := t.GetPosition().Get()
	return api.NewPoint(tileMapPosX-offsetX+tileMapOriginX, tileMapPosY-offsetY+tileMapOriginY)
}

func (t *TileMap) IsTileMapPosInside(tileMapPos *api.Point) bool {
	tileMapPosX, tileMapPosY := tileMapPos.Get()
	offsetX, offsetY := t.cameraOffset.Get()
	if (tileMapPosX < offsetX) || (tileMapPosY < offsetY) {
		return false
	}
	if (tileMapPosX >= t.GetSize().W) || (tileMapPosY >= t.GetSize().H) {
		return false
	}
	return true
}

func (t *TileMap) Draw(scene engine.IScene) {
	t.GetCanvas().RenderRectAt(scene.GetCamera(), t.cameraOffset, t.cameraSize, t.GetPosition())
}

func (t *TileMap) SetCameraOffset(offset *api.Point) bool {
	offsetX, offsetY := offset.Get()
	// camera offset can never go below tile map origin
	if (offsetX < 0) || (offsetY < 0) {
		return false
	}
	sizeW, sizeH := t.GetSize().Get()
	cameraW, cameraH := t.cameraSize.Get()
	// camera offset can never go further than right and bottom tile map edges.
	if (offsetX > (sizeW - cameraW)) || (offsetY > (sizeH - cameraH)) {
		return false
	}
	t.cameraOffset = offset
	return true
}

func (t *TileMap) SetCameraSize(size *api.Size) {
	t.cameraSize = size
}
