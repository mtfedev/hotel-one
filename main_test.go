package main

import (
	"reflect"
	"testing"
)

func TestEqualPlayer(t *testing.T) {
	expected := Player{
		name: "Tom",
		hp:   120,
	}
	far := Player{
		name: "Alice",
		hp:   32,
	}
	if !reflect.DeepEqual(expected, far) {
		t.Errorf("expected %+v but got %+v", expected, far)

	}
}

func TestCalculateValues(t *testing.T) {
	var (
		expected = 10
		a        = 5
		b        = 5
	)
	have := calculateValues(a, b)
	if have != expected {
		t.Errorf("expected %d but have %d", expected, have)
	}
}
