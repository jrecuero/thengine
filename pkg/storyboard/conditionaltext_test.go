package storyboard_test

import (
	"testing"

	"github.com/jrecuero/thengine/pkg/storyboard"
)

func TestNewConditionalText(t *testing.T) {
	// Create a condition instance
	condition := &Condition{conditionMet: true}
	// Create a new ConditionalText instance
	conditionalText := storyboard.NewConditionalText("This is a test text", condition)

	// Test the initial values
	if conditionalText.GetText() != "This is a test text" {
		t.Errorf("Expected text to be 'This is a test text', got '%s'", conditionalText.GetText())
	}
	if conditionalText.GetCondition() != condition {
		t.Errorf("Expected condition to match the given condition")
	}
}

func TestConditionalTextSetCondition(t *testing.T) {
	// Create an initial condition instance
	condition := &Condition{conditionMet: false}
	// Create a new ConditionalText instance
	conditionalText := storyboard.NewConditionalText("Text for condition", condition)

	// Set a new condition
	newCondition := &Condition{conditionMet: true}
	conditionalText.SetCondition(newCondition)

	// Test if the condition was updated
	if conditionalText.GetCondition() != newCondition {
		t.Errorf("Expected condition to be updated to newCondition")
	}
}

func TestConditionalTextSetText(t *testing.T) {
	// Create a condition instance
	condition := &Condition{conditionMet: true}
	// Create a new ConditionalText instance
	conditionalText := storyboard.NewConditionalText("Old text", condition)

	// Set a new text
	conditionalText.SetText("New text")

	// Test if the text was updated
	if conditionalText.GetText() != "New text" {
		t.Errorf("Expected text to be updated to 'New text', got '%s'", conditionalText.GetText())
	}
}
