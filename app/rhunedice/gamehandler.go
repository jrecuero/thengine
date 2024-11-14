package main

import (
	"github.com/gdamore/tcell/v2"
	"github.com/jrecuero/thengine/app/rhunedice/assets/tdice"
	"github.com/jrecuero/thengine/app/rhunedice/assets/tfaces"
	"github.com/jrecuero/thengine/app/rhunedice/assets/twidgets"
	"github.com/jrecuero/thengine/pkg/constants"
	"github.com/jrecuero/thengine/pkg/engine"
	"github.com/jrecuero/thengine/pkg/tools"
	"github.com/jrecuero/thengine/pkg/widgets"
)

var (
	theGameHandler *GameHandler
)

// -----------------------------------------------------------------------------
//
// GameHandler
//
// -----------------------------------------------------------------------------

type GameHandler struct {
	*engine.Entity
	playerDice       []*tdice.AnimBaseDie
	diceSelectWidget *twidgets.SelectWidget
}

func NewGameHandler() *GameHandler {
	if theGameHandler == nil {
		name := "handler/game/1"
		tools.Logger.WithField("module", "gamehandler").
			WithField("function", "NewGameHandler").
			Debugf(name)
		theGameHandler = &GameHandler{
			Entity:           engine.NewHandler(name),
			diceSelectWidget: nil,
		}
		theGameHandler.SetFocusType(engine.SingleFocus)
		theGameHandler.SetFocusEnable(true)
	}
	return theGameHandler
}

// -----------------------------------------------------------------------------
// GameHandler private methods
// -----------------------------------------------------------------------------

// -----------------------------------------------------------------------------
// GameHandler public methods
// -----------------------------------------------------------------------------

func (h *GameHandler) GetDiceSelectWidget() *twidgets.SelectWidget {
	return h.diceSelectWidget
}

func (h *GameHandler) GetPlayerDice() []*tdice.AnimBaseDie {
	return h.playerDice
}

func (h *GameHandler) SetPlayerDice(dice []*tdice.AnimBaseDie) {
	h.playerDice = dice
	var tmp []widgets.IWidget = make([]widgets.IWidget, len(dice))
	for i, d := range dice {
		tmp[i] = d
	}

	h.diceSelectWidget = twidgets.NewHorizontalSelectWidget("dice/widget",
		&constants.RedOverBlack,
		tmp,
		0)
}

func (h *GameHandler) Update(event tcell.Event, scene engine.IScene) {
	if !h.HasFocus() {
		return
	}
	switch ev := event.(type) {
	case *tcell.EventKey:
		switch ev.Key() {
		case tcell.KeyUp:
			for _, die := range h.playerDice {
				die.UnFreeze()
			}
		case tcell.KeyDown:
			for i, die := range h.playerDice {
				die.Freeze()
				frame := die.GetFrame().(*tfaces.RhuneFrame)
				rhune := frame.GetRhune()
				tools.Logger.WithField("module", "gamehandler").
					WithField("struct", "GameHandler").
					WithField("method", "Update").
					Debugf("[%d] rhune is %s", i, rhune.GetName())
			}
		case tcell.KeyLeft:
		case tcell.KeyRight:
			//canvas := h.playerDice[0].GetCanvas()
			//style := canvas.GetStyleAt(api.NewPoint(0, 0))
			//reverseStyle := tools.ReverseStyle(style)
			//for _, die := range h.playerDice {
			//    canvas := die.GetCanvas()
			//    canvas.SetStyleAt(nil, reverseStyle)
			//}
			//for _, die := range h.playerDice {
			//    canvas := die.GetCanvas()
			//    style := canvas.GetStyleAt(api.NewPoint(0, 0))
			//    tools.Logger.WithField("module", "gamehandler").
			//        WithField("struct", "GameHandler").
			//        WithField("method", "Update").
			//        Debugf("new style at (0, 0) is %s", tools.StyleToString(style))
			//}
		case tcell.KeyRune:
			switch ev.Rune() {
			case 'A', 'a':
			case '1':
			}
		}
	}
}
