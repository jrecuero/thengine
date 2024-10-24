package widgets

import (
	"github.com/gdamore/tcell/v2"
	"github.com/jrecuero/thengine/pkg/api"
	"github.com/jrecuero/thengine/pkg/engine"
)

var (
	BoxSingleLine []rune = []rune{tcell.RuneULCorner,
		tcell.RuneURCorner,
		tcell.RuneLLCorner,
		tcell.RuneLRCorner,
		tcell.RuneHLine,
		tcell.RuneVLine}
	BoxDoubleLine []rune = []rune{'╔', '╗', '╚', '╝', '═', '║'}
)

// -----------------------------------------------------------------------------
//
// Box
//
// -----------------------------------------------------------------------------

type Box struct {
	*Sprite
	pattern []rune
}

func NewBox(name string, position *api.Point, size *api.Size, style *tcell.Style, pattern []rune) *Box {
	box := &Box{
		Sprite:  NewSprite(name, position, nil),
		pattern: pattern,
	}
	box.SetPosition(position)
	box.SetSize(size)
	box.SetStyle(style)
	box.updateBox()
	return box
}

func (b *Box) updateBox() {
	var cell *engine.Cell
	var ul, ur, ll, lr, hl, vl rune
	if len(b.pattern) == 1 {
		ul = b.pattern[0]
		ur = b.pattern[0]
		ll = b.pattern[0]
		lr = b.pattern[0]
		hl = b.pattern[0]
		vl = b.pattern[0]
	} else if len(b.pattern) == 6 {
		ul = b.pattern[0]
		ur = b.pattern[1]
		ll = b.pattern[2]
		lr = b.pattern[3]
		hl = b.pattern[4]
		vl = b.pattern[5]
	} else {
		return
	}
	w, h := b.GetSize().Get()
	style := b.GetStyle()
	for col := 1; col < (w - 1); col++ {
		cell = engine.NewCellAt(style, hl, api.NewPoint(col, 0))
		b.AddCellAt(-1, cell)
		cell = engine.NewCellAt(style, hl, api.NewPoint(col, h-1))
		b.AddCellAt(-1, cell)
	}
	for row := 1; row < (h - 1); row++ {
		cell := engine.NewCell(style, vl)
		cell = engine.NewCellAt(style, vl, api.NewPoint(0, row))
		b.AddCellAt(-1, cell)
		cell = engine.NewCellAt(style, vl, api.NewPoint(w-1, row))
		b.AddCellAt(-1, cell)
	}
	cell = engine.NewCellAt(style, ul, api.NewPoint(0, 0))
	b.AddCellAt(-1, cell)
	cell = engine.NewCellAt(style, ur, api.NewPoint(w-1, 0))
	b.AddCellAt(-1, cell)
	cell = engine.NewCellAt(style, ll, api.NewPoint(0, h-1))
	b.AddCellAt(-1, cell)
	cell = engine.NewCellAt(style, lr, api.NewPoint(w-1, h-1))
	b.AddCellAt(-1, cell)
}
