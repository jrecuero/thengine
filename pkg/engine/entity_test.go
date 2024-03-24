package engine_test

import (
	"testing"

	"github.com/gdamore/tcell/v2"
	"github.com/jrecuero/thengine/pkg/api"
	"github.com/jrecuero/thengine/pkg/engine"
	"github.com/jrecuero/thengine/pkg/tools"
)

var (
	styleOne tcell.Style = tcell.StyleDefault.Foreground(tcell.ColorBlack).Background(tcell.ColorWhite)
)

func TestEntityNewEntity(t *testing.T) {
	cases := []struct {
		input struct {
			name     string
			position *api.Point
			size     *api.Size
			style    *tcell.Style
		}
		exp struct {
			name     string
			position *api.Point
			size     *api.Size
			style    *tcell.Style
		}
	}{
		{
			input: struct {
				name     string
				position *api.Point
				size     *api.Size
				style    *tcell.Style
			}{
				name:     "entity-one",
				position: api.NewPoint(0, 0),
				size:     api.NewSize(1, 1),
				style:    &styleOne,
			},
			exp: struct {
				name     string
				position *api.Point
				size     *api.Size
				style    *tcell.Style
			}{
				name:     "entity-one",
				position: api.NewPoint(0, 0),
				size:     api.NewSize(1, 1),
				style:    &styleOne,
			},
		},
	}
	for i, c := range cases {
		got := engine.NewEntity(c.input.name, c.input.position, c.input.size, c.input.style)
		if got == nil {
			t.Errorf("[%d] NewEntity Error exp:*Entitg got:nil", i)
			continue
		}
		if c.exp.name != got.GetName() {
			t.Errorf("[%d] NewEntity Error.Name exp:%s got:%s", i, c.exp.name, got.GetName())
		}
		if !c.exp.position.IsEqual(got.GetPosition()) {
			t.Errorf("[%d] NewEntity Error.Position exp:%s got:%s", i, c.exp.position.ToString(), got.GetPosition().ToString())
		}
		if !c.exp.size.IsEqual(got.GetSize()) {
			t.Errorf("[%d] NewEntity Error.Size exp:%s got:%s", i, c.exp.size.ToString(), got.GetSize().ToString())
		}
		if !tools.IsEqualStyle(c.exp.style, got.GetStyle()) {
			t.Errorf("[%d] NewEntity Error.Style exp:%+v got:%+v", i, c.exp.style, got.GetStyle())
		}
		gotCanvas := got.GetCanvas()
		if gotCanvas == nil {
			t.Errorf("[%d] NewEntity Error.Canvas exp:*Canvas got:%+v", i, gotCanvas)
		}
		if gotCanvas.Width() != c.exp.size.W {
			t.Errorf("[%d] NewEntity Error.Canvas.W exp:%d got:%d", i, c.exp.size.W, gotCanvas.Width())
		}
		if gotCanvas.Height() != c.exp.size.H {
			t.Errorf("[%d] NewEntity Error.Canvas.H exp:%d got:%d", i, c.exp.size.H, gotCanvas.Height())
		}
	}
}

func TestEntityNewEmptyEntity(t *testing.T) {
	cases := []struct {
		exp struct {
			name     string
			position *api.Point
			size     *api.Size
			style    *tcell.Style
		}
	}{
		{
			exp: struct {
				name     string
				position *api.Point
				size     *api.Size
				style    *tcell.Style
			}{
				name:     "",
				position: nil,
				size:     nil,
				style:    nil,
			},
		},
	}
	for i, c := range cases {
		got := engine.NewEmptyEntity()
		if got == nil {
			t.Errorf("[%d] NewEmptyEntity Error exp:*Entitg got:nil", i)
			continue
		}
		if c.exp.name != got.GetName() {
			t.Errorf("[%d] NewEmptyEntity Error.Name exp:%s got:%s", i, c.exp.name, got.GetName())
		}
		if c.exp.position != got.GetPosition() {
			t.Errorf("[%d] NewEmptyEntity Error.Position exp:%+v got:%+v", i, c.exp.position, got.GetPosition())
		}
		if c.exp.size != got.GetSize() {
			t.Errorf("[%d] NewEmptyEntity Error.Size exp:%+v got:%+v", i, c.exp.size, got.GetSize())
		}
		if c.exp.style != got.GetStyle() {
			t.Errorf("[%d] NewEmptyEntity Error.Style exp:%+v got:%+v", i, c.exp.style, got.GetStyle())
		}
		if got.GetCanvas() != nil {
			t.Errorf("[%d] NewEmptyEntity Error.Canvas exp:nil got:%+v", i, got.GetCanvas())
		}
	}
}

func TestEntityNewNamedEntity(t *testing.T) {
	cases := []struct {
		input string
		exp   struct {
			name     string
			position *api.Point
			size     *api.Size
			style    *tcell.Style
		}
	}{
		{
			input: "entity-one",
			exp: struct {
				name     string
				position *api.Point
				size     *api.Size
				style    *tcell.Style
			}{
				name:     "entity-one",
				position: nil,
				size:     nil,
				style:    nil,
			},
		},
	}
	for i, c := range cases {
		got := engine.NewNamedEntity(c.input)
		if got == nil {
			t.Errorf("[%d] NewNamedEntity Error exp:*Entitg got:nil", i)
			continue
		}
		if c.exp.name != got.GetName() {
			t.Errorf("[%d] NewNamedEntity Error.Name exp:%s got:%s", i, c.exp.name, got.GetName())
		}
		if c.exp.position != got.GetPosition() {
			t.Errorf("[%d] NewNamedEntity Error.Position exp:%+v got:%+v", i, c.exp.position, got.GetPosition())
		}
		if c.exp.size != got.GetSize() {
			t.Errorf("[%d] NewNamedEntity Error.Size exp:%+v got:%+v", i, c.exp.size, got.GetSize())
		}
		if c.exp.style != got.GetStyle() {
			t.Errorf("[%d] NewNamedEntity Error.Style exp:%+v got:%+v", i, c.exp.style, got.GetStyle())
		}
		if got.GetCanvas() != nil {
			t.Errorf("[%d] NewEmptyEntity Error.Canvas exp:nil got:%+v", i, got.GetCanvas())
		}
	}
}

func TestEntityProperties(t *testing.T) {
	cases := []struct {
		input struct {
			name     string
			position *api.Point
			size     *api.Size
			canvas   *engine.Canvas
			style    *tcell.Style
		}
		exp struct {
			name     string
			position *api.Point
			size     *api.Size
			style    *tcell.Style
		}
	}{
		{
			input: struct {
				name     string
				position *api.Point
				size     *api.Size
				canvas   *engine.Canvas
				style    *tcell.Style
			}{
				name:     "entity-one",
				position: api.NewPoint(0, 0),
				size:     api.NewSize(1, 1),
				canvas:   engine.NewCanvas(api.NewSize(1, 1)),
				style:    &styleOne,
			},
			exp: struct {
				name     string
				position *api.Point
				size     *api.Size
				style    *tcell.Style
			}{
				name:     "entity-one",
				position: api.NewPoint(0, 0),
				size:     api.NewSize(1, 1),
				style:    &styleOne,
			},
		},
	}
	for i, c := range cases {
		got := engine.NewEmptyEntity()
		if got == nil {
			t.Errorf("[%d] Properties Error exp:*Entitg got:nil", i)
			continue
		}
		got.SetName(c.input.name)
		if c.exp.name != got.GetName() {
			t.Errorf("[%d] Properties Error.Name exp:%s got:%s", i, c.exp.name, got.GetName())
		}
		got.SetPosition(c.input.position)
		if !c.exp.position.IsEqual(got.GetPosition()) {
			t.Errorf("[%d] Properties Error.Position exp:%s got:%s", i, c.exp.position.ToString(), got.GetPosition().ToString())
		}
		got.SetSize(c.input.size)
		if !c.exp.size.IsEqual(got.GetSize()) {
			t.Errorf("[%d] Properties Error.Size exp:%s got:%s", i, c.exp.size.ToString(), got.GetSize().ToString())
		}
		got.SetCanvas(c.input.canvas)
		gotCanvas := got.GetCanvas()
		if gotCanvas == nil {
			t.Errorf("[%d] Properties Error.Canvas exp:*Canvas got:%+v", i, gotCanvas)
		}
		if gotCanvas.Width() != c.exp.size.W {
			t.Errorf("[%d] Properties Error.Canvas.W exp:%d got:%d", i, c.exp.size.W, gotCanvas.Width())
		}
		if gotCanvas.Height() != c.exp.size.H {
			t.Errorf("[%d] Properties Error.Canvas.H exp:%d got:%d", i, c.exp.size.H, gotCanvas.Height())
		}
		got.SetStyle(c.input.style)
		if !tools.IsEqualStyle(c.exp.style, got.GetStyle()) {
			t.Errorf("[%d] Properties Error.Style exp:%+v got:%+v", i, c.exp.style, got.GetStyle())
		}
	}
}
