package tools_test

import (
	"testing"

	"github.com/jrecuero/thengine/pkg/tools"
)

func TestAbs(t *testing.T) {
	cases := []struct {
		input int
		exp   int
	}{
		{
			input: 0,
			exp:   0,
		},
		{
			input: 1,
			exp:   1,
		},
		{
			input: -2,
			exp:   2,
		},
	}
	for i, c := range cases {
		got := tools.Abs(c.input)
		if c.exp != got {
			t.Errorf("[%d] Abs Error exp:%d got:%d", i, c.exp, got)
		}
	}
}

func TestSign(t *testing.T) {
	cases := []struct {
		input int
		exp   int
	}{
		{
			input: 0,
			exp:   1,
		},
		{
			input: 1,
			exp:   1,
		},
		{
			input: -2,
			exp:   -1,
		},
	}
	for i, c := range cases {
		got := tools.Sign(c.input)
		if c.exp != got {
			t.Errorf("[%d] Sign Error exp:%d got:%d", i, c.exp, got)
		}
	}
}

func TestMax(t *testing.T) {
	cases := []struct {
		input []int
		exp   int
	}{
		{
			input: []int{0, 1, 10, 5},
			exp:   10,
		},
		{
			input: []int{10, 7},
			exp:   10,
		},
		{
			input: []int{1, 7, -2},
			exp:   7,
		},
	}
	for i, c := range cases {
		got := tools.Max(c.input...)
		if c.exp != got {
			t.Errorf("[%d] Max Error exp:%d got:%d", i, c.exp, got)
		}
	}
}

func TestMin(t *testing.T) {
	cases := []struct {
		input []int
		exp   int
	}{
		{
			input: []int{0, 1, 10, 5},
			exp:   0,
		},
		{
			input: []int{10, 7},
			exp:   7,
		},
		{
			input: []int{1, 7, -2},
			exp:   -2,
		},
	}
	for i, c := range cases {
		got := tools.Min(c.input...)
		if c.exp != got {
			t.Errorf("[%d] Max Error exp:%d got:%d", i, c.exp, got)
		}
	}
}
