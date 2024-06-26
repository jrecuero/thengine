// sprite.go contains all required feature to implement an sprite. Sprite does
// not have any canvas, it has a list of SpriteCells which contains the
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
// SpriteCell
//
// -----------------------------------------------------------------------------

// SpriteCell struct defines the basic structure of an Sprite with a position
// and a cell.
type SpriteCell struct {
	position *api.Point
	cell     *engine.Cell
}

// NewSpriteCell function creates a new SpriteCell instance.
func NewSpriteCell(position *api.Point, cell *engine.Cell) *SpriteCell {
	return &SpriteCell{
		position: position,
		cell:     cell,
	}
}

// -----------------------------------------------------------------------------
// SpriteCell public methods
// -----------------------------------------------------------------------------

// GetCell method returns the cell from the SpriteCell instance.
func (s *SpriteCell) GetCell() *engine.Cell {
	return s.cell
}

// GetPosition method returs the position from the SpriteCell instance.
func (s *SpriteCell) GetPosition() *api.Point {
	return s.position
}

// SetCell method sets the cell in a SpriteCell instance.
func (s *SpriteCell) SetCell(cell *engine.Cell) {
	s.cell = cell
}

// SetPosition method sets the position in a SpriteCell instance.
func (s *SpriteCell) SetPosition(position *api.Point) {
	s.position = position
}

// -----------------------------------------------------------------------------
//
// Sprite
//
// -----------------------------------------------------------------------------

// Sprite struture defines an Sprite widget which is represented by a list of
// SpriteCells drawn in the screeen.
type Sprite struct {
	*Widget
	spriteCells []*SpriteCell
}

// NewSprite function creates a new Sprite instance.
func NewSprite(name string, position *api.Point, spriteCells []*SpriteCell) *Sprite {
	sprite := &Sprite{
		Widget:      NewWidget(name, position, nil, nil),
		spriteCells: spriteCells,
	}
	sprite.SetSize(nil)
	sprite.SetCanvas(nil)
	return sprite
}

// -----------------------------------------------------------------------------
// Sprite public methods
// -----------------------------------------------------------------------------

func (s *Sprite) AddSpriteCellAt(atIndex int, spriteCell *SpriteCell) {
	spriteCells := make([]*SpriteCell, len(s.spriteCells)+1)
	// if the atIndex is equal to AtTheEnd(-1), add the sprite cell at the end.
	if atIndex == AtTheEnd {
		atIndex = len(s.spriteCells)
	}
	for i, cell := range s.spriteCells {
		index := i
		if i >= atIndex {
			index++
		}
		spriteCells[index] = cell
	}
	spriteCells[atIndex] = spriteCell
	s.spriteCells = spriteCells
}

func (s *Sprite) Draw(scene engine.IScene) {
	defer s.Entity.Draw(scene)
	if s.IsVisible() {
		for _, spriteCell := range s.spriteCells {
			position := api.ClonePoint(s.GetPosition())
			position.Add(spriteCell.GetPosition())
			scene.GetCamera().RenderCellAt(position, spriteCell.GetCell())
		}
	}
}

func (s *Sprite) GetCollider() *engine.Collider {
	points := []*api.Point{}
	for _, spriteCell := range s.spriteCells {
		position := api.ClonePoint(s.GetPosition())
		position.Add(spriteCell.GetPosition())
		points = append(points, position)
	}
	return engine.NewCollider(nil, points)
}

func (s *Sprite) GetSpriteCells() []*SpriteCell {
	return s.spriteCells
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
	for _, spriteCell := range s.GetSpriteCells() {
		pos := api.ClonePoint(spriteCell.position)
		if origin != nil {
			pos.Subtract(origin)
		}
		cell := spriteCell.cell
		ch := cell.Rune
		fg, bg, attrs := cell.Style.Decompose()
		sprite := map[string]any{
			"position": []int{pos.X, pos.Y},
			"size":     []int{1, 1},
			"style":    []string{fg.String(), bg.String(), strconv.Itoa(int(attrs))},
			"ch":       string(ch),
		}
		sprites = append(sprites, sprite)
		content["sprites"] = sprites
		tools.Logger.WithField("module", "sprite").
			WithField("method", "MarshalMap").
			Debugf("sprite %+#v", sprite)
	}
	return content, nil
}

// MarshalCode method is the custom marshal method to generate pseudocode for
// the the instance.
func (s *Sprite) MarshalCode(origin *api.Point) (string, error) {
	result := ""
	result += fmt.Sprintf("// sprite: Sprite:%s\n", s.GetName())
	result += fmt.Sprintf("sprite := NewSprite(%s, api.Point(0, 0), nil, nil)\n", s.GetName())
	for _, spriteCell := range s.GetSpriteCells() {
		pos := api.ClonePoint(spriteCell.position)
		if origin != nil {
			pos.Subtract(origin)
		}
		fg, bg, attrs := spriteCell.GetCell().Style.Decompose()
		result += fmt.Sprintf("style := tcell.StyleDefault.Foreground(tcell.GetColor(%s)).Background(tcell.GetColor(%s)).Attributes(tcell.AttrMask(%d))\n", fg, bg, attrs)
		result += fmt.Sprintf("cell := engine.NewCell(&style, %d)\n", spriteCell.GetCell().Rune)
		result += fmt.Sprintf("pos := engine.NewPoint(%d, %d)\n", pos.X, pos.Y)
		result += fmt.Sprintf("spriteCell := widgets.NewSpriteCell(pos, cell)\n")
		result += fmt.Sprintf("sprite.AddSpriteCellAt(AtTheEnd, spriteCell)\n")
		result += fmt.Sprintf("--\n")
	}
	result += fmt.Sprintf("\n")

	return result, nil
}

func (s *Sprite) RemoveSpriteCellAt(atIndex int) *SpriteCell {
	if atIndex == AtTheEnd {
		atIndex = len(s.spriteCells) - 1
	}
	if (atIndex < 0) || (atIndex >= len(s.spriteCells)) {
		return nil
	}
	spriteCell := s.spriteCells[atIndex]
	s.spriteCells = append(s.spriteCells[:atIndex], s.spriteCells[atIndex+1:]...)
	return spriteCell
}

func (s *Sprite) SetSpriteCells(spriteCells []*SpriteCell) {
	s.spriteCells = spriteCells
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
			cell := engine.NewCell(style, ch)
			pos := api.NewPoint(x+posX, y+posY)
			spriteCell := NewSpriteCell(pos, cell)
			s.AddSpriteCellAt(AtTheEnd, spriteCell)
		}
	}
}

func (s *Sprite) StringToSpriteAtEnd(str string, style *tcell.Style, opts ...any) {
	var pos *api.Point
	if lenSpriteCells := len(s.spriteCells); lenSpriteCells != 0 {
		lastSpriteCell := s.spriteCells[lenSpriteCells-1]
		pos = lastSpriteCell.position
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
		spriteCell := NewSpriteCell(nil, nil)
		cell := engine.NewCell(nil, 0)
		if position, ok := sprite["position"].([]any); ok {
			pos := api.NewPoint(int(position[0].(float64)), int(position[1].(float64)))
			//if origin != nil {
			//    pos.Add(origin)
			//}
			spriteCell.SetPosition(pos)
		}
		if style, ok := sprite["style"].([]any); ok {
			tcellStyle := tcell.StyleDefault.
				Foreground(tcell.GetColor(style[0].(string))).
				Background(tcell.GetColor(style[1].(string)))
			cell.Style = &tcellStyle
		}
		if ch, ok := sprite["ch"].(string); ok {
			cell.Rune = rune(ch[0])
		}
		spriteCell.SetCell(cell)
		tools.Logger.WithField("module", "sprite").
			WithField("method", "UnmarshalMap").
			Debugf("spriteccell %s %s %+#v", s.GetName(), spriteCell.GetPosition().ToString(), spriteCell.GetCell())
		s.AddSpriteCellAt(AtTheEnd, spriteCell)
	}
	return nil
}

var _ engine.IObject = (*Sprite)(nil)
var _ engine.IFocus = (*Sprite)(nil)
var _ engine.IEntity = (*Sprite)(nil)
