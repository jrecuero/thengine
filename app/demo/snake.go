package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/gdamore/tcell/v2"
	"github.com/jrecuero/thengine/pkg/api"
	"github.com/jrecuero/thengine/pkg/engine"
	"github.com/jrecuero/thengine/pkg/tools"
	"github.com/jrecuero/thengine/pkg/widgets"
)

const (
	SnakePointsTopic = "snake/points/topic"
)

type snake struct {
	*widgets.Sprite
	x         float64
	y         float64
	speed     float64
	direction string
	alive     bool
}

func NewSnake(position *api.Point, style *tcell.Style) *snake {
	cell := engine.NewCell(style, '#')
	snake := &snake{
		Sprite: widgets.NewSprite("snake/1", position,
			[]*widgets.SpriteCell{widgets.NewSpriteCell(api.NewPoint(5, 5), cell)}),
		speed:     10.0,
		direction: "none",
		alive:     true,
	}
	//snake.x = float64(snake.GetPosition().X)
	//snake.y = float64(snake.GetPosition().Y)
	snake.x = 5.0
	snake.y = 5.0
	snake.SetFocusType(engine.SingleFocus)
	snake.SetFocusEnable(true)
	return snake
}

func (s *snake) Move(args ...any) {
	direction := args[0].(string)
	// Don't allow to reverse vertical direction.
	if (direction == "up") && (s.direction == "down") ||
		(direction == "down") && (s.direction == "up") {
		return
	}
	// Don't allow to reverse horizontal direction.
	if (direction == "left") && (s.direction == "right") ||
		(direction == "right") && (s.direction == "left") {
		return
	}
	s.direction = args[0].(string)
}

func (s *snake) Update(event tcell.Event, scene engine.IScene) {
	if !s.HasFocus() {
		return
	}
	if !s.alive {
		return
	}
	actions := []*widgets.KeyboardAction{
		{Key: tcell.KeyUp, Callback: s.Move, Args: []any{"up"}},
		{Key: tcell.KeyDown, Callback: s.Move, Args: []any{"down"}},
		{Key: tcell.KeyRight, Callback: s.Move, Args: []any{"right"}},
		{Key: tcell.KeyLeft, Callback: s.Move, Args: []any{"left"}},
	}
	s.HandleKeyboardForActions(event, actions)
	var vx, vy int
	if s.direction != "none" {
		switch s.direction {
		case "up":
			vx = 0
			vy = -1
		case "down":
			vx = 0
			vy = 1
		case "left":
			vx = -1
			vy = 0
		case "right":
			vx = 1
			vy = 0
		}
		s.x += float64(vx) / s.speed
		s.y += float64(vy) / s.speed
		spriteCell := s.GetSpriteCells()[0]
		intX := int(s.x)
		intY := int(s.y)
		if (intX != spriteCell.GetPosition().X) || (intY != spriteCell.GetPosition().Y) {
			if (intX >= 80) || (intX < 0) || (intY >= 20) || (intY < 0) {
				s.alive = false
				tools.Logger.WithField("module", "snake").WithField("struct", "snake").WithField("function", "Update").Infof("snake is dead at %d,%d", intX, intY)
				return
			}
			// Remove the last entry and place first with position increased
			// based in the direction the snake is running.
			lastCell := s.RemoveSpriteCellAt(widgets.AtTheEnd)
			lastCell.SetPosition(api.NewPoint(intX, intY))
			s.AddSpriteCellAt(0, lastCell)
			tools.Logger.WithField("module", "snake").WithField("struct", "snake").WithField("function", "Update").Infof("rotate %s", lastCell.GetPosition().ToString())
		}
		// check collision with itself.
		spriteCells := s.GetSpriteCells()
		spriteCell = spriteCells[0]
		for i := 1; i < len(s.GetSpriteCells()); i++ {
			if spriteCell.GetPosition().IsEqual(spriteCells[i].GetPosition()) {
				s.alive = false
				tools.Logger.WithField("module", "snake").WithField("struct", "snake").WithField("function", "Update").Infof("snake is dead at %d", i)
				return
			}

		}

		collisions := scene.CheckCollisionWith(s)
		if len(collisions) != 0 {
			for _, ent := range collisions {
				spriteCell = s.GetSpriteCells()[0]
				scene.RemoveEntity(ent)
				// Update float64 position with new entry.
				s.x = float64(spriteCell.GetPosition().X + vx)
				s.y = float64(spriteCell.GetPosition().Y + vy)
				newSpriteCell := widgets.NewSpriteCell(api.NewPoint(int(s.x), int(s.y)), spriteCell.GetCell())
				s.AddSpriteCellAt(0, newSpriteCell)
				tools.Logger.WithField("module", "snake").WithField("struct", "snake").WithField("function", "Update").Infof("eat %s", newSpriteCell.GetPosition().ToString())
				message := &engine.Message{
					Topic:   SnakePointsTopic,
					Src:     s.GetName(),
					Dst:     "broadcast",
					Content: nil,
				}
				engine.GetMailbox().Publish(SnakePointsTopic, message)
			}
		}
	}
}

type TextPoints struct {
	*widgets.Text
	points int
}

func (t *TextPoints) Consume() {
	if message, _ := engine.GetMailbox().Consume(SnakePointsTopic, t.GetName()); message != nil {
		t.points += 10
		t.SetText(fmt.Sprintf("Points: %d", t.points))
	}
}

type SnakeGame struct {
	*engine.Engine
}

func NewSnakeGame() *SnakeGame {
	return &SnakeGame{
		Engine: engine.GetEngine(),
	}
}

type AppHandler struct {
	*engine.Entity
	Running bool
}

var (
	theAppHandler *AppHandler = &AppHandler{
		Entity:  engine.NewNamedEntity("app/handler/1"),
		Running: false,
	}
)

func demoSnake(dryRun bool) {
	tools.Logger.WithField("module", "main").WithField("dry-mode", dryRun).Infof("ThEngine demo-snake")
	fmt.Println("ThEngine demo-snake")
	//appEngine := engine.GetEngine()
	appEngine := NewSnakeGame()
	appEngine.InitResources()
	camera := engine.NewCamera(api.NewPoint(0, 0), api.NewSize(90, 30))
	styleOne := tcell.StyleDefault.Foreground(tcell.ColorRed).Background(tcell.ColorWhite)
	styleTwo := tcell.StyleDefault.Foreground(tcell.ColorRed).Background(tcell.ColorBlack)
	styleThree := tcell.StyleDefault.Foreground(tcell.ColorBlue).Background(tcell.ColorBlack)
	//styleFour := tcell.StyleDefault.Foreground(tcell.ColorBlack).Background(tcell.ColorBlue)
	//styleFive := tcell.StyleDefault.Foreground(tcell.ColorBlack).Background(tcell.ColorRed)
	scene := engine.NewScene("scene/snake/1", camera)

	textPoints := &TextPoints{
		Text: widgets.NewText("points/1", api.NewPoint(0, 0), api.NewSize(20, 1), &styleThree, "Points: 0"),
	}
	scene.AddEntity(textPoints)
	engine.GetMailbox().CreateTopic(SnakePointsTopic)
	engine.GetMailbox().Subscribe(SnakePointsTopic, textPoints.GetName())

	box := engine.NewEntity("box/1", api.NewPoint(0, 1), api.NewSize(80, 20), &styleTwo)
	box.GetCanvas().WriteRectangleInCanvasAt(nil, nil, &styleTwo, engine.CanvasRectSingleLine)
	scene.AddEntity(box)

	foodTimer := widgets.NewTimer("timer/food/1", 5*time.Second, widgets.ForeverTimer)
	foodTimer.SetWidgetCallback(func(entity engine.IEntity, args ...any) bool {
		x := rand.Intn(78) + 1
		y := rand.Intn(18) + 2
		food := engine.NewEntity("food/1", api.NewPoint(x, y), api.NewSize(1, 1), &styleOne)
		food.GetCanvas().WriteStringInCanvas(".", &styleOne)
		food.SetSolid(true)
		scene.AddEntity(food)
		return true
	})
	scene.AddEntity(foodTimer)

	snake := NewSnake(api.NewPoint(0, 1), &styleTwo)
	scene.AddEntity(snake)

	engineScene, err := appEngine.CreateEngineScene()
	if err != nil {
		panic(err)
	}
	engineScene.AddEntity(theAppHandler)
	appEngine.GetSceneManager().SetSceneAsActive(engineScene)
	appEngine.GetSceneManager().AddScene(scene)
	appEngine.GetSceneManager().SetSceneAsActive(scene)
	appEngine.GetSceneManager().SetSceneAsVisible(scene)
	//for _, sc := range appEngine.GetSceneManager().GetAllVisibleScenes() {
	//    tools.Logger.WithField("module", "main").WithField("dry-mode", dryRun).Infof("visible scenes %s", sc.GetName())
	//}
	gameOverScene := engine.NewScene("scene/game-over/1", camera)
	appEngine.GetSceneManager().AddScene(gameOverScene)

	appEngine.Init()
	appEngine.Start()
	theAppHandler.Running = true
	appEngine.Run(60.0)
}
