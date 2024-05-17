package main

import (
	"strings"

	"github.com/gdamore/tcell/v2"
	"github.com/jrecuero/thengine/pkg/api"
	"github.com/jrecuero/thengine/pkg/widgets"
)

type CommandLine struct {
	*widgets.Text
}

func NewCommandLine(name string, position *api.Point, size *api.Size, style *tcell.Style) *CommandLine {
	return &CommandLine{
		Text: widgets.NewText(name, position, size, style, ">"),
	}
}

func (t *CommandLine) AddText(str string) {
	maxLines := t.GetSize().H
	text := t.GetText()
	lines := strings.Split(text, "\n")
	if len(lines) >= maxLines {
		text = strings.Join(lines[1:], "\n")
	}
	newText := text + str
	t.SetText(newText)
}
