// sprite.go contains all required feature to implement an sprite. Sprite does
// not have any canvas, it has a list of Cell instance which contains the
// position and the cell to be render in such position.
package widgets

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/gdamore/tcell/v2"
	"github.com/jrecuero/thengine/pkg/api"
	"github.com/jrecuero/thengine/pkg/engine"
	"github.com/jrecuero/thengine/pkg/tools"
)

const (
	AtTheEnd int = -1
)

// -----------------------------------------------------------------------------
//
// Sprite
//
// -----------------------------------------------------------------------------

// Sprite struture defines an Sprite widget which is represented by a list of
// Cell instances drawn in the screen.
type Sprite struct {
	*Widget
	cells engine.CellGroup
}

// NewSprite function creates a new Sprite instance.
func NewSprite(name string, position *api.Point, cells engine.CellGroup) *Sprite {
	sprite := &Sprite{
		Widget: NewWidget(name, position, nil, nil),
		cells:  cells,
	}
	sprite.SetSize(nil)
	sprite.SetCanvas(nil)
	return sprite
}

// -----------------------------------------------------------------------------
// Sprite public methods
// -----------------------------------------------------------------------------

func (s *Sprite) AddCellAt(atIndex int, cell engine.ICell) {
	cells := make(engine.CellGroup, len(s.cells)+1)
	// if the atIndex is equal to AtTheEnd(-1), add the sprite cell at the end.
	if atIndex == AtTheEnd {
		atIndex = len(s.cells)
	}
	for i, cell := range s.cells {
		index := i
		if i >= atIndex {
			index++
		}
		cells[index] = cell
	}
	cells[atIndex] = cell
	s.cells = cells
}

func (s *Sprite) Draw(scene engine.IScene) {
	defer s.Entity.Draw(scene)
	if s.IsVisible() {
		for _, cell := range s.cells {
			position := api.ClonePoint(s.GetPosition())
			position.Add(cell.GetPosition())
			scene.GetCamera().RenderCellAt(position, cell)
		}
	}
}

func (s *Sprite) GetCells() engine.CellGroup {
	return s.cells
}

func (s *Sprite) GetCollider() *engine.Collider {
	points := []*api.Point{}
	for _, cell := range s.cells {
		position := api.ClonePoint(s.GetPosition())
		position.Add(cell.GetPosition())
		points = append(points, position)
	}
	return engine.NewCollider(nil, points)
}

// MarshalMap method is the custom marshal method to generate a map[string]any
// from an instance.
func (s *Sprite) MarshalMap(origin *api.Point) (map[string]any, error) {
	tools.Logger.WithField("module", "sprite").
		WithField("method", "MarshalMap").
		Debugf("sprite %s origin: %s", s.GetName(), origin.ToString())

	content := map[string]any{
		"class":    "Sprite",
		"name":     s.GetName(),
		"position": []int{0, 0},
		"sprites":  []map[string]any{},
	}
	sprites := []map[string]any{}
	for _, cell := range s.GetCells() {
		pos := api.ClonePoint(cell.GetPosition())
		if origin != nil {
			pos.Subtract(origin)
		}
		ch := cell.GetRune()
		fg, bg, attrs := cell.GetStyle().Decompose()
		sprite := map[string]any{
			"position": []int{pos.X, pos.Y},
			"size":     []int{1, 1},
			"style":    []string{fg.String(), bg.String(), strconv.Itoa(int(attrs))},
			"ch":       string(ch),
		}
		sprites = append(sprites, sprite)
		content["sprites"] = sprites
		tools.Logger.WithField("module", "sprite").
			WithField("struct", "Sprite").
			WithField("method", "MarshalMap").
			Tracef("sprite %+#v", sprite)
	}
	return content, nil
}

// MarshalCode method is the custom marshal method to generate pseudocode for
// the the instance.
func (s *Sprite) MarshalCode(origin *api.Point) (string, error) {
	result := ""
	result += fmt.Sprintf("// sprite: Sprite:%s\n", s.GetName())
	result += fmt.Sprintf("sprite := NewSprite(%s, api.Point(0, 0), nil, nil)\n", s.GetName())
	for _, cell := range s.GetCells() {
		pos := api.ClonePoint(cell.GetPosition())
		if origin != nil {
			pos.Subtract(origin)
		}
		fg, bg, attrs := cell.GetStyle().Decompose()
		result += fmt.Sprintf("style := tcell.StyleDefault.Foreground(tcell.GetColor(%s)).Background(tcell.GetColor(%s)).Attributes(tcell.AttrMask(%d))\n", fg, bg, attrs)
		result += fmt.Sprintf("cell := engine.NewCell(&style, %d)\n", cell.GetRune())
		result += fmt.Sprintf("pos := engine.NewPoint(%d, %d)\n", pos.X, pos.Y)
		result += fmt.Sprintf("sprite.AddCellAt(AtTheEnd, cell)\n")
		result += fmt.Sprintf("--\n")
	}
	result += fmt.Sprintf("\n")

	return result, nil
}

func (s *Sprite) RemoveCellAt(atIndex int) engine.ICell {
	if atIndex == AtTheEnd {
		atIndex = len(s.cells) - 1
	}
	if (atIndex < 0) || (atIndex >= len(s.cells)) {
		return nil
	}
	cellpos := s.cells[atIndex]
	s.cells = append(s.cells[:atIndex], s.cells[atIndex+1:]...)
	return cellpos
}

func (s *Sprite) SetCells(cells engine.CellGroup) {
	s.cells = cells
}

func (s *Sprite) StringToSprite(str string, style *tcell.Style, opts ...any) {
	s.StringToSpriteAt(str, nil, style, opts...)
}

// StringToSpriteAt method writes the given string in the sprite character by
// character.
// opts:
//
//	[0] skip-spaces: true will skill any space. false will write any spaces.
func (s *Sprite) StringToSpriteAt(str string, pos *api.Point, style *tcell.Style, opts ...any) {
	skipSpaces := true
	if len(opts) != 0 {
		skipSpaces = opts[0].(bool)
	}
	posX, posY := 0, 0
	if pos != nil {
		posX, posY = pos.Get()
	}
	lines := strings.Split(str, "\n")
	for y, line := range lines {
		for x, ch := range line {
			if skipSpaces && ch == ' ' {
				continue
			}
			pos := api.NewPoint(x+posX, y+posY)
			cellPos := engine.NewCellAt(style, ch, pos)
			s.AddCellAt(AtTheEnd, cellPos)
		}
	}
}

func (s *Sprite) StringToSpriteAtEnd(str string, style *tcell.Style, opts ...any) {
	var pos *api.Point
	if lenSpriteCells := len(s.cells); lenSpriteCells != 0 {
		lastSpriteCell := s.cells[lenSpriteCells-1]
		pos = lastSpriteCell.GetPosition()
	}
	s.StringToSpriteAt(str, pos, style, opts...)
}

// UnmarshalMap method is the custom method to unmarshal a map[string]any data
// into an instance.
func (s *Sprite) UnmarshalMap(content map[string]any, origin *api.Point) error {
	if name, ok := content["name"].(string); ok {
		s.SetName(name)
	}
	if position, ok := content["position"].([]any); ok {
		pos := api.NewPoint(int(position[0].(float64)), int(position[1].(float64)))
		if origin != nil {
			pos.Add(origin)
		}
		s.SetPosition(pos)
	}
	for _, spr := range content["sprites"].([]any) {
		sprite := spr.(map[string]any)
		cell := engine.NewEmptyCell()
		if position, ok := sprite["position"].([]any); ok {
			pos := api.NewPoint(int(position[0].(float64)), int(position[1].(float64)))
			//if origin != nil {
			//    pos.Add(origin)
			//}
			cell.SetPosition(pos)
		}
		if style, ok := sprite["style"].([]any); ok {
			tcellStyle := tcell.StyleDefault.
				Foreground(tcell.GetColor(style[0].(string))).
				Background(tcell.GetColor(style[1].(string)))
			cell.SetStyle(&tcellStyle)
		}
		if ch, ok := sprite["ch"].(string); ok {
			cell.SetRune(rune(ch[0]))
		}
		tools.Logger.WithField("module", "sprite").
			WithField("struct", "Sprite").
			WithField("method", "UnmarshalMap").
			Tracef("cell %s %s", s.GetName(), cell.ToString())
		s.AddCellAt(AtTheEnd, cell)
	}
	return nil
}

var _ engine.IObject = (*Sprite)(nil)
var _ engine.IFocus = (*Sprite)(nil)
var _ engine.IEntity = (*Sprite)(nil)
