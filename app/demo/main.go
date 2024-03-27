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
}

func NewPlayer(name string, position *api.Point, size *api.Size, style *tcell.Style) *Player {
	player := &Player{
		Widget: widgets.NewWidget(name, position, size, style),
	}
	player.SetFocusType(engine.SingleFocus)
	player.SetFocusEnable(true)
	return player

}
func (p *Player) Move(args ...any) {
	steps := args[0].(int)
	direction := args[1].(string)
	tools.Logger.WithField("module", "main").
		WithField("struct", "Player").
		WithField("function", "Move").
		Debugf("stets %d direction %s", steps, direction)
	if direction == "up" {
		x, y := p.GetPosition().Get()
		p.SetPosition(api.NewPoint(x, y-1))
	} else if direction == "down" {
		x, y := p.GetPosition().Get()
		p.SetPosition(api.NewPoint(x, y+1))
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
	}
	p.HandleKeyboardForActions(event, actions)
}

func demoOne() {
	fmt.Println("ThEngine demo-one")
	screen := engine.NewScreen(nil, api.NewSize(40, 80))
	defaultStyle := tcell.StyleDefault
	text := engine.NewCanvasFromString("Hello World", &defaultStyle)
	text.Render(screen)
	appEngine := engine.GetEngine()
	appEngine.Init()
	screen.Draw(true, appEngine.GetDisplay())
	appEngine.Run(60.0)
}

func demoTwo() {
	fmt.Println("ThEngine demo-two")
	screen := engine.NewScreen(nil, api.NewSize(40, 80))
	//defaultStyle := tcell.StyleDefault.Foreground(tcell.Color104).Background(tcell.ColorBlack).Attributes(tcell.AttrBlink)
	styleOne := tcell.StyleDefault.Foreground(tcell.Color100).Background(tcell.ColorBlack)
	styleTwo := tcell.StyleDefault.Foreground(tcell.Color101).Background(tcell.ColorBlack)
	scene := engine.NewScene("scene", screen)
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
	screen := engine.NewScreen(nil, api.NewSize(40, 80))
	styleOne := tcell.StyleDefault.Foreground(tcell.ColorRed).Background(tcell.ColorWhite)
	scene := engine.NewScene("scene", screen)
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
	screen := engine.NewScreen(nil, api.NewSize(40, 80))
	styleOne := tcell.StyleDefault.Foreground(tcell.ColorRed).Background(tcell.ColorWhite)
	scene := engine.NewScene("scene", screen)
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
	screen := engine.NewScreen(nil, api.NewSize(40, 80))
	styleOne := tcell.StyleDefault.Foreground(tcell.ColorRed).Background(tcell.ColorBlack)
	styleTwo := tcell.StyleDefault.Foreground(tcell.ColorRed).Background(tcell.ColorWhite)
	styleThree := tcell.StyleDefault.Foreground(tcell.ColorBlue).Background(tcell.ColorRed)
	scene := engine.NewScene("scene", screen)
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
	display := appEngine.GetDisplay()
	//style := tcell.StyleDefault
	display.SetStyle(styleThree)
	display.SetCursorStyle(tcell.CursorStyleBlinkingBlock)
	//display.SetStyle(style)
	display.ShowCursor(0, 0)
	appEngine.Run(60.0)
}

func demoSix(dryRun bool) {
	tools.Logger.WithField("module", "main").WithField("dry-mode", dryRun).Infof("ThEngine demo-six")
	fmt.Println("ThEngine demo-six")
	screen := engine.NewScreen(nil, api.NewSize(40, 80))
	styleOne := tcell.StyleDefault.Foreground(tcell.ColorRed).Background(tcell.ColorBlack)
	styleTwo := tcell.StyleDefault.Foreground(tcell.ColorRed).Background(tcell.ColorLightBlue)
	scene := engine.NewScene("scene", screen)
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
	display := appEngine.GetDisplay()
	//display.SetCursorStyle(tcell.CursorStyleBlinkingBlock)
	display.SetCursorStyle(tcell.CursorStyleSteadyBlock)
	appEngine.Run(60.0)
}

func demoSeven(dryRun bool) {
	tools.Logger.WithField("module", "main").WithField("dry-mode", dryRun).Infof("ThEngine demo-seven")
	fmt.Println("ThEngine demo-seven")
	screen := engine.NewScreen(nil, api.NewSize(40, 80))
	styleOne := tcell.StyleDefault.Foreground(tcell.ColorRed).Background(tcell.ColorWhite)
	scene := engine.NewScene("scene", screen)
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
