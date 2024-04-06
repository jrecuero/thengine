// menu.go module contains all attributes and methods required to implement a
// basic and generic menu.
package widgets

import (
	"fmt"
	"strings"

	"github.com/gdamore/tcell/v2"
	"github.com/jrecuero/thengine/pkg/api"
	"github.com/jrecuero/thengine/pkg/engine"
	"github.com/jrecuero/thengine/pkg/tools"
)

// -----------------------------------------------------------------------------
//
// Menu
//
// -----------------------------------------------------------------------------

// Menu structure defines a baseline for any menu widget.
type Menu struct {
	*Widget
	menuItems     []string
	menuItemIndex int
	scroller      *Scroller
}

// NewMenu function creates a new Menu instance.
func NewMenu(name string, position *api.Point, size *api.Size, style *tcell.Style, menuItems []string, menuItemIndex int) *Menu {
	numberOfMenuItems := len(menuItems)
	// Look for the menu item with the largest string.
	maxItemLength := 0
	for _, item := range menuItems {
		maxItemLength = tools.Max(maxItemLength, len(item))
	}
	// Reassign the maximum menu item length if the horizontal size is greater
	// than the number of items by the maximum number of character for any menu
	// item.
	if (maxItemLength * numberOfMenuItems) < (size.W - 2) {
		maxItemLength = (size.W - 2) / numberOfMenuItems
	}
	// Add padding for every menu item to fill the whole horizontal length.
	paddingMenuItems := make([]string, numberOfMenuItems)
	for i, s := range menuItems {
		paddingMenuItems[i] = fmt.Sprintf("%-*s", maxItemLength-1, s)
	}
	// Assign the total number of characters required to contains all menu
	// items.
	totalSelectionLength := numberOfMenuItems * maxItemLength
	tools.Logger.WithField("module", "menu").WithField("function", "NewMenu").Debugf("%s", name)
	menu := &Menu{
		Widget:        NewWidget(name, position, size, style),
		menuItems:     paddingMenuItems,
		menuItemIndex: menuItemIndex,
	}
	menu.scroller = NewScroller(totalSelectionLength, size.W-2, maxItemLength)
	menu.SetFocusType(engine.SingleFocus)
	menu.SetFocusEnable(true)
	menu.updateCanvas()
	return menu
}

// -----------------------------------------------------------------------------
// Menu private methods
// -----------------------------------------------------------------------------

func (m *Menu) execute(args ...any) {
	tools.Logger.WithField("module", "menu").WithField("function", "execute").Debugf("%s %+v", m.GetName(), args)
	switch args[0].(string) {
	case "up":
		fallthrough
	case "left":
		if m.menuItemIndex > 0 {
			m.menuItemIndex--
			m.updateCanvas()
		}
	case "down":
		fallthrough
	case "right":
		if m.menuItemIndex < (len(m.menuItems) - 1) {
			m.menuItemIndex++
			m.updateCanvas()
		}
	case "run":
	}
}

// updateCanvas method updates the list box canvas with proper menuItems to be
// displayed and the proper selected option.
func (m *Menu) updateCanvas() {
	// update the scroller with the selection index.
	m.scroller.Update(m.menuItemIndex)
	tools.Logger.WithField("module", "menu").WithField("function", "updateCanvas").Debugf("[%d, %d]", m.scroller.StartSelection, m.scroller.EndSelection)
	canvas := m.GetCanvas()
	m.scroller.CreateIter()
	for y := 1; m.scroller.IterHasNext(); {
		index, x := m.scroller.IterGetNext()
		selection := m.menuItems[index]
		tools.Logger.WithField("module", "menu").WithField("function", "updateCanvas").Debugf("selection %s", selection)
		if index == m.menuItemIndex {
			canvas.WriteStringInCanvasAt(selection, m.GetStyle(), api.NewPoint(x, y))
		} else {
			reverseStyle := tools.ReverseStyle(m.GetStyle())
			canvas.WriteStringInCanvasAt(selection, reverseStyle, api.NewPoint(x, y))
		}
	}
	//for index, x, y := m.scroller.StartSelection, 1, 1; index <= m.scroller.EndSelection; index, y = index+1, y+1 {
	//    selection := m.menuItems[index]
	//    tools.Logger.WithField("module", "menu").WithField("function", "updateCanvas").Debugf("selection %s", selection)
	//    if index == m.menuItemIndex {
	//        canvas.WriteStringInCanvasAt(selection, m.GetStyle(), api.NewPoint(x, y))
	//    } else {
	//        reverseStyle := tools.ReverseStyle(m.GetStyle())
	//        canvas.WriteStringInCanvasAt(selection, reverseStyle, api.NewPoint(x, y))
	//    }
	//}
}

// -----------------------------------------------------------------------------
// Menu public methods
// -----------------------------------------------------------------------------

// GetSelection method returns the option for the selected index.
func (m *Menu) GetSelection() string {
	return strings.TrimSpace(m.menuItems[m.menuItemIndex])
}

// Update method executes all listbox functionality every tick time. Keyboard
// inut is scanned in order to move the selection index and proceed to select
// any option.
func (m *Menu) Update(event tcell.Event) {
	if !m.HasFocus() {
		return
	}
	actions := []*KeyboardAction{
		{
			Key:      tcell.KeyDown,
			Callback: m.execute,
			Args:     []any{"down"},
		},
		{
			Key:      tcell.KeyUp,
			Callback: m.execute,
			Args:     []any{"up"},
		},
		{
			Key:      tcell.KeyLeft,
			Callback: m.execute,
			Args:     []any{"left"},
		},
		{
			Key:      tcell.KeyRight,
			Callback: m.execute,
			Args:     []any{"right"},
		},
		{
			Key:      tcell.KeyEnter,
			Callback: m.execute,
			Args:     []any{"run"},
		},
	}
	m.HandleKeyboardForActions(event, actions)
}

var _ engine.IEntity = (*Menu)(nil)
