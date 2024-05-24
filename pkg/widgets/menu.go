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
	enabled      bool
	menu         *Menu
	callback     WidgetCallback
	callbackArgs WidgetArgs
}

func NewMenuItem(label string) *MenuItem {
	return &MenuItem{
		label:   label,
		enabled: true,
	}
}

func NewExtendedMenuItem(label string, enabled bool, menu *Menu, callback WidgetCallback, args WidgetArgs) *MenuItem {
	return &MenuItem{
		label:        label,
		enabled:      enabled,
		menu:         menu,
		callback:     callback,
		callbackArgs: args,
	}
}

// -----------------------------------------------------------------------------
// MenuItem public methods
// -----------------------------------------------------------------------------

func (m *MenuItem) GetCallback() (WidgetCallback, WidgetArgs) {
	return m.callback, m.callbackArgs
}

func (m *MenuItem) GetLabel() string {
	return m.label
}

func (m *MenuItem) GetMenu() *Menu {
	return m.menu
}

func (m *MenuItem) IsEnabled() bool {
	return m.enabled
}

func (m *MenuItem) SetCallback(calback WidgetCallback, args WidgetArgs) *MenuItem {
	m.callback = calback
	m.callbackArgs = args
	return m
}

func (m *MenuItem) SetEnabled(enabled bool) *MenuItem {
	m.enabled = enabled
	return m
}

func (m *MenuItem) SetLabel(label string) *MenuItem {
	m.label = label
	return m
}

func (m *MenuItem) SetMenu(menu *Menu) *MenuItem {
	m.menu = menu
	return m
}

// -----------------------------------------------------------------------------
//
// Menu
//
// -----------------------------------------------------------------------------

// Menu structure defines a baseline for any menu widget.
type Menu struct {
	*Widget
	menuItems     []*MenuItem
	menuLabels    []string
	menuItemIndex int
	scroller      *Scroller
	parent        *Menu
}

// NewTopMenu function creates a new Menu instance.
func NewTopMenu(name string, position *api.Point, size *api.Size, style *tcell.Style, menuItems []*MenuItem, menuItemIndex int) *Menu {
	numberOfMenuItems := len(menuItems)
	// Look for the menu item with the largest string.
	maxItemLength := 0
	for _, item := range menuItems {
		maxItemLength = tools.Max(maxItemLength, len(item.GetLabel()))
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
		paddingMenuItems[i] = fmt.Sprintf("%-*s", maxItemLength-1, s.GetLabel())
	}
	// Assign the total number of characters required to contains all menu
	// items.
	totalSelectionLength := numberOfMenuItems * maxItemLength
	tools.Logger.WithField("module", "menu").WithField("function", "NewMenu").Debugf("%s", name)
	menu := &Menu{
		Widget:        NewWidget(name, position, size, style),
		menuItems:     menuItems,
		menuLabels:    paddingMenuItems,
		menuItemIndex: menuItemIndex,
		parent:        nil,
	}
	menu.scroller = NewScroller(totalSelectionLength, size.W-2, maxItemLength)
	menu.SetFocusType(engine.SingleFocus)
	menu.SetFocusEnable(true)
	menu.updateCanvas()
	return menu
}

func NewSubMenu(name string, position *api.Point, size *api.Size, style *tcell.Style, menuItems []*MenuItem, menuItemIndex int, parent *Menu) *Menu {
	selectionsLength := len(menuItems)
	// Add padding for every menu item to fill the whole horizontal length.
	paddingSelections := make([]string, selectionsLength)
	for i, s := range menuItems {
		paddingSelections[i] = fmt.Sprintf("%-*s", size.W-2, s.GetLabel())
	}
	tools.Logger.WithField("module", "menu").WithField("function", "NewSubMenu").Debugf("%s %+v", name, paddingSelections)
	menu := &Menu{
		Widget:        NewWidget(name, position, size, style),
		menuItems:     menuItems,
		menuLabels:    paddingSelections,
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
	return m.menuLabels[index]
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
		menuItem := m.menuItems[m.menuItemIndex]
		if callback, args := menuItem.GetCallback(); callback != nil {
			callback(m, args...)
		}
	}
}

func (m *Menu) nextMenuItem() {
	index := m.menuItemIndex
	for index < (len(m.menuItems) - 1) {
		index++
		if m.menuItems[index].IsEnabled() {
			m.menuItemIndex = index
			m.updateCanvas()
			return
		}
	}
}

func (m *Menu) prevMenuItem() {
	index := m.menuItemIndex
	for index > 0 {
		index--
		if m.menuItems[index].IsEnabled() {
			m.menuItemIndex = index
			m.updateCanvas()
			return
		}
	}
}

func (m *Menu) updateTopMenuCanvas() {
	m.scroller.Update(m.menuItemIndex)
	canvas := m.GetCanvas()
	m.scroller.CreateIter()
	for y := 1; m.scroller.IterHasNext(); {
		index, x := m.scroller.IterGetNext()
		selection := m.getMenuItemLabel(index)
		if index == m.menuItemIndex {
			canvas.WriteStringInCanvasAt(selection, m.GetStyle(), api.NewPoint(x+1, y))
		} else {
			reverseStyle := tools.ReverseStyle(m.GetStyle())
			if !m.menuItems[index].IsEnabled() {
				reverseStyle = tools.SetAttrToStyle(reverseStyle, tcell.AttrDim)
			}
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
		selection := m.getMenuItemLabel(index)
		if index == m.menuItemIndex {
			canvas.WriteStringInCanvasAt(selection, m.GetStyle(), api.NewPoint(x, y+1))
		} else {
			reverseStyle := tools.ReverseStyle(m.GetStyle())
			if !m.menuItems[index].IsEnabled() {
				reverseStyle = tools.SetAttrToStyle(reverseStyle, tcell.AttrDim)
			}
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

// DisableMenuItemForIndex method disables all menu items for given indexes.
func (m *Menu) DisableMenuItemForIndex(indexes ...int) error {
	for _, index := range indexes {
		if index < len(m.menuItems) {
			m.menuItems[index].SetEnabled(false)
		} else {
			return fmt.Errorf("Index %d out of range for menu %s", index, m.GetName())
		}
	}
	return nil
}

// DisableMenuItemForLabel method disables all menu item for given labels.
func (m *Menu) DisableMenuItemsForLabel(labels ...string) error {
	for _, label := range labels {
		if menuItem := m.FindMenuItemByLabel(label); menuItem != nil {
			menuItem.SetEnabled(false)
		} else {
			return fmt.Errorf("Label %s not found for menu %s", label, m.GetName())
		}
	}
	return nil
}

// EnableMenuItemForIndex method enables all menu items for given indexes.
func (m *Menu) EnableMenuItemForIndex(indexes ...int) error {
	for _, index := range indexes {
		if index < len(m.menuItems) {
			m.menuItems[index].SetEnabled(true)
		} else {
			return fmt.Errorf("Index %d out of range for menu %s", index, m.GetName())
		}
	}
	return nil
}

// EnableMenuItemForLabel method enables all menu item for given labels.
func (m *Menu) EnableMenuItemsForLabel(labels ...string) error {
	for _, label := range labels {
		if menuItem := m.FindMenuItemByLabel(label); menuItem != nil {
			menuItem.SetEnabled(true)
		} else {
			return fmt.Errorf("Label %s not found for menu %s", label, m.GetName())
		}
	}
	return nil
}

// FindMenuItemByLabel method finds the menu item for the given label.
func (m *Menu) FindMenuItemByLabel(label string) *MenuItem {
	for _, menuItem := range m.menuItems {
		if menuItem.GetLabel() == label {
			return menuItem
		}
	}
	return nil
}

// GetSelection method returns the option for the selected index.
func (m *Menu) GetSelection() string {
	return strings.TrimSpace(m.getMenuItemLabel(m.menuItemIndex))
}

func (m *Menu) Refresh() {
	m.updateCanvas()
}

// SetSelectionToIndex method sets the menu item index selected to the given
// index.
func (m *Menu) SetSelectionToIndex(index int) error {
	if index < len(m.menuItems) {
		m.menuItemIndex = index
		return nil
	}
	return fmt.Errorf("Index %d out of range for menu %s", index, m.GetName())
}

// SetSelectionToLabel method sets the menu item index selected to the given
// label.
func (m *Menu) SetSelectionToLabel(label string) error {
	for index, menuLabel := range m.menuLabels {
		if menuLabel == label {
			m.menuItemIndex = index
			return nil
		}
	}
	return fmt.Errorf("Label %s not found for menu %s", label, m.GetName())
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
