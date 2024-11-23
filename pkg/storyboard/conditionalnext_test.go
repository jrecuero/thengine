package storyboard_test

import (
	"testing"

	"github.com/jrecuero/thengine/pkg/storyboard"
)

// Regular condition struct for testing
type Condition struct {
	conditionMet bool
}

func (c *Condition) GetConditionMet() bool {
	return c.conditionMet
}

func (c *Condition) SetConditionMet(value bool) {
	c.conditionMet = value
}

func TestNewConditionalNext(t *testing.T) {
	// Create a condition instance
	condition := &Condition{conditionMet: true}
	// Create a new ConditionalNext instance
	conditionalNext := storyboard.NewConditionalNext("NodeA", condition)

	// Test the initial values
	if conditionalNext.GetNode() != "NodeA" {
		t.Errorf("Expected node to be 'NodeA', got '%s'", conditionalNext.GetNode())
	}
	if conditionalNext.GetCondition() != condition {
		t.Errorf("Expected condition to match the given condition")
	}
}

func TestConditinalNextSetCondition(t *testing.T) {
	// Create an initial condition instance
	condition := &Condition{conditionMet: false}
	// Create a new ConditionalNext instance
	conditionalNext := storyboard.NewConditionalNext("NodeA", condition)

	// Set a new condition
	newCondition := &Condition{conditionMet: true}
	conditionalNext.SetCondition(newCondition)

	// Test if the condition was updated
	if conditionalNext.GetCondition() != newCondition {
		t.Errorf("Expected condition to be updated to newCondition")
	}
}

func TestConditinalNextSetNode(t *testing.T) {
	// Create a condition instance
	condition := &Condition{conditionMet: true}
	// Create a new ConditionalNext instance
	conditionalNext := storyboard.NewConditionalNext("NodeA", condition)

	// Set a new node
	conditionalNext.SetNode("NodeB")

	// Test if the node was updated
	if conditionalNext.GetNode() != "NodeB" {
		t.Errorf("Expected node to be updated to 'NodeB', got '%s'", conditionalNext.GetNode())
	}
}