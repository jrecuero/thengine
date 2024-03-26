package engine_test

import (
	"testing"

	"github.com/jrecuero/thengine/pkg/engine"
)

func TestEObject(t *testing.T) {
	cases := []struct {
		input struct {
			name    string
			active  bool
			visible bool
		}
		exp struct {
			name    string
			active  bool
			visible bool
		}
	}{
		{
			input: struct {
				name    string
				active  bool
				visible bool
			}{
				name:    "test",
				active:  true,
				visible: true,
			},
			exp: struct {
				name    string
				active  bool
				visible bool
			}{
				name:    "test",
				active:  true,
				visible: true,
			},
		},
	}
	for i, c := range cases {
		got := engine.NewEObject(c.input.name)
		if got == nil {
			t.Errorf("[%d] NewEObject Error exp:*EObject got:nil", i)
			continue
		}
		if c.exp.name != got.GetName() {
			t.Errorf("[%d] NewEObject GetName Error exp:%s got:%s", i, c.exp.name, got.GetName())
		}
		if got.GetID() == "" {
			t.Errorf("[%d] NewEObject GetID Error exp:not('') got:'%s'", i, got.GetID())
		}
		if c.exp.active != got.IsActive() {
			t.Errorf("[%d] NewEObject IsActive Error exp:%t got:%t", i, c.exp.active, got.IsActive())
		}
		if c.exp.visible != got.IsVisible() {
			t.Errorf("[%d] NewEObject IsVisible Error exp:%t got:%t", i, c.exp.visible, got.IsVisible())
		}
	}
}

func TestEObjectSetName(t *testing.T) {
	cases := []struct {
		input string
		exp   string
	}{
		{
			input: "new-test",
			exp:   "new-test",
		},
	}

	for i, c := range cases {
		eobj := engine.NewEObject("")
		eobj.SetName(c.input)
		got := eobj.GetName()
		if c.exp != got {
			t.Errorf("[%d] SetName Error exp:%s got:%s", i, c.exp, got)
		}
	}
}

func TestEObjectSetActive(t *testing.T) {
	cases := []struct {
		input bool
		exp   bool
	}{
		{
			input: true,
			exp:   true,
		},
	}

	for i, c := range cases {
		eobj := engine.NewEObject("")
		eobj.SetActive(c.input)
		got := eobj.IsActive()
		if c.exp != got {
			t.Errorf("[%d] IsActive Error exp:%t got:%t", i, c.exp, got)
		}
	}
}

func TestEObjectSetVisible(t *testing.T) {
	cases := []struct {
		input bool
		exp   bool
	}{
		{
			input: true,
			exp:   true,
		},
	}

	for i, c := range cases {
		eobj := engine.NewEObject("")
		eobj.SetVisible(c.input)
		got := eobj.IsVisible()
		if c.exp != got {
			t.Errorf("[%d] IsVisible Error exp:%t got:%t", i, c.exp, got)
		}
	}
}
