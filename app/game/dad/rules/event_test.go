package rules_test

import (
	"fmt"
	"testing"

	"github.com/jrecuero/thengine/app/game/dad/rules"
)

type DialogAction struct {
	*rules.Action
	message string
	branch  rules.UniqueActionId
}

func NewDialogAction(message string, branch rules.UniqueActionId) *DialogAction {
	return &DialogAction{
		Action:  rules.NewAction(),
		message: message,
		branch:  branch,
	}
}

func (a *DialogAction) Execute(e rules.IEvent, args ...any) (rules.UniqueActionId, error) {
	fmt.Println(a.message)
	return a.branch, nil
}

type QuestionAction struct {
	*rules.Action
	message string
	branch1 rules.UniqueActionId
	branch2 rules.UniqueActionId
}

func NewQuestionAction(message string, b1, b2 rules.UniqueActionId) *QuestionAction {
	return &QuestionAction{
		Action:  rules.NewAction(),
		branch1: b1,
		branch2: b2,
		message: message,
	}
}

func (a *QuestionAction) Execute(e rules.IEvent, args ...any) (rules.UniqueActionId, error) {
	question := args[0].(bool)
	cache := e.GetCache()
	fmt.Println(a.message, question)
	if _, found := cache.Get("key"); found {
		fmt.Println("you already have the key")
		return a.branch1, nil
	}
	if question {
		e.GetCache().Set("key", true)
		return a.branch1, nil
	} else {
		return a.branch2, nil
	}
}

func TestEventRun(t *testing.T) {
	act1 := NewDialogAction("you can open the door", rules.UAIdEnd)
	act2 := NewDialogAction("you don't have the key", rules.UAIdEnd)
	question := NewQuestionAction("do you want the key?", act1.GetUniqueActionId(), act2.GetUniqueActionId())
	actions := []rules.IAction{
		NewDialogAction("hello", rules.UAIdNone),
		NewDialogAction("world", rules.UAIdNone),
		question,
		act1,
		act2,
		rules.NewEndAction(),
	}
	e := rules.NewEvent("test/event/1", actions)
	if got := e.Run(true); got != nil {
		t.Errorf("[0] Run error exp:nil got:%s", got.Error())
	}
	if got := e.Run(false); got != nil {
		t.Errorf("[0] Run error exp:nil got:%s", got.Error())
	}
	e = rules.NewEvent("test/event/2", actions)
	if got := e.Run(false); got != nil {
		t.Errorf("[0] Run error exp:nil got:%s", got.Error())
	}
	if got := e.Run(true); got != nil {
		t.Errorf("[0] Run error exp:nil got:%s", got.Error())
	}
}
