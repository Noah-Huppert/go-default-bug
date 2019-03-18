package main

import (
	"testing"

	"gopkg.in/mcuadros/go-defaults.v1"
	"gotest.tools/assert"
)

// example is a struct used to show the bug behavior
type example struct {
	// Foo is a field which defaults to true
	Foo bool `default:"true"`
}

// TestDefaults ensures example.Foo defaults to true when not set
func TestDefaults(t *testing.T) {
	// Set defaults
	actual := example{}

	defaults.SetDefaults(&actual)

	// Assert
	assert.Equal(t, actual.Foo, true)
}

// TestSetFalse ensures example.Foo == false when set to false
func TestSetFalse(t *testing.T) {
	// Set
	actual := example{
		Foo: false,
	}

	defaults.SetDefaults(&actual)

	// Assert
	assert.Equal(t, actual.Foo, false)
}
