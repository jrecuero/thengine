package engine_test

import (
	"encoding/json"
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
			zLevel   int
			pLevel   int
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
				zLevel   int
				pLevel   int
			}{
				name:     "entity-one",
				position: api.NewPoint(0, 0),
				size:     api.NewSize(1, 1),
				style:    &styleOne,
				zLevel:   0,
				pLevel:   0,
			},
		},
	}
	for i, c := range cases {
		got := engine.NewEntity(c.input.name, c.input.position, c.input.size, c.input.style)
		if got == nil {
			t.Errorf("[%d] NewEntity Error exp:*Entity got:nil", i)
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
		gotZLevel := got.GetZLevel()
		if gotZLevel != c.exp.zLevel {
			t.Errorf("[%d] NewEntity Error.ZLevel exp:%d got:%d", i, c.exp.zLevel, gotZLevel)
		}
		gotPLevel := got.GetPLevel()
		if gotPLevel != c.exp.pLevel {
			t.Errorf("[%d] NewEntity Error.PLevel exp:%d got:%d", i, c.exp.pLevel, gotPLevel)
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
			t.Errorf("[%d] NewEmptyEntity Error exp:*Entity got:nil", i)
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
			t.Errorf("[%d] NewNamedEntity Error exp:*Entity got:nil", i)
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
			t.Errorf("[%d] Properties Error exp:*Entity got:nil", i)
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

func TestEntityNewEmptyEntityVisible(t *testing.T) {
	cases := []struct {
		input bool
		exp   bool
	}{
		{
			input: false,
			exp:   false,
		},
		{
			input: true,
			exp:   true,
		},
	}
	for i, c := range cases {
		got := engine.NewEmptyEntity()
		if got == nil {
			t.Errorf("[%d] NewEmptyEntityVisible Error exp:*Entity got:nil", i)
			continue
		}
		got.SetVisible(c.input)
		if c.exp != got.IsVisible() {
			t.Errorf("[%d] NewEmptyEntityVisible Error exp:%t got:%t", i, c.exp, got.IsVisible())
		}
	}
}

func TestEntityMarshalJSON(t *testing.T) {
	name := "test/1"
	position := api.NewPoint(1, 2)
	size := api.NewSize(10, 5)
	input := engine.NewEntity(name, position, size, &styleOne)
	if input == nil {
		t.Errorf("[0] MarshalJSON NewEntity Error exp:*Entity got:nil")
	}
	output, err := json.Marshal(input)
	if err != nil {
		t.Errorf("[0] MarshalJSON Error:%s", err.Error())
	}
	var got map[string]any
	err = json.Unmarshal(output, &got)
	if err != nil {
		t.Errorf("[0] MarshalJSON Error unmarshal map:%s", err.Error())
	}
	if got["name"].(string) != name {
		t.Errorf("[0] MarshalJSON name exp:%s got:%s", name, got["name"])
	}
	gotX := got["position"].([]any)[0].(float64)
	gotY := got["position"].([]any)[1].(float64)
	gotPosition := api.NewPoint(int(gotX), int(gotY))
	if !position.IsEqual(gotPosition) {
		t.Errorf("[0] MarshalJSON position exp:%s got:%s", position.ToString(), gotPosition.ToString())
	}
	gotW := int(got["size"].([]any)[0].(float64))
	gotH := int(got["size"].([]any)[1].(float64))
	gotSize := api.NewSize(gotW, gotH)
	if !size.IsEqual(gotSize) {
		t.Errorf("[0] MarshalJSON size exp:%s got:%s", size.ToString(), gotSize.ToString())
	}
	gotFg := got["style"].([]any)[0].(string)
	gotBg := got["style"].([]any)[1].(string)
	gotAttrs := got["style"].([]any)[2].(string)
	fg, bg, _ := styleOne.Decompose()
	if fg.String() != gotFg {
		t.Errorf("[0] MarshalJSON foreground exp:%s got:%s", fg.String(), gotFg)
	}
	if bg.String() != gotBg {
		t.Errorf("[0] MarshalJSON background exp:%s got:%s", bg.String(), gotBg)
	}
	_ = gotAttrs
	gotEntity := engine.NewEmptyEntity()
	err = json.Unmarshal(output, &gotEntity)
	if err != nil {
		t.Errorf("[0] MarshalJSON Error unmarshal Entity:%s", err.Error())
	}
	if name != gotEntity.GetName() {
		t.Errorf("[0] UnmarshalJSON name error exp:%s got:%s", name, gotEntity.GetName())
	}
	if !position.IsEqual(gotEntity.GetPosition()) {
		t.Errorf("[0] UnmarshalJSON position exp:%s got:%s", position.ToString(), gotEntity.GetPosition().ToString())
	}
	if !size.IsEqual(gotEntity.GetSize()) {
		t.Errorf("[0] UnmarshalJSON size exp:%s got:%s", size.ToString(), gotEntity.GetSize().ToString())
	}
	gotEntityFg, gotEntityBg, _ := gotEntity.GetStyle().Decompose()
	if fg.String() != gotEntityFg.String() {
		t.Errorf("[0] UnmarshalJSON foreground exp:%s got:%s", fg.String(), gotEntityFg.String())
	}
	if bg.String() != gotEntityBg.String() {
		t.Errorf("[0] UnmarshalJSON background exp:%s got:%s", bg.String(), gotEntityBg.String())
	}
}
