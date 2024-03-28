package engine_test

import (
	"testing"

	"github.com/jrecuero/thengine/pkg/api"
	"github.com/jrecuero/thengine/pkg/engine"
)

func TestCamera(t *testing.T) {
	cases := []struct {
		input struct {
			origin *api.Point
			size   *api.Size
		}
		exp struct {
			origin *api.Point
			size   *api.Size
		}
	}{
		{
			input: struct {
				origin *api.Point
				size   *api.Size
			}{
				size: api.NewSize(1, 1),
			},
			exp: struct {
				origin *api.Point
				size   *api.Size
			}{
				origin: api.NewPoint(0, 0),
				size:   api.NewSize(1, 1),
			},
		},
		{
			input: struct {
				origin *api.Point
				size   *api.Size
			}{
				origin: api.NewPoint(10, 5),
				size:   api.NewSize(2, 3),
			},
			exp: struct {
				origin *api.Point
				size   *api.Size
			}{
				origin: api.NewPoint(10, 5),
				size:   api.NewSize(2, 3),
			},
		},
	}
	for i, c := range cases {
		got := engine.NewCamera(c.input.origin, c.input.size)
		if got == nil {
			t.Errorf("[%d] NewCamera Error exp:*Camera got:nil", i)
			continue
		}
		if c.exp.origin.X != got.GetOrigin().X {
			t.Errorf("[%d] NewCamera Canvas.X exp:%d got:%d", i, c.exp.origin.X, got.GetOrigin().X)
		}
		if c.exp.origin.Y != got.GetOrigin().Y {
			t.Errorf("[%d] NewCamera Canvas.Y exp:%d got:%d", i, c.exp.origin.Y, got.GetOrigin().Y)
		}
		if c.exp.size.W != got.GetSize().W {
			t.Errorf("[%d] NewCamera Canvas.Width exp:%d got:%d", i, c.exp.size.W, got.GetSize().W)
		}
		if c.exp.size.H != got.GetSize().H {
			t.Errorf("[%d] NewCamera Canvas.Height exp:%d got:%d", i, c.exp.size.H, got.GetSize().H)
		}
	}
}

func TestCameraGetRect(t *testing.T) {
	cases := []struct {
		input struct {
			origin *api.Point
			size   *api.Size
		}
		exp struct {
			origin *api.Point
			size   *api.Size
		}
	}{
		{
			input: struct {
				origin *api.Point
				size   *api.Size
			}{
				size: api.NewSize(1, 1),
			},
			exp: struct {
				origin *api.Point
				size   *api.Size
			}{
				origin: api.NewPoint(0, 0),
				size:   api.NewSize(1, 1),
			},
		},
		{
			input: struct {
				origin *api.Point
				size   *api.Size
			}{
				origin: api.NewPoint(10, 5),
				size:   api.NewSize(2, 3),
			},
			exp: struct {
				origin *api.Point
				size   *api.Size
			}{
				origin: api.NewPoint(10, 5),
				size:   api.NewSize(2, 3),
			},
		},
	}
	for i, c := range cases {
		camera := engine.NewCamera(c.input.origin, c.input.size)
		got := camera.GetSize()
		if got == nil {
			t.Errorf("[%d] GetRect Error exp:*Rect got:nil", i)
			continue
		}
		if !c.exp.size.IsEqual(got) {
			t.Errorf("[%d] GetRect Error.Size exp:%s, got:%s", i, c.exp.size.ToString(), got.ToString())
		}
	}
}

func TestCameraGetOrigin(t *testing.T) {
	cases := []struct {
		input struct {
			origin *api.Point
			size   *api.Size
		}
		exp *api.Point
	}{
		{
			input: struct {
				origin *api.Point
				size   *api.Size
			}{
				size: api.NewSize(1, 1),
			},
			exp: api.NewPoint(0, 0),
		},
		{
			input: struct {
				origin *api.Point
				size   *api.Size
			}{
				origin: api.NewPoint(10, 5),
				size:   api.NewSize(2, 3),
			},
			exp: api.NewPoint(10, 5),
		},
	}
	for i, c := range cases {
		camera := engine.NewCamera(c.input.origin, c.input.size)
		got := camera.GetOrigin()
		if got == nil {
			t.Errorf("[%d] GetOrigin Error exp:*Rect got:nil", i)
			continue
		}
		if c.exp.X != got.X {
			t.Errorf("[%d] GetOrigin Error.X exp:%d got:%d", i, c.exp.X, got.X)
		}
		if c.exp.Y != got.Y {
			t.Errorf("[%d] GetOrigin Error.Y exp:%d got:%d", i, c.exp.Y, got.Y)
		}
	}
}

// func TestCameraSize(t *testing.T) {
// 	cases := []struct {
// 		input *api.Size
// 		exp   *api.Size
// 	}{
// 		{
// 			input: api.NewSize(1, 1),
// 			exp:   api.NewSize(1, 1),
// 		},
// 		{
// 			input: api.NewSize(10, 5),
// 			exp:   api.NewSize(10, 5),
// 		},
// 	}
// 	for i, c := range cases {
// 		camera := engine.NewCamera(c.input)
// 		got := camera.Size()
// 		if got == nil {
// 			t.Errorf("[%d] Size Error exp:*Rect got:nil", i)
// 			continue
// 		}
// 		if !c.exp.IsEqual(got) {
// 			t.Errorf("[%d] Size Error exp:%s got:%s", i, c.exp.ToString(), got.ToString())
// 		}
// 	}
// }

// func TestCameraRenderCellAt(t *testing.T) {
// 	createCells()
// 	cases := []struct {
// 		input struct {
// 			size  *api.Size
// 			point *api.Point
// 			cell  *engine.Cell
// 		}
// 		exp struct {
// 			ok   bool // check if the cell was properly rendered
// 			cell *engine.Cell
// 		}
// 	}{
// 		{
// 			input: struct {
// 				size  *api.Size
// 				point *api.Point
// 				cell  *engine.Cell
// 			}{
// 				size:  api.NewSize(2, 2),
// 				point: api.NewPoint(1, 1),
// 				cell:  cells[0],
// 			},
// 			exp: struct {
// 				ok   bool
// 				cell *engine.Cell
// 			}{
// 				ok:   true,
// 				cell: cells[0],
// 			},
// 		},
// 		{
// 			input: struct {
// 				size  *api.Size
// 				point *api.Point
// 				cell  *engine.Cell
// 			}{
// 				size:  api.NewSize(2, 2),
// 				point: api.NewPoint(3, 1),
// 				cell:  cells[1],
// 			},
// 			exp: struct {
// 				ok   bool
// 				cell *engine.Cell
// 			}{
// 				ok:   false,
// 				cell: nil,
// 			},
// 		},
// 		{
// 			input: struct {
// 				size  *api.Size
// 				point *api.Point
// 				cell  *engine.Cell
// 			}{
// 				size:  api.NewSize(2, 2),
// 				point: api.NewPoint(1, 1),
// 				cell:  cells[1],
// 			},
// 			exp: struct {
// 				ok   bool
// 				cell *engine.Cell
// 			}{
// 				ok:   true,
// 				cell: cells[1],
// 			},
// 		},
// 	}
// 	for i, c := range cases {
// 		camera := engine.NewCamera(c.input.size)
// 		camera.Canvas.FillWithCell(cells[1])
// 		got := camera.RenderCellAt(c.input.point, c.input.cell)
// 		if c.exp.ok != got {
// 			t.Errorf("[%d] RenderCellAt Error exp:%t got:%t", i, c.exp.ok, got)
// 			continue
// 		}
// 		if c.exp.ok == false {
// 			continue
// 		}
// 		gotCell := camera.Canvas.GetCellAt(c.input.point)
// 		if !c.exp.cell.IsEqual(gotCell) {
// 			t.Errorf("[%d] RenderCellAt Cell Error exp:%s got:%s", i, c.exp.cell.ToString(), gotCell.ToString())
// 		}
// 		if gotEqual := camera.OldCanvas.IsEqual(camera.Canvas); gotEqual {
// 			t.Errorf("[%d] RenderCellAt OldCanvas Error exp:%t got:%t", i, false, gotEqual)
// 		}
// 	}
// }
// func TestCameraDraw(t *testing.T) {
// 	createCells()
// 	cases := []struct {
// 		input struct {
// 			size  *api.Size
// 			point *api.Point
// 			cell  *engine.Cell
// 		}
// 		exp struct {
// 			ok     bool // check if the cell was properly rendered
// 			canvas bool // check if the oldcanvas and canvas are equal
// 			cell   *engine.Cell
// 		}
// 	}{
// 		{
// 			input: struct {
// 				size  *api.Size
// 				point *api.Point
// 				cell  *engine.Cell
// 			}{
// 				size:  api.NewSize(2, 2),
// 				point: api.NewPoint(1, 1),
// 				cell:  cells[0],
// 			},
// 			exp: struct {
// 				ok     bool
// 				canvas bool
// 				cell   *engine.Cell
// 			}{
// 				ok:     true,
// 				canvas: false,
// 				cell:   cells[0],
// 			},
// 		},
// 		{
// 			input: struct {
// 				size  *api.Size
// 				point *api.Point
// 				cell  *engine.Cell
// 			}{
// 				size:  api.NewSize(2, 2),
// 				point: api.NewPoint(3, 1),
// 				cell:  cells[1],
// 			},
// 			exp: struct {
// 				ok     bool
// 				canvas bool
// 				cell   *engine.Cell
// 			}{
// 				ok:     false,
// 				canvas: true,
// 				cell:   nil,
// 			},
// 		},
// 		{
// 			input: struct {
// 				size  *api.Size
// 				point *api.Point
// 				cell  *engine.Cell
// 			}{
// 				size:  api.NewSize(2, 2),
// 				point: api.NewPoint(1, 1),
// 				cell:  cells[1],
// 			},
// 			exp: struct {
// 				ok     bool
// 				canvas bool
// 				cell   *engine.Cell
// 			}{
// 				ok:     true,
// 				canvas: true,
// 				cell:   cells[1],
// 			},
// 		},
// 	}
// 	for i, c := range cases {
// 		camera := engine.NewCamera(c.input.size)
// 		camera.DryRun = true // run in test mode.
// 		camera.Canvas.FillWithCell(cells[1])
// 		camera.Draw(true)
// 		got := camera.RenderCellAt(c.input.point, c.input.cell)
// 		if c.exp.ok != got {
// 			t.Errorf("[%d] RenderCellAt Error exp:%t got:%t", i, c.exp.ok, got)
// 			continue
// 		}
// 		if c.exp.ok == false {
// 			continue
// 		}
// 		gotCell := camera.Canvas.GetCellAt(c.input.point)
// 		if !c.exp.cell.IsEqual(gotCell) {
// 			t.Errorf("[%d] RenderCellAt Cell Error exp:%s got:%s", i, c.exp.cell.ToString(), gotCell.ToString())
// 		}
// 		gotEqual := camera.OldCanvas.IsEqual(camera.Canvas)
// 		if c.exp.canvas != gotEqual {
// 			t.Errorf("[%d] RenderCellAt OldCanvas Error exp:%t got:%t", i, c.exp.canvas, gotEqual)
// 		}
// 	}
// }
