// animSprite.go contains all resources requird to create an animates sprite.
package widgets

import (
	"github.com/gdamore/tcell/v2"
	"github.com/jrecuero/thengine/pkg/api"
	"github.com/jrecuero/thengine/pkg/engine"
	"github.com/jrecuero/thengine/pkg/tools"
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
	isfrozen      bool
	isshuffle     bool
}

// NewAnimSprite function creates a new AnimSprite instance.
func NewAnimSprite(name string, position *api.Point, frames []*SpriteFrame, initFrame int) *AnimSprite {
	widget := &AnimSprite{
		Sprite:        NewSprite(name, position, nil),
		frames:        frames,
		frameTraverse: initFrame,
		isfrozen:      false,
		isshuffle:     false,
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

func (w *AnimSprite) Freeze(index int) {
	w.isfrozen = true
	if index >= 0 && index < len(w.frames) {
		w.frameTraverse = index
		w.updateSprite()
	}
}

func (w *AnimSprite) Shuffle() {
	w.isshuffle = true
}

func (w *AnimSprite) UnFreeze() {
	w.isfrozen = false
}

func (w *AnimSprite) UnShuffle() {
	w.isshuffle = false
}

// Update method updates the entity instance.
func (w *AnimSprite) Update(event tcell.Event, scene engine.IScene) {
	if w.IsActive() {
		frame := w.frames[w.frameTraverse]
		if w.isfrozen {
			return
		}
		if frame.Inc() {
			frame.Reset()
			if w.isshuffle {
				w.frameTraverse = tools.RandomRing.Intn(len(w.frames))
			} else {
				w.frameTraverse = (w.frameTraverse + 1) % len(w.frames)
			}
			w.updateSprite()
		}
	}
}

var _ engine.IObject = (*AnimSprite)(nil)
var _ engine.IFocus = (*AnimSprite)(nil)
var _ engine.IEntity = (*AnimSprite)(nil)
