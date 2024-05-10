// animSprite.go contains all resources requird to create an animates sprite.
package widgets

import (
	"github.com/gdamore/tcell/v2"
	"github.com/jrecuero/thengine/pkg/api"
	"github.com/jrecuero/thengine/pkg/engine"
)

// -----------------------------------------------------------------------------
//
// SpriteFrame
//
// -----------------------------------------------------------------------------

// SpriteFrame structure defines the sprite frame to be used and how many ticks
// has to be mantained.
type SpriteFrame struct {
	spriteCells []*SpriteCell
	maxTicks    int
	ticks       int
}

// NewSpriteFrameInfo function creates a new SpriteFrame instance.
func NewSpriteFrame(spriteCells []*SpriteCell, maxTicks int) *SpriteFrame {
	return &SpriteFrame{
		spriteCells: spriteCells,
		maxTicks:    maxTicks,
		ticks:       0,
	}
}

// -----------------------------------------------------------------------------
// SpriteFrame public methods
// -----------------------------------------------------------------------------

// GetSpriteCells method returns the canvas instance number.
func (f *SpriteFrame) GetSpriteCells() []*SpriteCell {
	return f.spriteCells
}

// Inc method increase the actual counter for the sprite frame instance and
// returns if that counter has reached the maxTicks.
func (f *SpriteFrame) Inc() bool {
	f.ticks++
	return f.ticks >= f.maxTicks
}

// Reset method resets the sprite frame counter value.
func (f *SpriteFrame) Reset() {
	f.ticks = 0
}

// -----------------------------------------------------------------------------
//
// AnimSprite
//
// -----------------------------------------------------------------------------

// AnimSprite structure defines all attributes and methods for any basic
// animated widget.
type AnimSprite struct {
	*Sprite
	frames        []*SpriteFrame
	frameTraverse int
}

// NewAnimSprite function creates a new AnimSprite instance.
func NewAnimSprite(name string, position *api.Point, frames []*SpriteFrame, initFrame int) *AnimSprite {
	widget := &AnimSprite{
		Sprite:        NewSprite(name, position, nil),
		frames:        frames,
		frameTraverse: initFrame,
	}
	widget.updateSprite()
	return widget
}

// -----------------------------------------------------------------------------
// AnimSprite private methods
// -----------------------------------------------------------------------------

func (w *AnimSprite) updateSprite() {
	spriteCells := w.frames[w.frameTraverse].GetSpriteCells()
	w.SetSpriteCells(spriteCells)
}

// -----------------------------------------------------------------------------
// AnimSprite public methods
// -----------------------------------------------------------------------------

// Update method updates the entity instance.
func (w *AnimSprite) Update(event tcell.Event, scene engine.IScene) {
	if w.IsActive() {
		frame := w.frames[w.frameTraverse]
		if frame.Inc() {
			frame.Reset()
			w.frameTraverse = (w.frameTraverse + 1) % len(w.frames)
			w.updateSprite()
		}
	}
}

var _ engine.IObject = (*AnimSprite)(nil)
var _ engine.IFocus = (*AnimSprite)(nil)
var _ engine.IEntity = (*AnimSprite)(nil)
