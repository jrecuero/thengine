package main

import (
	"github.com/jrecuero/thengine/pkg/api"
	"github.com/jrecuero/thengine/pkg/constants"
	"github.com/jrecuero/thengine/pkg/engine"
)

// -----------------------------------------------------------------------------
// Module private variables
// -----------------------------------------------------------------------------

var (
	theCamera   = engine.NewCamera(TheStoryBoxOrigin, TheStoryBoxSize)
	theEngine   = engine.GetEngine()
	theBoxStyle = &constants.WhiteOverBlack
)

// -----------------------------------------------------------------------------
// Module public variables
// -----------------------------------------------------------------------------

var (
	TheStoryBoxOrigin = api.NewPoint(0, 0)
	TheStoryBoxSize   = api.NewSize(TheScreenWidth, TheScreenHeight)
)
