package tdice

import (
	"github.com/jrecuero/thengine/app/rhunedice/assets/tfaces"
	API "github.com/jrecuero/thengine/pkg/api"
	"github.com/jrecuero/thengine/pkg/widgets"
)

type AnimBaseDie struct {
	*widgets.AnimWidget
}

func NewAnimBaseDie(pos *API.Point, ticks int) *AnimBaseDie {
	frames := tfaces.NewAsciiFramesFromAllFaces(ticks)
	animDie := &AnimBaseDie{
		AnimWidget: widgets.NewAnimWidget("assets/tdice/anim-base-die/1", pos, tfaces.AsciiFaceSize, frames, 0),
	}
	animDie.Shuffle()
	return animDie
}
