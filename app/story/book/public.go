package book

import "github.com/jrecuero/thengine/pkg/storyboard"

var (
	TheStart   = storyboard.NewNode("storyboard/node/start/1")
	ThePreface = storyboard.NewNode("storyboard/node/preface/1")
	TheEnd     = storyboard.NewNode("storyboard/node/end/1")
)

const (
	TheNarrator = "Narrator"
	TheHero     = "Hero"
	ThePrincess = "Princess"
	TheVillain  = "Villain"
)
