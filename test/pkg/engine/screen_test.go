package engine_test

import (
	"testing"

	"github.com/jrecuero/thengine/pkg/api"
	"github.com/jrecuero/thengine/pkg/engine"
)

func TestScreen(t *testing.T) {
	cases := []struct {
		input *api.Size
		exp   *api.Size
	}{
		{
			input: api.NewSize(1, 1),
			exp:   api.NewSize(1, 1),
		},
		{
			input: api.NewSize(10, 5),
			exp:   api.NewSize(10, 5),
		},
	}
	for i, c := range cases {
		got := engine.NewScreen(c.input)
		if got == nil {
			t.Errorf("[%d] NewScreen Error exp:*Screen got:nil", i)
			continue
		}
		if got.Canvas == nil {
			t.Errorf("[%d] NewScreen Canvas exp:*Canvas got:nil", i)
		}
		if c.exp.W != got.Canvas.Width() {
			t.Errorf("[%d] NewScreen Canvas.Width exp:%d got:%d", i, c.exp.W, got.Canvas.Width())
		}
		if c.exp.H != got.Canvas.Height() {
			t.Errorf("[%d] NewScreen Canvas.Height exp:%d got:%d", i, c.exp.H, got.Canvas.Height())
		}
	}
}

func TestScreenGetRect(t *testing.T) {
	cases := []struct {
		input *api.Size
		exp   struct {
			origin *api.Point
			size   *api.Size
		}
	}{
		{
			input: api.NewSize(1, 1),
			exp: struct {
				origin *api.Point
				size   *api.Size
			}{
				origin: api.NewPoint(0, 0),
				size:   api.NewSize(1, 1),
			},
		},
		{
			input: api.NewSize(10, 5),
			exp: struct {
				origin *api.Point
				size   *api.Size
			}{
				origin: api.NewPoint(0, 0),
				size:   api.NewSize(10, 5),
			},
		},
	}
	for i, c := range cases {
		screen := engine.NewScreen(c.input)
		got := screen.GetRect()
		if got == nil {
			t.Errorf("[%d] GetRect Error exp:*Rect got:nil", i)
			continue
		}
		if !c.exp.origin.IsEqual(got.Origin) {
			t.Errorf("[%d] GetRect Error.Origin exp:%s got:%s", i, c.exp.origin.ToString(), got.Origin.ToString())
		}
		if !c.exp.size.IsEqual(got.Size) {
			t.Errorf("[%d] GetRect Error.Size exp:%s, got:%s", i, c.exp.size.ToString(), got.Size.ToString())
		}
	}
}

func TestScreenSize(t *testing.T) {
	cases := []struct {
		input *api.Size
		exp   *api.Size
	}{
		{
			input: api.NewSize(1, 1),
			exp:   api.NewSize(1, 1),
		},
		{
			input: api.NewSize(10, 5),
			exp:   api.NewSize(10, 5),
		},
	}
	for i, c := range cases {
		screen := engine.NewScreen(c.input)
		got := screen.Size()
		if got == nil {
			t.Errorf("[%d] Size Error exp:*Rect got:nil", i)
			continue
		}
		if !c.exp.IsEqual(got) {
			t.Errorf("[%d] Size Error exp:%s got:%s", i, c.exp.ToString(), got.ToString())
		}
	}
}

func TestScreenRenderCellAt(t *testing.T) {
	createCells()
	cases := []struct {
		input struct {
			size  *api.Size
			point *api.Point
			cell  *engine.Cell
		}
		exp struct {
			ok   bool // check if the cell was properly rendered
			cell *engine.Cell
		}
	}{
		{
			input: struct {
				size  *api.Size
				point *api.Point
				cell  *engine.Cell
			}{
				size:  api.NewSize(2, 2),
				point: api.NewPoint(1, 1),
				cell:  cells[0],
			},
			exp: struct {
				ok   bool
				cell *engine.Cell
			}{
				ok:   true,
				cell: cells[0],
			},
		},
		{
			input: struct {
				size  *api.Size
				point *api.Point
				cell  *engine.Cell
			}{
				size:  api.NewSize(2, 2),
				point: api.NewPoint(3, 1),
				cell:  cells[1],
			},
			exp: struct {
				ok   bool
				cell *engine.Cell
			}{
				ok:   false,
				cell: nil,
			},
		},
		{
			input: struct {
				size  *api.Size
				point *api.Point
				cell  *engine.Cell
			}{
				size:  api.NewSize(2, 2),
				point: api.NewPoint(1, 1),
				cell:  cells[1],
			},
			exp: struct {
				ok   bool
				cell *engine.Cell
			}{
				ok:   true,
				cell: cells[1],
			},
		},
	}
	for i, c := range cases {
		screen := engine.NewScreen(c.input.size)
		screen.Canvas.FillWithCell(cells[1])
		got := screen.RenderCellAt(c.input.point, c.input.cell)
		if c.exp.ok != got {
			t.Errorf("[%d] RenderCellAt Error exp:%t got:%t", i, c.exp.ok, got)
			continue
		}
		if c.exp.ok == false {
			continue
		}
		gotCell := screen.Canvas.GetCellAt(c.input.point)
		if !c.exp.cell.IsEqual(gotCell) {
			t.Errorf("[%d] RenderCellAt Cell Error exp:%s got:%s", i, c.exp.cell.ToString(), gotCell.ToString())
		}
		if gotEqual := screen.OldCanvas.IsEqual(screen.Canvas); gotEqual {
			t.Errorf("[%d] RenderCellAt OldCanvas Error exp:%t got:%t", i, false, gotEqual)
		}
	}
}
func TestScreenDraw(t *testing.T) {
	createCells()
	cases := []struct {
		input struct {
			size  *api.Size
			point *api.Point
			cell  *engine.Cell
		}
		exp struct {
			ok     bool // check if the cell was properly rendered
			canvas bool // check if the oldcanvas and canvas are equal
			cell   *engine.Cell
		}
	}{
		{
			input: struct {
				size  *api.Size
				point *api.Point
				cell  *engine.Cell
			}{
				size:  api.NewSize(2, 2),
				point: api.NewPoint(1, 1),
				cell:  cells[0],
			},
			exp: struct {
				ok     bool
				canvas bool
				cell   *engine.Cell
			}{
				ok:     true,
				canvas: false,
				cell:   cells[0],
			},
		},
		{
			input: struct {
				size  *api.Size
				point *api.Point
				cell  *engine.Cell
			}{
				size:  api.NewSize(2, 2),
				point: api.NewPoint(3, 1),
				cell:  cells[1],
			},
			exp: struct {
				ok     bool
				canvas bool
				cell   *engine.Cell
			}{
				ok:     false,
				canvas: true,
				cell:   nil,
			},
		},
		{
			input: struct {
				size  *api.Size
				point *api.Point
				cell  *engine.Cell
			}{
				size:  api.NewSize(2, 2),
				point: api.NewPoint(1, 1),
				cell:  cells[1],
			},
			exp: struct {
				ok     bool
				canvas bool
				cell   *engine.Cell
			}{
				ok:     true,
				canvas: true,
				cell:   cells[1],
			},
		},
	}
	for i, c := range cases {
		screen := engine.NewScreen(c.input.size)
		screen.DryRun = true // run in test mode.
		screen.Canvas.FillWithCell(cells[1])
		screen.Draw(true)
		got := screen.RenderCellAt(c.input.point, c.input.cell)
		if c.exp.ok != got {
			t.Errorf("[%d] RenderCellAt Error exp:%t got:%t", i, c.exp.ok, got)
			continue
		}
		if c.exp.ok == false {
			continue
		}
		gotCell := screen.Canvas.GetCellAt(c.input.point)
		if !c.exp.cell.IsEqual(gotCell) {
			t.Errorf("[%d] RenderCellAt Cell Error exp:%s got:%s", i, c.exp.cell.ToString(), gotCell.ToString())
		}
		gotEqual := screen.OldCanvas.IsEqual(screen.Canvas)
		if c.exp.canvas != gotEqual {
			t.Errorf("[%d] RenderCellAt OldCanvas Error exp:%t got:%t", i, c.exp.canvas, gotEqual)
		}
	}
}
