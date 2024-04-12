package widgets_test

import (
	"testing"
	"time"

	"github.com/jrecuero/thengine/pkg/engine"
	"github.com/jrecuero/thengine/pkg/widgets"
)

var (
	simRunning   bool = false
	timerCounter int  = 0
)

func simEngine(timer *widgets.Timer) {
	go func() {
		for simRunning {
			time.Sleep(10 * time.Millisecond)
			timer.Update(nil, nil)
		}
	}()
}

func timerCallback(entity engine.IEntity, args ...any) bool {
	timerCounter++
	return true
}

func TestTimer(t *testing.T) {
	got := widgets.NewTimer("timer/1", 100*time.Millisecond, widgets.OneTimeTimer)
	got.SetWidgetCallback(timerCallback, nil)
	if got == nil {
		t.Errorf("[1] NewTimer Error exp:*Timer got:nil")
		return
	}
	simRunning = true
	simEngine(got)
	got.StartTimer()
	time.Sleep(50 * time.Millisecond)
	if timerCounter != 0 {
		t.Errorf("[1] NewTimer Error.Counter.counting exp:%d got:%d", 0, timerCounter)
		return
	}
	time.Sleep(60 * time.Millisecond)
	if timerCounter != 1 {
		t.Errorf("[1] NewTimer Error.Counter.expired exp:%d got:%d", 1, timerCounter)
		return
	}
	time.Sleep(200 * time.Millisecond)
	if timerCounter != 1 {
		t.Errorf("[1] NewTimer Error.Counter.finished exp:%d got:%d", 1, timerCounter)
		return
	}
	simRunning = false

	timerCounter = 0
	got = widgets.NewTimer("timer/2", 100*time.Millisecond, widgets.ForeverTimer)
	if got == nil {
		t.Errorf("[2] NewTimer Error exp:*Timer got:nil")
		return
	}
	got.SetWidgetCallback(timerCallback, nil)
	simRunning = true
	simEngine(got)
	got.StartTimer()
	time.Sleep(20 * time.Millisecond)
	got.StopTimer()
	if timerCounter != 0 {
		t.Errorf("[2] NewTimer Error.Counter.counting exp:%d got:%d", 0, timerCounter)
		return
	}
	time.Sleep(300 * time.Millisecond)
	if timerCounter != 0 {
		t.Errorf("[2] NewTimer Error.Counter.stopped exp:%d got:%d", 0, timerCounter)
		return
	}
	got.StartTimer()
	time.Sleep(200 * time.Millisecond)
	if timerCounter != 2 {
		t.Errorf("[2] NewTimer Error.Counter.running exp:%d got:%d", 2, timerCounter)
		return
	}
	got.CancelTimer()
	time.Sleep(100 * time.Millisecond)
	simRunning = false

	timerCounter = 0
	got = widgets.NewTimer("timer/3", 100*time.Millisecond, 2)
	if got == nil {
		t.Errorf("[3] NewTimer Error exp:*Timer got:nil")
		return
	}
	got.SetWidgetCallback(timerCallback, nil)
	simRunning = true
	simEngine(got)
	got.StartTimer()
	time.Sleep(300 * time.Millisecond)
	if timerCounter != 2 {
		t.Errorf("[3] NewTimer Error.Counter.running exp:%d got:%d", 2, timerCounter)
		return
	}
	time.Sleep(100 * time.Millisecond)
	simRunning = false

	timerCounter = 0
	got = widgets.NewTimer("timer/3", 100*time.Millisecond, 2)
	if got == nil {
		t.Errorf("[4] NewTimer Error exp:*Timer got:nil")
		return
	}
	got.SetWidgetCallback(timerCallback, nil)
	simRunning = true
	simEngine(got)
	got.StartTimer()
	time.Sleep(150 * time.Millisecond)
	if timerCounter != 1 {
		t.Errorf("[4] NewTimer Error.Counter.running exp:%d got:%d", 1, timerCounter)
		return
	}
	got.CancelTimer()
	time.Sleep(200 * time.Millisecond)
	if timerCounter != 1 {
		t.Errorf("[4] NewTimer Error.Counter.cancelled exp:%d got:%d", 1, timerCounter)
		return
	}
	time.Sleep(100 * time.Millisecond)
	simRunning = false
}
