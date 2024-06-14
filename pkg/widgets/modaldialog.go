// modaldialog.go contains all required data and logic to create a modal dialog
// scene.
package widgets

import (
	"fmt"

	"github.com/jrecuero/thengine/pkg/engine"
)

// -----------------------------------------------------------------------------
//
// ModalDialog
//
// -----------------------------------------------------------------------------

type ModalDialog struct {
	parentScene engine.IScene
	dialogScene engine.IScene
	dialog      *Dialog
}

func NewModalDialog(parentScene engine.IScene) *ModalDialog {
	dialogSceneName := fmt.Sprintf("scene/dialog/%s", parentScene.GetName())
	d := &ModalDialog{
		parentScene: parentScene,
		dialogScene: engine.NewScene(dialogSceneName, parentScene.GetCamera()),
		dialog:      nil,
	}
	return d
}

// -----------------------------------------------------------------------------
// ModalDialog public methods
// -----------------------------------------------------------------------------

func (d *ModalDialog) Close() {

	sceneManager := engine.GetEngine().GetSceneManager()
	sceneManager.RemoveScene(d.dialogScene)
	sceneManager.SetSceneAsActive(d.parentScene)
	sceneManager.SetSceneAsVisible(d.parentScene)
	sceneManager.UpdateFocus()
}

func (d *ModalDialog) GetDialog() *Dialog {
	return d.dialog
}

func (d *ModalDialog) GetDialogScene() engine.IScene {
	return d.dialogScene
}

func (d *ModalDialog) GetParentScene() engine.IScene {
	return d.parentScene
}

func (d *ModalDialog) Open(dialog *Dialog) {
	d.dialog = dialog
	sceneManager := engine.GetEngine().GetSceneManager()
	sceneManager.RemoveSceneAsActive(d.parentScene)
	sceneManager.AddScene(d.dialogScene)
	sceneManager.SetSceneAsActive(d.dialogScene)
	sceneManager.SetSceneAsVisible(d.dialogScene)
	sceneManager.UpdateFocus()
}
