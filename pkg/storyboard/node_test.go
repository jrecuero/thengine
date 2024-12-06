package storyboard_test

import (
	"testing"

	"github.com/jrecuero/thengine/pkg/storyboard"
)

// Test case for Node's AddText method
func TestNodeAddText(t *testing.T) {
	// Create a new Node instance
	node := storyboard.NewNode("Node1")

	// Create mock ConditionalText
	// instances
	condText1 := storyboard.NewConditionalText("This is some text.", nil)
	condText2 := storyboard.NewConditionalText("This is some more text.", nil)

	// Add text entries to the
	// node
	node.AddText(condText1, condText2)

	// Test if the text
	// entries were
	// added correctly
	if len(node.GetText()) != 2 {
		t.Errorf("[Node] Expected 2 text entries, got %d", len(node.GetText()))
	}
	if node.GetText()[0].GetText() != "This is some text." {
		t.Errorf("[Node] Expected text 'This is some text.', got '%s'", node.GetText()[0].GetText())
	}
	if node.GetText()[1].GetText() != "This is some more text." {
		t.Errorf("[Node] Expected text 'This is some more text.', got '%s'", node.GetText()[1].GetText())
	}
}

// Test case for Node's AddNext method
func TestNodeAddNext(t *testing.T) {
	// Create a new Node instance
	node := storyboard.NewNode("Node1")
	node2 := storyboard.NewNode("Node2")
	node3 := storyboard.NewNode("Node3")

	// Create mock ConditionalNext instances
	condNext1 := storyboard.NewConditionalNext(node2, nil)
	condNext2 := storyboard.NewConditionalNext(node3, nil)

	// Add next nodes to the node
	node.AddNext(condNext1, condNext2)

	// Test if the next nodes were added correctly
	if len(node.GetNext()) != 2 {
		t.Errorf("[Node] Expected 2 next nodes, got %d", len(node.GetNext()))
	}
	if node.GetNext()[0].GetNode() != node2 {
		t.Errorf("[Node] Expected next node 'Node2', got '%s'", node.GetNext()[0].GetNode())
	}
	if node.GetNext()[1].GetNode() != node3 {
		t.Errorf("[Node] Expected next node 'Node3', got '%s'", node.GetNext()[1].GetNode())
	}
}

// Test case for Node's AddQuestion method
func TestNodeAddQuestion(t *testing.T) {
	// Create a new Node instance
	node := storyboard.NewNode("Node1")

	// Create mock Question instances (for simplicity, using the same struct as in the previous tests)
	question := storyboard.NewQuestion("Question")

	// Add questions to the node
	node.AddQuestion(question)

	// Test if the questions were added correctly
	if len(node.GetQuestions()) != 1 {
		t.Errorf("[Node] Expected 1 question, got %d", len(node.GetQuestions()))
	}
	if node.GetQuestions()[0] != question {
		t.Errorf("[Node] Expected question 'Question1', got '%v'", question)
	}
}

// Test case for Node's SetSpeaker and GetSpeaker methods
func TestNodeSetGetSpeaker(t *testing.T) {
	// Create a new Node instance
	node := storyboard.NewNode("Node1")

	// Set a speaker
	node.SetSpeaker("Narrator")

	// Test if the speaker is correctly set
	if node.GetSpeaker() != "Narrator" {
		t.Errorf("[Node] Expected speaker 'Narrator', got '%s'", node.GetSpeaker())
	}
}

// Test case for Node's GetID method
func TestNodeGetID(t *testing.T) {
	// Create a new Node instance with an ID
	node := storyboard.NewNode("Node1")

	// Test if the ID is correctly returned
	if node.GetID() != "Node1" {
		t.Errorf("[Node] Expected ID 'Node1', got '%s'", node.GetID())
	}
}
