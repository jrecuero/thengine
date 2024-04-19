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
// MenuItem
//
// -----------------------------------------------------------------------------

// MenuItem structure defines every item in a menu widget.
type MenuItem struct {
	label        string
	menu         *Menu
	callback     WidgetCallback
	callbackArgs WidgetArgs
}

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
	parent        *Menu
}

// NewTopMenu function creates a new Menu instance.
func NewTopMenu(name string, position *api.Point, size *api.Size, style *tcell.Style, menuItems []string, menuItemIndex int) *Menu {
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
		parent:        nil,
	}
	menu.scroller = NewScroller(totalSelectionLength, size.W-2, maxItemLength)
	menu.SetFocusType(engine.SingleFocus)
	menu.SetFocusEnable(true)
	menu.updateCanvas()
	return menu
}

func NewSubMenu(name string, position *api.Point, size *api.Size, style *tcell.Style, menuItems []string, menuItemIndex int, parent *Menu) *Menu {
	selectionsLength := len(menuItems)
	// Add padding for every menu item to fill the whole horizontal length.
	paddingSelections := make([]string, selectionsLength)
	for i, s := range menuItems {
		paddingSelections[i] = fmt.Sprintf("%-*s", size.W-2, s)
	}
	tools.Logger.WithField("module", "menu").WithField("function", "NewSubMenu").Debugf("%s %+v", name, paddingSelections)
	menu := &Menu{
		Widget:        NewWidget(name, position, size, style),
		menuItems:     paddingSelections,
		menuItemIndex: menuItemIndex,
		parent:        parent,
	}
	menu.scroller = NewVerticalScroller(selectionsLength, size.H-2)
	menu.SetFocusType(engine.SingleFocus)
	menu.SetFocusEnable(true)
	menu.updateCanvas()
	return menu
}

// -----------------------------------------------------------------------------
// Menu private methods
// -----------------------------------------------------------------------------

func (m *Menu) getMenuItemLabel(index int) string {
	return m.menuItems[index]
}

func (m *Menu) execute(args ...any) {
	tools.Logger.WithField("module", "menu").WithField("function", "execute").Debugf("%s %+v", m.GetName(), args)
	switch args[0].(string) {
	case "up":
		if m.parent != nil {
			m.prevMenuItem()
		}
	case "left":
		if m.parent == nil {
			m.prevMenuItem()
		}
	case "down":
		if m.parent != nil {
			m.nextMenuItem()
		}
	case "right":
		if m.parent == nil {
			m.nextMenuItem()
		}
	case "run":
	}
}

func (m *Menu) nextMenuItem() {
	if m.menuItemIndex < (len(m.menuItems) - 1) {
		m.menuItemIndex++
		m.updateCanvas()
	}
}

func (m *Menu) prevMenuItem() {
	if m.menuItemIndex > 0 {
		m.menuItemIndex--
		m.updateCanvas()
	}
}

func (m *Menu) updateTopMenuCanvas() {
	m.scroller.Update(m.menuItemIndex)
	canvas := m.GetCanvas()
	m.scroller.CreateIter()
	for y := 1; m.scroller.IterHasNext(); {
		index, x := m.scroller.IterGetNext()
		selection := m.menuItems[index]
		if index == m.menuItemIndex {
			canvas.WriteStringInCanvasAt(selection, m.GetStyle(), api.NewPoint(x+1, y))
		} else {
			reverseStyle := tools.ReverseStyle(m.GetStyle())
			canvas.WriteStringInCanvasAt(selection, reverseStyle, api.NewPoint(x+1, y))
		}
	}
}

func (m *Menu) updateSubMenuCanvas() {
	m.scroller.Update(m.menuItemIndex)
	canvas := m.GetCanvas()
	m.scroller.CreateIter()
	for x := 1; m.scroller.IterHasNext(); {
		index, y := m.scroller.IterGetNext()
		selection := m.menuItems[index]
		if index == m.menuItemIndex {
			canvas.WriteStringInCanvasAt(selection, m.GetStyle(), api.NewPoint(x, y+1))
		} else {
			reverseStyle := tools.ReverseStyle(m.GetStyle())
			canvas.WriteStringInCanvasAt(selection, reverseStyle, api.NewPoint(x, y+1))
		}
	}
}

// updateCanvas method updates the list box canvas with proper menuItems to be
// displayed and the proper selected option.
func (m *Menu) updateCanvas() {
	if m.parent == nil {
		m.updateTopMenuCanvas()
	} else {
		m.updateSubMenuCanvas()
	}
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
func (m *Menu) Update(event tcell.Event, scene engine.IScene) {
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

var _ engine.IObject = (*Menu)(nil)
var _ engine.IFocus = (*Menu)(nil)
var _ engine.IEntity = (*Menu)(nil)
