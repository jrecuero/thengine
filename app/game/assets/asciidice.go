package assets

import (
	"fmt"

	"github.com/gdamore/tcell/v2"
	"github.com/jrecuero/thengine/pkg/api"
	"github.com/jrecuero/thengine/pkg/engine"
	"github.com/jrecuero/thengine/pkg/tools"
	"github.com/jrecuero/thengine/pkg/widgets"
)

const (
	Ascii0 = "000000\n00  00\n00  00\n00  00\n000000"
	Ascii1 = "1111  \n  11  \n  11  \n  11  \n111111"
	Ascii2 = "222222\n     2\n222222\n2     \n222222"
	Ascii3 = "333333\n    33\n333333\n    33\n333333"
	Ascii4 = "44  44\n44  44\n444444\n    44\n    44"
	Ascii5 = "555555\n55    \n555555\n    55\n555555"
	Ascii6 = "666666\n66    \n666666\n66  66\n666666"
	Ascii7 = "777777\n    77\n    77\n    77\n    77"
	Ascii8 = "888888\n88  88\n888888\n88  88\n888888"
	Ascii9 = "999999\n99  99\n999999\n    99\n999999"
)

var (
	whiteOnBlack = tcell.StyleDefault.Foreground(tcell.ColorWhite).Background(tcell.ColorBlack)

	AsciiCanvas0 = engine.NewCanvasFromString(Ascii0, &whiteOnBlack)
	AsciiCanvas1 = engine.NewCanvasFromString(Ascii1, &whiteOnBlack)
	AsciiCanvas2 = engine.NewCanvasFromString(Ascii2, &whiteOnBlack)
	AsciiCanvas3 = engine.NewCanvasFromString(Ascii3, &whiteOnBlack)
	AsciiCanvas4 = engine.NewCanvasFromString(Ascii4, &whiteOnBlack)
	AsciiCanvas5 = engine.NewCanvasFromString(Ascii5, &whiteOnBlack)
	AsciiCanvas6 = engine.NewCanvasFromString(Ascii6, &whiteOnBlack)
	AsciiCanvas7 = engine.NewCanvasFromString(Ascii7, &whiteOnBlack)
	AsciiCanvas8 = engine.NewCanvasFromString(Ascii8, &whiteOnBlack)
	AsciiCanvas9 = engine.NewCanvasFromString(Ascii9, &whiteOnBlack)

	AsciiCanvasAllNumbers = []*engine.Canvas{
		AsciiCanvas0,
		AsciiCanvas1,
		AsciiCanvas2,
		AsciiCanvas3,
		AsciiCanvas4,
		AsciiCanvas5,
		AsciiCanvas6,
		AsciiCanvas7,
		AsciiCanvas8,
		AsciiCanvas9,
	}
)

func NewAsciiFrameForNumber(number int, ticks int) *widgets.Frame {
	if number < 0 || number > 9 {
		return nil
	}
	canvas := AsciiCanvasAllNumbers[number]
	return widgets.NewFrameWithCanvas(canvas, ticks)
}

func NewAsciiFramesForAllNumbers(ticks int) []widgets.IFrame {
	var frames []widgets.IFrame = make([]widgets.IFrame, 10)
	for i := 0; i <= 9; i++ {
		frames[i] = NewAsciiFrameForNumber(i, ticks)
	}
	return frames
}

type ShuffleDie struct {
	*widgets.Widget
	faces         int
	frameTraverse int
	ticks         int
	maxTicks      int
}

func NewShuffleDieWidget(name string, position *api.Point, style *tcell.Style, faces int, maxTicks int) *ShuffleDie {
	var size *api.Size
	if faces < 10 {
		size = api.NewSize(3, 3)
	} else {
		size = api.NewSize(4, 3)
	}
	die := &ShuffleDie{
		Widget:        widgets.NewWidget(name, position, size, style),
		faces:         faces,
		frameTraverse: 0,
		ticks:         0,
		maxTicks:      maxTicks,
	}
	die.GetCanvas().WriteRectangleInCanvasAt(nil, nil, die.GetStyle(), engine.CanvasRectSingleLine)
	cell := engine.NewCell(die.GetStyle(), '0')
	die.GetCanvas().SetCellAt(api.NewPoint(1, 1), cell)
	return die
}

func (w *ShuffleDie) Update(tcell.Event, engine.IScene) {
	w.ticks++
	if w.ticks >= w.maxTicks {
		var newface string
		if w.faces < 10 {
			newface = fmt.Sprintf("%d", tools.RandomRing.Intn(w.faces+1))
		} else {
			newface = fmt.Sprintf("%02d", tools.RandomRing.Intn(w.faces+1))
		}
		w.GetCanvas().WriteStringInCanvasAt(newface, w.GetStyle(), api.NewPoint(1, 1))
		w.ticks = 0
	}
}
