// sprite.go contains all required feature to implement an sprite. Sprite does
// not have any canvas, it has a list of SpriteCells which contains the
// position and the cell to be render in such position.
package widgets

import (
	"github.com/jrecuero/thengine/pkg/api"
	"github.com/jrecuero/thengine/pkg/engine"
)

const (
	AtTheEnd int = -1
)

// -----------------------------------------------------------------------------
//
// SpriteCell
//
// -----------------------------------------------------------------------------

// SpriteCell struct defines the basic structure of an Sprite with a position
// and a cell.
type SpriteCell struct {
	position *api.Point
	cell     *engine.Cell
}

// NewSpriteCell function creates a new SpriteCell instance.
func NewSpriteCell(position *api.Point, cell *engine.Cell) *SpriteCell {
	return &SpriteCell{
		position: position,
		cell:     cell,
	}
}

// -----------------------------------------------------------------------------
// SpriteCell public methods
// -----------------------------------------------------------------------------

// GetCell method returns the cell from the SpriteCell instance.
func (s *SpriteCell) GetCell() *engine.Cell {
	return s.cell
}

// GetPosition method returs the position from the SpriteCell instance.
func (s *SpriteCell) GetPosition() *api.Point {
	return s.position
}

// SetCell method sets the cell in a SpriteCell instance.
func (s *SpriteCell) SetCell(cell *engine.Cell) {
	s.cell = cell
}

// SetPosition method sets the position in a SpriteCell instance.
func (s *SpriteCell) SetPosition(position *api.Point) {
	s.position = position
}

// -----------------------------------------------------------------------------
//
// Sprite
//
// -----------------------------------------------------------------------------

// Sprite struture defines an Sprite widget which is represented by a list of
// SpriteCells drawn in the screeen.
type Sprite struct {
	*Widget
	spriteCells []*SpriteCell
}

// NewSprite function creates a new Sprite instance.
func NewSprite(name string, position *api.Point, spriteCells []*SpriteCell) *Sprite {
	sprite := &Sprite{
		Widget:      NewWidget(name, position, nil, nil),
		spriteCells: spriteCells,
	}
	sprite.SetCanvas(nil)
	return sprite
}

// -----------------------------------------------------------------------------
// Sprite public methods
// -----------------------------------------------------------------------------

func (s *Sprite) AddSpriteCellAt(atIndex int, spriteCell *SpriteCell) {
	spriteCells := make([]*SpriteCell, len(s.spriteCells)+1)
	// if the atIndex is equal to AtTheEnd(-1), add the sprite cell at the end.
	if atIndex == AtTheEnd {
		atIndex = len(s.spriteCells)
	}
	for i, cell := range s.spriteCells {
		index := i
		if i >= atIndex {
			index++
		}
		spriteCells[index] = cell
	}
	spriteCells[atIndex] = spriteCell
	s.spriteCells = spriteCells
}

func (s *Sprite) Draw(scene engine.IScene) {
	if s.IsVisible() {
		for _, spriteCell := range s.spriteCells {
			position := api.ClonePoint(s.GetPosition())
			position.Add(spriteCell.GetPosition())
			scene.GetCamera().RenderCellAt(position, spriteCell.GetCell())
		}
	}
}

func (s *Sprite) GetCollider() *engine.Collider {
	points := []*api.Point{}
	for _, spriteCell := range s.spriteCells {
		position := api.ClonePoint(s.GetPosition())
		position.Add(spriteCell.GetPosition())
		points = append(points, position)
	}
	return engine.NewCollider(nil, points)
}

func (s *Sprite) GetSpriteCells() []*SpriteCell {
	return s.spriteCells
}

func (s *Sprite) RemoveSpriteCellAt(atIndex int) *SpriteCell {
	if atIndex == AtTheEnd {
		atIndex = len(s.spriteCells) - 1
	}
	if (atIndex < 0) || (atIndex >= len(s.spriteCells)) {
		return nil
	}
	spriteCell := s.spriteCells[atIndex]
	s.spriteCells = append(s.spriteCells[:atIndex], s.spriteCells[atIndex+1:]...)
	return spriteCell
}

func (s *Sprite) SetSpriteCells(spriteCells []*SpriteCell) {
	s.spriteCells = spriteCells
}

var _ engine.IObject = (*Sprite)(nil)
var _ engine.IFocus = (*Sprite)(nil)
var _ engine.IEntity = (*Sprite)(nil)
