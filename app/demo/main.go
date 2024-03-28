package main

import (
	"fmt"

	"github.com/gdamore/tcell/v2"
	"github.com/jrecuero/thengine/pkg/api"
	"github.com/jrecuero/thengine/pkg/engine"
	"github.com/jrecuero/thengine/pkg/tools"
	"github.com/jrecuero/thengine/pkg/widgets"
)

type Player struct {
	*widgets.Widget
	origin *api.Point
}

func NewPlayer(name string, position *api.Point, size *api.Size, style *tcell.Style) *Player {
	player := &Player{
		Widget: widgets.NewWidget(name, position, size, style),
		origin: position,
	}
	player.SetFocusType(engine.SingleFocus)
	player.SetFocusEnable(true)
	return player

}

func (p *Player) Reset(args ...any) {
	p.SetPosition(p.origin)
}

func (p *Player) Move(args ...any) {
	steps := args[0].(int)
	direction := args[1].(string)
	tools.Logger.WithField("module", "main").
		WithField("struct", "Player").
		WithField("function", "Move").
		Debugf("stets %d direction %s", steps, direction)
	x, y := p.GetPosition().Get()
	if direction == "up" {
		p.SetPosition(api.NewPoint(x, y-1))
	} else if direction == "down" {
		p.SetPosition(api.NewPoint(x, y+1))
	} else if direction == "right" {
		p.SetPosition(api.NewPoint(x+1, y))
	} else if direction == "left" {
		p.SetPosition(api.NewPoint(x-1, y))
	}
}

func (p *Player) Update(event tcell.Event) {
	if !p.HasFocus() {
		return
	}
	actions := []*widgets.KeyboardAction{
		{
			Key:      tcell.KeyUp,
			Callback: p.Move,
			Args:     []any{1, "up"},
		},
		{
			Key:      tcell.KeyDown,
			Callback: p.Move,
			Args:     []any{1, "down"},
		},
		{
			Key:      tcell.KeyRight,
			Callback: p.Move,
			Args:     []any{1, "right"},
		},
		{
			Key:      tcell.KeyLeft,
			Callback: p.Move,
			Args:     []any{1, "left"},
		},
		{
			Rune:     'x',
			Callback: p.Reset,
			Args:     nil,
		},
	}
	p.HandleKeyboardForActions(event, actions)
}

func demoOne() {
	fmt.Println("ThEngine demo-one")
	camera := engine.NewCamera(nil, api.NewSize(40, 80))
	defaultStyle := tcell.StyleDefault
	text := engine.NewCanvasFromString("Hello World", &defaultStyle)
	text.Render(camera)
	appEngine := engine.GetEngine()
	appEngine.Init()
	camera.Draw(true, appEngine.GetScreen())
	appEngine.Run(60.0)
}

func demoTwo() {
	fmt.Println("ThEngine demo-two")
	camera := engine.NewCamera(nil, api.NewSize(40, 80))
	//defaultStyle := tcell.StyleDefault.Foreground(tcell.Color104).Background(tcell.ColorBlack).Attributes(tcell.AttrBlink)
	styleOne := tcell.StyleDefault.Foreground(tcell.Color100).Background(tcell.ColorBlack)
	styleTwo := tcell.StyleDefault.Foreground(tcell.Color101).Background(tcell.ColorBlack)
	scene := engine.NewScene("scene", camera)
	textOne := engine.NewEntity("text-one", api.NewPoint(0, 0), api.NewSize(1, 1), &styleOne)
	textOneCanvas := engine.NewCanvasFromString("Hello World!!!", &styleOne)
	textOne.SetCanvas(textOneCanvas)
	scene.AddEntity(textOne)
	textTwo := engine.NewEntity("text-two", api.NewPoint(0, 1), api.NewSize(1, 1), &styleTwo)
	textTwoCanvas := engine.NewCanvasFromString("Hello World******", &styleTwo)
	textTwo.SetCanvas(textTwoCanvas)
	scene.AddEntity(textTwo)
	appEngine := engine.GetEngine()
	if !appEngine.GetSceneManager().AddScene(scene) {
		panic(fmt.Sprintf("can not add scene %s", scene.GetName()))
	}
	if !appEngine.GetSceneManager().SetSceneAsActive(scene) {
		panic(fmt.Sprintf("can not set scene %s as active", scene.GetName()))
	}
	if !appEngine.GetSceneManager().SetSceneAsVisible(scene) {
		panic(fmt.Sprintf("can not set scene %s as visible", scene.GetName()))
	}
	appEngine.Init()
	appEngine.Run(60.0)
}

func demoThree() {
	tools.Logger.WithField("module", "main").Infof("ThEngine demo-three")
	fmt.Println("ThEngine demo-three")
	camera := engine.NewCamera(nil, api.NewSize(40, 80))
	styleOne := tcell.StyleDefault.Foreground(tcell.ColorRed).Background(tcell.ColorWhite)
	scene := engine.NewScene("scene", camera)
	textOne := engine.NewEntity("text-one", api.NewPoint(0, 0), api.NewSize(3, 3), &styleOne)
	textOneCanvas := engine.NewCanvasFromString("\\ /\n O \n/ \\", &styleOne)
	textOne.SetCanvas(textOneCanvas)
	scene.AddEntity(textOne)
	appEngine := engine.GetEngine()
	if !appEngine.GetSceneManager().AddScene(scene) {
		panic(fmt.Sprintf("can not add scene %s", scene.GetName()))
	}
	appEngine.GetSceneManager().SetSceneAsActive(scene)
	appEngine.GetSceneManager().SetSceneAsVisible(scene)
	appEngine.Init()
	appEngine.Run(60.0)
}

func demoFour(dryRun bool) {
	tools.Logger.WithField("module", "main").WithField("dry-mode", dryRun).Infof("ThEngine demo-four")
	fmt.Println("ThEngine demo-four")
	camera := engine.NewCamera(nil, api.NewSize(40, 80))
	styleOne := tcell.StyleDefault.Foreground(tcell.ColorRed).Background(tcell.ColorWhite)
	scene := engine.NewScene("scene", camera)
	textOne := widgets.NewText("text-one", api.NewPoint(0, 0), api.NewSize(3, 3), &styleOne, "Hello\nWorld!")
	scene.AddEntity(textOne)
	appEngine := engine.GetEngine()
	if !appEngine.GetSceneManager().AddScene(scene) {
		panic(fmt.Sprintf("can not add scene %s", scene.GetName()))
	}
	appEngine.GetSceneManager().SetSceneAsActive(scene)
	appEngine.GetSceneManager().SetSceneAsVisible(scene)
	// appEngine.SetDryRun(dryRun)
	appEngine.Init()
	appEngine.Run(60.0)
}

func demoFive(dryRun bool) {
	tools.Logger.WithField("module", "main").WithField("dry-mode", dryRun).Infof("ThEngine demo-five")
	fmt.Println("ThEngine demo-five")
	camera := engine.NewCamera(nil, api.NewSize(40, 80))
	styleOne := tcell.StyleDefault.Foreground(tcell.ColorRed).Background(tcell.ColorBlack)
	styleTwo := tcell.StyleDefault.Foreground(tcell.ColorRed).Background(tcell.ColorWhite)
	styleThree := tcell.StyleDefault.Foreground(tcell.ColorBlue).Background(tcell.ColorRed)
	scene := engine.NewScene("scene", camera)
	textOne := widgets.NewText("text-one", api.NewPoint(0, 0), api.NewSize(1, 1), &styleOne, "Name:")
	scene.AddEntity(textOne)
	inputOne := widgets.NewTextInput("input-one", api.NewPoint(6, 0), api.NewSize(30, 1), &styleTwo, "Jose Carlos Recuero")
	scene.AddEntity(inputOne)
	appEngine := engine.GetEngine()
	if !appEngine.GetSceneManager().AddScene(scene) {
		panic(fmt.Sprintf("can not add scene %s", scene.GetName()))
	}
	appEngine.GetSceneManager().SetSceneAsActive(scene)
	appEngine.GetSceneManager().SetSceneAsVisible(scene)
	// appEngine.SetDryRun(dryRun)
	appEngine.Init()
	screen := appEngine.GetScreen()
	//style := tcell.StyleDefault
	screen.SetStyle(styleThree)
	screen.SetCursorStyle(tcell.CursorStyleBlinkingBlock)
	//screen.SetStyle(style)
	screen.ShowCursor(0, 0)
	appEngine.Run(60.0)
}

func demoSix(dryRun bool) {
	tools.Logger.WithField("module", "main").WithField("dry-mode", dryRun).Infof("ThEngine demo-six")
	fmt.Println("ThEngine demo-six")
	camera := engine.NewCamera(nil, api.NewSize(40, 80))
	styleOne := tcell.StyleDefault.Foreground(tcell.ColorRed).Background(tcell.ColorBlack)
	styleTwo := tcell.StyleDefault.Foreground(tcell.ColorRed).Background(tcell.ColorLightBlue)
	scene := engine.NewScene("scene", camera)
	textFirstName := widgets.NewText("text-first-name", api.NewPoint(0, 0), api.NewSize(1, 1), &styleOne, "First Name:")
	scene.AddEntity(textFirstName)
	textLastName := widgets.NewText("text-last-name", api.NewPoint(0, 1), api.NewSize(1, 1), &styleOne, "Last Name:")
	scene.AddEntity(textLastName)
	textAge := widgets.NewText("text-age", api.NewPoint(0, 2), api.NewSize(1, 1), &styleOne, "Age:")
	scene.AddEntity(textAge)
	inputFirstName := widgets.NewTextInput("input-first-name", api.NewPoint(12, 0), api.NewSize(30, 1), &styleTwo, "Jose Carlos")
	scene.AddEntity(inputFirstName)
	inputLastName := widgets.NewTextInput("input-last-name", api.NewPoint(12, 1), api.NewSize(30, 1), &styleTwo, "Recuero Arias")
	scene.AddEntity(inputLastName)
	inputAge := widgets.NewTextInput("input-age", api.NewPoint(12, 2), api.NewSize(30, 1), &styleTwo, "57")
	scene.AddEntity(inputAge)
	appEngine := engine.GetEngine()
	if !appEngine.GetSceneManager().AddScene(scene) {
		panic(fmt.Sprintf("can not add scene %s", scene.GetName()))
	}
	appEngine.GetSceneManager().SetSceneAsActive(scene)
	appEngine.GetSceneManager().SetSceneAsVisible(scene)
	appEngine.Init()
	screen := appEngine.GetScreen()
	//screen.SetCursorStyle(tcell.CursorStyleBlinkingBlock)
	screen.SetCursorStyle(tcell.CursorStyleSteadyBlock)
	appEngine.Run(60.0)
}

func demoSeven(dryRun bool) {
	tools.Logger.WithField("module", "main").WithField("dry-mode", dryRun).Infof("ThEngine demo-seven")
	fmt.Println("ThEngine demo-seven")
	camera := engine.NewCamera(nil, api.NewSize(40, 80))
	styleOne := tcell.StyleDefault.Foreground(tcell.ColorRed).Background(tcell.ColorWhite)
	scene := engine.NewScene("scene", camera)
	player := NewPlayer("player", api.NewPoint(0, 0), api.NewSize(3, 3), &styleOne)
	playerCanvas := engine.NewCanvasFromString("\\ /\n O \n/ \\", &styleOne)
	player.SetCanvas(playerCanvas)
	scene.AddEntity(player)
	appEngine := engine.GetEngine()
	if !appEngine.GetSceneManager().AddScene(scene) {
		panic(fmt.Sprintf("can not add scene %s", scene.GetName()))
	}
	appEngine.GetSceneManager().SetSceneAsActive(scene)
	appEngine.GetSceneManager().SetSceneAsVisible(scene)
	appEngine.Init()
	appEngine.Run(60.0)
}

func main() {
	demoSeven(true)
}
