package engine_test

import (
	"testing"

	"github.com/jrecuero/thengine/pkg/api"
	"github.com/jrecuero/thengine/pkg/engine"
)

func TestSceneManager(t *testing.T) {
	cases := []struct {
		exp struct {
			scenes  int
			active  int
			visible int
		}
	}{
		{
			exp: struct {
				scenes  int
				active  int
				visible int
			}{
				scenes:  1,
				active:  0,
				visible: 0,
			},
		},
	}

	for i, c := range cases {
		got := engine.NewSceneManager(engine.NewScreen(nil, api.NewSize(0, 0)))
		if got == nil {
			t.Errorf("[%d] NewSceneManager Error exp:*SceneManager got:nil", i)
		}
		if c.exp.scenes != len(got.GetAllScenes()) {
			t.Errorf("[%d] NewSceneManager Error Scenes exp:%d got:%d",
				i, c.exp.scenes, len(got.GetAllScenes()))
		}
		if c.exp.active != len(got.GetAllActiveScenes()) {
			t.Errorf("[%d] NewSceneManager Error Scenes-Active exp:%d got:%d",
				i, c.exp.active, len(got.GetAllActiveScenes()))
		}
		if c.exp.visible != len(got.GetAllVisibleScenes()) {
			t.Errorf("[%d] NewSceneManager Error Scenes-Visible exp:%d got:%d",
				i, c.exp.visible, len(got.GetAllVisibleScenes()))
		}
	}
}

func TestSceneManagerAddScene(t *testing.T) {
}
