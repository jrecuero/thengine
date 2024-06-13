package main

import (
	"fmt"

	"github.com/jrecuero/thengine/pkg/api"
	"github.com/jrecuero/thengine/pkg/engine"
	"github.com/jrecuero/thengine/pkg/tools"
	"github.com/jrecuero/thengine/pkg/widgets"
)

const (
	SpriteHandlerName = "handler/sprite-handler/1"
)

type SpriteHandler struct {
	*engine.Entity
	counter int
	sprite  *widgets.Sprite
}

func NewSpriteHandler() *SpriteHandler {
	handler := &SpriteHandler{
		Entity:  engine.NewHandler(SpriteHandlerName),
		counter: 0,
		sprite:  nil,
	}
	handler.SetFocusEnable(false)
	return handler
}

func (h *SpriteHandler) AddPoint(point *api.Point, cell *engine.Cell) {
	if h.sprite == nil {
		return
	}
	spriteCell := widgets.NewSpriteCell(point, cell)
	h.sprite.AddSpriteCellAt(widgets.AtTheEnd, spriteCell)
	tools.Logger.WithField("module", "SpriteHandler").
		WithField("method", "AddPoint").
		Debugf("%s %+#v", h.sprite.GetName(), spriteCell)
}

func (h *SpriteHandler) EndSprite() {
	h.sprite = nil
}

func (h *SpriteHandler) GetSprite() *widgets.Sprite {
	return h.sprite
}

func (h *SpriteHandler) NewSprite(name string) *widgets.Sprite {
	h.counter++
	if name == "" {
		name = fmt.Sprintf("sprite_%02d", h.counter)
	}
	h.sprite = widgets.NewSprite(name, api.NewPoint(0, 0), nil)
	tools.Logger.WithField("module", "SpriteHandler").
		WithField("method", "NewSprite").
		Debugf("%s %+#v", h.sprite.GetName(), h.sprite)
	return h.sprite
}
