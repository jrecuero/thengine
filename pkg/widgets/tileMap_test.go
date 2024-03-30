package widgets_test

import (
	"testing"

	"github.com/gdamore/tcell/v2"
	"github.com/jrecuero/thengine/pkg/api"
	"github.com/jrecuero/thengine/pkg/tools"
	"github.com/jrecuero/thengine/pkg/widgets"
)

var (
	styleOne tcell.Style = tcell.StyleDefault.Foreground(tcell.ColorBlack).Background(tcell.ColorWhite)
)

func TestTileMapNewTileMap(t *testing.T) {
	cases := []struct {
		input struct {
			name         string
			origin       *api.Point
			size         *api.Size
			style        *tcell.Style
			cameraOffset *api.Point
			cameraSize   *api.Size
		}
		exp struct {
			name         string
			origin       *api.Point
			size         *api.Size
			style        *tcell.Style
			cameraOffset *api.Point
			cameraSize   *api.Size
		}
	}{
		{
			input: struct {
				name         string
				origin       *api.Point
				size         *api.Size
				style        *tcell.Style
				cameraOffset *api.Point
				cameraSize   *api.Size
			}{
				name:         "test-one",
				origin:       api.NewPoint(1, 2),
				size:         api.NewSize(10, 5),
				style:        &styleOne,
				cameraOffset: api.NewPoint(5, 6),
				cameraSize:   api.NewSize(3, 4),
			},
			exp: struct {
				name         string
				origin       *api.Point
				size         *api.Size
				style        *tcell.Style
				cameraOffset *api.Point
				cameraSize   *api.Size
			}{
				name:         "test-one",
				origin:       api.NewPoint(1, 2),
				size:         api.NewSize(10, 5),
				style:        &styleOne,
				cameraOffset: api.NewPoint(5, 6),
				cameraSize:   api.NewSize(3, 4),
			},
		},
	}
	for i, c := range cases {
		got := widgets.NewTileMap(c.input.name, c.input.origin, c.input.size, c.input.style, c.input.cameraOffset, c.input.cameraSize)
		if got == nil {
			t.Errorf("[%d] NewTileMap Error exp:*TileMap got:*nil", i)
		}
		gotName := got.GetName()
		if c.exp.name != gotName {
			t.Errorf("[%d] NewTileMap Error.GetName exp:%s got:%s", i, c.exp.name, gotName)
		}
		gotPosition := got.GetPosition()
		if !c.exp.origin.IsEqual(gotPosition) {
			t.Errorf("[%d] NewTileMap Error.GetPosition exp:%s got:%s", i, c.exp.origin.ToString(), gotPosition.ToString())
		}
		gotSize := got.GetSize()
		if !c.exp.size.IsEqual(gotSize) {
			t.Errorf("[%d] NewTileMap Error.GetSize exp:%s, got:%s", i, c.exp.size.ToString(), gotSize.ToString())
		}
		gotStyle := got.GetStyle()
		if !tools.IsEqualStyle(c.exp.style, gotStyle) {
			t.Errorf("[%d] NewTileMap Error.GetStyle exp:%+v got:%+v", i, c.exp.style, gotStyle)
		}
		gotCameraOffset := got.GetCameraOffset()
		if !c.exp.cameraOffset.IsEqual(gotCameraOffset) {
			t.Errorf("[%d] NewTileMap Error.GetCameraOffset exp:%s got:%s", i, c.exp.cameraOffset.ToString(), gotCameraOffset.ToString())
		}
		gotCameraSize := got.GetCameraSize()
		if !c.exp.cameraSize.IsEqual(gotCameraSize) {
			t.Errorf("[%d] NewTileMap Error.GetCameraSize exp:%s got:%s", i, c.exp.cameraSize.ToString(), gotCameraSize.ToString())
		}
	}
}

func TestTileMapSetCameraOffset(t *testing.T) {
	cases := []struct {
		input struct {
			origin       *api.Point
			size         *api.Size
			cameraOffset *api.Point
			cameraSize   *api.Size
			newOffset    *api.Point
		}
		exp struct {
			cameraOffset *api.Point
			ok           bool
		}
	}{
		{
			input: struct {
				origin       *api.Point
				size         *api.Size
				cameraOffset *api.Point
				cameraSize   *api.Size
				newOffset    *api.Point
			}{
				origin:       api.NewPoint(1, 2),
				size:         api.NewSize(20, 10),
				cameraOffset: api.NewPoint(0, 0),
				cameraSize:   api.NewSize(3, 4),
				newOffset:    api.NewPoint(2, 2),
			},
			exp: struct {
				cameraOffset *api.Point
				ok           bool
			}{
				cameraOffset: api.NewPoint(2, 2),
				ok:           true,
			},
		},
		{
			input: struct {
				origin       *api.Point
				size         *api.Size
				cameraOffset *api.Point
				cameraSize   *api.Size
				newOffset    *api.Point
			}{
				origin:       api.NewPoint(0, 0),
				size:         api.NewSize(20, 10),
				cameraOffset: api.NewPoint(0, 0),
				cameraSize:   api.NewSize(5, 5),
				newOffset:    api.NewPoint(16, 2),
			},
			exp: struct {
				cameraOffset *api.Point
				ok           bool
			}{
				cameraOffset: api.NewPoint(0, 0),
				ok:           false,
			},
		},
		{
			input: struct {
				origin       *api.Point
				size         *api.Size
				cameraOffset *api.Point
				cameraSize   *api.Size
				newOffset    *api.Point
			}{
				origin:       api.NewPoint(0, 0),
				size:         api.NewSize(20, 10),
				cameraOffset: api.NewPoint(0, 0),
				cameraSize:   api.NewSize(5, 5),
				newOffset:    api.NewPoint(0, 8),
			},
			exp: struct {
				cameraOffset *api.Point
				ok           bool
			}{
				cameraOffset: api.NewPoint(0, 0),
				ok:           false,
			},
		},
	}
	for i, c := range cases {
		tm := widgets.NewTileMap("test", c.input.origin, c.input.size, &styleOne, c.input.cameraOffset, c.input.cameraSize)
		got := tm.SetCameraOffset(c.input.newOffset)
		if got != c.exp.ok {
			t.Errorf("[%d] SetCameraOffset Error exp:%t got:%t", i, c.exp.ok, got)
		}
		gotCameraOffset := tm.GetCameraOffset()
		if !c.exp.cameraOffset.IsEqual(gotCameraOffset) {
			t.Errorf("[%d] SetCamera Error.GetCameraOffset exp:%s got:%s", i, c.exp.cameraOffset.ToString(), gotCameraOffset.ToString())
		}
	}
}

func TestTileMapGetTileMapPosFromScreenPos(t *testing.T) {
	cases := []struct {
		input struct {
			origin       *api.Point
			size         *api.Size
			cameraOffset *api.Point
			cameraSize   *api.Size
			screenPos    *api.Point
		}
		exp *api.Point
	}{
		{
			input: struct {
				origin       *api.Point
				size         *api.Size
				cameraOffset *api.Point
				cameraSize   *api.Size
				screenPos    *api.Point
			}{
				origin:       api.NewPoint(5, 5),
				size:         api.NewSize(20, 10),
				cameraOffset: api.NewPoint(2, 2),
				cameraSize:   api.NewSize(5, 5),
				screenPos:    api.NewPoint(16, 10),
			},
			exp: api.NewPoint(13, 7),
		},
		{
			input: struct {
				origin       *api.Point
				size         *api.Size
				cameraOffset *api.Point
				cameraSize   *api.Size
				screenPos    *api.Point
			}{
				origin:       api.NewPoint(5, 5),
				size:         api.NewSize(20, 10),
				cameraOffset: api.NewPoint(2, 2),
				cameraSize:   api.NewSize(5, 5),
				screenPos:    api.NewPoint(4, 10),
			},
			exp: nil,
		},
		{
			input: struct {
				origin       *api.Point
				size         *api.Size
				cameraOffset *api.Point
				cameraSize   *api.Size
				screenPos    *api.Point
			}{
				origin:       api.NewPoint(5, 5),
				size:         api.NewSize(20, 10),
				cameraOffset: api.NewPoint(2, 2),
				cameraSize:   api.NewSize(5, 5),
				screenPos:    api.NewPoint(10, 1),
			},
			exp: nil,
		},
		{
			input: struct {
				origin       *api.Point
				size         *api.Size
				cameraOffset *api.Point
				cameraSize   *api.Size
				screenPos    *api.Point
			}{
				origin:       api.NewPoint(5, 5),
				size:         api.NewSize(20, 10),
				cameraOffset: api.NewPoint(2, 2),
				cameraSize:   api.NewSize(5, 5),
				screenPos:    api.NewPoint(25, 10),
			},
			exp: nil,
		},
		{
			input: struct {
				origin       *api.Point
				size         *api.Size
				cameraOffset *api.Point
				cameraSize   *api.Size
				screenPos    *api.Point
			}{
				origin:       api.NewPoint(5, 5),
				size:         api.NewSize(20, 10),
				cameraOffset: api.NewPoint(2, 2),
				cameraSize:   api.NewSize(5, 5),
				screenPos:    api.NewPoint(15, 15),
			},
			exp: nil,
		},
	}
	for i, c := range cases {
		tm := widgets.NewTileMap("test", c.input.origin, c.input.size, &styleOne, c.input.cameraOffset, c.input.cameraSize)
		got := tm.GetTileMapPosFromScreenPos(c.input.screenPos)
		if (c.exp == nil) && (got != nil) {
			t.Errorf("[%d] GetTileMapPosFromScreenPos Error exp:%v got:%s", i, c.exp, got.ToString())
		}
		if c.exp != nil && !c.exp.IsEqual(got) {
			t.Errorf("[%d] GetTileMapPosFromScreenPos Error exp:%s got:%s", i, c.exp.ToString(), got.ToString())
		}
	}
}

func TestTileMapGetScreenPosFromTileMapPos(t *testing.T) {
	cases := []struct {
		input struct {
			origin       *api.Point
			size         *api.Size
			cameraOffset *api.Point
			cameraSize   *api.Size
			tileMapPos   *api.Point
		}
		exp *api.Point
	}{
		{
			input: struct {
				origin       *api.Point
				size         *api.Size
				cameraOffset *api.Point
				cameraSize   *api.Size
				tileMapPos   *api.Point
			}{
				origin:       api.NewPoint(5, 5),
				size:         api.NewSize(20, 10),
				cameraOffset: api.NewPoint(2, 2),
				cameraSize:   api.NewSize(5, 5),
				tileMapPos:   api.NewPoint(10, 5),
			},
			exp: api.NewPoint(13, 8),
		},
		{
			input: struct {
				origin       *api.Point
				size         *api.Size
				cameraOffset *api.Point
				cameraSize   *api.Size
				tileMapPos   *api.Point
			}{
				origin:       api.NewPoint(5, 5),
				size:         api.NewSize(20, 10),
				cameraOffset: api.NewPoint(2, 2),
				cameraSize:   api.NewSize(5, 5),
				tileMapPos:   api.NewPoint(20, 10),
			},
			exp: nil,
		},
		{
			input: struct {
				origin       *api.Point
				size         *api.Size
				cameraOffset *api.Point
				cameraSize   *api.Size
				tileMapPos   *api.Point
			}{
				origin:       api.NewPoint(5, 5),
				size:         api.NewSize(20, 10),
				cameraOffset: api.NewPoint(2, 2),
				cameraSize:   api.NewSize(5, 5),
				tileMapPos:   api.NewPoint(10, 10),
			},
			exp: nil,
		},
	}
	for i, c := range cases {
		tm := widgets.NewTileMap("test", c.input.origin, c.input.size, &styleOne, c.input.cameraOffset, c.input.cameraSize)
		got := tm.GetScreenPosFromTileMapPos(c.input.tileMapPos)
		if (c.exp == nil) && (got != nil) {
			t.Errorf("[%d] GetScreenPosFromTilePos Error exp:%v got:%s", i, c.exp, got.ToString())
		}
		if c.exp != nil && !c.exp.IsEqual(got) {
			t.Errorf("[%d] GetScreenPosFromTilePos Error exp:%s got:%s", i, c.exp.ToString(), got.ToString())
		}
	}
}

func TestTileMapDistanceToTileMapEdgesX(t *testing.T) {
	cases := []struct {
		input struct {
			origin       *api.Point
			size         *api.Size
			cameraOffset *api.Point
			cameraSize   *api.Size
			tileMapPos   *api.Point
		}
		exp struct {
			ok  bool
			one int
			two int
		}
	}{
		{
			input: struct {
				origin       *api.Point
				size         *api.Size
				cameraOffset *api.Point
				cameraSize   *api.Size
				tileMapPos   *api.Point
			}{
				origin:       api.NewPoint(5, 5),
				size:         api.NewSize(20, 10),
				cameraOffset: api.NewPoint(2, 2),
				cameraSize:   api.NewSize(5, 5),
				tileMapPos:   api.NewPoint(10, 5),
			},
			exp: struct {
				ok  bool
				one int
				two int
			}{
				ok:  true,
				one: 10,
				two: 10,
			},
		},
		{
			input: struct {
				origin       *api.Point
				size         *api.Size
				cameraOffset *api.Point
				cameraSize   *api.Size
				tileMapPos   *api.Point
			}{
				origin:       api.NewPoint(5, 5),
				size:         api.NewSize(20, 10),
				cameraOffset: api.NewPoint(2, 2),
				cameraSize:   api.NewSize(5, 5),
				tileMapPos:   api.NewPoint(20, 10),
			},
			exp: struct {
				ok  bool
				one int
				two int
			}{
				ok:  false,
				one: 0,
				two: 0,
			},
		},
	}
	for i, c := range cases {
		tm := widgets.NewTileMap("test", c.input.origin, c.input.size, &styleOne, c.input.cameraOffset, c.input.cameraSize)
		gotOk, gotOne, gotTwo := tm.DistanceToTileMapEdgesX(c.input.tileMapPos)
		if c.exp.ok != gotOk {
			t.Errorf("[%d] DistanceToTileMapEdgesX Error exp:%t got:%t", i, c.exp.ok, gotOk)
		}
		if c.exp.one != gotOne {
			t.Errorf("[%d] DistanceToTileMapEdgesX Error.one exp:%d got:%d", i, c.exp.one, gotOne)
		}
		if c.exp.two != gotTwo {
			t.Errorf("[%d] DistanceToTileMapEdgesX Error.two exp:%d got:%d", i, c.exp.two, gotTwo)
		}
	}
}

func TestTileMapDistancetoTileMapEdgesY(t *testing.T) {
	cases := []struct {
		input struct {
			origin       *api.Point
			size         *api.Size
			cameraOffset *api.Point
			cameraSize   *api.Size
			tileMapPos   *api.Point
		}
		exp struct {
			ok  bool
			one int
			two int
		}
	}{
		{
			input: struct {
				origin       *api.Point
				size         *api.Size
				cameraOffset *api.Point
				cameraSize   *api.Size
				tileMapPos   *api.Point
			}{
				origin:       api.NewPoint(5, 5),
				size:         api.NewSize(20, 10),
				cameraOffset: api.NewPoint(2, 2),
				cameraSize:   api.NewSize(5, 5),
				tileMapPos:   api.NewPoint(10, 5),
			},
			exp: struct {
				ok  bool
				one int
				two int
			}{
				ok:  true,
				one: 5,
				two: 5,
			},
		},
		{
			input: struct {
				origin       *api.Point
				size         *api.Size
				cameraOffset *api.Point
				cameraSize   *api.Size
				tileMapPos   *api.Point
			}{
				origin:       api.NewPoint(5, 5),
				size:         api.NewSize(20, 10),
				cameraOffset: api.NewPoint(2, 2),
				cameraSize:   api.NewSize(5, 5),
				tileMapPos:   api.NewPoint(20, 10),
			},
			exp: struct {
				ok  bool
				one int
				two int
			}{
				ok:  false,
				one: 0,
				two: 0,
			},
		},
	}
	for i, c := range cases {
		tm := widgets.NewTileMap("test", c.input.origin, c.input.size, &styleOne, c.input.cameraOffset, c.input.cameraSize)
		gotOk, gotOne, gotTwo := tm.DistanceToTileMapEdgesY(c.input.tileMapPos)
		if c.exp.ok != gotOk {
			t.Errorf("[%d] DistanceToTileMapEdgesY Error exp:%t got:%t", i, c.exp.ok, gotOk)
		}
		if c.exp.one != gotOne {
			t.Errorf("[%d] DistanceToTileMapEdgesY Error.one exp:%d got:%d", i, c.exp.one, gotOne)
		}
		if c.exp.two != gotTwo {
			t.Errorf("[%d] DistanceToTileMapEdgesY Error.two exp:%d got:%d", i, c.exp.two, gotTwo)
		}
	}
}

func TestTileMapDistanceToCameraEdgesX(t *testing.T) {
	cases := []struct {
		input struct {
			origin       *api.Point
			size         *api.Size
			cameraOffset *api.Point
			cameraSize   *api.Size
			tileMapPos   *api.Point
		}
		exp struct {
			ok  bool
			one int
			two int
		}
	}{
		{
			input: struct {
				origin       *api.Point
				size         *api.Size
				cameraOffset *api.Point
				cameraSize   *api.Size
				tileMapPos   *api.Point
			}{
				origin:       api.NewPoint(5, 5),
				size:         api.NewSize(20, 10),
				cameraOffset: api.NewPoint(2, 2),
				cameraSize:   api.NewSize(5, 5),
				tileMapPos:   api.NewPoint(5, 4),
			},
			exp: struct {
				ok  bool
				one int
				two int
			}{
				ok:  true,
				one: 3,
				two: 2,
			},
		},
		{
			input: struct {
				origin       *api.Point
				size         *api.Size
				cameraOffset *api.Point
				cameraSize   *api.Size
				tileMapPos   *api.Point
			}{
				origin:       api.NewPoint(5, 5),
				size:         api.NewSize(20, 10),
				cameraOffset: api.NewPoint(2, 2),
				cameraSize:   api.NewSize(5, 5),
				tileMapPos:   api.NewPoint(20, 10),
			},
			exp: struct {
				ok  bool
				one int
				two int
			}{
				ok:  false,
				one: 0,
				two: 0,
			},
		},
	}
	for i, c := range cases {
		tm := widgets.NewTileMap("test", c.input.origin, c.input.size, &styleOne, c.input.cameraOffset, c.input.cameraSize)
		gotOk, gotOne, gotTwo := tm.DistanceToCameraEdgesX(c.input.tileMapPos)
		if c.exp.ok != gotOk {
			t.Errorf("[%d] DistanceToCameraEdgesX Error exp:%t got:%t", i, c.exp.ok, gotOk)
		}
		if c.exp.one != gotOne {
			t.Errorf("[%d] DistanceToCameraEdgesX Error.one exp:%d got:%d", i, c.exp.one, gotOne)
		}
		if c.exp.two != gotTwo {
			t.Errorf("[%d] DistanceToCameraEdgesX Error.two exp:%d got:%d", i, c.exp.two, gotTwo)
		}
	}
}

func TestTileMapDistanceToCameraEdgesY(t *testing.T) {
	cases := []struct {
		input struct {
			origin       *api.Point
			size         *api.Size
			cameraOffset *api.Point
			cameraSize   *api.Size
			tileMapPos   *api.Point
		}
		exp struct {
			ok  bool
			one int
			two int
		}
	}{
		{
			input: struct {
				origin       *api.Point
				size         *api.Size
				cameraOffset *api.Point
				cameraSize   *api.Size
				tileMapPos   *api.Point
			}{
				origin:       api.NewPoint(5, 5),
				size:         api.NewSize(20, 10),
				cameraOffset: api.NewPoint(2, 2),
				cameraSize:   api.NewSize(5, 5),
				tileMapPos:   api.NewPoint(5, 4),
			},
			exp: struct {
				ok  bool
				one int
				two int
			}{
				ok:  true,
				one: 2,
				two: 3,
			},
		},
		{
			input: struct {
				origin       *api.Point
				size         *api.Size
				cameraOffset *api.Point
				cameraSize   *api.Size
				tileMapPos   *api.Point
			}{
				origin:       api.NewPoint(5, 5),
				size:         api.NewSize(20, 10),
				cameraOffset: api.NewPoint(2, 2),
				cameraSize:   api.NewSize(5, 5),
				tileMapPos:   api.NewPoint(20, 10),
			},
			exp: struct {
				ok  bool
				one int
				two int
			}{
				ok:  false,
				one: 0,
				two: 0,
			},
		},
	}
	for i, c := range cases {
		tm := widgets.NewTileMap("test", c.input.origin, c.input.size, &styleOne, c.input.cameraOffset, c.input.cameraSize)
		gotOk, gotOne, gotTwo := tm.DistanceToCameraEdgesY(c.input.tileMapPos)
		if c.exp.ok != gotOk {
			t.Errorf("[%d] DistanceToCameraEdgesY Error exp:%t got:%t", i, c.exp.ok, gotOk)
		}
		if c.exp.one != gotOne {
			t.Errorf("[%d] DistanceToCameraEdgesY Error.one exp:%d got:%d", i, c.exp.one, gotOne)
		}
		if c.exp.two != gotTwo {
			t.Errorf("[%d] DistanceToCameraEdgesY Error.two exp:%d got:%d", i, c.exp.two, gotTwo)
		}
	}
}
