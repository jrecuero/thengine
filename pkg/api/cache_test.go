package api_test

import (
	"testing"

	"github.com/jrecuero/thengine/pkg/api"
)

type cached struct {
	key  string
	data any
}

func TestCacheGetAndSet(t *testing.T) {
	cases := []struct {
		input *cached
		exp   any
	}{
		{
			input: &cached{
				key:  "one",
				data: 1,
			},
			exp: 1,
		},
		{
			input: &cached{
				key:  "two",
				data: 2,
			},
			exp: 2,
		},
	}

	cache := api.NewCache()
	if cache == nil {
		t.Errorf("NewCache Error exp:*Cache got:nil")
		return
	}
	for i, c := range cases {
		cache.Set(c.input.key, c.input.data)
		got, gotOk := cache.Get(c.input.key)
		if !gotOk {
			t.Errorf("[%d] Get Error exp:true got:%t", i, gotOk)
		}
		if c.exp != got {
			t.Errorf("[%d] Get Error.data exp:%d got %d", i, c.exp, got)
		}
	}
}

func TestCacheExists(t *testing.T) {
	cases := []struct {
		input *cached
		exp   *cached
	}{
		{
			input: &cached{
				key:  "one",
				data: 1,
			},
			exp: &cached{
				key:  "one",
				data: true,
			},
		},
		{
			input: &cached{
				key:  "two",
				data: 2,
			},
			exp: &cached{
				key:  "three",
				data: false,
			},
		},
	}

	cache := api.NewCache()
	if cache == nil {
		t.Errorf("NewCache Error exp:*Cache got:nil")
		return
	}
	for i, c := range cases {
		cache.Set(c.input.key, c.input.data)
		got := cache.Exists(c.exp.key)
		if c.exp.data != got {
			t.Errorf("[%d] Exists Error exp:%t got %t", i, c.exp.data, got)
		}
	}
}

func TestCacheClear(t *testing.T) {
	cache := api.NewCache()
	if cache == nil {
		t.Errorf("NewCache Error exp:*Cache got:nil")
		return
	}
	cache.Set("one", 1)
	cache.Set("two", 2)
	data := cache.GetCache()
	if data == nil {
		t.Errorf("GetCache Error exp:Cache got:nil")
	}
	if len(data) != 2 {
		t.Errorf("GetCache Error.len exp:%d got:%d", 2, len(data))
	}
	cache.Clear()
	data = cache.GetCache()
	if data == nil {
		t.Errorf("Cleat Error exp:Cache got:nil")
	}
	if len(data) != 0 {
		t.Errorf("Cleat Error.len exp:%d got:%d", 0, len(data))
	}
}

func TestCacheDelete(t *testing.T) {
	cache := api.NewCache()
	if cache == nil {
		t.Errorf("NewCache Error exp:*Cache got:nil")
		return
	}
	cache.Set("one", 1)
	cache.Set("two", 2)
	got, ok := cache.Get("one")
	if !ok {
		t.Errorf("Get Error exp:true got:%t", ok)
	}
	if got != 1 {
		t.Errorf("Get Error.data exp:%d got:%d", 1, got)
	}
	cache.Delete("one")
	got, ok = cache.Get("one")
	if ok {
		t.Errorf("Get Error exp:false got:%t", ok)
	}
	if got != nil {
		t.Errorf("Get Error.data exp:nil got:%d", got)
	}
}
