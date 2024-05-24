package main

import (
	"github.com/jrecuero/thengine/pkg/api"
	"github.com/jrecuero/thengine/pkg/engine"
	"github.com/jrecuero/thengine/pkg/widgets"
)

const (
	CursorName = "widget/cursor/1"
)

type Cursor struct {
	*widgets.Widget
}

func NewCursor(position *api.Point) *Cursor {
	cursor := &Cursor{
		Widget: widgets.NewWidget(CursorName, position, api.NewSize(1, 1), &TheStyleWhiteOverBlack),
	}
	cell := engine.NewCell(&TheStyleWhiteOverBlack, 'x')
	cursor.GetCanvas().SetCellAt(nil, cell)
	return cursor
}
