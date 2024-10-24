package engine_test

import (
	"testing"

	"github.com/gdamore/tcell/v2"
	"github.com/jrecuero/thengine/pkg/api"
	"github.com/jrecuero/thengine/pkg/engine"
	"github.com/jrecuero/thengine/pkg/tools"
)

func TestRow(t *testing.T) {
	cases := []struct {
		input int
		exp   int
	}{
		{
			input: 1,
			exp:   1,
		},
		{
			input: 3,
			exp:   3,
		},
	}
	for i, c := range cases {
		got := engine.NewRow(c.input)
		if got == nil {
			t.Errorf("[%d] NewRow Error exp:*Row got:nil", i)
			continue
		}
		if c.exp != len(got.Cols) {
			t.Errorf("[%d] NewRow Cols exp:%d got:%d", i, c.exp, len(got.Cols))
		}
	}
}

// func TestRowSaveToDict(t *testing.T) {
// 	createCells()
// 	cases := []struct {
// 		input []*engine.Cell
// 		exp   map[string]any
// 	}{
// 		{
// 			input: []*engine.Cell{cells[0], cells[1], cells[2]},
// 			exp: map[string]any{
// 				"cols": []map[string]any{
// 					{
// 						"color": map[string]any{
// 							"fg": tcell.ColorBlack,
// 							"bg": tcell.ColorWhite,
// 						},
// 						"rune": '0',
// 					},
// 					{
// 						"color": map[string]any{
// 							"fg": tcell.ColorBlue,
// 							"bg": tcell.ColorRed,
// 						},
// 						"rune": '1',
// 					},
// 					{
// 						"color": map[string]any{
// 							"fg": tcell.ColorLightBlue,
// 							"bg": tcell.ColorDefault,
// 						},
// 						"rune": '2',
// 					},
// 				},
// 			},
// 		},
// 	}
// 	for i, c := range cases {
// 		row := engine.NewRow(len(c.input))
// 		copy(row.Cols, c.input)
// 		got := row.SaveToDict()
// 		if len(c.exp) != len(got) {
// 			t.Errorf("[%d] SaveToDict Length Error exp:%d got:%d", i, len(c.exp), len(got))
// 		}
// 		if len(c.exp["cols"].([]map[string]any)) != len(got["cols"].([]map[string]any)) {
// 			t.Errorf("[%d] SaveToDict Cols-Length Error exp:%d got:%d",
// 				i, len(c.exp["cols"].([]map[string]any)), len(got["cols"].([]map[string]any)))
// 		}
// 		for x, entry := range c.exp["cols"].([]map[string]any) {
// 			var expRune rune = entry["rune"].(rune)
// 			var gotRune rune = got["cols"].([]map[string]any)[x]["rune"].(rune)
// 			if expRune != gotRune {
// 				t.Errorf("[%d] SaveToDict Cols:Rune (%d) Error exp%c got%c", i, x, expRune, gotRune)
// 			}
// 			var expColorFg tcell.Color = entry["color"].(map[string]any)["fg"].(tcell.Color)
// 			var gotColorFg tcell.Color = got["cols"].([]map[string]any)[x]["color"].(map[string]any)["fg"].(tcell.Color)
// 			if expColorFg != gotColorFg {
// 				t.Errorf("[%d] SaveToDict Cols:Color:Fg (%d) Error exp:%d got:%d", i, x, expColorFg, gotColorFg)
// 			}
// 			var expColorBg tcell.Color = entry["color"].(map[string]any)["bg"].(tcell.Color)
// 			var gotColorBg tcell.Color = got["cols"].([]map[string]any)[x]["color"].(map[string]any)["bg"].(tcell.Color)
// 			if expColorBg != gotColorBg {
// 				t.Errorf("[%d] SaveToDict Cols:Color:Bg (%d) Error exp:%d got:%d", i, x, expColorBg, gotColorBg)
// 			}
// 		}
// 	}
// }

func TestCanvas(t *testing.T) {
	cases := []struct {
		input *api.Size
		exp   *api.Size
	}{
		{
			input: api.NewSize(1, 10),
			exp:   api.NewSize(1, 10),
		},
		{
			input: api.NewSize(5, 1),
			exp:   api.NewSize(5, 1),
		},
		{
			input: api.NewSize(3, 3),
			exp:   api.NewSize(3, 3),
		},
	}
	for i, c := range cases {
		got := engine.NewCanvas(c.input)
		if got == nil {
			t.Errorf("[%d] NewCanvas exp:*Canvas got:nil", i)
			continue
		}
		if c.exp.H != len(got.Rows) {
			t.Errorf("[%d] NewCanvas Rows Error exp:%d got:%d", i, c.exp.H, len(got.Rows))
		}
		for j, r := range got.Rows {
			if c.exp.W != len(r.Cols) {
				t.Errorf("[%d][%d] NewCanvas Cols Error exp:%d got:%d", i, j, c.exp.W, len(r.Cols))
			}
		}
	}
}

func TestCanvasCloneCanvas(t *testing.T) {
	createCells()
	cases := []struct {
		input [][]engine.ICell
		exp   *api.Size
	}{
		{
			input: [][]engine.ICell{{cells[0], cells[1]}, {cells[2], cells[3]}, {cells[4], cells[5]}},
			exp:   api.NewSize(2, 3),
		},
		{
			input: [][]engine.ICell{{cells[0], cells[1], cells[2]}, {cells[3], cells[4], cells[5]}},
			exp:   api.NewSize(3, 2),
		},
	}
	for i, c := range cases {
		toClone := engine.NewCanvas(api.NewSize(len(c.input[0]), len(c.input)))
		for x, rows := range c.input {
			copy(toClone.Rows[x].Cols, rows)
		}
		got := engine.CloneCanvas(toClone)
		got.Clone(toClone)
		if c.exp.H != len(got.Rows) {
			t.Errorf("[%d] CloneCanvas Rows Error exp:%d got:%d", i, c.exp.H, len(got.Rows))
		}
		for j, r := range got.Rows {
			if c.exp.W != len(r.Cols) {
				t.Errorf("[%d][%d] CloneCanvas Cols Error exp:%d got:%d", i, j, c.exp.W, len(r.Cols))
			}
		}
		for x, rows := range got.Rows {
			for y, col := range rows.Cols {
				if !got.Rows[x].Cols[y].IsEqual(col) {
					t.Errorf("[%d] CloneCanvas (%d,%d) Error exp:%s got:%s",
						i, x, y, got.Rows[x].Cols[y].ToString(), col.ToString())
				}
			}
		}
	}
}

func TestCanvasNewCanvasFromString(t *testing.T) {
	cases := []struct {
		input struct {
			str   string
			style *tcell.Style
		}
		exp struct {
			width  int
			height int
			str    []string
			style  *tcell.Style
		}
	}{
		{
			input: struct {
				str   string
				style *tcell.Style
			}{
				str:   "Hello",
				style: engine.NewStyle(tcell.ColorBlue, tcell.ColorRed, 0),
			},
			exp: struct {
				width  int
				height int
				str    []string
				style  *tcell.Style
			}{
				width:  5,
				height: 1,
				str:    []string{"Hello"},
				style:  engine.NewStyle(tcell.ColorBlue, tcell.ColorRed, 0),
			},
		},
		{
			input: struct {
				str   string
				style *tcell.Style
			}{
				str:   "Hello\nDeveloper",
				style: engine.NewStyle(tcell.ColorLightBlue, tcell.ColorYellow, 0),
			},
			exp: struct {
				width  int
				height int
				str    []string
				style  *tcell.Style
			}{
				width:  9,
				height: 2,
				str:    []string{"Hello", "Developer"},
				style:  engine.NewStyle(tcell.ColorLightBlue, tcell.ColorYellow, 0),
			},
		},
		{
			input: struct {
				str   string
				style *tcell.Style
			}{
				str:   "Have\nA great\n day",
				style: nil,
			},
			exp: struct {
				width  int
				height int
				str    []string
				style  *tcell.Style
			}{
				width:  7,
				height: 3,
				str:    []string{"Have", "A great", " day"},
				style:  engine.NewStyle(tcell.ColorDefault, tcell.ColorDefault, 0),
			},
		},
	}
	for i, c := range cases {
		got := engine.NewCanvasFromString(c.input.str, c.input.style)
		if c.exp.width != got.Width() {
			t.Errorf("[%d] NewCanvasFromString Width Error exp:%d, got:%d", i, c.exp.width, got.Width())
		}
		if c.exp.height != got.Height() {
			t.Errorf("[%d] NewCanvasFromString Height Error exp:%d got:%d", i, c.exp.height, got.Height())
		}
		for row, line := range c.exp.str {
			for col, ch := range line {
				cell := got.GetCellAt(api.NewPoint(col, row))
				if cell == nil {
					t.Errorf("[%d] NewCanvasFromString (%d,%d) Cell exp:*Cell got:nil", i, col, row)
					continue
				}
				if !engine.CompareStyle(c.exp.style, cell.GetStyle()) {
					t.Errorf("[%d] NewCanvasFromString (%d,%d) Style exp:%s got:%s",
						i, col, row, engine.StyleToString(c.exp.style), engine.StyleToString(cell.GetStyle()))
				}
				if ch != cell.GetRune() {
					t.Errorf("[%d] NewCanvasFromString (%d,%d) Rune exp:%c got:%c", i, col, row, ch, cell.GetRune())
				}
			}
		}
	}
}

func TestCanvasWriteStringInCanvas(t *testing.T) {
	cases := []struct {
		input struct {
			size  *api.Size
			str   string
			style *tcell.Style
		}
		exp struct {
			width  int
			height int
			str    []string
			style  *tcell.Style
		}
	}{
		{
			input: struct {
				size  *api.Size
				str   string
				style *tcell.Style
			}{
				size:  api.NewSize(10, 1),
				str:   "Hello",
				style: engine.NewStyle(tcell.ColorBlue, tcell.ColorRed, 0),
			},
			exp: struct {
				width  int
				height int
				str    []string
				style  *tcell.Style
			}{
				width:  10,
				height: 1,
				str:    []string{"Hello"},
				style:  engine.NewStyle(tcell.ColorBlue, tcell.ColorRed, 0),
			},
		},
		{
			input: struct {
				size  *api.Size
				str   string
				style *tcell.Style
			}{
				size:  api.NewSize(20, 2),
				str:   "Hello\nDeveloper",
				style: engine.NewStyle(tcell.ColorLightBlue, tcell.ColorYellow, 0),
			},
			exp: struct {
				width  int
				height int
				str    []string
				style  *tcell.Style
			}{
				width:  20,
				height: 2,
				str:    []string{"Hello", "Developer"},
				style:  engine.NewStyle(tcell.ColorLightBlue, tcell.ColorYellow, 0),
			},
		},
		{
			input: struct {
				size  *api.Size
				str   string
				style *tcell.Style
			}{
				size:  api.NewSize(20, 3),
				str:   "Have\nA great\n day",
				style: nil,
			},
			exp: struct {
				width  int
				height int
				str    []string
				style  *tcell.Style
			}{
				width:  20,
				height: 3,
				str:    []string{"Have", "A great", " day"},
				style:  engine.NewStyle(tcell.ColorDefault, tcell.ColorDefault, 0),
			},
		},
		{
			input: struct {
				size  *api.Size
				str   string
				style *tcell.Style
			}{
				size:  api.NewSize(10, 1),
				str:   "Have A great day",
				style: nil,
			},
			exp: struct {
				width  int
				height int
				str    []string
				style  *tcell.Style
			}{
				width:  10,
				height: 1,
				str:    []string{"Have A gre"},
				style:  engine.NewStyle(tcell.ColorDefault, tcell.ColorDefault, 0),
			},
		},
		{
			input: struct {
				size  *api.Size
				str   string
				style *tcell.Style
			}{
				size:  api.NewSize(20, 2),
				str:   "Have\nA great\n day",
				style: nil,
			},
			exp: struct {
				width  int
				height int
				str    []string
				style  *tcell.Style
			}{
				width:  20,
				height: 2,
				str:    []string{"Have", "A great"},
				style:  engine.NewStyle(tcell.ColorDefault, tcell.ColorDefault, 0),
			},
		},
		{
			input: struct {
				size  *api.Size
				str   string
				style *tcell.Style
			}{
				size:  api.NewSize(3, 2),
				str:   "Have\nA great\n day",
				style: nil,
			},
			exp: struct {
				width  int
				height int
				str    []string
				style  *tcell.Style
			}{
				width:  3,
				height: 2,
				str:    []string{"Hav", "A g"},
				style:  engine.NewStyle(tcell.ColorDefault, tcell.ColorDefault, 0),
			},
		},
	}
	for i, c := range cases {
		got := engine.NewCanvas(c.input.size)
		got.WriteStringInCanvas(c.input.str, c.input.style)
		if c.exp.width != got.Width() {
			t.Errorf("[%d] NewCanvasFromString Width Error exp:%d, got:%d", i, c.exp.width, got.Width())
		}
		if c.exp.height != got.Height() {
			t.Errorf("[%d] NewCanvasFromString Height Error exp:%d got:%d", i, c.exp.height, got.Height())
		}
		for row, line := range c.exp.str {
			for col, ch := range line {
				cell := got.GetCellAt(api.NewPoint(col, row))
				if cell == nil {
					t.Errorf("[%d] NewCanvasFromString (%d,%d) Cell exp:*Cell got:nil", i, col, row)
					continue
				}
				if !engine.CompareStyle(c.exp.style, cell.GetStyle()) {
					t.Errorf("[%d] NewCanvasFromString (%d,%d) Style exp:%s got:%s",
						i, col, row, engine.StyleToString(c.exp.style), engine.StyleToString(cell.GetStyle()))
				}
				if ch != cell.GetRune() {
					t.Errorf("[%d] NewCanvasFromString (%d,%d) Rune exp:%c got:%c", i, col, row, ch, cell.GetRune())
				}
			}
		}
	}
}

func TestCanvasNewCanvasFromFile(t *testing.T) {
	cases := []struct {
		input struct {
			filename string
			style    *tcell.Style
		}
		exp struct {
			width  int
			height int
			str    []string
			style  *tcell.Style
		}
	}{
		{
			input: struct {
				filename string
				style    *tcell.Style
			}{
				filename: "assets/test/canvasString.01",
				style:    engine.NewStyle(tcell.ColorBlue, tcell.ColorRed, 0),
			},
			exp: struct {
				width  int
				height int
				str    []string
				style  *tcell.Style
			}{
				width:  5,
				height: 1,
				str:    []string{"Hello"},
				style:  engine.NewStyle(tcell.ColorBlue, tcell.ColorRed, 0),
			},
		},
		{
			input: struct {
				filename string
				style    *tcell.Style
			}{
				filename: "assets/test/canvasString.02",
				style:    engine.NewStyle(tcell.ColorDarkCyan, tcell.ColorYellow, 0),
			},
			exp: struct {
				width  int
				height int
				str    []string
				style  *tcell.Style
			}{
				width:  6,
				height: 2,
				str:    []string{"Hello", "World!"},
				style:  engine.NewStyle(tcell.ColorDarkCyan, tcell.ColorYellow, 0),
			},
		},
		//{
		//    input: struct {
		//        filename string
		//        style    *tcell.Style
		//    }{
		//        filename: "assets/test/canvasString.03",
		//        style:    engine.NewStyle(tcell.ColorRed, tcell.ColorBlue, 0),
		//    },
		//    exp: struct {
		//        width  int
		//        height int
		//        str    []string
		//        style  *tcell.Style
		//    }{
		//        width:  10,
		//        height: 3,
		//        str:    []string{"Hi", "My name is", "Developer"},
		//        style:  engine.NewStyle(tcell.ColorRed, tcell.ColorBlue, 0),
		//    },
		//},
	}
	for i, c := range cases {
		got := engine.NewCanvasFromFile(c.input.filename, c.input.style)
		if c.exp.width != got.Width() {
			t.Errorf("[%d] NewCanvasFromString Width Error exp:%d, got:%d", i, c.exp.width, got.Width())
		}
		if c.exp.height != got.Height() {
			t.Errorf("[%d] NewCanvasFromString Height Error exp:%d got:%d", i, c.exp.height, got.Height())
		}
		for row, line := range c.exp.str {
			for col, ch := range line {
				cell := got.GetCellAt(api.NewPoint(col, row))
				if cell == nil {
					t.Errorf("[%d] NewCanvasFromString (%d,%d) Cell exp:*Cell got:nil", i, col, row)
					continue
				}
				if !engine.CompareStyle(c.exp.style, cell.GetStyle()) {
					t.Errorf("[%d] NewCanvasFromString (%d,%d) Color exp:%s got:%s",
						i, col, row, engine.StyleToString(c.exp.style), engine.StyleToString(cell.GetStyle()))
				}
				if ch != cell.GetRune() {
					t.Errorf("[%d] NewCanvasFromString (%d,%d) Rune exp:%c got:%c", i, col, row, ch, cell.GetRune())
				}
			}
		}
	}
}

func TestCanvasHeight(t *testing.T) {
	cases := []struct {
		input *api.Size
		exp   int
	}{
		{
			input: api.NewSize(1, 10),
			exp:   10,
		},
		{
			input: api.NewSize(5, 1),
			exp:   1,
		},
		{
			input: api.NewSize(3, 7),
			exp:   7,
		},
	}
	for i, c := range cases {
		canvas := engine.NewCanvas(c.input)
		got := canvas.Height()
		if c.exp != got {
			t.Errorf("[%d] Height Error exp:%d got:%d", i, c.exp, got)
		}
	}
}

func TestCanvasSize(t *testing.T) {
	cases := []struct {
		input *api.Size
		exp   *api.Size
	}{
		{
			input: api.NewSize(1, 10),
			exp:   api.NewSize(1, 10),
		},
		{
			input: api.NewSize(5, 1),
			exp:   api.NewSize(5, 1),
		},
		{
			input: api.NewSize(3, 7),
			exp:   api.NewSize(3, 7),
		},
	}
	for i, c := range cases {
		canvas := engine.NewCanvas(c.input)
		got := canvas.Size()
		if c.exp.W != got.W {
			t.Errorf("[%d] Len Width Error exp:%d got:%d", i, c.exp.W, got.W)
		}
		if c.exp.H != got.H {
			t.Errorf("[%d] Len Height Error exp:%d got:%d", i, c.exp.H, got.H)
		}
	}
}

func TestCanvasClone(t *testing.T) {
	createCells()
	cases := []struct {
		input [][]engine.ICell
		exp   *api.Size
	}{
		{
			input: [][]engine.ICell{{cells[0], cells[1]}, {cells[2], cells[3]}, {cells[4], cells[5]}},
			exp:   api.NewSize(2, 3),
		},
		{
			input: [][]engine.ICell{{cells[0], cells[1], cells[2]}, {cells[3], cells[4], cells[5]}},
			exp:   api.NewSize(3, 2),
		},
	}
	for i, c := range cases {
		toClone := engine.NewCanvas(api.NewSize(len(c.input[0]), len(c.input)))
		for x, rows := range c.input {
			copy(toClone.Rows[x].Cols, rows)
		}
		got := engine.NewCanvas(api.NewSize(len(c.input[0]), len(c.input)))
		got.Clone(toClone)
		if c.exp.H != len(got.Rows) {
			t.Errorf("[%d] Clone Rows Error exp:%d got:%d", i, c.exp.H, len(got.Rows))
		}
		for j, r := range got.Rows {
			if c.exp.W != len(r.Cols) {
				t.Errorf("[%d][%d] Clone Cols Error exp:%d got:%d", i, j, c.exp.W, len(r.Cols))
			}
		}
		for x, rows := range got.Rows {
			for y, col := range rows.Cols {
				if !got.Rows[x].Cols[y].IsEqual(col) {
					t.Errorf("[%d] Clone (%d,%d) Error exp:%s got:%s",
						i, x, y, got.Rows[x].Cols[y].ToString(), col.ToString())
				}
			}
		}
	}
}

func TestCanvasIsEqual(t *testing.T) {
	createCells()
	cases := []struct {
		input [][]engine.ICell
		idata [][]engine.ICell
		exp   bool
	}{
		{
			input: [][]engine.ICell{{cells[0], cells[1]}, {cells[2], cells[3]}, {cells[4], cells[5]}},
			idata: [][]engine.ICell{{cells[0], cells[1]}, {cells[2], cells[3]}, {cells[4], cells[5]}},
			exp:   true,
		},
		{
			input: [][]engine.ICell{{cells[0], cells[1], cells[2]}, {cells[3], cells[4], cells[5]}},
			idata: [][]engine.ICell{{cells[0], cells[1], cells[2]}, {cells[3], cells[4], cells[5]}},
			exp:   true,
		},
		{
			input: [][]engine.ICell{{cells[0], cells[1]}, {cells[2], cells[3]}, {cells[4], cells[5]}},
			idata: [][]engine.ICell{{cells[0]}, {cells[2]}, {cells[4]}},
			exp:   false,
		},
		{
			input: [][]engine.ICell{{cells[0], cells[1]}, {cells[2], cells[3]}, {cells[4], cells[5]}},
			idata: [][]engine.ICell{{cells[0], cells[1]}, {cells[2], cells[3]}},
			exp:   false,
		},
		{
			input: [][]engine.ICell{{cells[0], cells[1]}, {cells[2], cells[3]}, {cells[4], cells[5]}},
			idata: [][]engine.ICell{{cells[0], cells[0]}, {cells[2], cells[3]}, {cells[4], cells[5]}},
			exp:   false,
		},
		{
			input: [][]engine.ICell{{cells[0], cells[1], cells[2]}, {cells[3], cells[4], cells[5]}},
			idata: [][]engine.ICell{{cells[0], cells[1], cells[2]}, {cells[3], cells[4], cells[4]}},
			exp:   false,
		},
	}
	for i, c := range cases {
		toClone := engine.NewCanvas(api.NewSize(len(c.idata[0]), len(c.idata)))
		for x, rows := range c.idata {
			copy(toClone.Rows[x].Cols, rows)
		}
		canvas := engine.NewCanvas(api.NewSize(len(c.input[0]), len(c.input)))
		for x, rows := range c.input {
			copy(canvas.Rows[x].Cols, rows)
		}
		got := canvas.IsEqual(toClone)
		if c.exp != got {
			t.Errorf("[%d] IsEqual Error exp:%t got:%t", i, c.exp, got)
		}
	}
}

func TestCanvasIsInside(t *testing.T) {
	createCells()
	cases := []struct {
		input [][]engine.ICell
		point *api.Point
		exp   bool
	}{
		{
			input: [][]engine.ICell{{cells[0], cells[1]}, {cells[2], cells[3]}, {cells[4], cells[5]}},
			point: api.NewPoint(0, 0),
			exp:   true,
		},
		{
			input: [][]engine.ICell{{cells[0], cells[1], cells[2]}, {cells[3], cells[4], cells[5]}},
			point: api.NewPoint(2, 1),
			exp:   true,
		},
		{
			input: [][]engine.ICell{{cells[0], cells[1]}, {cells[2], cells[3]}, {cells[4], cells[5]}},
			point: api.NewPoint(3, 3),
			exp:   false,
		},
	}
	for i, c := range cases {
		canvas := engine.NewCanvas(api.NewSize(len(c.input[0]), len(c.input)))
		for x, rows := range c.input {
			copy(canvas.Rows[x].Cols, rows)
		}
		got := canvas.IsInside(c.point)
		if c.exp != got {
			t.Errorf("[%d] IsInside Error exp:%t got:%t", i, c.exp, got)
			continue
		}
	}
}

func TestCanvasGetCellAt(t *testing.T) {
	createCells()
	cases := []struct {
		input [][]engine.ICell
		point *api.Point
		exp   *engine.Cell
	}{
		{
			input: [][]engine.ICell{{cells[0], cells[1]}, {cells[2], cells[3]}, {cells[4], cells[5]}},
			point: api.NewPoint(0, 0),
			exp:   cells[0],
		},
		{
			input: [][]engine.ICell{{cells[0], cells[1], cells[2]}, {cells[3], cells[4], cells[5]}},
			point: api.NewPoint(2, 1),
			exp:   cells[5],
		},
		{
			input: [][]engine.ICell{{cells[0], cells[1]}, {cells[2], cells[3]}, {cells[4], cells[5]}},
			point: api.NewPoint(3, 3),
			exp:   nil,
		},
	}
	for i, c := range cases {
		canvas := engine.NewCanvas(api.NewSize(len(c.input[0]), len(c.input)))
		for x, rows := range c.input {
			copy(canvas.Rows[x].Cols, rows)
		}
		got := canvas.GetCellAt(c.point)
		if (c.exp == nil) && (got == nil) {
			continue
		}
		if (c.exp != nil) && (got == nil) {
			t.Errorf("[%d] GetCellAt Error exp:%v got:nil", i, c.exp)
			continue
		}
		if (c.exp == nil) && (c.exp != got) {
			t.Errorf("[%d] GetCellAt Error exp:nil got:%v", i, got)
			continue
		}
		if (got != nil) && !c.exp.IsEqual(got) {
			t.Errorf("[%d] GetCellAt Error exp:%s got:%s", i, c.exp.ToString(), got.ToString())
		}
	}
}

// func TestCanvasGetColorAt(t *testing.T) {
// 	createCells()
// 	cases := []struct {
// 		input [][]*engine.Cell
// 		point *api.Point
// 		exp   *api.Color
// 	}{
// 		{
// 			input: [][]*engine.Cell{{cells[0], cells[1]}, {cells[2], cells[3]}, {cells[4], cells[5]}},
// 			point: api.NewPoint(0, 0),
// 			exp:   cells[0].Color,
// 		},
// 		{
// 			input: [][]*engine.Cell{{cells[0], cells[1], cells[2]}, {cells[3], cells[4], cells[5]}},
// 			point: api.NewPoint(2, 1),
// 			exp:   cells[5].Color,
// 		},
// 		{
// 			input: [][]*engine.Cell{{cells[0], cells[1]}, {cells[2], cells[3]}, {cells[4], cells[5]}},
// 			point: api.NewPoint(3, 3),
// 			exp:   nil,
// 		},
// 	}
// 	for i, c := range cases {
// 		canvas := engine.NewCanvas(api.NewSize(len(c.input[0]), len(c.input)))
// 		for x, rows := range c.input {
// 			copy(canvas.Rows[x].Cols, rows)
// 		}
// 		got := canvas.GetColorAt(c.point)
// 		if (c.exp == nil) && (got == nil) {
// 			continue
// 		}
// 		if (c.exp != nil) && (got == nil) {
// 			t.Errorf("[%d] GetColorAt Error exp:%v got:nil", i, c.exp)
// 			continue
// 		}
// 		if (c.exp == nil) && (c.exp != got) {
// 			t.Errorf("[%d] GetColorAt Error exp:nil got:%v", i, got)
// 			continue
// 		}
// 		if (got != nil) && !c.exp.IsEqual(got) {
// 			t.Errorf("[%d] GetColorAt Error exp:%s got:%s", i, c.exp.ToString(), got.ToString())
// 		}
// 	}
// }

func TestCanvasGetRuneAt(t *testing.T) {
	createCells()
	cases := []struct {
		input [][]engine.ICell
		point *api.Point
		exp   rune
	}{
		{
			input: [][]engine.ICell{{cells[0], cells[1]}, {cells[2], cells[3]}, {cells[4], cells[5]}},
			point: api.NewPoint(0, 0),
			exp:   cells[0].GetRune(),
		},
		{
			input: [][]engine.ICell{{cells[0], cells[1], cells[2]}, {cells[3], cells[4], cells[5]}},
			point: api.NewPoint(2, 1),
			exp:   cells[5].GetRune(),
		},
		{
			input: [][]engine.ICell{{cells[0], cells[1]}, {cells[2], cells[3]}, {cells[4], cells[5]}},
			point: api.NewPoint(3, 3),
			exp:   0,
		},
	}
	for i, c := range cases {
		canvas := engine.NewCanvas(api.NewSize(len(c.input[0]), len(c.input)))
		for x, rows := range c.input {
			copy(canvas.Rows[x].Cols, rows)
		}
		got := canvas.GetRuneAt(c.point)
		if c.exp != got {
			t.Errorf("[%d] GetRuneAt Error exp:%c got:%c", i, c.exp, got)
		}
	}
}

// func TestCanvasSetCellAt(t *testing.T) {
// 	createCells()
// 	cases := []struct {
// 		input [][]*engine.Cell
// 		point *api.Point
// 		cell  *engine.Cell
// 		exp   struct {
// 			result bool
// 			cell   *engine.Cell
// 		}
// 	}{
// 		{
// 			input: [][]*engine.Cell{{cells[0], cells[1]}, {cells[2], cells[3]}, {cells[4], cells[5]}},
// 			point: api.NewPoint(0, 0),
// 			cell:  engine.NewCell(api.NewColor(api.ColorDefault, api.ColorBlack), 'a'),
// 			exp: struct {
// 				result bool
// 				cell   *engine.Cell
// 			}{true, engine.NewCell(api.NewColor(api.ColorDefault, api.ColorBlack), 'a')},
// 		},
// 		{
// 			input: [][]*engine.Cell{{cells[0], cells[1], cells[2]}, {cells[3], cells[4], cells[5]}},
// 			point: api.NewPoint(2, 1),
// 			cell:  engine.NewCell(api.NewColor(api.ColorDefault, api.ColorBlue), 'b'),
// 			exp: struct {
// 				result bool
// 				cell   *engine.Cell
// 			}{true, engine.NewCell(api.NewColor(api.ColorDefault, api.ColorBlue), 'b')},
// 		},
// 		{
// 			input: [][]*engine.Cell{{cells[0], cells[1]}, {cells[2], cells[3]}, {cells[4], cells[5]}},
// 			point: api.NewPoint(3, 3),
// 			cell:  engine.NewCell(api.NewColor(api.ColorDefault, api.ColorRed), 'c'),
// 			exp: struct {
// 				result bool
// 				cell   *engine.Cell
// 			}{false, nil},
// 		},
// 	}
// 	for i, c := range cases {
// 		canvas := engine.NewCanvas(api.NewSize(len(c.input[0]), len(c.input)))
// 		for x, rows := range c.input {
// 			copy(canvas.Rows[x].Cols, rows)
// 		}
// 		got := canvas.SetCellAt(c.point, c.cell)
// 		if c.exp.result != got {
// 			t.Errorf("[%d] SetCellAt Error exp:%t got:%t", i, c.exp.result, got)
// 			continue
// 		}
// 		if c.exp.result == false {
// 			continue
// 		}
// 		cell := canvas.GetCellAt(c.point)
// 		if !c.exp.cell.IsEqual(cell) {
// 			t.Errorf("[%d] SetCellAt Cell exp:%s got:%s", i, c.exp.cell.ToString(), cell.ToString())
// 		}
// 	}
// }

// func TestCanvasSetColorAt(t *testing.T) {
// 	createCells()
// 	cases := []struct {
// 		input [][]*engine.Cell
// 		point *api.Point
// 		color *api.Color
// 		exp   struct {
// 			result bool
// 			color  *api.Color
// 		}
// 	}{
// 		{
// 			input: [][]*engine.Cell{{cells[0], cells[1]}, {cells[2], cells[3]}, {cells[4], cells[5]}},
// 			point: api.NewPoint(0, 0),
// 			color: api.NewColor(api.ColorDefault, api.ColorBlack),
// 			exp: struct {
// 				result bool
// 				color  *api.Color
// 			}{true, api.NewColor(api.ColorDefault, api.ColorBlack)},
// 		},
// 		{
// 			input: [][]*engine.Cell{{cells[0], cells[1], cells[2]}, {cells[3], cells[4], cells[5]}},
// 			point: api.NewPoint(2, 1),
// 			color: api.NewColor(api.ColorDefault, api.ColorBlue),
// 			exp: struct {
// 				result bool
// 				color  *api.Color
// 			}{true, api.NewColor(api.ColorDefault, api.ColorBlue)},
// 		},
// 		{
// 			input: [][]*engine.Cell{{cells[0], cells[1]}, {cells[2], cells[3]}, {cells[4], cells[5]}},
// 			point: api.NewPoint(3, 3),
// 			color: api.NewColor(api.ColorDefault, api.ColorRed),
// 			exp: struct {
// 				result bool
// 				color  *api.Color
// 			}{false, nil},
// 		},
// 	}
// 	for i, c := range cases {
// 		canvas := engine.NewCanvas(api.NewSize(len(c.input[0]), len(c.input)))
// 		for x, rows := range c.input {
// 			copy(canvas.Rows[x].Cols, rows)
// 		}
// 		got := canvas.SetColorAt(c.point, c.color)
// 		if c.exp.result != got {
// 			t.Errorf("[%d] SetColorAt Error exp:%t got:%t", i, c.exp.result, got)
// 			continue
// 		}
// 		if c.exp.result == false {
// 			continue
// 		}
// 		color := canvas.GetColorAt(c.point)
// 		if !c.exp.color.IsEqual(color) {
// 			t.Errorf("[%d] SetColorAt Color exp:%s got:%s", i, c.exp.color.ToString(), color.ToString())
// 		}
// 	}
// }

func TestCanvasSetRuneAt(t *testing.T) {
	createCells()
	cases := []struct {
		input [][]engine.ICell
		point *api.Point
		ch    rune
		exp   struct {
			result bool
			ch     rune
		}
	}{
		{
			input: [][]engine.ICell{{cells[0], cells[1]}, {cells[2], cells[3]}, {cells[4], cells[5]}},
			point: api.NewPoint(0, 0),
			ch:    'a',
			exp: struct {
				result bool
				ch     rune
			}{true, 'a'},
		},
		{
			input: [][]engine.ICell{{cells[0], cells[1], cells[2]}, {cells[3], cells[4], cells[5]}},
			point: api.NewPoint(2, 1),
			ch:    'b',
			exp: struct {
				result bool
				ch     rune
			}{true, 'b'},
		},
		{
			input: [][]engine.ICell{{cells[0], cells[1]}, {cells[2], cells[3]}, {cells[4], cells[5]}},
			point: api.NewPoint(3, 3),
			ch:    'c',
			exp: struct {
				result bool
				ch     rune
			}{false, 0},
		},
	}
	for i, c := range cases {
		canvas := engine.NewCanvas(api.NewSize(len(c.input[0]), len(c.input)))
		for x, rows := range c.input {
			copy(canvas.Rows[x].Cols, rows)
		}
		got := canvas.SetRuneAt(c.point, c.ch)
		if c.exp.result != got {
			t.Errorf("[%d] SetRuneAt Error exp:%t got:%t", i, c.exp.result, got)
			continue
		}
		if c.exp.result == false {
			continue
		}
		ch := canvas.GetRuneAt(c.point)
		if c.exp.ch != ch {
			t.Errorf("[%d] SetRuneAt Rune exp:%c got:%c", i, c.exp.ch, ch)
		}
	}
}

func TestCanvasGetRect(t *testing.T) {
	cases := []struct {
		input *api.Size
		exp   *api.Rect
	}{
		{
			input: api.NewSize(1, 10),
			exp:   api.NewRect(api.NewPoint(0, 0), api.NewSize(1, 10)),
		},
		{
			input: api.NewSize(5, 1),
			exp:   api.NewRect(api.NewPoint(0, 0), api.NewSize(5, 1)),
		},
		{
			input: api.NewSize(3, 3),
			exp:   api.NewRect(api.NewPoint(0, 0), api.NewSize(3, 3)),
		},
	}
	for i, c := range cases {
		canvas := engine.NewCanvas(api.NewSize(c.input.W, c.input.H))
		got := canvas.GetRect()
		if !c.exp.IsEqual(got) {
			t.Errorf("[%d] GetRect Error exp:%s got:%s", i, c.exp.ToString(), got.ToString())
		}
	}
}

func TestCanvasIterator(t *testing.T) {
	createCells()
	size := api.NewSize(2, 3)
	canvas := engine.NewCanvas(size)
	index := 0
	for row := 0; row < size.H; row++ {
		for col := 0; col < size.W; col++ {
			canvas.SetCellAt(api.NewPoint(col, row), cells[index])
			index++
		}
	}
	canvas.CreateIter()
	index = 0
	icol := 0
	irow := 0
	for canvas.IterHasNext() {
		col, row, cell := canvas.IterGetNext()
		if icol != col {
			t.Errorf("[1] CanvasIterator Error.Col exp:%d got:%d", icol, col)
		}
		if irow != row {
			t.Errorf("[1] CanvasIterator Error.Row exp:%d got:%d", irow, row)
		}
		if cell != cells[index] {
			t.Errorf("[1] CanvasIterator Error.Cell exp:%s got:%s", cells[index].ToString(), cell.ToString())
		}
		// Increase all counters.
		index++
		icol++
		if icol >= size.W {
			icol = 0
			irow++
		}
	}
}

func TestCanvasGetStylAt(t *testing.T) {
	createCells()
	cases := []struct {
		input [][]engine.ICell
		exp   []*tcell.Style
	}{
		{
			input: [][]engine.ICell{{cells[0], cells[1]}, {cells[2], cells[3]}, {cells[4], cells[5]}},
			exp:   []*tcell.Style{cells[0].GetStyle(), cells[1].GetStyle(), cells[2].GetStyle(), cells[3].GetStyle(), cells[4].GetStyle(), cells[5].GetStyle()},
		},
	}
	for i, c := range cases {
		canvas := engine.NewCanvas(api.NewSize(len(c.input[0]), len(c.input)))
		for x, rows := range c.input {
			copy(canvas.Rows[x].Cols, rows)
		}
		canvas.CreateIter()
		index := 0
		for canvas.IterHasNext() {
			col, row, _ := canvas.IterGetNext()
			point := api.NewPoint(col, row)
			got := canvas.GetStyleAt(point)
			if !tools.IsEqualStyle(got, c.exp[index]) {
				t.Errorf("[%d] GetStyleAt Error exp:%+v got:%+v", i, c.exp[index], got)
			}
			index++
		}
	}
}

func TestCanvasSetStyleAt(t *testing.T) {
	createCells()
	cases := []struct {
		input [][]engine.ICell
		cell  *tcell.Style
		exp   *tcell.Style
	}{
		{
			input: [][]engine.ICell{{cells[0], cells[1]}, {cells[2], cells[3]}},
			cell:  cells[5].GetStyle(),
			exp:   cells[5].GetStyle(),
		},
	}
	for i, c := range cases {
		canvas := engine.NewCanvas(api.NewSize(len(c.input[0]), len(c.input)))
		for x, rows := range c.input {
			copy(canvas.Rows[x].Cols, rows)
		}
		canvas.CreateIter()
		for canvas.IterHasNext() {
			col, row, _ := canvas.IterGetNext()
			point := api.NewPoint(col, row)
			canvas.SetStyleAt(point, c.cell)
		}
		for canvas.CreateIter(); canvas.IterHasNext(); {
			_, _, cell := canvas.IterGetNext()
			if !tools.IsEqualStyle(c.exp, cell.GetStyle()) {
				t.Errorf("[%d] SetStyleAt Error exp:%+v got:%+v", i, c.exp, cell.GetStyle())
			}
		}
	}
}

func TestCanvasFillWithCell(t *testing.T) {
	createCells()
	cases := []struct {
		input [][]engine.ICell
		cell  *engine.Cell
		exp   *engine.Cell
	}{
		{
			input: [][]engine.ICell{{cells[0], cells[1]}, {cells[2], cells[3]}},
			cell:  cells[5],
			exp:   cells[5],
		},
	}
	for i, c := range cases {
		canvas := engine.NewCanvas(api.NewSize(len(c.input[0]), len(c.input)))
		for x, rows := range c.input {
			copy(canvas.Rows[x].Cols, rows)
		}
		canvas.FillWithCell(c.cell)
		canvas.CreateIter()
		for canvas.IterHasNext() {
			_, _, cell := canvas.IterGetNext()
			if !c.exp.IsEqual(cell) {
				t.Errorf("[%d] FillWithCell Error.Cell exp:%s got:%s", i, c.exp.ToString(), cell.ToString())
			}
		}
	}
}

// func TestCanvasFillWithCell(t *testing.T) {
// 	createCells()
// 	cases := []struct {
// 		input [][]*engine.Cell
// 		cell  *engine.Cell
// 		exp   *engine.Cell
// 	}{
// 		{
// 			input: [][]*engine.Cell{{cells[0], cells[1]}, {cells[2], cells[3]}, {cells[4], cells[5]}},
// 			cell:  engine.NewCell(api.NewColor(api.ColorDefault, api.ColorBlack), 'a'),
// 			exp:   engine.NewCell(api.NewColor(api.ColorDefault, api.ColorBlack), 'a'),
// 		},
// 		{
// 			input: [][]*engine.Cell{{cells[0], cells[1], cells[2]}, {cells[3], cells[4], cells[5]}},
// 			cell:  engine.NewCell(api.NewColor(api.ColorDefault, api.ColorBlue), 'b'),
// 			exp:   engine.NewCell(api.NewColor(api.ColorDefault, api.ColorBlue), 'b'),
// 		},
// 	}
// 	for i, c := range cases {
// 		canvas := engine.NewCanvas(api.NewSize(len(c.input[0]), len(c.input)))
// 		for x, rows := range c.input {
// 			copy(canvas.Rows[x].Cols, rows)
// 		}
// 		canvas.FillWithCell(c.cell)
// 		for x, rows := range canvas.Rows {
// 			for y, col := range rows.Cols {
// 				if !col.IsEqual(c.exp) {
// 					t.Errorf("[%d] FillWithCell (%d,%d) Error exp:%s got:%s",
// 						i, x, y, c.exp.ToString(), col.ToString())
// 				}
// 			}
// 		}
// 	}
// }

// func TestCanvasToString(t *testing.T) {
// 	cases := []struct {
// 		input struct {
// 			str   string
// 			color *api.Color
// 		}
// 		exp string
// 	}{
// 		{
// 			input: struct {
// 				str   string
// 				color *api.Color
// 			}{
// 				str:   "Hello",
// 				color: api.NewColor(api.ColorBlue, api.ColorRed),
// 			},
// 			exp: "[H][blue:red]\n[e][blue:red]\n[l][blue:red]\n[l][blue:red]\n[o][blue:red]\n",
// 		},
// 		{
// 			input: struct {
// 				str   string
// 				color *api.Color
// 			}{
// 				str:   "Hello\nyou",
// 				color: api.NewColor(api.ColorCyan, api.ColorYellow),
// 			},
// 			exp: "[H][cyan:yellow]\n[e][cyan:yellow]\n[l][cyan:yellow]\n[l][cyan:yellow]\n[o][cyan:yellow]\n[y][cyan:yellow]\n[o][cyan:yellow]\n[u][cyan:yellow]\n",
// 		},
// 		{
// 			input: struct {
// 				str   string
// 				color *api.Color
// 			}{
// 				str:   "Have\nyou\nman",
// 				color: nil,
// 			},
// 			exp: "[H][black:white]\n[a][black:white]\n[v][black:white]\n[e][black:white]\n[y][black:white]\n[o][black:white]\n[u][black:white]\n[m][black:white]\n[a][black:white]\n[n][black:white]\n",
// 		},
// 	}
// 	for i, c := range cases {
// 		canvas := engine.NewCanvasFromString(c.input.str, c.input.color)
// 		got := canvas.ToString()
// 		if c.exp != got {
// 			t.Errorf("[%d] ToString Error exp:%s got:%s", i, c.exp, got)
// 		}
// 	}
// }

// func TestCanvasSaveToDict(t *testing.T) {
// 	createCells()
// 	cases := []struct {
// 		input [][]*engine.Cell
// 		exp   map[string]any
// 	}{
// 		{
// 			input: [][]*engine.Cell{{cells[0]}, {cells[1]}, {cells[2]}},
// 			exp: map[string]any{
// 				"rows": []map[string]any{
// 					{
// 						"cols": []map[string]any{
// 							{
// 								"color": map[string]any{
// 									"fg": tcell.ColorBlack,
// 									"bg": tcell.ColorWhite,
// 								},
// 								"rune": '0',
// 							},
// 						},
// 					},
// 					{
// 						"cols": []map[string]any{
// 							{
// 								"color": map[string]any{
// 									"fg": tcell.ColorBlue,
// 									"bg": tcell.ColorRed,
// 								},
// 								"rune": '1',
// 							},
// 						},
// 					},
// 					{
// 						"cols": []map[string]any{
// 							{
// 								"color": map[string]any{
// 									"fg": tcell.ColorLightBlue,
// 									"bg": tcell.ColorDefault,
// 								},
// 								"rune": '2',
// 							},
// 						},
// 					},
// 				},
// 			},
// 		},
// 	}
// 	for i, c := range cases {
// 		canvas := engine.NewCanvas(api.NewSize(len(c.input[0]), len(c.input)))
// 		for x, rows := range c.input {
// 			copy(canvas.Rows[x].Cols, rows)
// 		}
// 		got := canvas.SaveToDict()
// 		if len(c.exp) != len(got) {
// 			t.Errorf("[%d] SaveToDict Length Error exp:%+v got:%+v", i, c.exp, got)
// 		}
// 		if len(c.exp["rows"].([]map[string]any)) != len(got["rows"].([]map[string]any)) {
// 			t.Errorf("[%d] SaveToDict Rows-Length Error exp:%d got:%d",
// 				i, len(c.exp["rows"].([]map[string]any)), len(got["rows"].([]map[string]any)))
// 		}
// 		for r, rentry := range c.exp["rows"].([]map[string]any) {
// 			gotRow := got["rows"].([]map[string]any)[r]
// 			if len(rentry) != len(gotRow) {
// 				t.Errorf("[%d] SaveToDict Rows-Cols-Length (%d) Error exp:%d  got:%d", i, r, len(rentry), len(gotRow))
// 			}
// 			for x, entry := range rentry["cols"].([]map[string]any) {
// 				var expRune rune = entry["rune"].(rune)
// 				var gotRune rune = gotRow["cols"].([]map[string]any)[x]["rune"].(rune)
// 				if expRune != gotRune {
// 					t.Errorf("[%d] SaveToDict Rows-Cols:Rune (%d,%d) Error exp:%c got:%c", i, r, x, expRune, gotRune)
// 				}
// 				var expColorFg api.Attr = entry["color"].(map[string]any)["fg"].(api.Attr)
// 				var gotColorFg api.Attr = gotRow["cols"].([]map[string]any)[x]["color"].(map[string]any)["fg"].(api.Attr)
// 				if expColorFg != gotColorFg {
// 					t.Errorf("[%d] SaveToDict Rows-Cols:Color:Fg (%d,%d) Error exp:%d got:%d", i, r, x, expColorFg, gotColorFg)
// 				}
// 				var expColorBg api.Attr = entry["color"].(map[string]any)["bg"].(api.Attr)
// 				var gotColorBg api.Attr = gotRow["cols"].([]map[string]any)[x]["color"].(map[string]any)["bg"].(api.Attr)
// 				if expColorBg != gotColorBg {
// 					t.Errorf("[%d] SaveToDict Rows-Cols:Color:Bg (%d,%d) Error exp:%d got:%d", i, r, x, expColorBg, gotColorBg)
// 				}
// 			}
// 		}
// 	}
// }
