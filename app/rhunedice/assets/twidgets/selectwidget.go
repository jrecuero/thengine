package twidgets

import (
	"github.com/gdamore/tcell/v2"
	"github.com/jrecuero/thengine/app/rhunedice/assets/tdice"
	"github.com/jrecuero/thengine/app/rhunedice/assets/tfaces"
	"github.com/jrecuero/thengine/pkg/api"
	"github.com/jrecuero/thengine/pkg/engine"
	"github.com/jrecuero/thengine/pkg/tools"
	"github.com/jrecuero/thengine/pkg/widgets"
)

// -----------------------------------------------------------------------------
//
// SelectWidget
//
// -----------------------------------------------------------------------------

// SelectWidget struct defines a new Widget that contains a list of widgets
// that could be horizontal or vertical aligned, all with the same size.
// Selection process will update the widget color to the given color and one or
// more widgets can be selected.
type SelectWidget struct {
	*widgets.Widget
	infocus        bool
	originalStyles []*tcell.Style
	selected       []int
	selectionIndex int
	widgets        []widgets.IWidget
}

// -----------------------------------------------------------------------------
// package [SelectWidget] public functions
// -----------------------------------------------------------------------------

func NewHorizontalSelectWidget(name string, style *tcell.Style, w []widgets.IWidget, s int) *SelectWidget {
	// Get origin position and size from the first widget in the list.
	origin := w[0].GetPosition()
	selectWidget := &SelectWidget{
		Widget:         widgets.NewWidget(name, origin, nil, style),
		infocus:        false,
		originalStyles: make([]*tcell.Style, len(w)),
		selected:       []int{},
		selectionIndex: s,
		widgets:        w,
	}
	selectWidget.SetFocusType(engine.SingleFocus)
	selectWidget.SetFocusEnable(true)
	//selectWidget.updateCanvas()
	return selectWidget
}

// -----------------------------------------------------------------------------
// SelectWidget private methods
// -----------------------------------------------------------------------------

func (w *SelectWidget) execute(args ...any) {
	switch args[0].(string) {
	case "left":
		if w.selectionIndex > 0 {
			w.updateCanvasForIndex(w.selectionIndex)
			w.selectionIndex -= 1
			w.updateCanvas()
		}
	case "right":
		if w.selectionIndex < len(w.widgets)-1 {
			w.updateCanvasForIndex(w.selectionIndex)
			w.selectionIndex += 1
			w.updateCanvas()
		}
	case "select":
		if _, ok := tools.Contains(w.selected, w.selectionIndex); ok {
			w.selected = tools.Removes(w.selected, w.selectionIndex)
		} else {
			w.selected = append(w.selected, w.selectionIndex)
		}
		widget := w.widgets[w.selectionIndex].(*tdice.AnimBaseDie)
		frame := widget.GetFrame().(*tfaces.RhuneFrame)
		tools.Logger.WithField("module", "selectwidget").
			WithField("struct", "SelectWidget").
			WithField("method", "execute").
			Debugf("[%d] widget selected %s selected-list: %v", w.selectionIndex, frame.GetRhune().GetName(), w.selected)
		w.updateCanvasForSelect()
	default:
	}
}

func (w *SelectWidget) updateCanvasForSelect() {
	var style *tcell.Style
	index := w.selectionIndex
	widget := w.widgets[index]
	canvas := widget.GetCanvas()
	if _, ok := tools.Contains(w.selected, index); ok {
		// if the entry is being selected, it is in the reversed style, so it
		// have to be reversed to the original style before being stored.
		reversed := canvas.GetStyleAt(api.NewPoint(0, 0))
		w.originalStyles[index] = tools.ReverseStyle(reversed)
		style = w.GetStyle()
	} else {
		style = w.originalStyles[index]
	}
	//tools.Logger.WithField("module", "selectwidget").
	//    WithField("struct", "SelectWidget").
	//    WithField("method", "updateCanvasForSelect").
	//    Debugf("style is %s", tools.StyleToString(style))
	reverseStyle := tools.ReverseStyle(style)
	canvas.SetStyleAt(nil, reverseStyle)
	observerManager := engine.GetEngine().GetObserverManager()
	observerManager.NotifyObservers(w.GetName(), w.selected)
}

func (w *SelectWidget) updateCanvasForIndex(index int) {
	var style *tcell.Style
	widget := w.widgets[index]
	canvas := widget.GetCanvas()
	//if _, ok := tools.Contains(w.selected, index); ok {
	//    style = w.GetStyle()
	//} else {
	//    style = canvas.GetStyleAt(api.NewPoint(0, 0))
	//}
	style = canvas.GetStyleAt(api.NewPoint(0, 0))
	reverseStyle := tools.ReverseStyle(style)
	canvas.SetStyleAt(nil, reverseStyle)
}

// updateCanvas method reverse the style for the selected widget.
func (w *SelectWidget) updateCanvas() {
	w.updateCanvasForIndex(w.selectionIndex)
}

// -----------------------------------------------------------------------------
// SelectWidget public methods
// -----------------------------------------------------------------------------

func (w *SelectWidget) AcquireFocus() (bool, error) {
	if ok, err := w.Widget.Focus.AcquireFocus(); !ok {
		return ok, err
	}
	w.infocus = true
	return true, nil
}

func (w *SelectWidget) GetSelected() []int {
	return w.selected
}

func (w *SelectWidget) GetSelectIndex() int {
	return w.selectionIndex
}

func (w *SelectWidget) GetWidgets() []widgets.IWidget {
	return w.widgets
}

//func (w *SelectWidget) ReleaseFocus() (bool, error) {
//    w.infocus = false
//    return w.Widget.Focus.AcquireFocus()
//}

func (w *SelectWidget) Update(event tcell.Event, scene engine.IScene) {
	defer w.Entity.Update(event, scene)
	if !w.HasFocus() {
		return
	}
	if w.infocus {
		w.updateCanvas()
		w.infocus = false
	}
	actions := []*widgets.KeyboardAction{
		widgets.NewKeyboardActionForKey(tcell.KeyLeft, w.execute, []any{"left"}),
		widgets.NewKeyboardActionForKey(tcell.KeyRight, w.execute, []any{"right"}),
		widgets.NewKeyboardActionForKey(tcell.KeyEnter, w.execute, []any{"select"}),
	}
	w.HandleKeyboardForActions(event, actions)
}

var _ engine.IObject = (*SelectWidget)(nil)
var _ engine.IFocus = (*SelectWidget)(nil)
var _ engine.IEntity = (*SelectWidget)(nil)
