package main

import (
	"github.com/jrecuero/thengine/pkg/api"
	"github.com/jrecuero/thengine/pkg/engine"
)

// -----------------------------------------------------------------------------
// Module private constants
// -----------------------------------------------------------------------------

var (
	theCamera = engine.NewCamera(TheGameBoxOrigin, TheGameBoxSize)
	theEngine = engine.GetEngine()
	theFPS    = 60.0
)

// -----------------------------------------------------------------------------
// Module public constants
// -----------------------------------------------------------------------------

var (
	TheScreenWidth          = 180
	TheScreenHeight         = 48
	TheGameBoxOrigin        = api.NewPoint(0, 0)
	TheGameBoxSize          = api.NewSize(TheScreenWidth, TheScreenHeight)
	TheStageBoxOrigin       = api.NewPoint(1, 1)
	TheStageBoxSize         = api.NewSize(130, 30)
	TheDiceBoxOrigin        = api.NewPoint(1, 31)
	TheDiceBoxSize          = api.NewSize(TheScreenWidth-2, 8)
	ThePlayerBoxOrigin      = api.NewPoint(131, 1)
	ThePlayerBoxSize        = api.NewSize(48, 15)
	TheEnemyBoxOrigin       = api.NewPoint(131, 16)
	TheEnemyBoxSize         = api.NewSize(48, 15)
	TheKeysBoxOrigin        = api.NewPoint(1, 39)
	TheKeysBoxSize          = api.NewSize(TheScreenWidth-2, 4)
	TheCommandLineBoxOrigin = api.NewPoint(1, 43)
	TheCommandLineBoxSize   = api.NewSize(TheScreenWidth-2, 4)
)
