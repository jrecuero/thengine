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
// TileMap public methods
// -----------------------------------------------------------------------------

func (t *TileMap) DistanceToTileMapEdgesX(tileMapPos *api.Point) (int, int) {
	return tileMapPos.X, t.GetSize().W - tileMapPos.X
}

func (t *TileMap) DistanceToTileMapEdgesY(tileMapPos *api.Point) (int, int) {
	return tileMapPos.Y, t.GetSize().H - tileMapPos.Y
}

func (t *TileMap) DistanceToCameraEdgesX(tileMapPos *api.Point) (int, int) {
	return tileMapPos.X - t.cameraOffset.X, t.cameraOffset.X + t.cameraSize.W - tileMapPos.X
}

func (t *TileMap) DistanceToCameraEdgesY(tileMapPos *api.Point) (int, int) {
	return tileMapPos.Y - t.cameraOffset.Y, t.cameraOffset.Y + t.cameraSize.H - tileMapPos.Y
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
	offsetX, offsetY := t.cameraOffset.Get()
	return api.NewPoint(screenX-tileMapOriginX+offsetX, screenY-tileMapOriginY+offsetY)
}

func (t *TileMap) GetScreenPosFromTileMapPos(position *api.Point) *api.Point {
	tileMapPosX, tileMapPosY := position.Get()
	offsetX, offsetY := t.cameraOffset.Get()
	if (tileMapPosX < offsetX) || (tileMapPosY < offsetY) {
		return nil
	}
	tileMapOriginX, tileMapOriginY := t.GetPosition().Get()
	return api.NewPoint(tileMapPosX-offsetX+tileMapOriginX, tileMapPosY-offsetY+tileMapOriginY)
}

func (t *TileMap) Draw(camera engine.ICamera) {
	t.GetCanvas().RenderRectAt(camera, t.cameraOffset, t.cameraSize, t.GetPosition())
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
