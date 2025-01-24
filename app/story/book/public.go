package book

import "github.com/jrecuero/thengine/pkg/storyboard"

var (
	TheStart   = storyboard.NewNode("storyboard/node/start/1")
	ThePreface = storyboard.NewNode("storyboard/node/preface/1")
	TheEnd     = storyboard.NewNode("storyboard/node/end/1")

	TheBook = storyboard.NewStoryBoard("The Book")
)

const (
	TheNarrator = "Narrator"
	TheHero     = "Hero"
	ThePrincess = "Princess"
	TheVillain  = "Villain"
)

func init() {
	TheBook.SetStart(createStart())
}

func aboutThePrince() storyboard.INode {
	node := storyboard.NewNode("storyboard/node/about/prince/1")
	node.SetSpeaker(TheNarrator)
	node.AddText("The prince was a valient young man")
	node.AddNext(TheEnd)
	return node
}

func aboutThePrincess() storyboard.INode {
	node := storyboard.NewNode("storyboard/node/about/princess/1")
	node.SetSpeaker(TheNarrator)
	node.AddText("The princess was a beautiful and cheerful girl")
	node.AddNext(TheEnd)
	return node
}

func createStart() storyboard.INode {
	TheStart.SetSpeaker(TheNarrator)
	TheStart.AddText("Once upon a time there was prince in a shiny kingdom")
	TheStart.AddText("Prince and his family were happy and the kingdom love them")
	TheStart.AddText("And the prince was in love with the princess of a neighbord kingdom")
	TheStart.AddText("Do you want to know more about?")
	q1 := storyboard.NewQuestion("The Prince")
	q1.AddNext(aboutThePrince())
	q2 := storyboard.NewQuestion("The Princess")
	q2.AddNext(aboutThePrincess)
	TheStart.AddQuestion(q1, q2)
	return TheStart
}
