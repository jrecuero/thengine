package storyboard_test

import (
	"testing"

	"github.com/jrecuero/thengine/pkg/storyboard"
)

// Mock ICondition type for testing purposes
type MockCondition struct {
	ConditionValue bool
}

func (m *MockCondition) Evaluate() bool {
	return m.ConditionValue
}

// Test for Node, Question, ConditionalText, and ConditionalNext interaction
func TestStoryBoard(t *testing.T) {
	// Create a mock condition that evaluates to true
	conditionTrue := &MockCondition{ConditionValue: true}
	conditionFalse := &MockCondition{ConditionValue: false}

	// Create ConditionalNext and ConditionalText instances
	condNext1 := storyboard.NewConditionalNext("NextNode1", nil)
	condNext2 := storyboard.NewConditionalNext("NextNode2", nil)
	condText1 := storyboard.NewConditionalText("Text for True Condition", nil)
	condText2 := storyboard.NewConditionalText("Text for False Condition", nil)

	// Create a Question and associate it with ConditionalText and ConditionalNext
	question := storyboard.NewQuestion()
	question.AddText(condText1)
	question.AddText(condText2)
	question.AddNext(condNext1)
	question.AddNext(condNext2)

	// Create a Node and set a speaker and next conditional behavior
	node := storyboard.NewNode("Node1")
	node.SetSpeaker("Narrator")
	node.AddText(condText1)
	node.AddText(condText2)
	node.AddQuestion(question)
	node.AddNext(condNext1)

	// Test if the is correctly set
	if node.GetSpeaker() != "Narrator" {
		t.Errorf("[Node] Expected speaker 'Narrator', got '%s'", node.GetSpeaker())
	}

	// Test if the correct text is associated with the Node based on condition
	if len(node.GetText()) != 2 {
		t.Errorf("[Node] Expected 2 text entries, got %d", len(node.GetText()))
	}

	// Test the flow of questions and conditional texts
	if len(node.GetQuestions()) != 1 {
		t.Errorf("[Node] Expected 1 question, got %d", len(node.GetQuestions()))
	}

	// Check conditional text based on the condition
	if node.GetQuestions()[0].GetText()[0].GetText() != "Text for True Condition" {
		t.Errorf("[Node] Expected text 'Text for True Condition', got '%s'", node.GetQuestions()[0].GetText()[0].GetText())
	}

	// Test the flow of next nodes based on conditions
	if len(node.GetNext()) != 1 {
		t.Errorf("[Node] Expected 1 next node, got %d", len(node.GetNext()))
	}
	if node.GetNext()[0].GetNode() != "NextNode1" {
		t.Errorf("[Node] Expected next node 'NextNode1', got '%s'", node.GetNext()[0].GetNode())
	}

	// Test conditional evaluation for a question
	if conditionTrue.Evaluate() != true {
		t.Errorf("[Condition] Expected condition to evaluate to true, got %v", conditionTrue.Evaluate())
	}
	if conditionFalse.Evaluate() != false {
		t.Errorf("[Condition] Expected condition to evaluate to false, got %v", conditionFalse.Evaluate())
	}
}
