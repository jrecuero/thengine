package engine_test

import (
	"testing"

	"github.com/jrecuero/thengine/pkg/engine"
)

func TestFocusNewFocus(t *testing.T) {
	cases := []struct {
		input engine.FocusType
		exp   struct {
			focus     bool
			enable    bool
			focusType engine.FocusType
		}
	}{
		{
			input: engine.SingleFocus,
			exp: struct {
				focus     bool
				enable    bool
				focusType engine.FocusType
			}{
				focus:     false,
				enable:    true,
				focusType: engine.SingleFocus,
			},
		},
	}
	for i, c := range cases {
		got := engine.NewFocus(c.input)
		if got == nil {
			t.Errorf("[%d] NewFocus Error exp:*Focus got:nil", i)
		}
		if got.HasFocus() != c.exp.focus {
			t.Errorf("[%d] NewFocus Error.HasFocus exp:%t got:%t", i, c.exp.focus, got.HasFocus())
		}
		if got.IsFocusEnable() != c.exp.enable {
			t.Errorf("[%d] NewFocus Error.IsFocusEnable exp:%t got:%t", i, c.exp.enable, got.IsFocusEnable())
		}
		if got.GetFocusType() != c.exp.focusType {
			t.Errorf("[%d] NewFocus Error.GetFocusType exp%d got:%d", i, c.exp.focusType, got.GetFocusType())
		}
	}
}

func TestFocusNewDisableFocus(t *testing.T) {
	cases := []struct {
		exp struct {
			focus     bool
			enable    bool
			focusType engine.FocusType
		}
	}{
		{
			exp: struct {
				focus     bool
				enable    bool
				focusType engine.FocusType
			}{
				focus:     false,
				enable:    false,
				focusType: engine.NoFocus,
			},
		},
	}
	for i, c := range cases {
		got := engine.NewDisableFocus()
		if got == nil {
			t.Errorf("[%d] NewDisableFocus Error exp:*Focus got:nil", i)
		}
		if got.HasFocus() != c.exp.focus {
			t.Errorf("[%d] NewDisableFocus Error.HasFocus exp:%t got:%t", i, c.exp.focus, got.HasFocus())
		}
		if got.IsFocusEnable() != c.exp.enable {
			t.Errorf("[%d] NewFocus Error.IsFocusEnable exp:%t got:%t", i, c.exp.enable, got.IsFocusEnable())
		}
		if got.GetFocusType() != c.exp.focusType {
			t.Errorf("[%d] NewDisableFocus Error.GetFocusType exp%d got:%d", i, c.exp.focusType, got.GetFocusType())
		}
	}
}

func TestFocusHasFocus(t *testing.T) {
	cases := []struct {
		input struct {
			focus     bool
			focusType engine.FocusType
		}
		exp struct {
			focus     bool
			focusType engine.FocusType
		}
	}{
		{
			input: struct {
				focus     bool
				focusType engine.FocusType
			}{
				focus:     true,
				focusType: engine.SingleFocus,
			},
			exp: struct {
				focus     bool
				focusType engine.FocusType
			}{
				focus:     false,
				focusType: engine.SingleFocus,
			},
		},
	}
	for i, c := range cases {
		focus := engine.NewFocus(c.input.focusType)
		got := focus.HasFocus()
		if got != c.exp.focus {
			t.Errorf("[%d] HasFocus Error exp%t got:%t", i, c.exp.focus, got)
		}
	}
}

func TestFocusIsFocusEnable(t *testing.T) {
	cases := []struct {
		input engine.FocusType
		exp   bool
	}{
		{
			input: engine.SingleFocus,
			exp:   true,
		},
	}
	for i, c := range cases {
		focus := engine.NewFocus(c.input)
		got := focus.IsFocusEnable()
		if got != c.exp {
			t.Errorf("[%d] IsFocusEnable Error exp%t got:%t", i, c.exp, got)
		}
	}
}

func TestFocusGetFocusType(t *testing.T) {
	cases := []struct {
		input engine.FocusType
		exp   engine.FocusType
	}{
		{
			input: engine.SingleFocus,
			exp:   engine.SingleFocus,
		},
		{
			input: engine.NoFocus,
			exp:   engine.NoFocus,
		},
		{
			input: engine.MultiFocus,
			exp:   engine.MultiFocus,
		},
	}
	for i, c := range cases {
		focus := engine.NewFocus(c.input)
		got := focus.GetFocusType()
		if got != c.exp {
			t.Errorf("[%d] GetFocusType Error exp%d got:%d", i, c.exp, got)
		}
	}
}

func TestFocusSetFocusEnable(t *testing.T) {
	cases := []struct {
		input bool
		exp   bool
	}{
		{
			input: true,
			exp:   true,
		},
		{
			input: false,
			exp:   false,
		},
	}
	for i, c := range cases {
		focus := engine.NewFocus(engine.NoFocus)
		focus.SetFocusEnable(c.input)
		got := focus.IsFocusEnable()
		if got != c.exp {
			t.Errorf("[%d] SetFocusEnable Error exp%t got:%t", i, c.exp, got)
		}
	}
}

func TestFocusAcquireFocus(t *testing.T) {
	cases := []struct {
		input bool
		exp   []bool
	}{
		{
			input: true,
			exp:   []bool{true, true},
		},
		{
			input: false,
			exp:   []bool{false, false},
		},
	}
	for i, c := range cases {
		focus := engine.NewFocus(engine.NoFocus)
		focus.SetFocusEnable(c.input)
		gotOK, err := focus.AcquireFocus()
		got := focus.HasFocus()
		if c.exp[0] && err != nil {
			t.Errorf("[%d] AcquireFocus Error exp:nil got:%s", i, err.Error())
		}
		if c.exp[0] != gotOK {
			t.Errorf("[%d] AcquireFocus Error.return exp%t got:%t", i, c.exp[0], gotOK)
		}
		if c.exp[1] != got {
			t.Errorf("[%d] AcquireFocus Error.Focus exp%t got:%t", i, c.exp[1], got)
		}
	}
}

func TestFocusReleaseFocus(t *testing.T) {
	cases := []struct {
		input bool
		exp   []bool
	}{
		{
			input: true,
			exp:   []bool{true, false},
		},
		{
			input: false,
			exp:   []bool{false, false},
		},
	}
	for i, c := range cases {
		focus := engine.NewFocus(engine.NoFocus)
		focus.SetFocusEnable(c.input)
		_, _ = focus.AcquireFocus()
		gotOK, err := focus.ReleaseFocus()
		got := focus.HasFocus()
		if c.exp[0] && err != nil {
			t.Errorf("[%d] ReleaseFocus Error exp:nil got:%s", i, err.Error())
		}
		if c.exp[0] != gotOK {
			t.Errorf("[%d] ReleaseFocus Error.return exp:%t got:%t", i, c.exp[0], gotOK)
		}
		if c.exp[1] != got {
			t.Errorf("[%d] ReleaseFocus Error.Focus exp:%t got:%t", i, c.exp[1], got)
		}
	}
}
