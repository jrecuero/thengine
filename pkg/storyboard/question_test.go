package storyboard_test

import (
	"testing"

	"github.com/jrecuero/thengine/pkg/storyboard"
)

func TestQuestionAddNext(t *testing.T) {
	// Create a new Question instance
	question := storyboard.NewQuestion("Question")

	nodeOne := storyboard.NewNode("NodeOne")
	nodeTwo := storyboard.NewNode("NodeTwo")
	// Create mock ConditionalNext instances
	condNext1 := storyboard.NewConditionalNext(nodeOne, nil)
	condNext2 := storyboard.NewConditionalNext(nodeTwo, nil)

	// Add next nodes to the question
	question.AddNext(condNext1, condNext2)

	// Test if the next nodes were added correctly
	if len(question.GetNext()) != 2 {
		t.Errorf("[Question] Expected 2 next nodes, got %d", len(question.GetNext()))
	}
	if question.GetNext()[0].GetNode() != nodeOne {
		t.Errorf("[Question] Expected next node 'NodeOne', got '%s'", question.GetNext()[0].GetNode())
	}
	if question.GetNext()[1].GetNode() != nodeTwo {
		t.Errorf("[Question] Expected next node 'NodeTwo', got '%s'", question.GetNext()[1].GetNode())
	}
}

func TestQuestionAddText(t *testing.T) {
	// Create a new Question instance
	question := storyboard.NewQuestion("Question")

	// Create mock ConditionalText instances
	condText1 := storyboard.NewConditionalText("What do you want to do?", nil)
	condText2 := storyboard.NewConditionalText("Choose wisely.", nil)

	// Add text to the question
	question.AddText(condText1, condText2)

	// Test if the text entries were added correctly
	if len(question.GetText()) != 2 {
		t.Errorf("[Question] Expected 2 text entries, got %d", len(question.GetText()))
	}
	if question.GetText()[0].GetText() != "What do you want to do?" {
		t.Errorf("[Question] Expected text 'What do you want to do?', got '%s'", question.GetText()[0].GetText())
	}
	if question.GetText()[1].GetText() != "Choose wisely." {
		t.Errorf("[Question] Expected text 'Choose wisely.', got '%s'", question.GetText()[1].GetText())
	}
}

func TestQuestionGetCondition(t *testing.T) {
	// Create a new Question instance
	question := storyboard.NewQuestion("Question")

	// Create a mock condition instance and set it on the question
	condition := &Condition{conditionMet: true}
	question.SetCondition(condition)

	// Test if the condition is returned correctly
	if question.GetCondition() != condition {
		t.Errorf("[Question] Expected condition to be set correctly")
	}
}
