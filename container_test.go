package framework

import (
	"testing"
)

func TestNewApp(t *testing.T) {
	c := NewContainer()
	if c == nil {
		t.Error("Expected to get a pointer to the container, but got nil instead")
	}
}
