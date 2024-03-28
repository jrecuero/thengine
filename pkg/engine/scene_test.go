package engine_test

import (
	"testing"

	"github.com/jrecuero/thengine/pkg/api"
	"github.com/jrecuero/thengine/pkg/engine"
)

func TestScene(t *testing.T) {
	cases := []struct {
		input string
		exp   string
	}{
		{
			input: "text",
			exp:   "text",
		},
	}

	for i, c := range cases {
		got := engine.NewScene(c.input, engine.NewCamera(nil, api.NewSize(0, 0)))
		if got == nil {
			t.Errorf("[%d] NewScene Error exp:*Scene got:nil", i)
			continue
		}
		if c.exp != got.GetName() {
			t.Errorf("[%d] NewScene Name Error exp:%s got:%s", i, c.exp, got.GetName())
		}
	}
}
