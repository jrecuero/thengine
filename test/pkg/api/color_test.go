package api_test

import (
	"testing"

	"github.com/jrecuero/thengine/pkg/api"
)

func TestColorsToString(t *testing.T) {
	cases := []struct {
		input api.Attr
		exp   string
	}{
		{api.ColorBlack, "black"},
		{api.ColorRed, "red"},
		{api.ColorGreen, "green"},
		{api.ColorYellow, "yellow"},
		{api.ColorBlue, "blue"},
		{api.ColorMagenta, "magenta"},
		{api.ColorCyan, "cyan"},
		{api.ColorWhite, "white"},
		{api.ColorDarkGray, "darkgray"},
		{api.ColorLightRed, "lightred"},
		{api.ColorLightGreen, "lightgreen"},
		{api.ColorLightYellow, "lightyellow"},
		{api.ColorLightBlue, "lightblue"},
		{api.ColorLightMagenta, "lightmagenta"},
		{api.ColorLightCyan, "lightcyan"},
		{api.ColorLightGray, "lightgray"},
		{api.ColorDefault, "default"},
	}
	for i, c := range cases {
		got := api.ColorToString(c.input)
		if c.exp != got {
			t.Errorf("[%d] ColorToString Error exp:%s got:%s", i, c.exp, got)
		}
	}
}

func TestColorsToAttr(t *testing.T) {
	cases := []struct {
		input string
		exp   api.Attr
	}{
		{"black", api.ColorBlack},
		{"red", api.ColorRed},
		{"green", api.ColorGreen},
		{"yellow", api.ColorYellow},
		{"blue", api.ColorBlue},
		{"magenta", api.ColorMagenta},
		{"cyan", api.ColorCyan},
		{"white", api.ColorWhite},
		{"darkgray", api.ColorDarkGray},
		{"lightred", api.ColorLightRed},
		{"lightgreen", api.ColorLightGreen},
		{"lightyellow", api.ColorLightYellow},
		{"lightblue", api.ColorLightBlue},
		{"lightmagenta", api.ColorLightMagenta},
		{"lightcyan", api.ColorLightCyan},
		{"lightgray", api.ColorLightGray},
		{"default", api.ColorDefault},
	}
	for i, c := range cases {
		got := api.ColorToAttr(c.input)
		if c.exp != got {
			t.Errorf("[%d] ColorToAttr Error exp:%d got:%d", i, c.exp, got)
		}
	}
}

func TestColor(t *testing.T) {
	cases := []struct {
		input []api.Attr
		exp   []api.Attr
	}{
		{
			input: []api.Attr{api.ColorBlack, api.ColorRed},
			exp:   []api.Attr{api.ColorBlack, api.ColorRed},
		},
		{
			input: []api.Attr{api.ColorGreen, api.ColorYellow},
			exp:   []api.Attr{api.ColorGreen, api.ColorYellow},
		},
		{
			input: []api.Attr{api.ColorBlue, api.ColorMagenta},
			exp:   []api.Attr{api.ColorBlue, api.ColorMagenta},
		},
		{
			input: []api.Attr{api.ColorCyan, api.ColorWhite},
			exp:   []api.Attr{api.ColorCyan, api.ColorWhite},
		},
		{
			input: []api.Attr{api.ColorDarkGray, api.ColorLightRed},
			exp:   []api.Attr{api.ColorDarkGray, api.ColorLightRed},
		},
		{
			input: []api.Attr{api.ColorLightGreen, api.ColorLightYellow},
			exp:   []api.Attr{api.ColorLightGreen, api.ColorLightYellow},
		},
		{
			input: []api.Attr{api.ColorLightBlue, api.ColorLightMagenta},
			exp:   []api.Attr{api.ColorLightBlue, api.ColorLightMagenta},
		},
		{
			input: []api.Attr{api.ColorLightCyan, api.ColorLightGray},
			exp:   []api.Attr{api.ColorLightCyan, api.ColorLightGray},
		},
	}
	for i, c := range cases {
		got := api.NewColor(c.input[0], c.input[1])
		if got == nil {
			t.Errorf("[%d] NewColor Error exp:*Color got:nil", i)
			continue
		}
		if c.exp[0] != got.Fg {
			t.Errorf("[%d] NewColor Fg exp:%d got:%d", i, c.exp[0], got.Fg)
		}
		if c.exp[1] != got.Bg {
			t.Errorf("[%d] NewColor Bg exp:%d got:%d", i, c.exp[1], got.Bg)
		}
	}
}

func TestColorCloneColor(t *testing.T) {
	cases := []struct {
		input []api.Attr
		exp   []api.Attr
	}{
		{
			input: []api.Attr{api.ColorBlack, api.ColorRed},
			exp:   []api.Attr{api.ColorBlack, api.ColorRed},
		},
		{
			input: []api.Attr{api.ColorGreen, api.ColorYellow},
			exp:   []api.Attr{api.ColorGreen, api.ColorYellow},
		},
		{
			input: []api.Attr{api.ColorBlue, api.ColorMagenta},
			exp:   []api.Attr{api.ColorBlue, api.ColorMagenta},
		},
		{
			input: []api.Attr{api.ColorCyan, api.ColorWhite},
			exp:   []api.Attr{api.ColorCyan, api.ColorWhite},
		},
		{
			input: []api.Attr{api.ColorDarkGray, api.ColorLightRed},
			exp:   []api.Attr{api.ColorDarkGray, api.ColorLightRed},
		},
		{
			input: []api.Attr{api.ColorLightGreen, api.ColorLightYellow},
			exp:   []api.Attr{api.ColorLightGreen, api.ColorLightYellow},
		},
		{
			input: []api.Attr{api.ColorLightBlue, api.ColorLightMagenta},
			exp:   []api.Attr{api.ColorLightBlue, api.ColorLightMagenta},
		},
		{
			input: []api.Attr{api.ColorLightCyan, api.ColorLightGray},
			exp:   []api.Attr{api.ColorLightCyan, api.ColorLightGray},
		},
	}
	for i, c := range cases {
		toClone := api.NewColor(c.input[0], c.input[1])
		got := api.CloneColor(toClone)
		if got == nil {
			t.Errorf("[%d] CloneColor Error exp:*Color got:nil", i)
			continue
		}
		if c.exp[0] != got.Fg {
			t.Errorf("[%d] CloneColor Fg exp:%d got:%d", i, c.exp[0], got.Fg)
		}
		if c.exp[1] != got.Bg {
			t.Errorf("[%d] CloneColor Bg exp:%d got:%d", i, c.exp[1], got.Bg)
		}
	}
}

func TestColorNewFromString(t *testing.T) {
	cases := []struct {
		input []string
		exp   []api.Attr
	}{
		{
			[]string{"black", "red"},
			[]api.Attr{api.ColorBlack, api.ColorRed},
		},
		{
			[]string{"green", "yellow"},
			[]api.Attr{api.ColorGreen, api.ColorYellow},
		},
		{
			[]string{"blue", "magenta"},
			[]api.Attr{api.ColorBlue, api.ColorMagenta},
		},
		{
			[]string{"cyan", "white"},
			[]api.Attr{api.ColorCyan, api.ColorWhite},
		},
		{
			[]string{"darkgray", "lightred"},
			[]api.Attr{api.ColorDarkGray, api.ColorLightRed},
		},
		{
			[]string{"lightgreen", "lightyellow"},
			[]api.Attr{api.ColorLightGreen, api.ColorLightYellow},
		},
		{
			[]string{"lightblue", "lightmagenta"},
			[]api.Attr{api.ColorLightBlue, api.ColorLightMagenta},
		},
		{
			[]string{"lightcyan", "lightgray"},
			[]api.Attr{api.ColorLightCyan, api.ColorLightGray},
		},
	}
	for i, c := range cases {
		got := api.NewColorFromString(c.input[0], c.input[1])
		if got == nil {
			t.Errorf("[%d] NewColorFromString Error exp:*Color got:nil", i)
			continue
		}
		if c.exp[0] != got.Fg {
			t.Errorf("[%d] NewColorFromString Fg exp:%d got:%d", i, c.exp[0], got.Fg)
		}
		if c.exp[1] != got.Bg {
			t.Errorf("[%d] NewColorFromString Bg exp:%d got:%d", i, c.exp[1], got.Bg)
		}
	}
}

func TestColorClone(t *testing.T) {
	cases := []struct {
		input []api.Attr
		exp   []api.Attr
	}{
		{
			input: []api.Attr{api.ColorBlack, api.ColorRed},
			exp:   []api.Attr{api.ColorBlack, api.ColorRed},
		},
		{
			input: []api.Attr{api.ColorGreen, api.ColorYellow},
			exp:   []api.Attr{api.ColorGreen, api.ColorYellow},
		},
		{
			input: []api.Attr{api.ColorBlue, api.ColorMagenta},
			exp:   []api.Attr{api.ColorBlue, api.ColorMagenta},
		},
		{
			input: []api.Attr{api.ColorCyan, api.ColorWhite},
			exp:   []api.Attr{api.ColorCyan, api.ColorWhite},
		},
		{
			input: []api.Attr{api.ColorDarkGray, api.ColorLightRed},
			exp:   []api.Attr{api.ColorDarkGray, api.ColorLightRed},
		},
		{
			input: []api.Attr{api.ColorLightGreen, api.ColorLightYellow},
			exp:   []api.Attr{api.ColorLightGreen, api.ColorLightYellow},
		},
		{
			input: []api.Attr{api.ColorLightBlue, api.ColorLightMagenta},
			exp:   []api.Attr{api.ColorLightBlue, api.ColorLightMagenta},
		},
		{
			input: []api.Attr{api.ColorLightCyan, api.ColorLightGray},
			exp:   []api.Attr{api.ColorLightCyan, api.ColorLightGray},
		},
	}
	for i, c := range cases {
		toClone := api.NewColor(c.input[0], c.input[1])
		got := api.NewColor(api.ColorDefault, api.ColorDefault)
		got.Clone(toClone)
		if got == nil {
			t.Errorf("[%d] Clone Error exp:*Color got:nil", i)
			continue
		}
		if c.exp[0] != got.Fg {
			t.Errorf("[%d] Clone Fg exp:%d got:%d", i, c.exp[0], got.Fg)
		}
		if c.exp[1] != got.Bg {
			t.Errorf("[%d] Clone Bg exp:%d got:%d", i, c.exp[1], got.Bg)
		}
	}
}

func TestColorGet(t *testing.T) {
	cases := []struct {
		input []api.Attr
		exp   []api.Attr
	}{
		{
			input: []api.Attr{api.ColorBlack, api.ColorRed},
			exp:   []api.Attr{api.ColorBlack, api.ColorRed},
		},
		{
			input: []api.Attr{api.ColorGreen, api.ColorYellow},
			exp:   []api.Attr{api.ColorGreen, api.ColorYellow},
		},
		{
			input: []api.Attr{api.ColorBlue, api.ColorMagenta},
			exp:   []api.Attr{api.ColorBlue, api.ColorMagenta},
		},
		{
			input: []api.Attr{api.ColorCyan, api.ColorWhite},
			exp:   []api.Attr{api.ColorCyan, api.ColorWhite},
		},
		{
			input: []api.Attr{api.ColorDarkGray, api.ColorLightRed},
			exp:   []api.Attr{api.ColorDarkGray, api.ColorLightRed},
		},
		{
			input: []api.Attr{api.ColorLightGreen, api.ColorLightYellow},
			exp:   []api.Attr{api.ColorLightGreen, api.ColorLightYellow},
		},
		{
			input: []api.Attr{api.ColorLightBlue, api.ColorLightMagenta},
			exp:   []api.Attr{api.ColorLightBlue, api.ColorLightMagenta},
		},
		{
			input: []api.Attr{api.ColorLightCyan, api.ColorLightGray},
			exp:   []api.Attr{api.ColorLightCyan, api.ColorLightGray},
		},
	}
	for i, c := range cases {
		p := api.NewColor(c.input[0], c.input[1])
		gotFg, gotBg := p.Get()
		if c.exp[0] != gotFg {
			t.Errorf("[%d] Get Fg exp:%d got:%d", i, c.exp[0], gotFg)
		}
		if c.exp[1] != gotBg {
			t.Errorf("[%d] Get Bg exp:%d got:%d", i, c.exp[1], gotBg)
		}
	}
}

func TestColorSet(t *testing.T) {
	cases := []struct {
		input []api.Attr
		exp   []api.Attr
	}{
		{
			input: []api.Attr{api.ColorBlack, api.ColorRed},
			exp:   []api.Attr{api.ColorBlack, api.ColorRed},
		},
		{
			input: []api.Attr{api.ColorGreen, api.ColorYellow},
			exp:   []api.Attr{api.ColorGreen, api.ColorYellow},
		},
		{
			input: []api.Attr{api.ColorBlue, api.ColorMagenta},
			exp:   []api.Attr{api.ColorBlue, api.ColorMagenta},
		},
		{
			input: []api.Attr{api.ColorCyan, api.ColorWhite},
			exp:   []api.Attr{api.ColorCyan, api.ColorWhite},
		},
		{
			input: []api.Attr{api.ColorDarkGray, api.ColorLightRed},
			exp:   []api.Attr{api.ColorDarkGray, api.ColorLightRed},
		},
		{
			input: []api.Attr{api.ColorLightGreen, api.ColorLightYellow},
			exp:   []api.Attr{api.ColorLightGreen, api.ColorLightYellow},
		},
		{
			input: []api.Attr{api.ColorLightBlue, api.ColorLightMagenta},
			exp:   []api.Attr{api.ColorLightBlue, api.ColorLightMagenta},
		},
		{
			input: []api.Attr{api.ColorLightCyan, api.ColorLightGray},
			exp:   []api.Attr{api.ColorLightCyan, api.ColorLightGray},
		},
	}
	for i, c := range cases {
		got := api.NewColor(api.ColorDefault, api.ColorDefault)
		got.Set(c.input[0], c.input[1])
		if c.exp[0] != got.Fg {
			t.Errorf("[%d] Set Fg exp:%d got:%d", i, c.exp[0], got.Fg)
		}
		if c.exp[1] != got.Bg {
			t.Errorf("[%d] Set Bg exp:%d got:%d", i, c.exp[1], got.Bg)
		}
	}
}

func TestColorToStringAttr(t *testing.T) {
	cases := []struct {
		input []api.Attr
		exp   string
	}{
		{
			input: []api.Attr{api.ColorBlack, api.ColorRed},
			exp:   "[1:2]",
		},
		{
			input: []api.Attr{api.ColorGreen, api.ColorYellow},
			exp:   "[3:4]",
		},
		{
			input: []api.Attr{api.ColorBlue, api.ColorMagenta},
			exp:   "[5:6]",
		},
		{
			input: []api.Attr{api.ColorCyan, api.ColorWhite},
			exp:   "[7:8]",
		},
		{
			input: []api.Attr{api.ColorDarkGray, api.ColorLightRed},
			exp:   "[9:10]",
		},
		{
			input: []api.Attr{api.ColorLightGreen, api.ColorLightYellow},
			exp:   "[11:12]",
		},
		{
			input: []api.Attr{api.ColorLightBlue, api.ColorLightMagenta},
			exp:   "[13:14]",
		},
		{
			input: []api.Attr{api.ColorLightCyan, api.ColorLightGray},
			exp:   "[15:16]",
		},
	}
	for i, c := range cases {
		color := api.NewColor(c.input[0], c.input[1])
		got := color.ToStringAttr()
		if c.exp != got {
			t.Errorf("[%d] ToStringAttr Fg exp:%s got:%s", i, c.exp, got)
		}
	}
}

func TestColorEquals(t *testing.T) {
	cases := []struct {
		input []api.Attr
		exp   bool
	}{
		{
			input: []api.Attr{api.ColorBlack, api.ColorRed, api.ColorBlack, api.ColorRed},
			exp:   true,
		},
		{
			input: []api.Attr{api.ColorGreen, api.ColorYellow, api.ColorGreen, api.ColorYellow},
			exp:   true,
		},
		{
			input: []api.Attr{api.ColorBlue, api.ColorMagenta, api.ColorBlue, api.ColorMagenta},
			exp:   true,
		},
		{
			input: []api.Attr{api.ColorBlack, api.ColorRed, api.ColorBlack, api.ColorWhite},
			exp:   false,
		},
		{
			input: []api.Attr{api.ColorGreen, api.ColorYellow, api.ColorDefault, api.ColorYellow},
			exp:   false,
		},
		{
			input: []api.Attr{api.ColorBlue, api.ColorMagenta, api.ColorBlue, api.ColorBlack},
			exp:   false,
		},
	}
	for i, c := range cases {
		toEquals := api.NewColor(c.input[0], c.input[1])
		color := api.NewColor(c.input[2], c.input[3])
		got := color.IsEqual(toEquals)
		if c.exp != got {
			t.Errorf("[%d] IsEqual exp:%t got:%t", i, c.exp, got)
		}
	}
}

func TestColorToString(t *testing.T) {
	cases := []struct {
		input []api.Attr
		exp   string
	}{
		{
			input: []api.Attr{api.ColorBlack, api.ColorRed},
			exp:   "[black:red]",
		},
		{
			input: []api.Attr{api.ColorGreen, api.ColorYellow},
			exp:   "[green:yellow]",
		},
		{
			input: []api.Attr{api.ColorBlue, api.ColorMagenta},
			exp:   "[blue:magenta]",
		},
		{
			input: []api.Attr{api.ColorCyan, api.ColorWhite},
			exp:   "[cyan:white]",
		},
		{
			input: []api.Attr{api.ColorDarkGray, api.ColorLightRed},
			exp:   "[darkgray:lightred]",
		},
		{
			input: []api.Attr{api.ColorLightGreen, api.ColorLightYellow},
			exp:   "[lightgreen:lightyellow]",
		},
		{
			input: []api.Attr{api.ColorLightBlue, api.ColorLightMagenta},
			exp:   "[lightblue:lightmagenta]",
		},
		{
			input: []api.Attr{api.ColorLightCyan, api.ColorLightGray},
			exp:   "[lightcyan:lightgray]",
		},
	}
	for i, c := range cases {
		color := api.NewColor(c.input[0], c.input[1])
		got := color.ToString()
		if c.exp != got {
			t.Errorf("[%d] ToString Fg exp:%s got:%s", i, c.exp, got)
		}
	}
}

func TestColorSaveToDict(t *testing.T) {
	cases := []struct {
		input []api.Attr
		exp   map[string]any
	}{
		{
			input: []api.Attr{api.ColorBlack, api.ColorRed},
			exp:   map[string]any{"fg": api.ColorBlack, "bg": api.ColorRed},
		},
		{
			input: []api.Attr{api.ColorBlue, api.ColorWhite},
			exp:   map[string]any{"fg": api.ColorBlue, "bg": api.ColorWhite},
		},
	}
	for i, c := range cases {
		color := api.NewColor(c.input[0], c.input[1])
		got := color.SaveToDict()
		if c.exp["fg"] != got["fg"] {
			t.Errorf("[%d] SaveToDict Fg exp:%+v got:%+v", i, c.exp, got)
		}
		if c.exp["bg"] != got["bg"] {
			t.Errorf("[%d] SaveToDict Bg exp:%+v got:%+v", i, c.exp, got)
		}
	}
}
