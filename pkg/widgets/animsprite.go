// animSprite.go contains all resources requird to create an animates sprite.
// AnimSprite is based in a Sprite, where every frame is an Sprite which can
// have a different size. That is the difference with AnimWidget which has a
// size size because Canvas is used for every frame.
package widgets

import (
	"github.com/gdamore/tcell/v2"
	"github.com/jrecuero/thengine/pkg/api"
	"github.com/jrecuero/thengine/pkg/engine"
	"github.com/jrecuero/thengine/pkg/tools"
)

// -----------------------------------------------------------------------------
//
// AnimSprite
//
// -----------------------------------------------------------------------------

// AnimSprite structure defines all attributes and methods for any basic
// animated widget.
type AnimSprite struct {
	*Sprite
	frames        engine.CellFrames
	frameTraverse int
	isfrozen      bool
	isshuffle     bool
}

// NewAnimSprite function creates a new AnimSprite instance.
func NewAnimSprite(name string, position *api.Point, frames engine.CellFrames, initFrame int) *AnimSprite {
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
	defer w.Entity.Update(event, scene)
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
