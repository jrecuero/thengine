package api_test

import (
	"testing"

	"github.com/jrecuero/thengine/pkg/api"
)

func TestRect(t *testing.T) {
	cases := []struct {
		input []int
		exp   []int
	}{
		{
			input: []int{0, 0, 10, 10},
			exp:   []int{0, 0, 10, 10},
		},
		{
			input: []int{5, 5, 7, 7},
			exp:   []int{5, 5, 7, 7},
		},
	}
	for i, c := range cases {
		origin := api.NewPoint(c.input[0], c.input[1])
		size := api.NewSize(c.input[2], c.input[3])
		got := api.NewRect(origin, size)
		if c.exp[0] != got.Origin.X {
			t.Errorf("[%d] NewRect Origin.X Error exp:%d got:%d", i, c.exp[0], got.Origin.X)
		}
		if c.exp[1] != got.Origin.Y {
			t.Errorf("[%d] NewRect Origin.Y Error exp:%d got:%d", i, c.exp[1], got.Origin.Y)
		}
		if c.exp[2] != got.Size.W {
			t.Errorf("[%d] NewRect Size.W Error exp:%d got:%d", i, c.exp[2], got.Size.W)
		}
		if c.exp[3] != got.Size.H {
			t.Errorf("[%d] NewRect Size.H Error exp:%d got:%d", i, c.exp[3], got.Size.H)
		}
	}
}

func TestRectCloneRect(t *testing.T) {
	cases := []struct {
		input []int
		exp   []int
	}{
		{
			input: []int{0, 0, 10, 10},
			exp:   []int{0, 0, 10, 10},
		},
		{
			input: []int{5, 5, 7, 7},
			exp:   []int{5, 5, 7, 7},
		},
	}
	for i, c := range cases {
		origin := api.NewPoint(c.input[0], c.input[1])
		size := api.NewSize(c.input[2], c.input[3])
		toClone := api.NewRect(origin, size)
		got := api.CloneRect(toClone)
		if c.exp[0] != got.Origin.X {
			t.Errorf("[%d] CloneRect Origin.X Error exp:%d got:%d", i, c.exp[0], got.Origin.X)
		}
		if c.exp[1] != got.Origin.Y {
			t.Errorf("[%d] CloneRect Origin.Y Error exp:%d got:%d", i, c.exp[1], got.Origin.Y)
		}
		if c.exp[2] != got.Size.W {
			t.Errorf("[%d] CloneRect Size.W Error exp:%d got:%d", i, c.exp[2], got.Size.W)
		}
		if c.exp[3] != got.Size.H {
			t.Errorf("[%d] CloneRect Size.H Error exp:%d got:%d", i, c.exp[3], got.Size.H)
		}
	}
}

func TestRectClone(t *testing.T) {
	cases := []struct {
		input []int
		exp   []int
	}{
		{
			input: []int{0, 0, 10, 10},
			exp:   []int{0, 0, 10, 10},
		},
		{
			input: []int{5, 5, 7, 7},
			exp:   []int{5, 5, 7, 7},
		},
	}
	for i, c := range cases {
		origin := api.NewPoint(c.input[0], c.input[1])
		size := api.NewSize(c.input[2], c.input[3])
		toClone := api.NewRect(origin, size)
		got := api.NewRect(api.NewPoint(0, 0), api.NewSize(0, 0))
		got.Clone(toClone)
		if c.exp[0] != got.Origin.X {
			t.Errorf("[%d] Clone Origin.X Error exp:%d got:%d", i, c.exp[0], got.Origin.X)
		}
		if c.exp[1] != got.Origin.Y {
			t.Errorf("[%d] Clone Origin.Y Error exp:%d got:%d", i, c.exp[1], got.Origin.Y)
		}
		if c.exp[2] != got.Size.W {
			t.Errorf("[%d] Clone Size.W Error exp:%d got:%d", i, c.exp[2], got.Size.W)
		}
		if c.exp[3] != got.Size.H {
			t.Errorf("[%d] Clone Size.H Error exp:%d got:%d", i, c.exp[3], got.Size.H)
		}
	}
}

func TestRectSet(t *testing.T) {
	cases := []struct {
		input []int
		exp   []int
	}{
		{
			input: []int{0, 0, 10, 10},
			exp:   []int{0, 0, 10, 10},
		},
		{
			input: []int{5, 5, 7, 7},
			exp:   []int{5, 5, 7, 7},
		},
	}
	for i, c := range cases {
		origin := api.NewPoint(c.input[0], c.input[1])
		size := api.NewSize(c.input[2], c.input[3])
		got := api.NewRect(api.NewPoint(0, 0), api.NewSize(0, 0))
		got.Set(origin, size)
		if c.exp[0] != got.Origin.X {
			t.Errorf("[%d] Set Origin.X Error exp:%d got:%d", i, c.exp[0], got.Origin.X)
		}
		if c.exp[1] != got.Origin.Y {
			t.Errorf("[%d] Set Origin.Y Error exp:%d got:%d", i, c.exp[1], got.Origin.Y)
		}
		if c.exp[2] != got.Size.W {
			t.Errorf("[%d] Set Size.W Error exp:%d got:%d", i, c.exp[2], got.Size.W)
		}
		if c.exp[3] != got.Size.H {
			t.Errorf("[%d] Set Size.H Error exp:%d got:%d", i, c.exp[3], got.Size.H)
		}
	}
}

func TestRectSetOrigin(t *testing.T) {
	cases := []struct {
		input []int
		exp   []int
	}{
		{
			input: []int{10, 10},
			exp:   []int{10, 10, 0, 0},
		},
		{
			input: []int{5, 5},
			exp:   []int{5, 5, 0, 0},
		},
	}
	for i, c := range cases {
		origin := api.NewPoint(c.input[0], c.input[1])
		got := api.NewRect(api.NewPoint(0, 0), api.NewSize(0, 0))
		got.SetOrigin(origin)
		if c.exp[0] != got.Origin.X {
			t.Errorf("[%d] SetOrigin Origin.X Error exp:%d got:%d", i, c.exp[0], got.Origin.X)
		}
		if c.exp[1] != got.Origin.Y {
			t.Errorf("[%d] SetOrigin Origin.Y Error exp:%d got:%d", i, c.exp[1], got.Origin.Y)
		}
		if c.exp[2] != got.Size.W {
			t.Errorf("[%d] SetOrigin Size.W Error exp:%d got:%d", i, c.exp[2], got.Size.W)
		}
		if c.exp[3] != got.Size.H {
			t.Errorf("[%d] SetOrigin Size.H Error exp:%d got:%d", i, c.exp[3], got.Size.H)
		}
	}
}

func TestRectSetSize(t *testing.T) {
	cases := []struct {
		input []int
		exp   []int
	}{
		{
			input: []int{10, 10},
			exp:   []int{0, 0, 10, 10},
		},
		{
			input: []int{5, 5},
			exp:   []int{0, 0, 5, 5},
		},
	}
	for i, c := range cases {
		size := api.NewSize(c.input[0], c.input[1])
		got := api.NewRect(api.NewPoint(0, 0), api.NewSize(0, 0))
		got.SetSize(size)

		if c.exp[0] != got.Origin.X {
			t.Errorf("[%d] SetSize Origin.X Error exp:%d got:%d", i, c.exp[0], got.Origin.X)
		}
		if c.exp[1] != got.Origin.Y {
			t.Errorf("[%d] SetSize Origin.Y Error exp:%d got:%d", i, c.exp[1], got.Origin.Y)
		}
		if c.exp[2] != got.Size.W {
			t.Errorf("[%d] SetSize Size.W Error exp:%d got:%d", i, c.exp[2], got.Size.W)
		}
		if c.exp[3] != got.Size.H {
			t.Errorf("[%d] SetSize Size.H Error exp:%d got:%d", i, c.exp[3], got.Size.H)
		}
	}
}

func TestRectGet(t *testing.T) {
	cases := []struct {
		input []int
		exp   []int
	}{
		{
			input: []int{0, 0, 10, 10},
			exp:   []int{0, 0, 10, 10},
		},
		{
			input: []int{5, 5, 7, 7},
			exp:   []int{5, 5, 7, 7},
		},
	}
	for i, c := range cases {
		origin := api.NewPoint(c.input[0], c.input[1])
		size := api.NewSize(c.input[2], c.input[3])
		got := api.NewRect(origin, size)
		gotOrigin, gotSize := got.Get()
		if c.exp[0] != gotOrigin.X {
			t.Errorf("[%d] Get Origin.X Error exp:%d got:%d", i, c.exp[0], gotOrigin.X)
		}
		if c.exp[1] != gotOrigin.Y {
			t.Errorf("[%d] Get Origin.Y Error exp:%d got:%d", i, c.exp[1], gotOrigin.Y)
		}
		if c.exp[2] != gotSize.W {
			t.Errorf("[%d] Get Size.W Error exp:%d got:%d", i, c.exp[2], gotSize.W)
		}
		if c.exp[3] != gotSize.H {
			t.Errorf("[%d] Get Size.H Error exp:%d got:%d", i, c.exp[3], gotSize.H)
		}
	}
}

func TestRectGetOrigin(t *testing.T) {
	cases := []struct {
		input []int
		exp   []int
	}{
		{
			input: []int{0, 0, 10, 10},
			exp:   []int{0, 0},
		},
		{
			input: []int{5, 5, 7, 7},
			exp:   []int{5, 5},
		},
	}
	for i, c := range cases {
		origin := api.NewPoint(c.input[0], c.input[1])
		size := api.NewSize(c.input[2], c.input[3])
		got := api.NewRect(origin, size)
		gotOrigin := got.GetOrigin()
		if c.exp[0] != gotOrigin.X {
			t.Errorf("[%d] GetOrigin Origin.X Error exp:%d got:%d", i, c.exp[0], gotOrigin.X)
		}
		if c.exp[1] != gotOrigin.Y {
			t.Errorf("[%d] GetOrigin Origin.Y Error exp:%d got:%d", i, c.exp[1], gotOrigin.Y)
		}
	}
}

func TestRectGetSize(t *testing.T) {
	cases := []struct {
		input []int
		exp   []int
	}{
		{
			input: []int{0, 0, 10, 10},
			exp:   []int{0, 0, 10, 10},
		},
		{
			input: []int{5, 5, 7, 7},
			exp:   []int{5, 5, 7, 7},
		},
	}
	for i, c := range cases {
		origin := api.NewPoint(c.input[0], c.input[1])
		size := api.NewSize(c.input[2], c.input[3])
		got := api.NewRect(origin, size)
		gotSize := got.GetSize()
		if c.exp[2] != gotSize.W {
			t.Errorf("[%d] GetSize Size.W Error exp:%d got:%d", i, c.exp[2], gotSize.W)
		}
		if c.exp[3] != gotSize.H {
			t.Errorf("[%d] GetSize Size.H Error exp:%d got:%d", i, c.exp[3], gotSize.H)
		}
	}
}
func TestRectIsEqual(t *testing.T) {
	cases := []struct {
		input []int
		exp   bool
	}{
		{
			input: []int{0, 0, 10, 10, 0, 0, 10, 10},
			exp:   true,
		},
		{
			input: []int{5, 5, 7, 7, 5, 5, 7, 7},
			exp:   true,
		},
		{
			input: []int{0, 0, 10, 10, 1, 0, 10, 10},
			exp:   false,
		},
		{
			input: []int{0, 0, 10, 10, 0, 1, 10, 10},
			exp:   false,
		},
		{
			input: []int{5, 5, 7, 7, 5, 5, 10, 7},
			exp:   false,
		},
		{
			input: []int{5, 5, 7, 7, 5, 5, 7, 10},
			exp:   false,
		},
	}
	for i, c := range cases {
		origin1 := api.NewPoint(c.input[0], c.input[1])
		size1 := api.NewSize(c.input[2], c.input[3])
		rect1 := api.NewRect(origin1, size1)
		origin2 := api.NewPoint(c.input[4], c.input[5])
		size2 := api.NewSize(c.input[6], c.input[7])
		rect2 := api.NewRect(origin2, size2)
		got := rect1.IsEqual(rect2)
		if c.exp != got {
			t.Errorf("[%d] Equals Error exp:%t got:%t", i, c.exp, got)
		}
		got = rect2.IsEqual(rect1)
		if c.exp != got {
			t.Errorf("[%d] Equals (reverse) Error exp:%t got:%t", i, c.exp, got)
		}
	}
}

func TestRectIn(t *testing.T) {
	cases := []struct {
		input []int
		exp   bool
	}{
		{
			input: []int{0, 0, 10, 10, 5, 5},
			exp:   true,
		},
		{
			input: []int{0, 0, 10, 10, 0, 0},
			exp:   true,
		},
		{
			input: []int{0, 0, 10, 10, 0, 9},
			exp:   true,
		},
		{
			input: []int{0, 0, 10, 10, 9, 0},
			exp:   true,
		},
		{
			input: []int{0, 0, 10, 10, 9, 9},
			exp:   true,
		},
		{
			input: []int{0, 0, 10, 10, 15, 15},
			exp:   false,
		},
	}
	for i, c := range cases {
		origin := api.NewPoint(c.input[0], c.input[1])
		size := api.NewSize(c.input[2], c.input[3])
		rect := api.NewRect(origin, size)
		point := api.NewPoint(c.input[4], c.input[5])
		got := rect.IsIn(point)
		if c.exp != got {
			t.Errorf("[%d] IsIn Error exp:%t got: %t", i, c.exp, got)
		}
	}
}

func TestRectInside(t *testing.T) {
	cases := []struct {
		input []int
		exp   bool
	}{
		{
			input: []int{0, 0, 10, 10, 5, 5},
			exp:   true,
		},
		{
			input: []int{0, 0, 10, 10, 1, 1},
			exp:   true,
		},
		{
			input: []int{0, 0, 10, 10, 8, 8},
			exp:   true,
		},
		{
			input: []int{0, 0, 10, 10, 0, 0},
			exp:   false,
		},
		{
			input: []int{0, 0, 10, 10, 0, 9},
			exp:   false,
		},
		{
			input: []int{0, 0, 10, 10, 9, 0},
			exp:   false,
		},
		{
			input: []int{0, 0, 10, 10, 9, 9},
			exp:   false,
		},
		{
			input: []int{0, 0, 10, 10, 15, 15},
			exp:   false,
		},
	}
	for i, c := range cases {
		origin := api.NewPoint(c.input[0], c.input[1])
		size := api.NewSize(c.input[2], c.input[3])
		rect := api.NewRect(origin, size)
		point := api.NewPoint(c.input[4], c.input[5])
		got := rect.IsInside(point)
		if c.exp != got {
			t.Errorf("[%d] IsInside Error %+v exp:%t got: %t", i, c.input, c.exp, got)
		}
	}
}

func TestRectIsRectIntersect(t *testing.T) {
	cases := []struct {
		origin1 *api.Point
		size1   *api.Size
		origin2 *api.Point
		size2   *api.Size
		exp     bool
	}{
		{
			origin1: api.NewPoint(0, 0),
			size1:   api.NewSize(10, 10),
			origin2: api.NewPoint(5, 5),
			size2:   api.NewSize(10, 10),
			exp:     true,
		},
		{
			origin1: api.NewPoint(3, 3),
			size1:   api.NewSize(10, 10),
			origin2: api.NewPoint(0, 0),
			size2:   api.NewSize(10, 10),
			exp:     true,
		},
		{
			origin1: api.NewPoint(0, 0),
			size1:   api.NewSize(5, 5),
			origin2: api.NewPoint(6, 0),
			size2:   api.NewSize(10, 10),
			exp:     false,
		},
		{
			origin1: api.NewPoint(0, 0),
			size1:   api.NewSize(5, 5),
			origin2: api.NewPoint(0, 6),
			size2:   api.NewSize(10, 10),
			exp:     false,
		},
		{
			origin1: api.NewPoint(0, 0),
			size1:   api.NewSize(5, 5),
			origin2: api.NewPoint(6, 6),
			size2:   api.NewSize(10, 10),
			exp:     false,
		},
	}
	for i, c := range cases {
		rect1 := api.NewRect(c.origin1, c.size1)
		rect2 := api.NewRect(c.origin2, c.size2)
		got := rect1.IsRectIntersect(rect2)
		if c.exp != got {
			t.Errorf("[%d] IsRectIntersect Error exp:%t got:%t", i, c.exp, got)
		}
	}
}

func TestRectIsBorder(t *testing.T) {
	cases := []struct {
		input []int
		exp   bool
	}{
		{
			input: []int{0, 0, 10, 10, 0, 0},
			exp:   true,
		},
		{
			input: []int{0, 0, 10, 10, 9, 9},
			exp:   true,
		},
		{
			input: []int{0, 0, 10, 10, 0, 5},
			exp:   true,
		},
		{
			input: []int{0, 0, 10, 10, 9, 5},
			exp:   true,
		},
		{
			input: []int{0, 0, 10, 10, 5, 0},
			exp:   true,
		},
		{
			input: []int{0, 0, 10, 10, 5, 9},
			exp:   true,
		},
		{
			input: []int{0, 0, 10, 10, 1, 1},
			exp:   false,
		},
		{
			input: []int{0, 0, 10, 10, 10, 10},
			exp:   false,
		},
		{
			input: []int{0, 0, 10, 10, 1, 5},
			exp:   false,
		},
		{
			input: []int{0, 0, 10, 10, 10, 5},
			exp:   false,
		},
		{
			input: []int{0, 0, 10, 10, 5, 1},
			exp:   false,
		},
		{
			input: []int{0, 0, 10, 10, 5, 10},
			exp:   false,
		},
	}
	for i, c := range cases {
		origin := api.NewPoint(c.input[0], c.input[1])
		size := api.NewSize(c.input[2], c.input[3])
		rect := api.NewRect(origin, size)
		point := api.NewPoint(c.input[4], c.input[5])
		got := rect.IsBorder(point)
		if c.exp != got {
			t.Errorf("[%d] IsBorder Error exp:%t got: %t", i, c.exp, got)
		}
	}
}

func TestRectToString(t *testing.T) {
	cases := []struct {
		input []int
		exp   string
	}{
		{
			input: []int{0, 0, 10, 10},
			exp:   "(0,0)/(10-10)",
		},
		{
			input: []int{5, 5, 7, 7},
			exp:   "(5,5)/(7-7)",
		},
	}
	for i, c := range cases {
		origin := api.NewPoint(c.input[0], c.input[1])
		size := api.NewSize(c.input[2], c.input[3])
		rect := api.NewRect(origin, size)
		got := rect.ToString()
		if c.exp != got {
			t.Errorf("[%d] ToString Error exp:%s got:%s", i, c.exp, got)
		}
	}
}

func TestRectSaveToDict(t *testing.T) {
	cases := []struct {
		input []int
		exp   map[string]any
	}{
		{
			input: []int{0, 0, 10, 10},
			exp:   map[string]any{"origin": map[string]any{"x": 0, "y": 0}, "size": map[string]any{"w": 10, "h": 10}},
		},
		{
			input: []int{5, 5, 7, 7},
			exp:   map[string]any{"origin": map[string]any{"x": 5, "y": 5}, "size": map[string]any{"w": 7, "h": 7}},
		},
	}
	for i, c := range cases {
		origin := api.NewPoint(c.input[0], c.input[1])
		size := api.NewSize(c.input[2], c.input[3])
		rect := api.NewRect(origin, size)
		got := rect.SaveToDict()
		if c.exp["origin"].(map[string]any)["x"] != got["origin"].(map[string]any)["x"] {
			t.Errorf("[%d] SaveToDict Origin.X Error exp:%d got:%d",
				i, c.exp["origin"].(map[string]any)["x"], got["origin"].(map[string]any)["x"])
		}
		if c.exp["origin"].(map[string]any)["y"] != got["origin"].(map[string]any)["y"] {
			t.Errorf("[%d] SaveToDict Origin.Y Error exp:%d got:%d",
				i, c.exp["origin"].(map[string]any)["y"], got["origin"].(map[string]any)["y"])
		}
		if c.exp["size"].(map[string]any)["w"] != got["size"].(map[string]any)["w"] {
			t.Errorf("[%d] SaveToDict Size.W Error exp:%d got:%d",
				i, c.exp["size"].(map[string]any)["w"], got["size"].(map[string]any)["w"])
		}
		if c.exp["size"].(map[string]any)["h"] != got["size"].(map[string]any)["h"] {
			t.Errorf("[%d] SaveToDict Size.H Error exp:%d got:%d",
				i, c.exp["size"].(map[string]any)["h"], got["size"].(map[string]any)["h"])
		}
	}
}
