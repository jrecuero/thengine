package widgets_test

import (
	"testing"

	"github.com/gdamore/tcell/v2"
	"github.com/jrecuero/thengine/pkg/api"
	"github.com/jrecuero/thengine/pkg/engine"
	"github.com/jrecuero/thengine/pkg/widgets"
)

var (
	widgetCbCounter int = 0
	widgetCbArgs    []any
)

func widgetCb(entity engine.IEntity, args ...any) bool {
	widgetCbCounter++
	widgetCbArgs = args
	return true
}

func TestWidgetCallback(t *testing.T) {
	style := tcell.StyleDefault.Foreground(tcell.ColorRed).Background(tcell.ColorWhite)
	widget := widgets.NewWidget("test", api.NewPoint(0, 0), api.NewSize(1, 1), &style)
	if widget == nil {
		t.Errorf("[1] NewWidget Error exp:*Widget got:nil")
		return
	}
	widget.SetWidgetCallback(widgetCb, "test", "one")
	gotCb := widget.GetWidgetCallback()
	if gotCb == nil {
		t.Errorf("[1] GetWidgetCallback Error exp:WidgetCallback got:nil")
	}
	gotArgs := widget.GetWidgetCallbackArgs()
	if len(gotArgs) != 2 {
		t.Errorf("[1] GetWidgetCallbackArgs Error.Len exp:2 got:%d", len(gotArgs))
		return
	}
	if gotArgs[0].(string) != "test" {
		t.Errorf("[1] GetWidgetCallbackArgs Error.Args exp:%s got:%s", "test", gotArgs[0].(string))
	}
	if gotArgs[1].(string) != "one" {
		t.Errorf("[1] GetWidgetCallbackArgs Error.Args exp:%s got:%s", "one", gotArgs[1].(string))
	}
	widgetCbCounter = 0
	widgetCbArgs = []any{}
	gotOk := widget.RunCallback()
	if !gotOk {
		t.Errorf("[1] RunCallback Error exp:true got:%t", gotOk)
	}
	if widgetCbCounter != 1 {
		t.Errorf("[1] RunCallback Error.Callback exp:%d got:%d", 1, widgetCbCounter)
	}
	if len(widgetCbArgs) != 2 {
		t.Errorf("[1] GetWidgetCallbackArgs Error.Callback.Args.Len exp:%d got:%d", 2, len(widgetCbArgs))
		return
	}
	if widgetCbArgs[0].(string) != "test" {
		t.Errorf("[1] GetWidgetCallbackArgs Error.Callback.Args exp:%s got:%s", "test", widgetCbArgs[0].(string))
	}
	if widgetCbArgs[1].(string) != "one" {
		t.Errorf("[1] GetWidgetCallbackArgs Error.Callback.Args exp:%s got:%s", "one", widgetCbArgs[1].(string))
	}

	gotOk = widget.RunCallback("two")
	if !gotOk {
		t.Errorf("[2] RunCallback Error exp:true got:%t", gotOk)
	}
	if widgetCbCounter != 2 {
		t.Errorf("[2] RunCallback Error.Callback exp:%d got:%d", 1, widgetCbCounter)
	}
	if len(widgetCbArgs) != 1 {
		t.Errorf("[2] GetWidgetCallbackArgs Error.Callback.Args.Len exp:%d got:%d", 1, len(widgetCbArgs))
		return
	}
	if widgetCbArgs[0].(string) != "two" {
		t.Errorf("[2] GetWidgetCallbackArgs Error.Callback.Args exp:%s got:%s", "two", widgetCbArgs[0].(string))
	}

	widget.SetWidgetCallbackArgs("three", 4, 5)
	gotOk = widget.RunCallback()
	if !gotOk {
		t.Errorf("[3] RunCallback Error exp:true got:%t", gotOk)
	}
	if widgetCbCounter != 3 {
		t.Errorf("[3] RunCallback Error.Callback exp:%d got:%d", 3, widgetCbCounter)
	}
	if len(widgetCbArgs) != 3 {
		t.Errorf("[3] GetWidgetCallbackArgs Error.Callback.Args.Len exp:%d got:%d", 3, len(widgetCbArgs))
		return
	}
	if widgetCbArgs[0].(string) != "three" {
		t.Errorf("[3] GetWidgetCallbackArgs Error.Callback.Args exp:%s got:%s", "three", widgetCbArgs[0].(string))
	}
	if widgetCbArgs[1].(int) != 4 {
		t.Errorf("[3] GetWidgetCallbackArgs Error.Callback.Args exp:%d got:%d", 4, widgetCbArgs[1].(int))
	}
	if widgetCbArgs[2].(int) != 5 {
		t.Errorf("[3] GetWidgetCallbackArgs Error.Callback.Args exp:%d got:%d", 5, widgetCbArgs[2].(int))
	}

	widget.SetWidgetCallback(nil)

	widget.SetWidgetCallbackArgs("four")
	gotOk = widget.RunCallback()
	if !gotOk {
		t.Errorf("[4] RunCallback Error exp:true got:%t", gotOk)
	}
	if widgetCbCounter != 3 {
		t.Errorf("[4] RunCallback Error.Callback exp:%d got:%d", 3, widgetCbCounter)
	}
	if len(widgetCbArgs) != 3 {
		t.Errorf("[4] GetWidgetCallbackArgs Error.Callback.Args.Len exp:%d got:%d", 3, len(widgetCbArgs))
		return
	}
	if widgetCbArgs[0].(string) != "three" {
		t.Errorf("[4] GetWidgetCallbackArgs Error.Callback.Args exp:%s got:%s", "three", widgetCbArgs[0].(string))
	}
	if widgetCbArgs[1].(int) != 4 {
		t.Errorf("[4] GetWidgetCallbackArgs Error.Callback.Args exp:%d got:%d", 4, widgetCbArgs[1].(int))
	}
	if widgetCbArgs[2].(int) != 5 {
		t.Errorf("[4] GetWidgetCallbackArgs Error.Callback.Args exp:%d got:%d", 5, widgetCbArgs[2].(int))
	}
}
